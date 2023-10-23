package scope

import (
	"strings"

	"github.com/hexcraft-biz/misc/basic"
)

const (
	Delimiter = " "
)

type Scopes map[string]bool

func New(input interface{}) Scopes {
	items := []string{}
	if basic.IsSlice(input) {
		items = input.([]string)
	} else {
		items = strings.Split(input.(string), Delimiter)
	}

	scopes := Scopes{}
	for _, i := range items {
		scopes.Set(i)
	}

	return scopes
}

func (s *Scopes) Set(item string) {
	(*s)[item] = true
}

func (s Scopes) HasOneOf(sub Scopes) bool {
	for item, has := range sub {
		if has {
			if val, ok := s[item]; ok && val {
				return true
			}
		}
	}
	return false
}

func (s Scopes) Contains(sub Scopes) bool {
	for item, has := range sub {
		if has {
			if val, ok := s[item]; !ok || !val {
				return false
			}
		}
	}
	return true
}
