package main

import (
	"fmt"
	"log"

	"github.com/gomorpheus/morpheus-go-sdk"
)

func main() {
	client := morpheus.NewClient("https://yourmorpheus.com")
	client.SetUsernameAndPassword("username", "password")
	resp, err := client.Login()
	if err != nil {
		fmt.Println("LOGIN ERROR: ", err)
	}
	fmt.Println("LOGIN RESPONSE:", resp)

	// List integrations
	req := &morpheus.Request{}
	integrationsResponse, err := client.ListIntegrations(req)
	if err != nil {
		log.Fatal(err)
	}
	result := integrationsResponse.Result.(*morpheus.ListIntegrationsResult)
	log.Println(result.Integrations)
}
