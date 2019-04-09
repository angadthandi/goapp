package dbconnect

import (
	log "github.com/angadthandi/goapp/log"
	"github.com/angadthandi/goapp/util/config"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

func Conn(configData config.ConfigStruct) *gorm.DB {
	var dbRef *gorm.DB

	databasePath := configData.SQLiteDBPath
	databasePath += "?_busy_timeout=5000"
	dbRef, err := gorm.Open("sqlite3", databasePath)
	dbRef.DB().SetMaxIdleConns(25)
	if err != nil {
		log.Fatal("failed to connect database")
	}

	return dbRef
}
