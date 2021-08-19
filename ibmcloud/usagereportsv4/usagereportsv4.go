package main

import (
	"fmt"
	"github.com/IBM/go-sdk-core/v5/core"
	"github.com/IBM/platform-services-go-sdk/usagereportsv4"
)

func getResourceUsageAccount(creds core.BearerTokenAuthenticator) (bool, error) {
	usageReport, err := usagereportsv4.NewUsageReportsV4UsingExternalConfig(
		&usagereportsv4.UsageReportsV4Options{
			Authenticator: &creds,
		},
		)
	if err != nil {
		return false, err
	}

	resourceUsageOption := usageReport.NewGetResourceUsageAccountOptions("27212a64558a41a9b19577d4232fc3f2", "2021-08")
	instancesUsage, response, err := usageReport.GetResourceUsageAccount(resourceUsageOption)
	if err != nil {
		return false, err
	}

	fmt.Printf("GetResourceUsageAccount response is %s", response)
	fmt.Printf("GetResourceUsageAccount result count is %d", instancesUsage.Count)

	accountSummaryOptions := usageReport.NewGetAccountSummaryOptions("27212a64558a41a9b19577d4232fc3f2", "2021-08")
	accountSummary, response, err := usageReport.GetAccountSummary(accountSummaryOptions)
	if err != nil {
		return false, err
	}

	fmt.Printf("GetAccountSummary response is %s", response)
	fmt.Printf("GetAccountSummary result count is %d", accountSummary.Resources)

	accountUsageOptions := usageReport.NewGetAccountUsageOptions("27212a64558a41a9b19577d4232fc3f2", "2021-08")
	accountUsage, response, err := usageReport.GetAccountUsage(accountUsageOptions)
	if err != nil {
		return false, err
	}

	fmt.Printf("GetAccountSummary response is %s", response)
	fmt.Printf("GetAccountSummary result count is %d", accountUsage.Resources)

	return true, nil
}

func main() {
	creds := core.BearerTokenAuthenticator {
		//BearerToken: "xxx",
		BearerToken: "xxx",
	}

	_, err := getResourceUsageAccount(creds)
	if err != nil {
		panic(err.Error())
	}
}