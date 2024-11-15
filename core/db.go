package core

import (
	"database/sql"
	"fmt"
	"github.com/zze326/zpm/util"
	"log"
)

func initDB() *sql.DB {
	db, err := sql.Open("sqlite3", getZpmFilePath())
	if err != nil {
		log.Fatalln("init db failed", err)
	}

	exists := util.TableExists(db, MAIN_TABLE)
	if !exists {
		if _, err := db.Exec(fmt.Sprintf(`CREATE TABLE IF NOT EXISTS %s (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    addr varchar(512) not null,
    pwd varchar(256),
    desc varchar(256)
);`, MAIN_TABLE)); err != nil {
			log.Fatalln("create table failed", err)
		}
	}
	return db
}
