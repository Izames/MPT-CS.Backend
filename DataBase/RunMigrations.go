package DataBase

import "MPT-CS/models"

func RunMigrations() error {
	err := DB.AutoMigrate(
		&models.User{},
		&models.ResetPassword{},
	)
	if err != nil {
		return err
	}
	return nil
}
