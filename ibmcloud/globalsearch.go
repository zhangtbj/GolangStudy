package main

import (
	"encoding/json"
	"fmt"
	"github.com/IBM/go-sdk-core/v5/core"
	"github.com/IBM/platform-services-go-sdk/globalsearchv2"
)

func getResources(creds core.BearerTokenAuthenticator) (bool, error) {
	globalSearch, err := globalsearchv2.NewGlobalSearchV2UsingExternalConfig(
		&globalsearchv2.GlobalSearchV2Options{
			Authenticator: &creds,
		},
	)
	searchOptions := globalSearch.NewSearchOptions()
	searchOptions.SetLimit(1000)
	searchOptions.SetQuery("type:resource-instance OR type:cf-application")
	searchOptions.SetFields([]string{"name", "crn", "region", "tags", "service_name", "account_id", "type", "service_instance", "doc.resource_group_id"})
	//searchOptions.SetQuery("*")
	//searchOptions.SetFields([]string{"*"})
	scanResult, response, err := globalSearch.Search(searchOptions)
	if err != nil {
		return false, err
	}
	b, _ := json.MarshalIndent(scanResult, "", "  ")
	fmt.Printf("\nSearch() result:\n%s\n", string(b))
	fmt.Printf("\nSearch() response:\n%s\n", response.String())
	return true, nil
}

func getSupportedTypes(creds core.BearerTokenAuthenticator) (globalsearchv2.SupportedTypesList, error) {
	globalSearch, err := globalsearchv2.NewGlobalSearchV2UsingExternalConfig(
		&globalsearchv2.GlobalSearchV2Options{
			Authenticator: &creds,
		},
	)
	if err != nil {
		return globalsearchv2.SupportedTypesList{}, err
	}

	result, response, err := globalSearch.GetSupportedTypes(globalSearch.NewGetSupportedTypesOptions())
	if err != nil {
		return globalsearchv2.SupportedTypesList{}, err
	}

	fmt.Printf("\nSearch() result:\n%s\n", result)
	fmt.Printf("\nSearch() response:\n%s\n", response.String())

	return *result, nil
}

func main() {
	creds := core.BearerTokenAuthenticator {
		BearerToken: "xxx",
	}

	result, err := getSupportedTypes(creds)
	if err != nil {
		panic(err.Error())
	}

	for i:=0; i< len(result.SupportedTypes); i++ {
		fmt.Printf("Type: %s \n", result.SupportedTypes[i])
	}
}
