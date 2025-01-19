package model

type User struct {
	Email     string `gorm:"not null;unique"`
	FirstName string
	LastName  string
	IsOnline  bool `gorm:"default:false"`
}

func (User) TableName() string {
	return "accounts"
}

func Seed() []User {
	return []User{
		{Email: "satu@mailinator.com", FirstName: "John", LastName: "Doe"},
		{Email: "dua@mailinator.com", FirstName: "John", LastName: "Doe"},
		{Email: "tiga@mailinator.com", FirstName: "John", LastName: "Doe"},
		{Email: "empat@mailinator.com", FirstName: "John", LastName: "Doe"},
		{Email: "lima@mailinator.com", FirstName: "John", LastName: "Doe"},
		{Email: "enam@mailinator.com", FirstName: "John", LastName: "Doe"},
		{Email: "tujuh@mailinator.com", FirstName: "John", LastName: "Doe", IsOnline: true},
		{Email: "delapan@mailinator.com", FirstName: "John", LastName: "Doe", IsOnline: true},
	}
}
