package models

type Feedback struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	PhoneNumber string `json:"phone_number"`
	Email       string `json:"email"`
	Text        string `json:"text"`
}

type UpdateFeedbackInput struct {
	Name        *string `json:"name"`
	PhoneNumber *string `json:"phone_number"`
	Email       *string `json:"email"`
	Text        *string `json:"text"`
}
