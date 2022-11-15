package models

type Feedback struct {
	ID          int    `json:"id"`
	UserId      int    `json:"user_id"`
	PhoneNumber string `json:"phone_number"`
	Email       string `json:"email"`
	Text        string `json:"text"`
	ProductId   int    `json:"product_id"`
}

type UpdateFeedbackInput struct {
	UserId      *int    `json:"user_id"`
	PhoneNumber *string `json:"phone_number"`
	Email       *string `json:"email"`
	Text        *string `json:"text"`
	ProductId   *int    `json:"product_id"`
}
