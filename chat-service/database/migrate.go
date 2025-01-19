package database

import (
	"project/chat-service/model"

	"gorm.io/gorm"
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

	return createViews(db)
}

func autoMigrates(db *gorm.DB) error {
	return db.AutoMigrate(
		&model.Room{},
		&model.RoomParticipant{},
		&model.Message{},
	)
}

func dropTables(db *gorm.DB) error {
	return db.Migrator().DropTable(
		&model.Message{},
		&model.RoomParticipant{},
		&model.Room{},
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
