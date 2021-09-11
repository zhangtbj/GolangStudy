package main

import (
	"errors"
	"fmt"
	"github.com/IBM/go-sdk-core/v5/core"
	"github.com/IBM/platform-services-go-sdk/globalsearchv2"
	"strconv"
)

const (
	globalSearchLimit              int64 = 1000
	globalSearchQuery                    = "*"
)

func getResources(creds core.BearerTokenAuthenticator) (bool, error) {
	authenticator := &core.BearerTokenAuthenticator{
		BearerToken: creds.BearerToken,
	}
	globalSearch, err := globalsearchv2.NewGlobalSearchV2UsingExternalConfig(
		&globalsearchv2.GlobalSearchV2Options{
			Authenticator: authenticator,
		},
	)
	if err != nil {
		return false, err
	}
	searchOptions := globalSearch.NewSearchOptions()

	s := ""
	var searchCursor *string = &s
	var allResources []globalsearchv2.ResultItem

	for searchCursor != nil {
		searchOptions.SetLimit(globalSearchLimit)
		searchOptions.SetQuery(globalSearchQuery)
		searchOptions.SetFields([]string{"*"})

		if searchCursor != nil {
			if *searchCursor != "" {
				searchOptions.SetSearchCursor(*searchCursor)
			}
		} else {
			break
		}

		scanResult, response, err := globalSearch.Search(searchOptions)
		if err != nil {
			return false, err
		}

		if response.StatusCode < 200 || response.StatusCode >= 300 {
			return false, errors.New("the IBM Global Search API response is not successful, statusCode is "+
				strconv.Itoa(response.StatusCode)+", result is "+response.GetResult().(string))
		}

		allResources = append(allResources, scanResult.Items...)
		searchCursor = scanResult.SearchCursor
	}

	fmt.Printf("resource count is %d", len(allResources))

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
		// BearerToken: "xxx",
		BearerToken: "xxx",
	}

	//result, err := getSupportedTypes(creds)
	//if err != nil {
	//	panic(err.Error())
	//}
	//
	//for i:=0; i< len(result.SupportedTypes); i++ {
	//	fmt.Printf("Type: %s \n", result.SupportedTypes[i])
	//}

	_, err := getResources(creds)
	if err != nil {
		panic(err.Error())
	}
}

