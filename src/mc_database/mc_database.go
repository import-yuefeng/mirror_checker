package mc_database

import (
	"database/sql"
	"log"
	"os"

	"github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

const (
	port = ":5759"
)

var mysqlConfig = &mysql.Config{
	User:                 "root",
	Passwd:               "eibZLC123",
	Net:                  "tcp",
	Addr:                 "94.237.73.226",
	DBName:               "mirror_checker",
	AllowNativePasswords: true,
}
var mysqlDSN = mysqlConfig.FormatDSN()

func init() {
	err := godotenv.Load("/Users/zhuluchi/go/mirror_checker/src/.env")
	if err != nil {
		return // no .env file, just use the default mysql config
	}

	mysqlConfig.User = os.Getenv("DB_USER")
	mysqlConfig.Passwd = os.Getenv("DB_PASSWD")
	mysqlConfig.Addr = os.Getenv("DB_ADDR")
	mysqlConfig.DBName = os.Getenv("DB_NAME")
	mysqlDSN = mysqlConfig.FormatDSN()
}

type server struct{}

func ConnDatabase() {
	db, err := sql.Open("mysql", mysqlDSN)
	if err != nil {
		log.Fatalf("failed to create init mysql: %v", err)
	}
	log.Println("Initialized the MySQL connection pool")
	db.Close()
}

// func InsertDatabase()
