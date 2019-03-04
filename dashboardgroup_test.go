package signalfx

import (
	"net/http"
	"net/url"
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

// TODO Create simple dashboard group

func TestCreateDashboardGroup(t *testing.T) {
	teardown := setup()
	defer teardown()

	mux.HandleFunc("/v2/dashboardgroup", verifyRequest(t, "POST", http.StatusOK, nil, "dashboardgroup/create_success.json"))

	result, err := client.CreateDashboardGroup(&DashboardGroup{
		Name: "string",
	})
	assert.NoError(t, err, "Unexpected error creating dashboard group")
	assert.Equal(t, "string", result.Name, "Name does not match")
}

func TestDeleteDashboardGroup(t *testing.T) {
	teardown := setup()
	defer teardown()

	mux.HandleFunc("/v2/dashboardgroup/string", verifyRequest(t, "DELETE", http.StatusOK, nil, ""))

	err := client.DeleteDashboardGroup("string")
	assert.NoError(t, err, "Unexpected error getting dashboard group")
}

func TestGetDashboardGroup(t *testing.T) {
	teardown := setup()
	defer teardown()

	mux.HandleFunc("/v2/dashboardgroup/string", verifyRequest(t, "GET", http.StatusOK, nil, "dashboardgroup/get_success.json"))

	result, err := client.GetDashboardGroup("string")
	assert.NoError(t, err, "Unexpected error getting dashboard group")
	assert.Equal(t, result.Name, "string", "Name does not match")
}

func TestSearchDashboardGroup(t *testing.T) {
	teardown := setup()
	defer teardown()

	limit := 10
	name := "foo"
	offset := 2
	tags := "bar"
	params := url.Values{}
	params.Add("limit", strconv.Itoa(limit))
	params.Add("name", name)
	params.Add("offset", strconv.Itoa(offset))
	params.Add("tags", tags)

	mux.HandleFunc("/v2/dashboardgroup", verifyRequest(t, "GET", http.StatusOK, params, "dashboardgroup/get_success.json"))

	results, err := client.SearchDashboardGroup(limit, name, offset, tags)
	assert.NoError(t, err, "Unexpected error search dashboard group")
	assert.Equal(t, int64(0), results.Count, "Incorrect number of results")
}

func TestUpdateDashboardGroup(t *testing.T) {
	teardown := setup()
	defer teardown()

	mux.HandleFunc("/v2/dashboardgroup/string", verifyRequest(t, "PUT", http.StatusOK, nil, "dashboardgroup/update_success.json"))

	result, err := client.UpdateDashboardGroup("string", &DashboardGroup{
		Name: "string",
	})
	assert.NoError(t, err, "Unexpected error updating dashboard group")
	assert.Equal(t, "string", result.Name, "Name does not match")
}
