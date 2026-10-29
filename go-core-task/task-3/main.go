package main

import (
	"fmt"
	"strings"
)

type StringIntMap struct {
	m map[string]int
}

func NewStringIntMap() StringIntMap {
	return StringIntMap{
		m: make(map[string]int),
	}
}

func (s StringIntMap) String() string {
	builder := strings.Builder{}
	for key, value := range s.m {
		builder.WriteString(fmt.Sprintf("{%s: %d}\n", key, value))
	}

	return builder.String()
}

func (s StringIntMap) Add(key string, value int) {
	s.m[key] = value
}

func (s StringIntMap) Remove(key string) {
	delete(s.m, key)
}

func (s StringIntMap) Copy() map[string]int {
	result := make(map[string]int, len(s.m))
	for key, value := range s.m {
		result[key] = value
	}
	return result
}

func (s StringIntMap) Exists(key string) bool {
	_, ok := s.m[key]
	return ok
}

func (s StringIntMap) Get(key string) (int, bool) {
	value, ok := s.m[key]
	return value, ok
}

func main() {
	m := NewStringIntMap()
	m.Add("key", 1)
	fmt.Println(m)
}
