package gsheet

import (
	"fmt"
	"io/ioutil"
	"log"

	"golang.org/x/net/context"
	"google.golang.org/api/option"
	"google.golang.org/api/sheets/v4"
)

const API_KEY = "AIzaSyCo3_7miV2h1_2BpD9zK0nBy4YxvuxAopU"
const FILE_ID = "1ptowWZ8FuWWJOOTRHnQvVdb9Z3WPNv7mxNMPLK8Es6o"

func GetGoogleSheet() {
	ctx := context.Background()

	b, err := ioutil.ReadFile("secret/gdrive.json")
	if err != nil {
		log.Fatalf("Unable to read client secret file: %v", err)
	}

	svc, err := sheets.NewService(ctx, option.WithCredentialsJSON(b))

	readRange := "'links'!A1:B2"
	resp, err := svc.Spreadsheets.Values.Get(FILE_ID, readRange).Context(ctx).Do()

	fmt.Printf("%#v ", resp)
}
