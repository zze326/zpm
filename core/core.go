package core

import (
	"database/sql"
	"github.com/zze326/zpm/util"
	"log"
	"os"
	"path"
)

var Db *sql.DB
var Key string

const MAIN_TABLE = "main"

func Init() {
	Key = os.Getenv("ZPM_KEY")
	if util.StringIsEmpty(Key) {
		Key = "zzedaefeageaafeawgewagacageagewg"
	} else {
		if len(Key) != 32 {
			log.Fatalln("The ZPM_KEY must be 32 bits in length.")
		}
	}

	Db = initDB()
}

func DeferFunc() {
	Db.Close()
}

func getZpmFilePath() string {
	zpmFilePath := os.Getenv("ZPM_FILE_PATH")
	if util.StringIsEmpty(zpmFilePath) {
		homeDir, err := os.UserHomeDir()
		if err != nil {
			log.Fatalln(err)
		}
		zpmFilePath = path.Join(homeDir, ".zpmdata")
	} else {
		if !util.PathIsAbsolute(zpmFilePath) {
			log.Fatalln("ZPM_FILE_PATH environment variable is not an absolute path")
		}
	}

	err, exists, isDir := util.CheckPathInfo(zpmFilePath)
	if err != nil {
		log.Fatal(err)
	}
	if exists && isDir {
		zpmFilePath = path.Join(zpmFilePath, ".zpmdata")
	}

	return zpmFilePath
}
