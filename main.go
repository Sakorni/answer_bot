package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"
)

type AnswerRecord struct {
	Week     int
	Question string
	Answer   string
}

func main() {
	source, err := os.Open("v1.json")
	if err != nil {
		log.Fatal(err)
	}
	answers := readFromJson(source)
	for _, ans := range answers{
		if err := AddAnswer(ans); err != nil{
			fmt.Println(err)
		}
	}
}

func readFromJson(r io.Reader) []AnswerRecord {
	res := []AnswerRecord{}
	json.NewDecoder(r).Decode(&res)
	//fmt.Println(res)
	return res
}
