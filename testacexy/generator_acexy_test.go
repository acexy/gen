package testacexy

import (
	"github.com/acexy/gen"
	"github.com/acexy/golang-toolkit/util/str"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"testing"
)

func TestCreateModel(t *testing.T) {

	g := gen.NewGenerator(gen.Config{
		OutPath: "./testacexy",
		Mode:    gen.WithDefaultQuery,
	})
	gormdb, _ := gorm.Open(mysql.Open("root:root@(127.0.0.1:13306)/test?charset=utf8mb4&parseTime=True&loc=Local"))
	g.UseDB(gormdb)

	g.WithJSONTagNameStrategy(func(c string) string { return str.SnakeToCamel(c) })
	g.DisableGormTag()
	//g.DisableModelTableNameMethod()

	g.WithTableNameStrategy(func(tableName string) string {
		return tableName
	})
	g.GenerateModelAs("demo_teacher", "Teacher",
		gen.FieldTypeReg("^(create_time|update_time)$", "gormstarter.Timestamp"),
		gen.FieldTypeReg("^id$", "gormstarter.BaseModel[uint64]"),
	)
	//g.GenerateModelAs("demo_student", "Student", gen.FieldTypeReg("^(create_time|update_time)$", "gormstarter.Timestamp"))
	g.Execute()
}
