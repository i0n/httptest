// Package formatjson provides JSON pretty print.
package formatjson

import (
	"encoding/json"
	"fmt"
	"sort"
	"strconv"
	"strings"

	"github.com/fatih/color"
	"github.com/i0n/httptest/utils"
)

// Formatter is a struct to format JSON data. `color` is github.com/fatih/color: https://github.com/fatih/color
type Formatter struct {

	// JSON color for passing test
	PassColor *color.Color

	// JSON color for failing test
	FailColor *color.Color

	// Max length of JSON string value. When the value is 1 and over, string is truncated to length of the value. Default is 0 (not truncated).
	StringMaxLength int

	// Boolean to diesable color. Default is false.
	DisabledColor bool

	// Indent space number. Default is 2.
	Indent int
}

// StructWithFailures interface for any struct that returns an array of failure strings
type StructWithFailures interface {
	Failures() []map[string]string
}

// NewFormatter returns a new formatter with following default values.
func NewFormatter() *Formatter {
	return &Formatter{
		PassColor:       color.New(color.FgGreen),
		FailColor:       color.New(color.FgRed),
		StringMaxLength: 0,
		DisabledColor:   false,
		Indent:          2,
	}
}

// Marshal and formats JSON data.
func (f *Formatter) Marshal(v StructWithFailures) ([]byte, error) {
	data, err := json.Marshal(v)

	if err != nil {
		return nil, err
	}

	return f.Format(data, v.Failures())
}

// Format JSON string.
func (f *Formatter) Format(data []byte, failures []map[string]string) ([]byte, error) {
	var v interface{}
	err := json.Unmarshal(data, &v)

	if err != nil {
		return nil, err
	}

	s := f.pretty(v, 1, failures)

	return []byte(s), nil
}

func (f *Formatter) sprintfColor(c *color.Color, format string, args ...interface{}) string {
	if f.DisabledColor || c == nil {
		return fmt.Sprintf(format, args...)
	}
	return c.SprintfFunc()(format, args...)
}

func (f *Formatter) pretty(v interface{}, depth int, failures []map[string]string) string {
	switch val := v.(type) {
	case nil:
		return "null"
	case map[string]interface{}:
		return f.processMap(val, depth, failures)
	}
	return ""
}

func (f *Formatter) formatValue(v interface{}, state string) string {
	switch val := v.(type) {
	case string:
		return f.processString(val, state)
	case float64:
		return f.processFloat(val, state)
	case bool:
		return f.processBool(val, state)
	case nil:
		return "null"
	}
	return ""
}

func (f *Formatter) processString(s string, state string) string {
	r := []rune(s)

	if f.StringMaxLength != 0 && len(r) >= f.StringMaxLength {
		s = string(r[0:f.StringMaxLength]) + "..."
	}

	b, _ := json.Marshal(s)

	if state == "fail" {
		return f.sprintfColor(f.FailColor, string(b))
	}
	return f.sprintfColor(f.PassColor, string(b))
}

func (f *Formatter) processFloat(val float64, state string) string {
	if state == "fail" {
		return f.sprintfColor(f.FailColor, strconv.FormatFloat(val, 'f', -1, 64))
	}
	return f.sprintfColor(f.PassColor, strconv.FormatFloat(val, 'f', -1, 64))
}

func (f *Formatter) processBool(val bool, state string) string {
	if state == "fail" {
		return f.sprintfColor(f.FailColor, strconv.FormatBool(val))
	}
	return f.sprintfColor(f.PassColor, strconv.FormatBool(val))
}

func (f *Formatter) processMap(m map[string]interface{}, depth int, failures []map[string]string) string {
	currentIndent := f.generateIndent(depth - 1)
	nextIndent := f.generateIndent(depth)
	rows := []string{}

	if len(m) == 0 {
		return "{}"
	}

	for key, val := range m {
		v := val
		if utils.StringInSliceMapKey(key, failures, "failure") {
			v = f.formatValue(val, "fail")
		} else {
			v = f.formatValue(val, "pass")
		}
		row := fmt.Sprintf("%s%s: %s", nextIndent, key, v)
		rows = append(rows, row)
	}

	sort.Strings(rows)

	return fmt.Sprintf("{\n%s\n%s}", strings.Join(rows, ",\n"), currentIndent)
}

func (f *Formatter) generateIndent(depth int) string {
	return strings.Join(make([]string, f.Indent*depth+1), " ")
}

// Marshal JSON data with default options.
func Marshal(v StructWithFailures) ([]byte, error) {
	return NewFormatter().Marshal(v)
}
