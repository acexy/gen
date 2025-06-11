module gorm.io/gen/tests

go 1.23.5


require (
	gorm.io/driver/mysql v1.6.0
	gorm.io/driver/sqlite v1.5.0
	gorm.io/gen v0.3.19
	gorm.io/gorm v1.30.0
	gorm.io/plugin/dbresolver v1.6.0
)

require github.com/mattn/go-sqlite3 v1.14.16 // indirect

replace gorm.io/gen => ../
