package util

import (
	"database/sql"
	"fmt"
)

// TableExists 判断 SQLite 数据库中一个表是否存在
func TableExists(db *sql.DB, tableName string) bool {
	var count int
	err := db.QueryRow("SELECT count(*) FROM sqlite_master WHERE type='table' AND name=?", tableName).Scan(&count)
	if err != nil {
		fmt.Printf("查询表是否存在出错: %v\n", err)
		return false
	}
	return count > 0
}
