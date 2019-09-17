package main

import (
	"fmt"
	"testing"

	"github.com/l-lin/doggy-cli/add"
	"github.com/l-lin/doggy-cli/dog"
	"github.com/l-lin/doggy-cli/list"
	"github.com/pact-foundation/pact-go/dsl"
)

const (
	dogsPath = "/dogs"
)

func TestGet_Dog(t *testing.T) {
	// test case
	id := int64(88)
	var test = func() error {
		s := list.Service{
			URL: fmt.Sprintf("http://%s:%d%s", pact.Host, pact.Server.Port, dogsPath),
		}
		if _, err := s.Get(id); err != nil {
			return err
		}
		return nil
	}
	// set up expected interactions
	pact.
		AddInteraction().
		Given(fmt.Sprintf("there is a dog with an id %d", id)).
		UponReceiving(fmt.Sprintf("a request to get dog id %d", id)).
		WithRequest(dsl.Request{
			Method: "GET",
			Path:   dsl.String(fmt.Sprintf("%s/%d", dogsPath, id)),
		}).
		WillRespondWith(dsl.Response{
			Status:  200,
			Headers: dsl.MapMatcher{"Content-Type": dsl.String("application/json;charset=UTF-8")},
			Body:    dsl.Match(dog.Dog{}),
		})

	// verify
	if err := pact.Verify(test); err != nil {
		t.Error(err)
	}
}

func TestGet_DogNotExists(t *testing.T) {
	// test case
	id := int64(888)
	var test = func() error {
		s := list.Service{
			URL: fmt.Sprintf("http://%s:%d%s", pact.Host, pact.Server.Port, dogsPath),
		}
		if c, err := s.Get(id); err == nil {
			return fmt.Errorf("expected an error, got a dog %v instead", c)
		}
		return nil
	}
	// set up expected interactions
	pact.
		AddInteraction().
		Given(fmt.Sprintf("there is no dog with an id %d", id)).
		UponReceiving(fmt.Sprintf("a request to get dog id %d", id)).
		WithRequest(dsl.Request{
			Method: "GET",
			Path:   dsl.String(fmt.Sprintf("%s/%d", dogsPath, id)),
		}).
		WillRespondWith(dsl.Response{
			Status:  404,
			Headers: dsl.MapMatcher{"Content-Type": dsl.String("application/json;charset=UTF-8")},
		})

	// verify
	if err := pact.Verify(test); err != nil {
		t.Error(err)
	}
}

func TestAdd_Dog(t *testing.T) {
	// test case
	dogName := "Chico"
	dogType := "Shiba Inu"
	var test = func() error {
		s := add.Service{
			URL: fmt.Sprintf("http://%s:%d%s", pact.Host, pact.Server.Port, dogsPath),
		}
		if _, err := s.Add(dog.Dog{Name: dogName, Type: dogType}); err != nil {
			return err
		}
		return nil
	}
	// set up expected interactions
	pact.
		AddInteraction().
		Given(fmt.Sprintf("creating a %s dog whose name is %s", dogType, dogName)).
		UponReceiving(fmt.Sprintf("a request to add a %s dog whose name is %s", dogType, dogName)).
		WithRequest(dsl.Request{
			Method:  "POST",
			Headers: dsl.MapMatcher{"Content-Type": dsl.String("application/json;charset=UTF-8")},
			Path:    dsl.String(dogsPath),
			Body:    dsl.Match(dog.Dog{}),
		}).
		WillRespondWith(dsl.Response{
			Status:  201,
			Headers: dsl.MapMatcher{"Content-Type": dsl.String("application/json;charset=UTF-8")},
			Body:    dsl.Match(dog.Dog{}),
		})

	// verify
	if err := pact.Verify(test); err != nil {
		t.Error(err)
	}
}
