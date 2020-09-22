package main

import (
	"github.com/captain-blue210/youtube-sample/youtube-manager-go/databases"
	"github.com/captain-blue210/youtube-sample/youtube-manager-go/models"
	"github.com/sirupsen/logrus"
)

func main() {
	db, err := databases.Connect()
	defer db.Close()

	if err != nil {
		logrus.Fatal(err)
	}

	db.Debug().AutoMigrate(&models.User{})
	db.Debug().AutoMigrate(&models.Favorite{})
}
