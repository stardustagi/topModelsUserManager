package common

import (
	"fmt"
	"logistics/libs/databases"
	"logistics/libs/logs"
	"logistics/models"
	"time"

	"github.com/go-xorm/xorm/migrate"
	"go.uber.org/zap"
)

var counter int

func generateMigrationID(tableName string) string {
	counter++
	return fmt.Sprintf("%s_%s_%02d", time.Now().Format("20060102"), tableName, counter)
}

func DatabaseSync() bool {
	//在这里处理数据库同步的逻辑
	// Define migration options
	options := &migrate.Options{
		TableName:    "migrations",
		IDColumnName: "id",
	}

	tables := []map[string]interface{}{
		{"id": generateMigrationID("users"), "name": new(models.Users)},
	}
	// 清空Migration表
	dao := databases.GetDao()
	ok, _ := dao.Exists("migrations")
	// if err != nil {
	// 	logs.Error("check migrations table exists error: ", zap.Error(err))
	// 	panic(err)
	// }
	if !ok {
		// 创建表
		if err := dao.Native().CreateTables("migrations", &models.Migrations{}); err != nil {
			logs.Error("create migrations table error: ", zap.Error(err))
			panic(err)
		}
	} else {
		// 清空表
		if _, err := dao.Native().Exec("TRUNCATE TABLE migrations"); err != nil {
			logs.Error("truncate migrations table error: ", zap.Error(err))
			return false
		}
	}

	if err := dao.Migrations(options, tables); err != nil {
		logs.Error("sync databases error: ", zap.Error(err))
		panic(err)
	}
	return true
}
