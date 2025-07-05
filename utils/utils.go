package utils

import (
	"strconv"
	"strings"
)

func NumberFixer(numb string) string {
	if numb[:2] != "08" {
		if numb[:3] == "628" {
			var nomer []string
			nomer = append(nomer, "0")
			nomer = append(nomer, numb[2:])
			numb = strings.Join(nomer, "")
		} else {
			return ""
		}
	}
	if len(numb) > 5 {
		numb = numb[:5]
	}
	return numb
}

func QuerySupport(query string) string {
	count := strings.Count(query, "?")
	for i := 0; i < count; i++ {
		query = strings.Replace(query, "?", "$"+strconv.Itoa(i+1), 1)
	}
	return query
}
