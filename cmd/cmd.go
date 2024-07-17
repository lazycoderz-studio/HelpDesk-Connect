package cmd

import (
	"fmt"
	"lazycoderz-studio/helpdesk-connect/config"
	"lazycoderz-studio/helpdesk-connect/connect/repository"
	"lazycoderz-studio/helpdesk-connect/db"

	"gorm.io/gorm"
)

func Execute() {

	InitDB()
	defer db.CloseDB()

	gormDb := db.GetDB()

	InitRoutes(gormDb)

}

func InitDB() {
	cfg := config.GetConfig()

	if err := db.InitDB(&cfg.Database); err != nil {
		panic(err)
	}
	db.Migrate()
}

func InitRoutes(db *gorm.DB) {

	repo := repository.NewRepositories(db)
	fmt.Println(repo)

}
