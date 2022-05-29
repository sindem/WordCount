package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"sort"
	"strings"
)

func main() {
	http.HandleFunc("/wordcounts", handler)
	log.Fatal(http.ListenAndServe(":8000", nil))
}

func handler(w http.ResponseWriter, r *http.Request) {

	switch r.Method {
	case "GET":
		fmt.Fprintf(w, "Sorry,  GET request processing is under maintenance")
	case "POST":
		if err := r.ParseForm(); err != nil {
			fmt.Fprintf(w, "ParseForm() err: %v", err)
			return
		}

		textstring := r.FormValue("textstring")
		var count []WordCount
		count = FrequencyCountWords(textstring)

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(count)

	default:
		fmt.Fprintf(w, "Sorry, only GET and POST methods are supported.")
	}
}

func FrequencyCountWords(text string) []WordCount {

	words := CountOfWordsInMap(text)
	var wordFreqs []WordCount
	for k, v := range words {
		wordFreqs = append(wordFreqs, WordCount{k, v})
	}

	sort.Slice(wordFreqs, func(i, j int) bool {
		return wordFreqs[i].Count > wordFreqs[j].Count
	})

	lengthToPrint := 0

	if len(wordFreqs) >= 10 {
		lengthToPrint = 10
	} else {
		lengthToPrint = len(wordFreqs)
	}

	return wordFreqs[:lengthToPrint]
}

type WordCount struct {
	Word  string `json:"word"`
	Count int    `json:"count"`
}

func (wc WordCount) String() string {
	return fmt.Sprintf("{\n word:	%s"+",\n count:"+"	%d \n}", wc.Word, wc.Count)
}

func CountOfWordsInMap(st string) map[string]int {
	input := strings.Fields(st)
	wc := make(map[string]int)
	for _, word := range input {
		_, matched := wc[word]
		if matched {
			wc[word] += 1
		} else {
			wc[word] = 1
		}
	}

	return wc
}
