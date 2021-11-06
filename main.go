package main

import (
	"CAOS/utils"
	"fmt"
	"os"
)


func main() {
	res, err := utils.AskForAnswer(os.Args[1])
	if err != nil{
		fmt.Println(err)
	}else {
		fmt.Println(res)
	}
}
