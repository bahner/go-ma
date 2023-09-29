package coll

import (
	"encoding/json"
	"errors"
	"fmt"
	"strings"

	log "github.com/sirupsen/logrus"
)

// Collections is a type that can be either a single element or a set of elements
// It's because parts of the DI can be either a string or a list of strings.

// It's nice to have a single interface for dealing with both cases, so
// we don't have to check continuously.
type Collection interface {
	Add(string) error
	Remove(string) error
	Contains(string) bool
	String() string
	Get() interface{}
}

type CollectionSet struct {
	elements map[string]struct{}
	isSingle bool
}

func (s *CollectionSet) String() string {
	if s.isSingle && len(s.elements) == 1 {
		for k := range s.elements {
			return fmt.Sprintf("Elements{element: %s}", k)
		}
	}

	keys := make([]string, 0, len(s.elements))
	for k := range s.elements {
		keys = append(keys, k)
	}
	return fmt.Sprintf("Elements{elements: [%s]}", strings.Join(keys, ", "))
}

func (s *CollectionSet) Get() interface{} {
	if s.isSingle && len(s.elements) == 1 {
		for k := range s.elements {
			return k
		}
	}

	keys := make([]string, 0, len(s.elements))
	for k := range s.elements {
		keys = append(keys, k)
	}
	return keys
}

func (s *CollectionSet) Add(element string) error {
	s.elements[element] = struct{}{}
	if len(s.elements) == 1 {
		s.isSingle = true
	} else {
		s.isSingle = false
	}
	return nil
}

func (s *CollectionSet) Remove(element string) error {
	delete(s.elements, element)
	if len(s.elements) == 1 {
		s.isSingle = true
	} else {
		s.isSingle = false
	}
	return nil
}

func (s *CollectionSet) Contains(element string) bool {
	_, exists := s.elements[element]
	return exists
}

func (s *CollectionSet) MarshalJSON() ([]byte, error) {
	if s.isSingle && len(s.elements) == 1 {
		for k := range s.elements {
			return json.Marshal(k)
		}
	}

	keys := make([]string, 0, len(s.elements))
	for k := range s.elements {
		keys = append(keys, k)
	}
	return json.Marshal(keys)
}

func (s *CollectionSet) UnmarshalJSON(data []byte) error {
	var single string
	var multiple []string

	if err := json.Unmarshal(data, &single); err == nil {
		log.Debug(fmt.Sprintf("coll: UnmarshalJSON: single: %s", single))
		s.elements = map[string]struct{}{single: {}}
		s.isSingle = true
		return nil
	}

	if err := json.Unmarshal(data, &multiple); err == nil {
		log.Debug(fmt.Sprintf("coll: UnmarshalJSON: multiple: %s", multiple))
		s.elements = make(map[string]struct{}, len(multiple))
		for _, k := range multiple {
			s.elements[k] = struct{}{}
		}
		s.isSingle = len(s.elements) == 1
		return nil
	}

	return fmt.Errorf("could not unmarshal as either single or multiple elements: single unmarshal error: %v, multiple unmarshal error: %v", json.Unmarshal(data, &single), json.Unmarshal(data, &multiple))
}

// Factory function to create the appropriate Elements
func New(data interface{}) (Collection, error) {
	elements := make(map[string]struct{})
	switch v := data.(type) {
	case string:
		elements[v] = struct{}{}
		return &CollectionSet{elements: elements, isSingle: true}, nil
	case []string:
		for _, element := range v {
			elements[element] = struct{}{}
		}
		return &CollectionSet{elements: elements, isSingle: len(elements) == 1}, nil
	default:
		return nil, errors.New("invalid data type for creating Elements")
	}
}
