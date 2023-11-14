package newrelicimpl

import (
	"fmt"
	"log"
	"os"

	"github.com/newrelic/go-agent/v3/integrations/logcontext-v2/logWriter"
	"github.com/newrelic/go-agent/v3/newrelic"
)

var newrelicApp *newrelic.Application
var newrelicLog *log.Logger
var newrelicWriter logWriter.LogWriter

func Initialize(newrelicAppName, newrelicLicenseKey string) {
	//if newrelicAppName is empty skip rest of the program
	if newrelicAppName != "" {

		fmt.Println("Newrelic load")
		app, err := newrelic.NewApplication(
			newrelic.ConfigAppName(newrelicAppName),
			newrelic.ConfigLicense(newrelicLicenseKey))

		if err != nil {
			fmt.Printf("Error initialize new relic %s\n", err.Error())
		} else {
			newrelicApp = app
		}

	} else {
		fmt.Println("Newrelic not specific")
	}

	newrelicWriter = logWriter.New(os.Stdout, newrelicApp)
	newrelicLog = log.New(&newrelicWriter, "", log.Default().Flags())

}

func GetNewrelicApp() *newrelic.Application {
	return newrelicApp
}

func GetNewrelicLog() *log.Logger {
	return newrelicLog
}
