package codegen

import "strings"

func varName(name string, t Tags) string {
	return strings.ToLower(name)
}
