package utils_test

import (
	"testing"

	"github.com/i0n/httptest/utils"
)

var stringInSliceTrueTests = []struct {
	key  string
	list []string
}{
	{"hello", []string{"hello"}},
	{"hello", []string{"world", "hello", "yo"}},
	{"", []string{"yes", "no", ""}},
}

var stringInSliceFalseTests = []struct {
	key  string
	list []string
}{
	{"hello", []string{}},
	{"", []string{"yes", "no"}},
}

func TestStringInSlice(t *testing.T) {
	for _, tt := range stringInSliceTrueTests {
		expected := true
		actual := utils.StringInSlice(tt.key, tt.list)
		if expected != actual {
			t.Errorf("Expected %v, actual %v", expected, actual)
		}
	}
	for _, tt := range stringInSliceFalseTests {
		expected := false
		actual := utils.StringInSlice(tt.key, tt.list)
		if expected != actual {
			t.Errorf("Expected %v, actual %v", expected, actual)
		}
	}
}

var stringInSliceMapKeyTrueTests = []struct {
	a   string
	m   []map[string]string
	key string
}{
	{"method", []map[string]string{{"failure": "method", "expected": "GET", "result": "POST"}}, "failure"},
	{"GET", []map[string]string{{"failure": "method", "expected": "GET", "result": "POST"}}, "expected"},
}

var stringInSliceMapKeyFalseTests = []struct {
	a   string
	m   []map[string]string
	key string
}{
	{"nope", []map[string]string{{"failure": "method", "expected": "GET", "result": "POST"}}, "failure"},
	{"not", []map[string]string{{"failure": "method", "expected": "GET", "result": "POST"}}, "expected"},
}

func TestStringInSliceMapKey(t *testing.T) {
	for _, tt := range stringInSliceMapKeyTrueTests {
		expected := true
		actual := utils.StringInSliceMapKey(tt.a, tt.m, tt.key)
		if expected != actual {
			t.Errorf("Expected %v, actual %v", expected, actual)
		}
	}
	for _, tt := range stringInSliceMapKeyFalseTests {
		expected := false
		actual := utils.StringInSliceMapKey(tt.a, tt.m, tt.key)
		if expected != actual {
			t.Errorf("Expected %v, actual %v", expected, actual)
		}
	}
}

var stringInSliceInterfaceKeyTrueTests = []struct {
	a   interface{}
	m   []map[string]string
	key string
}{
	{"method", []map[string]string{{"failure": "method", "expected": "GET", "result": "POST"}}, "failure"},
	{"GET", []map[string]string{{"failure": "method", "expected": "GET", "result": "POST"}}, "expected"},
}

var stringInSliceInterfaceKeyFalseTests = []struct {
	a   interface{}
	m   []map[string]string
	key string
}{
	{"nope", []map[string]string{{"failure": "method", "expected": "GET", "result": "POST"}}, "failure"},
	{"not", []map[string]string{{"failure": "method", "expected": "GET", "result": "POST"}}, "expected"},
}

func TestStringInSliceInterfaceKey(t *testing.T) {
	for _, tt := range stringInSliceMapKeyTrueTests {
		expected := true
		actual := utils.StringInSliceMapKey(tt.a, tt.m, tt.key)
		if expected != actual {
			t.Errorf("Expected %v, actual %v", expected, actual)
			t.Log(tt.a)
			t.Log(tt.key)
		}
	}
	for _, tt := range stringInSliceInterfaceKeyTrueTests {
		expected := true
		actual := utils.StringInSliceInterfaceKey(tt.a, tt.m, tt.key)
		if expected != actual {
			t.Errorf("Expected %v, actual %v", expected, actual)
			t.Log(tt.a)
			t.Log(tt.key)
		}
	}
	for _, tt := range stringInSliceMapKeyFalseTests {
		expected := false
		actual := utils.StringInSliceInterfaceKey(tt.a, tt.m, tt.key)
		if expected != actual {
			t.Errorf("Expected %v, actual %v", expected, actual)
		}
	}
	for _, tt := range stringInSliceInterfaceKeyFalseTests {
		expected := false
		actual := utils.StringInSliceInterfaceKey(tt.a, tt.m, tt.key)
		if expected != actual {
			t.Errorf("Expected %v, actual %v", expected, actual)
		}
	}
}
