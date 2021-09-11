package main

import (
	"fmt"
	"github.com/IBM/go-sdk-core/v5/core"
	"github.com/IBM/platform-services-go-sdk/usagereportsv4"
	"time"
)

func getResourceUsageData(creds core.BearerTokenAuthenticator) (bool, error) {
	usageReport, err := usagereportsv4.NewUsageReportsV4UsingExternalConfig(
		&usagereportsv4.UsageReportsV4Options{
			Authenticator: &creds,
		},
	)
	if err != nil {
		return false, err
	}

	timeUnix := time.Now().AddDate(0, 0, 1).Unix()
	billingMonth := time.Unix(timeUnix, 0).Format("2006-01")

	resourceUsageOption := usageReport.NewGetResourceUsageAccountOptions("27212a64558a41a9b19577d4232fc3f2", billingMonth)
	resourceUsageOption.SetNames(true)
	resourceUsageOption.SetLimit(100)

	var instanceUsages []usagereportsv4.InstanceUsage
	for {
		instancesUsage, _, err := usageReport.GetResourceUsageAccount(resourceUsageOption)
		if err != nil {
			return false, err
		}
		instanceUsages = append(instanceUsages, instancesUsage.Resources...)
		if instancesUsage.Next == nil {
			break
		}
		resourceUsageOption.SetStart(*instancesUsage.Next.Offset)
	}
	fmt.Printf("GetResourceUsageAccount result count is %d", len(instanceUsages))

	accountUsageOptions := usageReport.NewGetAccountUsageOptions("27212a64558a41a9b19577d4232fc3f2", billingMonth)
	accountUsageOptions.SetNames(true)
	accountUsage, response, err := usageReport.GetAccountUsage(accountUsageOptions)
	if err != nil {
		return false, err
	}
	fmt.Printf("GetAccountSummary response is %s", response)
	fmt.Printf("GetAccountSummary result count is %d", accountUsage.Resources)

	accountSummaryOptions := usageReport.NewGetAccountSummaryOptions("27212a64558a41a9b19577d4232fc3f2", billingMonth)
	accountSummary, response, err := usageReport.GetAccountSummary(accountSummaryOptions)
	if err != nil {
		return false, err
	}
	fmt.Printf("GetAccountSummary response is %s", response)
	fmt.Printf("GetAccountSummary result count is %d", accountSummary.Resources)

	return true, nil
}

func main() {
	creds := core.BearerTokenAuthenticator{
		// BearerToken: "xxx",
		BearerToken: "xxx",
	}

	_, err := getResourceUsageData(creds)
	if err != nil {
		panic(err.Error())
	}
}
