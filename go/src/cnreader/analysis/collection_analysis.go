/*
CollectionAResults type for vocabulary analysis of a collection of texts
*/
package analysis

import (
	"cnreader/dictionary"
	"cnreader/ngram"
)

// A struct to hold the analysis results for the collection
type CollectionAResults struct {
	Vocab map[string]int
	Usage map[string]string
	BigramFrequencies ngram.BigramFreqMap
	Collocations ngram.CollocationMap
	WC int
	UnknownChars map[string]int
}

// Add more results to this set of results
func (results *CollectionAResults) AddResults(more CollectionAResults) {

	for k, v := range more.Vocab {
    	results.Vocab[k] += v
	}

	for k, v := range more.Usage {
    	results.Usage[k] = v
	}

	results.BigramFrequencies.Merge(more.BigramFrequencies)

	results.Collocations.MergeCollocationMap(more.Collocations)

	results.WC += more.WC

	for k, v := range more.UnknownChars {
    	results.UnknownChars[k] += v
	}
}

// Returns the subset of words that are lexical (content) words
func (results *CollectionAResults) GetLexicalWordFreq(sortedWords []SortedWordItem) []WFResult {

	wfResults := make([]WFResult, 0)
	maxWFOutput := len(sortedWords)
	if maxWFOutput > MAX_WF_OUTPUT {
		maxWFOutput = MAX_WF_OUTPUT
	}
	for _, value := range sortedWords[:maxWFOutput] {
		ws, _ := dictionary.GetWordSense(value.Word)
		if !ws.IsFunctionWord() {
			wfResults = append(wfResults, WFResult{
				Freq: value.Freq,
				HeadwordId: ws.HeadwordId,
				Chinese: value.Word,
				Pinyin: ws.Pinyin, 
				English: ws.English, 
				Usage: results.Usage[value.Word]})
		}
	}
	return wfResults
}


// Returns the subset of words that are lexical (content) words
func (results *CollectionAResults) GetWordFreq(sortedWords []SortedWordItem) []WFResult {

	wfResults := make([]WFResult, 0)
	maxWFOutput := len(sortedWords)
	if maxWFOutput > MAX_WF_OUTPUT {
		maxWFOutput = MAX_WF_OUTPUT
	}
	for _, value := range sortedWords[:maxWFOutput] {
		ws, _ := dictionary.GetWordSense(value.Word)
		wfResults = append(wfResults, WFResult{
			Freq: value.Freq,
			HeadwordId: ws.HeadwordId,
			Chinese: value.Word,
			Pinyin: ws.Pinyin, 
			English: ws.English, 
			Usage: results.Usage[value.Word]})
	}
	return wfResults
}