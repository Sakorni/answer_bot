package utils

import (
	"CAOS/caos_db"
	"encoding/json"
	"io"
	"strings"
)

func ReadFromJsonAndMakeUniq(r io.Reader) []caos_db.AnswerRecord {
	filter := make(map[caos_db.AnswerRecord]bool)
	var decoded []caos_db.AnswerRecord
	var res []caos_db.AnswerRecord
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
	return res
	}
