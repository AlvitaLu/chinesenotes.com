/*
Library for Chinese vocabulary analysis based on literary genre
*/
package analysis

import (
	"log"
)

// Sorted list of word frequencies
type SortedByGenre struct {
	Genre string
	SWI []WFResult
}

// Word frequency by genre
type WordFreqByGenre struct {
	Genre string
	WF map[string]int
}

// An array of word frequency maps by genre
type WFArrayByGenre []WordFreqByGenre

// Gets the matching word frequency map
func  (wfArray WFArrayByGenre) Get(genre string) map[string]int {
	for _, wf := range wfArray {
		if wf.Genre == genre {
			return wf.WF
		}
	}
	return map[string]int{}
}

// Merge the argument into the word frequency map for the matching genre
// more: a word frequency map for a given genre
func MergeByGenre(wfArray WFArrayByGenre, more WordFreqByGenre) WFArrayByGenre {
	log.Printf("analysis.Merge: len(wfArray): '%d' more.Genre: '%s'\n",
		len(wfArray), more.Genre)
	found := false
	for _, wf := range wfArray {
		if wf.Genre == more.Genre {
			for k, v := range more.WF {
    			wf.WF[k] += v
			}
			found = true
		}
	}
	if !found {
		wfArray = append(wfArray, more)
	}
	log.Printf("analysis.MergeByGenre: len(wfArray): '%d' found: '%v'\n",
		len(wfArray), found)
	return wfArray
}

// Constructor
func NewWordFreqByGenre(genre string) WordFreqByGenre {
	return WordFreqByGenre{
		Genre: genre,
		WF: map[string]int{},
	}
}