package aur

import (
	"encoding/json"
	"io"
	"net/http"
	"os"
)

// Editor gives the default system editor, uses vi in last case
var Editor = "vi"

func init() {
	if os.Getenv("EDITOR") != "" {
		Editor = os.Getenv("EDITOR")
	}
}

// getJSON handles JSON retrieval and decoding to struct
func getJSON(url string, target interface{}) error {
	r, err := http.Get(url)
	if err != nil {
		return err
	}
	defer r.Body.Close()

	return json.NewDecoder(r.Body).Decode(target)
}

func downloadFile(filepath string, url string) (err error) {
	// Create the file
	out, err := os.Create(filepath)
	if err != nil {
		return err
	}
	defer out.Close()

	// Get the data
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// Writer the body to file
	_, err = io.Copy(out, resp.Body)
	return err
}
