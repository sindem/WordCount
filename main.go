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
		// Call ParseForm() to parse the raw query and update r.PostForm and r.Form.
		if err := r.ParseForm(); err != nil {
			fmt.Fprintf(w, "ParseForm() err: %v", err)
			return
		}
		//fmt.Fprintf(w, "Post from website! r.PostFrom = %v\n", r.PostForm)
		textstring := r.FormValue("textstring")

		var count []WordCount
		count = FrequencyCountWords(textstring)
		// for _, data := range count {
		// 	fmt.Printf("%v", data)
		// }

		//count := repetition(textstring)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		// for _, data := range count {
		// 	json.NewEncoder(w).Encode(data)
		// }
		json.NewEncoder(w).Encode(count)
		// //json.NewEncoder(w).Encode(&Response{"success", 200, count})
		// return

		//b, _ := json.Marshal(count)
		//fmt.Printf("%v", string(b)) // ["1", "2", "A", "B", "Hack\"er"]

		// wc, err := json.Marshal(count)
		// if err == nil {
		// 	w.Write(wc)
		// 	return
		//fmt.Fprintf(w, string(b))
		// }

	default:
		fmt.Fprintf(w, "Sorry, only GET and POST methods are supported.")
	}
}

func FrequencyCountWords(text string) []WordCount {

	//reg := regexp.MustCompile("[a-zA-Z0-9:']+")
	//matches := reg.FindAllString(text, -1)
	words := CountOfWordsInMap(text)

	// words := make(map[string]int)

	// for _, match := range matches {
	// 	words[match]++
	// }

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
