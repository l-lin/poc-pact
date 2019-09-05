package list

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/l-lin/kitty-cli/cat"
)

// Service to list cats
type Service struct {
	URL string
}

// Get from a given ID
func (s Service) Get(id int64) (*cat.Cat, error) {
	url := fmt.Sprintf("%s/%d", s.URL, id)
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("Could not get the cat with id %d. Status code was: %d. Error was: '%v'", id, resp.StatusCode, string(body[:]))
	}
	var c *cat.Cat
	err = json.Unmarshal(body, &c)
	if err != nil {
		return nil, err
	}
	return c, nil
}
