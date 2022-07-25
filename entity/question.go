package entity

type Question struct {
	Id       int    `json:"id" db:"id"`
	No       int    `json:"no" db:"no"`
	Question string `json:"question" db:"question"`
	Answer   string `json:"answer" db:"answer"`
}
