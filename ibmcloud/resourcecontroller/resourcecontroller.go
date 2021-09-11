package main

import (
	"errors"
	"fmt"
	"github.com/IBM/go-sdk-core/v5/core"
	"github.com/IBM/platform-services-go-sdk/globalsearchv2"
	"github.com/IBM/platform-services-go-sdk/resourcecontrollerv2"
	"strconv"
	"strings"
	"time"
)

const (
	globalSearchLimit              int64 = 1000
	globalSearchQuery                    = "*"
	globalSearchTypeField                = "type"
	defaultIBMAccountRegion              = "global"
	costAPIRetryCount                    = 5
	costAPIRetryDelay                    = 10 * time.Second
	resourceControllerPagerKeyWord       = "start="
)

// Resource represents basic information about IBM Cloud resource.
type Resource struct {
	ID       string `json:"id"`
	Type     string `json:"type"`
	Location string `json:"location"`
}

// getResourcesFromGlobalSearch searches IBM Cloud resources by using IBM Global Search API.
func getResourcesFromGlobalSearch(creds core.BearerTokenAuthenticator) ([]globalsearchv2.ResultItem, error) {
	authenticator := &core.BearerTokenAuthenticator{
		BearerToken: creds.BearerToken,
	}
	globalSearch, err := globalsearchv2.NewGlobalSearchV2UsingExternalConfig(
		&globalsearchv2.GlobalSearchV2Options{
			Authenticator: authenticator,
		},
	)
	if err != nil {
		return nil, err
	}
	searchOptions := globalSearch.NewSearchOptions()

	s := ""
	var searchCursor *string = &s
	var allResources []globalsearchv2.ResultItem

	for searchCursor != nil {
		searchOptions.SetLimit(globalSearchLimit)
		searchOptions.SetQuery(globalSearchQuery)
		searchOptions.SetFields([]string{"*"})
		globalSearch.EnableRetries(costAPIRetryCount, costAPIRetryDelay)

		if searchCursor != nil {
			if *searchCursor != "" {
				searchOptions.SetSearchCursor(*searchCursor)
			}
		} else {
			break
		}

		scanResult, response, err := globalSearch.Search(searchOptions)
		if err != nil {
			return nil, err
		}

		if response.StatusCode < 200 || response.StatusCode >= 300 {
			return nil, errors.New("the IBM Global Search API response is not successful, statusCode is " +
				strconv.Itoa(response.StatusCode) + ", result is " + response.GetResult().(string))
		}

		allResources = append(allResources, scanResult.Items...)
		searchCursor = scanResult.SearchCursor
	}

	return allResources, nil
}

// getResourcesFromResourceController searches IBM Cloud resources by using IBM Resource Controller API.
func getResourcesFromResourceController(creds core.BearerTokenAuthenticator) ([]resourcecontrollerv2.ResourceInstance, error) {
	authenticator := &core.BearerTokenAuthenticator{
		BearerToken: creds.BearerToken,
	}

	resourceController, err := resourcecontrollerv2.NewResourceControllerV2UsingExternalConfig(
		&resourcecontrollerv2.ResourceControllerV2Options{
			Authenticator: authenticator,
		},
	)
	if err != nil {
		return nil, err
	}
	resourceOptions := resourceController.NewListResourceInstancesOptions()
	resourceOptions.SetLimit(500)
	var allResult []resourcecontrollerv2.ResourceInstance

	for {
		scanResult, response, err := resourceController.ListResourceInstances(resourceOptions)
		if err != nil {
			return nil, err
		}

		if response.StatusCode < 200 || response.StatusCode >= 300 {
			return nil, errors.New("the IBM Resource Controller API response is not successful, statusCode is " +
				strconv.Itoa(response.StatusCode) + ", result is " + response.GetResult().(string))
		}

		allResult = append(allResult, scanResult.Resources...)
		if scanResult.NextURL == nil {
			break
		}
		nextURL := *scanResult.NextURL
		if strings.Contains(*scanResult.NextURL, resourceControllerPagerKeyWord) {
			nextURL = (*scanResult.NextURL)[strings.Index(*scanResult.NextURL, resourceControllerPagerKeyWord)+6:]
		}
		resourceOptions.SetStart(nextURL)

	}

	return allResult, nil
}

func main() {
	authenticator := core.BearerTokenAuthenticator{
		// BearerToken: "xxx",
		BearerToken: "xxx",
	}

	resourcesFromGlobalSearch, err := getResourcesFromGlobalSearch(authenticator)
	if err != nil {
		fmt.Printf(err.Error())
	}
	fmt.Printf("the resource count from global search api is %d\n\n", len(resourcesFromGlobalSearch))

	resourcesFromResourceController, err := getResourcesFromResourceController(authenticator)
	if err != nil {
		fmt.Printf(err.Error())
	}
	fmt.Printf("the resource count from resource controller api is %d\n\n", len(resourcesFromResourceController))

	//	var firstOut = make([]Resource, len(resourcesFromGlobalSearch))
	resourcesMapFromGlobalSearch := make(map[string]globalsearchv2.ResultItem, len(resourcesFromGlobalSearch))

	resourceCount := len(resourcesFromGlobalSearch)
	//	resourcesFromGlobalSearchCount := resourceCount
	for _, gsr := range resourcesFromGlobalSearch {
		resourcesMapFromGlobalSearch[*gsr.CRN] = gsr
	}

	for _, rcr := range resourcesFromResourceController {
		if _, ok := resourcesMapFromGlobalSearch[*rcr.CRN]; !ok {
			resourceCount++
		}
	}
	fmt.Printf("the resource count after merge from two apis is %d\n", resourceCount)
}
