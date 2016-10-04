// Test bigram functions
package ngram

import (
	"cnreader/dictionary"
	"fmt"
	"testing"
)

// Test basic Bigram functions
func TestBigram(t *testing.T) {
	fmt.Printf("TestBigram: Begin unit test\n")
	s1 := "蓝"
	s2 := "藍"
	hw1 := dictionary.HeadwordDef{
		Id: 1,
		Simplified: &s1, 
		Traditional: &s2,
		Pinyin: []string{},
		WordSenses: &[]dictionary.WordSenseEntry{},
	}
	s3 := "天"
	s4 := "\\N"
	hw2 := dictionary.HeadwordDef{
		Id: 2,
		Simplified: &s3, 
		Traditional: &s4,
		Pinyin: []string{},
		WordSenses: &[]dictionary.WordSenseEntry{},
	}
	example := ""
	exFile := ""
	exDocTitle := ""
	exColTitle := ""
	b := NewBigram(hw1, hw2, example, exFile, exDocTitle, exColTitle)
	r := b.Traditional()
	e := "藍天"
	if r != e {
		t.Error("TestBigram, expected ", e, " got, ", r)
	}
}
