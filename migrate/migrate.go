package migrate

import (
	"fmt"

	"github.com/AudiWu/mtg-backend-project/db"
	"github.com/AudiWu/mtg-backend-project/model"
)

var models = []interface{}{&model.User{}, &model.Card{}}

func Check() {
	db := db.GetDB()

	fmt.Println("Check Database table ...")

	for _, model := range models {
		if db.Migrator().HasTable(model) {
			continue
		} else {
			db.AutoMigrate(model)
		}
	}

	fmt.Println("Finished Migrating!")
}
