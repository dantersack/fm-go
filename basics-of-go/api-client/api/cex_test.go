package api_test

import (
	"testing"

	"dantejs.com/go/api-client/api"
)

func TestApiCall(t *testing.T) {
	// this should fail
	_, err := api.GetRate("")
	if err == nil {
		t.Error("error should not be nil with empty strings")
	}
}
