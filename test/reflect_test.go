package test

import (
	"fmt"
	"regexp"
	"testing"
)

func TestRep(t *testing.T) {
	sampleRegexp := regexp.MustCompile(`(:[^/.]*)`)

	input := "GET_/custom/group/list/:model/:id"

	result := sampleRegexp.ReplaceAllString(input, "(:[^/.]*)")
	fmt.Println(string(result))

	reg := "GET_/custom/group/list/([^/.]*)/([^/.]*)"
	route := "^GET_/custom/group/list/student/1$"

	match, err := regexp.MatchString(reg, route)

	fmt.Printf("%+v, %+v", match, err)

}
