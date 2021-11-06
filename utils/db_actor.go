package utils

import (
	"CAOS/caos_db"
	"fmt"
	"strings"
)
const noResult = "an empty set received"

func AskForAnswer(input string)(answer string, err error){
	res, err := caos_db.GetAnswer(0, input)
	if err != nil{
		err = fmt.Errorf("AN ERROR! %s", err.Error())
		return
	}
	if len(res) == 0{
		return noResult, nil
	}
	builder := strings.Builder{}
	for _, r := range res{
		builder.WriteString(r.String())
	}
	return builder.String(), nil
}

