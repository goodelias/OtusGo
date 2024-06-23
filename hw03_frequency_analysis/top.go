package hw03frequencyanalysis

import (
	"sort"
	"strings"
)

func Top10(s string) []string {
	words := strings.Fields(s)
	wordsCounter := make(map[string]int)

	for _, word := range words {
		wordsCounter[word]++
	}

	type WordFrequency struct {
		Word string
		Freq int
	}

	wordsSlice := make([]WordFrequency, 0, len(wordsCounter))
	for word, freq := range wordsCounter {
		wf := WordFrequency{
			Word: word,
			Freq: freq,
		}
		wordsSlice = append(wordsSlice, wf)
	}

	sort.Slice(wordsSlice, func(i, j int) bool {
		if wordsSlice[i].Freq == wordsSlice[j].Freq {
			return wordsSlice[i].Word < wordsSlice[j].Word
		}
		return wordsSlice[i].Freq > wordsSlice[j].Freq
	})

	ws10 := wordsSlice
	if len(wordsSlice) > 10 {
		ws10 = wordsSlice[0:10]
	}

	res := make([]string, 0, 10)

	for _, wf := range ws10 {
		res = append(res, wf.Word)
	}

	return res
}
