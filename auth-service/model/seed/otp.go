package seed

import (
	"github.com/google/uuid"
	"project/auth-service/helper"
	"project/auth-service/model"
)

func OTP() []model.Otp {
	return []model.Otp{
		{ID: uuid.MustParse("1cb7b89a-6627-43a8-a0a2-9df769804f83"), Otp: "1111", UserID: 1, ValidatedAt: helper.Ptr(helper.DateTime("2024-12-31 11:11:11"))},
		{ID: uuid.MustParse("3262d24b-e0b2-42e5-b08f-b694bfcd3171"), Otp: "1111", UserID: 1},
		{ID: uuid.MustParse("f7e8399d-bb15-4f78-9639-e9711c2bf061"), Otp: "2222", UserID: 2},
		{ID: uuid.MustParse("3c6fd69c-671f-4568-bbf7-72de13fb4379"), Otp: "3333", UserID: 3},
		{ID: uuid.MustParse("344be09c-208a-4a9b-a502-e60c9324eb5e"), Otp: "4444", UserID: 4},
		{ID: uuid.MustParse("cb8bee51-892a-44f5-a008-ca54331bc9e2"), Otp: "5555", UserID: 5},
		{ID: uuid.MustParse("74a7ea30-96e6-4cdf-8800-e118fb55a2c2"), Otp: "6666", UserID: 6},
		{ID: uuid.MustParse("03310322-721c-4f18-a905-ec1935878d2b"), Otp: "7777", UserID: 7},
		{ID: uuid.MustParse("00930308-5efd-4d2c-a24c-868a483171aa"), Otp: "8888", UserID: 8},
		{ID: uuid.MustParse("1529c576-9b6b-4054-a99d-6cec771bc5bd"), Otp: "9999", UserID: 9},
	}
}
