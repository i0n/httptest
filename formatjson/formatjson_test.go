package formatjson_test

import (
	"testing"

	httptest "github.com/i0n/httptest/cmd/cli"
	"github.com/i0n/httptest/formatjson"
)

var MarshalTests = []httptest.Path{
	{Path: "/hello", Method: "GET", ContentType: "application/json", ResponseStatus: 200, ResponseBody: ""},
}

func TestMarshal(t *testing.T) {
	for _, tt := range MarshalTests {
		expected := "{\n  content_type: \"application/json\",\n  method: \"GET\",\n  path: \"/hello\",\n  response_body: \"\",\n  response_status: 200\n}"
		actual, err := formatjson.Marshal(tt)
		if err != nil {
			t.Errorf("Returned Error %v", err)
		}
		if expected != string(actual) {
			t.Errorf("Expected %v, actual %v", expected, string(actual))
			t.Errorf("%q", string(actual))
		}
	}
}
