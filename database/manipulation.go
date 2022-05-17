package database

import (
	"fmt"
	"log"

	"github.com/raflynagachi/crowdfunding-web/app"
	"github.com/raflynagachi/crowdfunding-web/app/config"
)

func MigrateDB(dbConfig config.DBConfig, fresh bool) {
	db := app.OpenDB(dbConfig)

	fmt.Println("Database migrating...")
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

func TruncateDB(dbConfig config.DBConfig) {
	db := app.OpenDB(dbConfig)

	fmt.Println("Database truncate start...")
	for _, model := range RegisterModel() {
		query := fmt.Sprintf("DELETE FROM %s", model.TableName)

		err := db.Debug().Exec(query).Error
		if err != nil {
			log.Fatal(err)
		}
	}

	fmt.Println("Database truncate successfully")
}
