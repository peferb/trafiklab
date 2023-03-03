package trafiklab

import (
	"net/url"
	"testing"
)

// TestGetSLRealtimeDeparturesV4 calls trafiklab.TestGetSLRealtimeDeparturesV4 siteId 100o and wants status code "200 OK"
func TestGetSLRealtimeDeparturesV4(t *testing.T) {
	q := url.Values{}
	q.Set("format", "json")
	q.Set("Siteid", "1000")
	res, err := GetSLRealtimeDeparturesV4(q)
	if res.Status != "200 OK" || err != nil {
		t.Fatalf(`TestGetSLRealtimeDeparturesV4(siteid=1000) returns "%s" with "%s""`, res.Status, err)
	}
}
