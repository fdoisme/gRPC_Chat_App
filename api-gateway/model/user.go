package model

type User struct {
	Email     string  `json:"email,omitempty"`
	FirstName *string `json:"firstName,omitempty"`
	LastName  *string `json:"lastName,omitempty"`
	IsOnline  bool    `json:"isOnline"`
}
