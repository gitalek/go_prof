package hw03frequencyanalysis

import (
	"sort"
	"strings"
)

const excluded = " !\"#$%&'()*+,-./:;<=>?@[]^_`{|}~"

type record struct {
	word      string
	frequency int
}

func Top10(text string) []string {
	words := strings.Fields(text)

	counter := make(map[string]int)
	for _, word := range words {
		word = normalize(word)
		if word == "" {
			continue
		}
		counter[word]++
	}

	limit := 10
	if len(counter) < limit {
		limit = len(counter)
	}

	records := make([]record, 0, limit)
	for word, frequency := range counter {
		records = append(records, record{word, frequency})
	}

	sort.Slice(records, func(i, j int) bool {
		word1, frequency1 := records[i].word, records[i].frequency
		word2, frequency2 := records[j].word, records[j].frequency
		if frequency1 == frequency2 {
			i := strings.Compare(word1, word2)
			return i == -1
		}
		return frequency1 > frequency2
	})

	result := make([]string, 0, limit)
	for _, r := range records[:limit] {
		result = append(result, r.word)
	}
	return result
}

func normalize(word string) string {
	word = strings.ToLower(word)
	word = strings.Trim(word, excluded)
	return word
}
