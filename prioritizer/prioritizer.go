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

	var ids int = 1
	for _, v := range links {

		if v.Knowledge != "" {
			parsed, err := strconv.ParseInt(v.Knowledge, 10, 32)
			if err != nil {
				log.Fatal(err)
			}
			if int(parsed) == questions.Knowledge {
				v.ID = ids
				ids++
				result = append(result, v)
			}
		}
	}
	return result
}
