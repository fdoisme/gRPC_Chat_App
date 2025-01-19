package database

import (
	"gorm.io/gorm"
	"project/auth-service/model"
)

func Migrate(db *gorm.DB) error {
	var err error

	if err = dropTables(db); err != nil {
		return err
	}

	if err = setupJoinTables(db); err != nil {
		return err
	}

	if err = autoMigrates(db); err != nil {
		return err
	}

	if err = createCompositeIndexes(db); err != nil {
		return err
	}

	return createViews(db)
}

func autoMigrates(db *gorm.DB) error {
	return db.AutoMigrate(
		&model.User{},
		&model.Otp{},
	)
}

func dropTables(db *gorm.DB) error {
	return db.Migrator().DropTable(
		&model.Otp{},
		&model.User{},
	)
}

func setupJoinTables(db *gorm.DB) error {
	var err error

	return err
}

func createViews(db *gorm.DB) error {
	var err error

	return err
}

func createCompositeIndexes(db *gorm.DB) error {
	var err error
	if err = db.Exec(model.UserEmailUniqueIndex()).Error; err != nil {
		return err
	}
	return err
}
