package add

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/l-lin/kitty-cli/cat"
)

// Service to add cats
type Service struct {
	URL string
}

// Add to the petstore
func (s *Service) Add(c cat.Cat) (*cat.Cat, error) {
	url := fmt.Sprintf("%s", s.URL)
	reqBody, err := json.Marshal(c)
	if err != nil {
		log.Fatal(err)
	}
	resp, err := http.Post(url, "application/json;charset=UTF-8", bytes.NewReader(reqBody))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != http.StatusCreated {
		return nil, fmt.Errorf("Could not create the pet. Status code was: %d. Error was: '%v'", resp.StatusCode, string(body[:]))
	}
	var result *cat.Cat
	err = json.Unmarshal(body, &result)
	if err != nil {
		return nil, err
	}

	return result, nil
}
