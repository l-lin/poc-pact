package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
	"testing"
	"time"

	"github.com/pact-foundation/pact-go/dsl"
	"github.com/pact-foundation/pact-go/types"
)

var (
	pact    dsl.Pact
	dir, _  = os.Getwd()
	pactDir = fmt.Sprintf("%s/pacts", dir)
	logDir  = fmt.Sprintf("%s/logs", dir)
)

const (
	consumer = "kitty-cli"
	provider = "petstore"
)

func TestMain(m *testing.M) {
	// setup pacts
	setup()

	// run all tests
	code := m.Run()

	// shutdown the mock service and write the pact files to disk
	pact.WritePact()
	pact.Teardown()

	// publish pact if env variable is set
	if os.Getenv("PACT_PUBLISH") != "" {
		version := fmt.Sprintf("0.0.1-%d", time.Now().Unix())

		p := dsl.Publisher{}
		err := p.Publish(types.PublishRequest{
			PactURLs:        []string{filepath.FromSlash(fmt.Sprintf("%s/%s-%s.json", pactDir, consumer, provider))},
			PactBroker:      os.Getenv("PACT_BROKER_HOST"),
			ConsumerVersion: version,
			Tags:            strings.Split(strings.Trim(os.Getenv("PACT_TAGS"), " "), ","),
			BrokerUsername:  os.Getenv("PACT_BROKER_USERNAME"),
			BrokerPassword:  os.Getenv("PACT_BROKER_PASSWORD"),
		})
		if err != nil {
			log.Println(err)
			os.Exit(1)
		}
	} else {
		log.Println("Skipping publishing pact result to pact broker")
	}
	os.Exit(code)
}

func setup() {
	pact = dsl.Pact{
		Consumer: consumer,
		Provider: provider,
		PactDir:  pactDir,
		LogDir:   logDir,
	}
}
