package morpheus_test

import (
	"testing"

	"github.com/gomorpheus/morpheus-go-sdk"
)

func TestMonitoringApps(t *testing.T) {
	client := getTestClient(t)
	req := &morpheus.Request{}
	resp, err := client.ListMonitoringApps(req)
	assertResponse(t, resp, err)
}

func TestGetMonitoringApp(t *testing.T) {
	client := getTestClient(t)
	req := &morpheus.Request{}
	resp, err := client.ListMonitoringApps(req)
	assertResponse(t, resp, err)

	// parse JSON and fetch the first one by ID

	result := resp.Result.(*morpheus.ListMonitoringAppsResult)
	recordCount := result.Meta.Total
	t.Logf("Found %d Monitoring Apps.", recordCount)
	if recordCount != 0 {
		// Get by ID
		record := (*result.MonitoringApps)[0]
		resp, err = client.GetMonitoringApp(record.ID, &morpheus.Request{})
		assertResponse(t, resp, err)
	}
}
