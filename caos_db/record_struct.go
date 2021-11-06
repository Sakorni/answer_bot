package caos_db

import "fmt"

type AnswerRecord struct {
	Week     int
	Question string
	Answer   string
}



func (r AnswerRecord) String() string{
	return fmt.Sprintf("Week: %d\nFull question: {%s},\nAnswer is: {%s}\n" , r.Week, r.Question, r.Answer)
}

