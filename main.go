package main

import (
	"context"
	"fmt"
	"os"

	openapiclient "github.com/leaseweb/leaseweb-go-sdk/abuse"
)

func main() {
	limit := int32(20)              // int32 | Limit the number of results returned. (optional)
	offset := int32(10)             // int32 | Return results starting from the given offset. (optional)
	status := "OPEN,WAITING,CLOSED" // string | Comma separated list of report statuses to filter on.  (optional) (default to "OPEN,WAITING,CLOSED")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AbuseAPI.GetReportList(context.Background()).Limit(limit).Offset(offset).Status(status).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AbuseAPI.GetReportList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetReportList`: GetReportListResult
	fmt.Fprintf(os.Stdout, "Response from `AbuseAPI.GetReportList`: %v\n", resp)
}
