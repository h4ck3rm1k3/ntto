package ntto

import (
	"encoding/xml"
	"errors"
	"fmt"
	"strings"
)

const AppVersion = "0.2"

type Triple struct {
	XMLName   xml.Name `json:"-" xml:"t"`
	Subject   string   `json:"s" xml:"s"`
	Predicate string   `json:"s" xml:"s"`
	Object    string   `json:"s" xml:"s"`
}

type Rule struct {
	Prefix   string
	Shortcut string
}

func (r Rule) String() string {
	return fmt.Sprintf("%s: %s", r.Shortcut, r.Prefix)
}

func IsURIRef(s string) bool {
	return strings.HasPrefix(s, "<") && strings.HasSuffix(s, ">")
}

func IsLiteral(s string) bool {
	return strings.HasPrefix(s, "\"")
}

func IsLiteralLanguage(s, language string) bool {
	if !IsLiteral(s) {
		return false
	}
	if !strings.Contains(s, "@") {
		return true
	} else {
		return strings.Contains(s, "@"+language)
	}
}

func IsNamedNode(s string) bool {
	return strings.HasPrefix(s, "_:")
}

// ParseAbbreviations takes a string, parse the abbreviations and returns them as slice
func ParseRules(s string) ([]Rule, error) {
	var rules []Rule
	var err error
	for _, line := range strings.Split(s, "\n") {
		line = strings.TrimSpace(line)
		if len(line) == 0 || strings.HasPrefix(line, "#") || strings.HasPrefix(line, "//") {
			continue
		}
		fields := strings.Fields(line)
		if len(fields) < 2 {
			err = errors.New(fmt.Sprintf("Broken rule: %s", line))

		}
		rules = append(rules, Rule{Prefix: fields[1], Shortcut: fields[0]})
	}
	return rules, err
}
