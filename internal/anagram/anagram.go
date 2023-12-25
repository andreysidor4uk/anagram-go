package anagram

import "strings"

type Anagram struct {
	anagram  string
	solution string
}

func (anagram Anagram) String() string {
	return anagram.anagram
}

func (anagram Anagram) Check(solution string) bool {
	return anagram.solution == strings.Trim(strings.ToLower(solution), " ")
}

type AnagramStorage []Anagram

func NewTestAnagramStorage() AnagramStorage {
	return AnagramStorage{
		Anagram{
			anagram:  "бората",
			solution: "работа",
		},
	}
}
