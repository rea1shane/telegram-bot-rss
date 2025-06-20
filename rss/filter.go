package rss

import (
	"fmt"
	"strings"
)

type filter struct {
	title    filterUnit
	content  filterUnit
	category filterUnit
}

type filterMode string

const (
	filterModeNone       = ""
	filterModeInclude    = "INCLUDE"
	filterModeNotInclude = "NOT_INCLUDE"
	filterModeMatch      = "MATCH"
	filterModeMismatch   = "MISMATCH"
)

// The item that hits the rule will be retained.
// If the mode is filterModeNone or the keywords is empty, retain will always return true.
type filterUnit struct {
	mode     filterMode
	keywords []string
}

func (u filterUnit) retain(input string) bool {
	if len(u.keywords) == 0 {
		return true
	}

	switch u.mode {
	case filterModeNone:
		return true

	case filterModeInclude:
		for _, keyword := range u.keywords {
			if strings.Contains(input, keyword) {
				return true
			}
		}
		return false

	case filterModeNotInclude:
		for _, keyword := range u.keywords {
			if strings.Contains(input, keyword) {
				return false
			}
		}
		return true

	case filterModeMatch:
		for _, keyword := range u.keywords {
			if input == keyword {
				return true
			}
		}
		return false

	case filterModeMismatch:
		for _, keyword := range u.keywords {
			if input == keyword {
				return false
			}
		}
		return true

	default:
		panic(fmt.Errorf("invalid mode: %s", u.mode))
	}
}
