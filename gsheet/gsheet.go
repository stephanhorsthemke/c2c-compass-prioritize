package gsheet

import (
	"fmt"
	"io/ioutil"
	"log"
	"time"

	"github.com/go-co-op/gocron"
	"golang.org/x/net/context"
	"google.golang.org/api/option"
	"google.golang.org/api/sheets/v4"
)

const API_KEY = "AIzaSyCo3_7miV2h1_2BpD9zK0nBy4YxvuxAopU"
const FILE_ID = "1ptowWZ8FuWWJOOTRHnQvVdb9Z3WPNv7mxNMPLK8Es6o"

// The range the gsheet is read up to
const READ_RANGE = "'links'!A3:Z1000"

var links []Link

type Link struct {
	// was passiert mit mails?
	ID          int    `json:"id"`
	Link        string `json:"link"`
	Language    string `json:"language"`
	Description string `json:"description"`
	Knowledge   string `json:"knowledge"`
	Position    string `json:"position"`
	Sector      string `json:"sector"`
}

func getGoogleSheet() [][]interface{} {
	ctx := context.Background()

	b, err := ioutil.ReadFile("secret/gdrive.json")
	if err != nil {
		log.Fatalf("Unable to read client secret file: %v", err)
	}

	svc, err := sheets.NewService(ctx, option.WithCredentialsJSON(b))

	resp, err := svc.Spreadsheets.Values.Get(FILE_ID, READ_RANGE).Context(ctx).Do()

	return resp.Values
}

func parseLinks(values [][]interface{}) []Link {

	links := make([]Link, len(values))
	for i, row := range values {
		l := Link{}

		length := len(row)
		switch {
		case length >= 6:
			l.Sector = row[5].(string)
			fallthrough
		case length >= 5:
			l.Position = row[4].(string)
			fallthrough
		case length >= 4:
			l.Knowledge = row[3].(string)
			fallthrough
		case length >= 3:
			l.Description = row[2].(string)
			fallthrough
		case length >= 2:
			l.Language = row[1].(string)
			fallthrough
		case length >= 1:
			l.Link = row[0].(string)
		}
		links[i] = l
	}
	return links
}

var getAndParse = func() {
	values := getGoogleSheet()
	links = parseLinks(values)
	fmt.Println("Links updated: ", links)
}

func UpdateLinks() {
	// Update the links hourly
	s := gocron.NewScheduler(time.UTC)
	job, err := s.Every(1).Hour().Do(getAndParse)
	if err != nil {
		log.Fatal("could not set up scheduler: ", err)
	}
	job.SingletonMode()
	s.StartAsync()
}

func GetLinks() []Link {
	return links
}
