package main

import (
	"fmt"
	"testing"

	"github.com/l-lin/kitty-cli/add"
	"github.com/l-lin/kitty-cli/cat"
	"github.com/l-lin/kitty-cli/list"
	"github.com/pact-foundation/pact-go/dsl"
)

const (
	catsPath = "/cats"
)

func TestGet_Cat(t *testing.T) {
	// test case
	id := int64(88)
	var test = func() error {
		s := list.Service{
			URL: fmt.Sprintf("http://%s:%d%s", pact.Host, pact.Server.Port, catsPath),
		}
		if _, err := s.Get(id); err != nil {
			return err
		}
		return nil
	}
	// set up expected interactions
	pact.
		AddInteraction().
		Given(fmt.Sprintf("there is a cat with an id %d", id)).
		UponReceiving(fmt.Sprintf("a request to get cat id %d", id)).
		WithRequest(dsl.Request{
			Method: "GET",
			Path:   dsl.String(fmt.Sprintf("%s/%d", catsPath, id)),
		}).
		WillRespondWith(dsl.Response{
			Status:  200,
			Headers: dsl.MapMatcher{"Content-Type": dsl.String("application/json")},
			Body:    dsl.Match(cat.Cat{}),
		})

	// verify
	if err := pact.Verify(test); err != nil {
		t.Error(err)
	}
}

func TestGet_CatNotExists(t *testing.T) {
	// test case
	id := int64(888)
	var test = func() error {
		s := list.Service{
			URL: fmt.Sprintf("http://%s:%d%s", pact.Host, pact.Server.Port, catsPath),
		}
		if c, err := s.Get(id); err == nil {
			return fmt.Errorf("expected an error, got a cat %v instead", c)
		}
		return nil
	}
	// set up expected interactions
	pact.
		AddInteraction().
		Given(fmt.Sprintf("there is no cat with an id %d", id)).
		UponReceiving(fmt.Sprintf("a request to get cat id %d", id)).
		WithRequest(dsl.Request{
			Method: "GET",
			Path:   dsl.String(fmt.Sprintf("%s/%d", catsPath, id)),
		}).
		WillRespondWith(dsl.Response{
			Status:  404,
			Headers: dsl.MapMatcher{"Content-Type": dsl.String("application/json")},
		})

	// verify
	if err := pact.Verify(test); err != nil {
		t.Error(err)
	}
}

func TestAdd_Cat(t *testing.T) {
	// test case
	catName := "Grumpy cat"
	catType := "Tardar Sauce"
	var test = func() error {
		s := add.Service{
			URL: fmt.Sprintf("http://%s:%d%s", pact.Host, pact.Server.Port, catsPath),
		}
		if _, err := s.Add(cat.Cat{Name: catName, Type: catType}); err != nil {
			return err
		}
		return nil
	}
	// set up expected interactions
	pact.
		AddInteraction().
		Given(fmt.Sprintf("creating a %s cat whose name is %s", catType, catName)).
		UponReceiving(fmt.Sprintf("a request to add a %s cat whose name is %s", catType, catName)).
		WithRequest(dsl.Request{
			Method:  "POST",
			Headers: dsl.MapMatcher{"Content-Type": dsl.String("application/json")},
			Path:    dsl.String(catsPath),
			Body:    dsl.Match(cat.Cat{}),
		}).
		WillRespondWith(dsl.Response{
			Status:  201,
			Headers: dsl.MapMatcher{"Content-Type": dsl.String("application/json")},
			Body:    dsl.Match(cat.Cat{}),
		})

	// verify
	if err := pact.Verify(test); err != nil {
		t.Error(err)
	}
}
