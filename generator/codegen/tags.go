package codegen

import (
	"reflect"
	"strings"
)

const (
	strictRules = "strict"
)

type Tags map[string][]string

func (t Tags) jsonName() string {
	if v := t["json"]; len(v) > 0 {
		return v[0]
	}
	return ""
}

func (t Tags) defaultValue() string {
	if v := t["default"]; len(v) > 0 {
		return v[0]
	}
	return ""
}

func parseTags(tag string) Tags {
	var result = make(map[string][]string)
	result["json"] = strings.Split(structTag(tag).Get("json"), ",")
	result["default"] = strings.Split(structTag(tag).Get("default"), ",")
	return result
}

func structTag(tag string) reflect.StructTag {
	if len(tag) < 2 {
		return ""
	}
	if tag[0] == '`' && tag[len(tag)-1] == '`' {
		return reflect.StructTag(tag[1 : len(tag)-1])
	}
	return ""
}

type StructTags []string

func (t StructTags) Has(s string) bool {
	for _, tag := range t {
		if tag == s {
			return true
		}
	}
	return false
}
