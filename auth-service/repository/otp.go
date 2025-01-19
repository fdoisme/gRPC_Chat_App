package repository

import (
	"go.uber.org/zap"
	"gorm.io/gorm"
	"project/auth-service/model"
)

type OtpRepository struct {
	db  *gorm.DB
	log *zap.Logger
}

func NewOtpRepository(db *gorm.DB, log *zap.Logger) *OtpRepository {
	return &OtpRepository{db, log}
}

func (repo *OtpRepository) Create(otp *model.Otp) error {
	return repo.db.Create(&otp).Error
}

func (repo *OtpRepository) Update(criteria model.Otp) (*model.Otp, error) {
	result := repo.db.Model(&model.Otp{}).
		Where("otps.id = ?", criteria.ID).
		Where("otps.otp = ?", criteria.Otp).
		Where("validated_at IS NULL").
		Update("validated_at", gorm.Expr("now()"))

	if result.Error != nil {
		return nil, result.Error
	}

	if result.RowsAffected == 0 {
		return nil, gorm.ErrRecordNotFound
	}

	otp := model.Otp{}
	repo.db.Preload("User").First(&otp, "id=?", criteria.ID)
	return &otp, nil
}
