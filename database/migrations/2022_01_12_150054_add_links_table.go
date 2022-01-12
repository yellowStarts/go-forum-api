package migrations

import (
	"database/sql"
	"huango/app/models"
	"huango/pkg/migrate"

	"gorm.io/gorm"
)

func init() {

	type Link struct {
		models.BaseModel

		Name string `gorm:"type:varchar(255);not null"`
        URL  string `gorm:"type:varchar(255);default:null"`

		models.CommonTimestampsFiels
	}

	up := func(migrator gorm.Migrator, DB *sql.DB) {
		migrator.AutoMigrate(&Link{})
	}

	down := func(migrator gorm.Migrator, DB *sql.DB) {
		migrator.DropTable(&Link{})
	}

	migrate.Add("2022_01_12_150054_add_links_table", up, down)
}
