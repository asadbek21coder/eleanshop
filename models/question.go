package models

type Question struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	PhoneNumber string `json:"phone_number"`
	Time        string `json:"time"`
	Text        string `json:"text"`
}

type UpdateQuestionInput struct {
	Name        *string `json:"name"`
	PhoneNumber *string `json:"phone_number"`
	Time        *string `json:"email"`
	Text        *string `json:"text"`
}
