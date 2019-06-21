package project

import (
	"bytes"
	"log"

	"github.com/BurntSushi/toml"
	"github.com/typical-go/typical-rest-server/typical"
)

// ContextDetail return context detail string
func ContextDetail() string {
	buf := new(bytes.Buffer)
	if err := toml.NewEncoder(buf).Encode(typical.Context); err != nil {
		log.Fatal(err)
	}
	return buf.String()
}
