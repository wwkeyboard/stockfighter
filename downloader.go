package venue

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

// Downloader actually calls the API and returns the unprocessed response.
// It should handle authentication and recoverable HTTP errors.
type Downloader interface {
	GetJSON(path string) ([]byte, error)
}

// HTTPDownloader implements Downloader by calling the production API
type HTTPDownloader struct {
	BaseURL string
}

// GetJSON calls the path and populates the obj with the response.
func (d HTTPDownloader) GetJSON(path string) (r []byte, err error) {
	url := fmt.Sprintf("%s%s", d.BaseURL, path)
	resp, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}
