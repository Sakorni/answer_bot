package main

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"strings"
)

type AnswerRecord struct {
	Week     int
	Question string
	Answer   string
}

const NO_RESULT = "an empty set received"


func (r AnswerRecord) String() string{
	return fmt.Sprintf("Week: %d\nFull question: {%s},\nAnswer is: {%s}\n" , r.Week, r.Question, r.Answer)
}

func main() {
	res, err := askForAnswer(os.Args[1])
	if err != nil{
		fmt.Println(err)
	}else {
		fmt.Println(res)
	}
}

func askForAnswer(input string)(answer string, err error){
	res, err := GetAnswer(0, input)
	if err != nil{
		err = fmt.Errorf("AN ERROR! %s", err.Error())
		return
	}
	if len(res) == 0{
		err = fmt.Errorf(NO_RESULT)
		return
	}
	builder := strings.Builder{}
	for _, r := range res{
		builder.WriteString(r.String())
	}
	return builder.String(), nil
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
