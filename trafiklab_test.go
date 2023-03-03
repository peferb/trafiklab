package trafiklab

import (
	"net/url"
	"strings"
	"testing"
)

// TestGetSLRealtimeDeparturesV4 calls trafiklab.TestGetSLRealtimeDeparturesV4 siteId 9710 and wants status 0
func TestGetSLRealtimeDeparturesV4(t *testing.T) {
	q := url.Values{}
	q.Set("Siteid", "9710")
	data, err := GetSLRealtimeDeparturesV4(q)

	if strings.Contains(data, `"StatusCode":0`) {
		t.Fatalf(`TestGetSLRealtimeDeparturesV4(siteid=1000) data did not contain "StatusCode:0" returned "%s""`, data)
	}
	if err != nil {
		t.Fatalf(`TestGetSLRealtimeDeparturesV4(siteid=1000) threw error "%s""`, err)
	}
}
