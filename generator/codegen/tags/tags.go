package tags

import (
	"reflect"
	"strings"
)

const (
	StrictRules = "strict"
)

type Tags map[string][]string

func (t Tags) JsonName() string {
	if v := t["json"]; len(v) > 0 {
		return v[0]
	}
	return ""
}

func (t Tags) JsonAppendix() string {
	if v := t["json"]; len(v) > 1 {
		return v[1]
	}
	return ""
}

func (t Tags) JsonTags() StructTags {
	if v := t["json"]; len(v) > 0 {
		return v
	}
	return nil
}

func (t Tags) DefaultValue() string {
	if v := t["default"]; len(v) > 0 {
		return v[0]
	}
	return ""
}

func (t Tags) Layout() string {
	if v := t["layout"]; len(v) > 0 {
		return v[0]
	}
	return ""
}

func Parse(tag string) Tags {
	var result = make(map[string][]string)
	result["json"] = strings.Split(StructTag(tag).Get("json"), ",")
	result["layout"] = strings.Split(StructTag(tag).Get("layout"), ",")
	result["default"] = strings.Split(StructTag(tag).Get("default"), ",")
	return result
}

func StructTag(tag string) reflect.StructTag {
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

func (t StructTags) StrictRules() bool {
	return t.Has(StrictRules)
}
