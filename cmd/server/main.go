package main

import (
	"log"
	"os"

	"github.com/go-openapi/loads"
	flags "github.com/jessevdk/go-flags"

	"server/internal/server/restapi"
	"server/internal/server/restapi/operations"
)

func main() {

	swaggerSpec, err := loads.Embedded(restapi.SwaggerJSON, restapi.FlatSwaggerJSON)
	if err != nil {
		log.Fatal(err.Error())
	}

	api := operations.NewFlightsAPI(swaggerSpec)
	api.Logger = log.Printf
	server := restapi.NewServer(api)
	defer server.Shutdown() // nolint

	parser := flags.NewParser(server, flags.Default)
	parser.ShortDescription = "Flights API"
	parser.LongDescription = "#### API for flights management\"\n"
	server.ConfigureFlags()
	for _, optsGroup := range api.CommandLineOptionsGroups {
		_, err := parser.AddGroup(optsGroup.ShortDescription, optsGroup.LongDescription, optsGroup.Options)
		if err != nil {
			log.Fatal(err.Error())
		}
	}

	if _, err := parser.Parse(); err != nil {
		code := 1
		if fe, ok := err.(*flags.Error); ok {
			if fe.Type == flags.ErrHelp {
				code = 0
			}
		}
		os.Exit(code)
	}

	server.ConfigureAPI()

	if err := server.Serve(); err != nil {
		log.Fatal(err.Error())
	}
}
