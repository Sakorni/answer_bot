package utils

import (
	"CAOS/caos_db"
	"fmt"
)

// NoResult It has to be a constant, but you can't easily define a const error in go
var NoResult = fmt.Errorf("an empty set received")

func AskForAnswer(input string)(answer []string, err error){
	res, err := caos_db.GetAnswer(0, input)
	if err != nil{
		err = fmt.Errorf("AN ERROR! %s", err.Error())
		return
	}
	if len(res) == 0{
		return nil, NoResult
	}
	builder := make([]string, 0, len(res))
	for _, r := range res{
		builder = append(builder, r.String())
	}
	return builder, nil
}

