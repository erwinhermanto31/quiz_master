package util

import (
	"encoding/base64"
	"fmt"
	"strconv"
	"time"
)

const (
	DateFormatRFC3339 = time.RFC3339
)

var Command = map[string]string{
	"create_question": "create_question",
	"update_question": "update_question",
	"delete_question": "delete_question",
	"question":        "question",
	"questions":       "questions",
	"answer_question": "answer_question",
}

func GetRedisKey(domain, id string) string {
	return domain + ":" + id
}

func FormatDateToRFC3339(t time.Time) string {
	return t.Format(DateFormatRFC3339)
}

func StringToInt(s string) int {
	i, _ := strconv.Atoi(s)
	return i
}

func StringToInt64(s string) int64 {
	i, _ := strconv.ParseInt(s, 10, 64)
	return i
}

func EncodeBase64(s string) string {
	return base64.StdEncoding.EncodeToString([]byte(s))
}

func PrintHelp() {
	fmt.Println("Command                                  | Description")
	fmt.Println("create_question <no> <question> <answer> | Creates a question")
	fmt.Println("update_question <no> <question> <answer> | Updates a question")
	fmt.Println("delete_question <no>                     | Delete a question")
	fmt.Println("question <no>                            | Shows a question")
	fmt.Println("answer_question <no>                     | Answer a question by no")
	fmt.Println("questions                                | Shows question list")
}
