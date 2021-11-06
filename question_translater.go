package main

import (
	"bytes"
	"database/sql"
	"fmt"
	"github.com/go-sql-driver/mysql"
	"log"
	"os"
)

const EMPTY_RESULT_CAPTION = "sql: no rows in result set"

var db *sql.DB

func init() {
	cfg := mysql.Config{
		User:   os.Getenv("DBUSER"),
		Passwd: os.Getenv("DBPASS"),
		Net:    "tcp",
		Addr:   "localhost:3306",
		DBName: "caos",
	}
	var err error
	db, err = sql.Open("mysql", cfg.FormatDSN())
	if err != nil {
		log.Fatal(err)
	}
	pingerr := db.Ping()
	if pingerr != nil {
		log.Fatal(pingerr)
	}
	fmt.Println("Connected to database")
}

func AddAnswer(record AnswerRecord) error {
	_, err := db.Exec("INSERT INTO answers (week, question, answer) VALUES (?, ?, ?)",
		record.Week, record.Question, record.Answer)
	if err != nil {
		return err
	}
	return nil
}

func GetAnswer(week int, question string) (res []AnswerRecord, err error){
	res = []AnswerRecord{}
	buf := bytes.NewBuffer(nil)
	buf.WriteString("SELECT * FROM ANSWERS WHERE ")
	if week > 0{
		buf.WriteString(fmt.Sprintf("WEEK = %d and", week))
	}
	buf.WriteString("QUESTION LIKE ?")
	rows,err := db.Query(buf.String(), "%" + question + "%")
	if err != nil{
		return
	}
	defer rows.Close()

	for rows.Next(){
		rec := AnswerRecord{}
		if err = rows.Scan(new(interface{}), &rec.Week, &rec.Question, &rec.Answer); err != nil{
			return []AnswerRecord{}, err
		}
		res = append(res, rec)
	}
	return
}
