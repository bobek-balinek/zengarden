package zengarden

import (
	"strings"
	"text/template"
	"time"
)

var funcMap = template.FuncMap{
	"downcase":     strings.ToLower,
	"upcase":       strings.ToUpper,
	"date":         date,
	"dateToString": dateToString,
	"filter":       filter,
	"slice":        slice,
	"trim":         trim,
}

func date(format string, date time.Time) string {
	return date.Format(format)
}

func dateToString(date time.Time) string {
	return date.Format("2 Jan 2006")
}

func filter(key string, val interface{}, data []Context) []Context {
	var result []Context

	for _, ctx := range data {
		if ctx[key] == val {
			result = append(result, ctx)
		}
	}

	return result
}

func slice(offset, count int, data []Context) []Context {
	return data[offset:count]
}

func trim(needle, str string) string {
	return strings.TrimSuffix(str, needle)
}
