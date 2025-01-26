package auth

import (
	"net/http"
	"reflect"
	"testing"
)

func TestAuth(t *testing.T) {
	var headers map[string]http.Header = map[string]http.Header{
		"valid":     {"Authorization": {"ApiKey supersecretkey"}},
		"malformed": {"Authorization": {"malformedDeformity"}},
		"no_auth":   {"OtherHeader": {"No auth here"}},
	}
	tests := map[string]struct {
		input           http.Header
		expected_output string
		expected_error  error
	}{
		"valid response":     {input: headers["valid"], expected_output: "supersecretkey", expected_error: nil},
		"malformed response": {input: headers["malformed"], expected_output: "", expected_error: ErrMalformedAuthHeader},
		"no auth header":     {input: headers["no_auth"], expected_output: "", expected_error: ErrNoAuthHeaderIncluded},
	}
	// TC as in test case.
	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			response, error_received := GetAPIKey(tc.input)
			if !reflect.DeepEqual(tc.expected_output, response) {
				t.Fatalf("expected response: %v, got: %v", tc.expected_output, response)
			}
			if error_received != tc.expected_error {
				t.Fatalf("expected error: %v, got: %v", tc.expected_error, error_received)
			}
		})
	}
}
