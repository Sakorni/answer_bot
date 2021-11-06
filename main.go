package main

import (
	"encoding/json"
	"fmt"
	"io"
	"strings"
)

type AnswerRecord struct {
	Week     int
	Question string
	Answer   string
}

func (r AnswerRecord) String() string{
	return fmt.Sprintf("Week: %d\nFullQuestion:{%s}\n, Answer is:{%s}\n" , r.Week, r.Question, r.Answer)
}

func main() {
}

func askForAnswer(){
	res, err := GetAnswer(0, "Если UNIX-подобной операционной системой")
	if err != nil{
		fmt.Println("AN ERROR!", err)
	}
	if len(res) == 0{
		fmt.Println("An empty set")
	}
	for _, r := range res{
		fmt.Println(r)
	}
}

func readFromJsonAndMakeUniq(r io.Reader) []AnswerRecord {
	filter := make(map[AnswerRecord]bool)
	decoded := []AnswerRecord{}
	res := []AnswerRecord{}
	json.NewDecoder(r).Decode(&decoded)
	for _, rec := range decoded{
		rec.Question = strings.Replace(rec.Question, "\n", "", -1)
		rec.Answer = strings.Replace(rec.Answer, "\n", "", -1)
		if _, cont := filter[rec]; cont{
			continue
		}
		filter[rec] = true
		res = append(res, rec)
	}

	//fmt.Println(res)
	return res
}
