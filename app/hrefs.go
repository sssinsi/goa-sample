// Code generated by goagen v1.2.0, DO NOT EDIT.
//
// API "cellar": Application Resource Href Factories
//
// Command:
// $ goagen
// --design=github.com/sssinsi/goa-sample/cellar/design
// --out=$(GOPATH)/src/github.com/sssinsi/goa-sample
// --version=v1.2.0

package app

import (
	"fmt"
	"strings"
)

// BottleHref returns the resource href.
func BottleHref(bottleID interface{}) string {
	parambottleID := strings.TrimLeftFunc(fmt.Sprintf("%v", bottleID), func(r rune) bool { return r == '/' })
	return fmt.Sprintf("/bottles/%v", parambottleID)
}
