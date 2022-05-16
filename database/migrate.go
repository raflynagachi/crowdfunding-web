package database

import (
	"fmt"
	"log"

	"github.com/raflynagachi/crowdfunding-web/app"
	"github.com/raflynagachi/crowdfunding-web/app/config"
)

func GetDBConnection() {

}

func MigrateDB(fresh bool) {
	dbConfig, _ := config.ConfigEnv()
	db := app.OpenDB(dbConfig)

	defer fmt.Println("Database migrating...")
	for _, model := range RegisterModel() {
		if fresh {
			err := db.Debug().Migrator().DropTable(model.Model)
			if err != nil {
				log.Fatal(err)
			}
		}

		err := db.Debug().Migrator().AutoMigrate(model.Model)
		if err != nil {
			log.Fatal(err)
		}

	}
	if fresh {
		fmt.Println("Database fresh migration successfully")
	} else {
		fmt.Println("Database migration successfully")
	}
}
