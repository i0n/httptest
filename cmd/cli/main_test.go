package main_test

import (
	"reflect"
	"testing"

	httptest "github.com/i0n/httptest/cmd/cli"
	"github.com/oleiade/reflections"
)

var pathFailuresTestStruct = &httptest.Path{
	Path:           "/hello",
	Method:         "INVALID",
	ContentType:    "application/json",
	ResponseStatus: 200,
	ResponseBody:   "",
}

var failures = []map[string]string{{"failure": "method"}}
var _ = reflections.SetField(pathFailuresTestStruct, "failures", failures)

func TestMarshal(t *testing.T) {
	expected := failures
	actual := pathFailuresTestStruct.Failures()
	for i, v := range actual {
		eq := reflect.DeepEqual(v, expected[i])
		if !eq {
			t.Errorf("Expected %v, actual %v", expected, actual)
		}
	}
}
