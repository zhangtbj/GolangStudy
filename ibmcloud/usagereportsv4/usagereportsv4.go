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

	// FID
	accountID := "27212a64558a41a9b19577d4232fc3f2"
	// GBS
	//accountID := "17a1a02adc9f4e1393a1cb9d1a5aec8a"

	resourceUsageOption := usageReport.NewGetResourceUsageAccountOptions(accountID, billingMonth)
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
	fmt.Printf("GetResourceUsageAccount result count is %d\n\n", len(instanceUsages))

	accountUsageOptions := usageReport.NewGetAccountUsageOptions(accountID, billingMonth)
	accountUsageOptions.SetNames(true)
	accountUsage, _, err := usageReport.GetAccountUsage(accountUsageOptions)
	if err != nil {
		return false, err
	}

	var sumAccountUsage float64
	for i:=0;i<len(accountUsage.Resources);i++{
		sumAccountUsage += *accountUsage.Resources[i].BillableCost
	}
	//fmt.Printf("GetAccountSummary response is %s", response)
	fmt.Printf("GetAccountUsage result count is %d\n", len(accountUsage.Resources))
	fmt.Printf("GetAccountUsage sum cost is %f\n\n", sumAccountUsage)

	accountSummaryOptions := usageReport.NewGetAccountSummaryOptions(accountID, billingMonth)
	accountSummary, _, err := usageReport.GetAccountSummary(accountSummaryOptions)
	if err != nil {
		return false, err
	}
	fmt.Printf("GetAccountSummary sum cost is %f\n\n", *accountSummary.Resources.BillableCost)

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
