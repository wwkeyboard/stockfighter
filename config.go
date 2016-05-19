package stockfighter

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"os"
)

// Configuration contains everything needed to run the program
type Configuration struct {
	VenueName    string
	VenueAccount string
	Token        string
	Stock        string
}

// ReadConfig builds a Configuration from the file at filename
func ReadConfig(filename string) (*Configuration, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}

	body, err := ioutil.ReadAll(file)
	if err != nil {
		return nil, err
	}
	// Trim the UTF-8 BOM, if it's there
	prepedBody := bytes.TrimPrefix(body, []byte{239, 187, 191})

	configuration := Configuration{}
	err = json.Unmarshal(prepedBody, &configuration)
	if err != nil {
		return nil, err
	}

	return &configuration, nil
}
