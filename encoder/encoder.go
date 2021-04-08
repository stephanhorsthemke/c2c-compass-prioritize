package encoder

import (
	"encoding/json"
	"log"

	"github.com/stephanhorsthemke/c2c-compass-prioritize/gsheet"
)

func LinksToJson(l []gsheet.Link) []byte {

	j, err := json.Marshal(l)
	if err != nil {
		log.Fatal(err)
	}
	return j
}
