package utils

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"log"
	"sync"
	"time"
)

type DB struct {
	DbMysql *gorm.DB
	DbPostgres *gorm.DB
}

var (
	onceDbMysql     sync.Once
	onceDbPostgres  sync.Once
	instanceDB *DB
)

// This connection for L4 application database (read only)
func GetInstanceMysqlDb() *gorm.DB {
	onceDbMysql.Do(func() {
		mysqlInfo := Config.Database.Mysql
		logs := fmt.Sprintf("[INFO] Connected to MYSQL TYPE = %s | LogMode = %+v", mysqlInfo.Host, mysqlInfo.LogMode)

		dbConfig := mysqlInfo.Username + ":" + mysqlInfo.Password + "@tcp(" + fmt.Sprintf("%s:%d", mysqlInfo.Host, +mysqlInfo.Port) + ")/" + mysqlInfo.Name
		dbConnection, err := gorm.Open("mysql", dbConfig+"?charset=utf8&parseTime=True&loc=Local")
		if err != nil {
			logs = "[ERROR] Failed to connect to MYSQL. Config=" + mysqlInfo.Host
			log.Fatalln(logs)
		}
		fmt.Println(logs)
		instanceDB = &DB{DbMysql: dbConnection}
		dbConnection.LogMode(mysqlInfo.LogMode)
		dbConnection.SingularTable(true)
		dbConnection.DB().SetMaxIdleConns(10)
		dbConnection.DB().SetMaxOpenConns(20)
		dbConnection.DB().SetConnMaxLifetime(10 * time.Minute)
	})
	return instanceDB.DbMysql
}

// This connection for L4 application database (read only)
func GetInstancePostgresDb() *gorm.DB {
	onceDbPostgres.Do(func() {
		psqlInfo := Config.Database.Postgres
		logs := fmt.Sprintf("[INFO] Connected to Postgre TYPE = %s | LogMode = %+v", psqlInfo.Host, psqlInfo.LogMode)
		
		dbConfig := "host=" + psqlInfo.Host + " port=" + fmt.Sprintf("%d", psqlInfo.Port)  + " user=" + psqlInfo.Username + " dbname=" + psqlInfo.Name + " sslmode=" + psqlInfo.SslMode + " fallback_application_name=gohai-service"
		
		if psqlInfo.Password != "" {
			dbConfig += " password=" + psqlInfo.Password
		}

		dbConnection, err := gorm.Open("postgres", dbConfig)
		if err != nil {
			logs = "[ERROR] Failed to connect to Postgre. Config=" + psqlInfo.Host
			log.Fatalln(logs)
		}
		fmt.Println(logs)
		instanceDB = &DB{DbPostgres: dbConnection}
		dbConnection.LogMode(psqlInfo.LogMode)
		dbConnection.SingularTable(true)
		dbConnection.DB().SetMaxIdleConns(10)
		dbConnection.DB().SetMaxOpenConns(20)
		dbConnection.DB().SetConnMaxLifetime(10 * time.Minute)
	})
	return instanceDB.DbPostgres
}
