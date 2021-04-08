package prioritizer

import (
	"log"
	"strconv"

	"github.com/stephanhorsthemke/c2c-compass-prioritize/decoder"
	"github.com/stephanhorsthemke/c2c-compass-prioritize/gsheet"
)

/*
	This function sorts the links depending on the answered questions
*/
func Prioritize(questions decoder.Questions, links []gsheet.Link) []gsheet.Link {

	var result []gsheet.Link

	for _, v := range links {
		parsed, err := strconv.ParseInt(v.Knowledge, 10, 32)
		if err != nil {
			log.Fatal(err)
		}
		if int(parsed) == questions.Knowledge {
			result = append(result, v)
		}
	}
	return result
}
