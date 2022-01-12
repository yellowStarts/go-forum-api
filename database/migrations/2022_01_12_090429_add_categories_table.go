package migrations

import (
	"database/sql"
	"huango/app/models"
	"huango/pkg/migrate"

	"gorm.io/gorm"
)

func init() {

	type Category struct {
		models.BaseModel

		Name        string `gorm:"type:varchar(255);not null;index"`
		Description string `gorm:"type:varchar(255);default:''"`
		ParentID    uint64 `gorm:"column:parent_id;type:uint;default:0;index"`
		Order       int    `gorm:"type:int(11)"`

		models.CommonTimestampsFiels
	}

	up := func(migrator gorm.Migrator, DB *sql.DB) {
		migrator.AutoMigrate(&Category{})
	}

	down := func(migrator gorm.Migrator, DB *sql.DB) {
		migrator.DropTable(&Category{})
	}

	migrate.Add("2022_01_12_090429_add_categories_table", up, down)
}
