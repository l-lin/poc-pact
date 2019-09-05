package add

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/l-lin/doggy-cli/dog"
)

// Service to add dogs
type Service struct {
	URL string
}

// Add to the petstore
func (s *Service) Add(c dog.Dog) (*dog.Dog, error) {
	url := fmt.Sprintf("%s", s.URL)
	reqBody, err := json.Marshal(c)
	if err != nil {
		log.Fatal(err)
	}
	resp, err := http.Post(url, "application/json", bytes.NewReader(reqBody))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != http.StatusCreated {
		return nil, fmt.Errorf("Could not create the dog. Status code was: %d. Error was: '%v'", resp.StatusCode, string(body[:]))
	}
	var result *dog.Dog
	err = json.Unmarshal(body, &result)
	if err != nil {
		return nil, err
	}

	return result, nil
}
