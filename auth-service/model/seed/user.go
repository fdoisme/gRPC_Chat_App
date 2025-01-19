package seed

import (
	"project/auth-service/helper"
	"project/auth-service/model"
)

func User() []model.User {
	return []model.User{
		{Email: "satu@mailinator.com", VerifiedAt: helper.Ptr(helper.DateTime("2024-12-25 09:00:01"))},
		{Email: "dua@mailinator.com", VerifiedAt: helper.Ptr(helper.DateTime("2024-12-26 10:00:02"))},
		{Email: "tiga@mailinator.com", VerifiedAt: helper.Ptr(helper.DateTime("2024-12-27 11:00:03"))},
		{Email: "empat@mailinator.com", VerifiedAt: helper.Ptr(helper.DateTime("2024-12-28 12:00:04"))},
		{Email: "lima@mailinator.com", VerifiedAt: helper.Ptr(helper.DateTime("2024-12-29 16:00:05"))},
		{Email: "enam@mailinator.com"},
		{Email: "tujuh@mailinator.com"},
		{Email: "delapan@mailinator.com"},
		{Email: "sembilan@mailinator.com"},
		{Email: "sepuluh@mailinator.com"},
	}
}
