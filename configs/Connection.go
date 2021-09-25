package configs

import (
	"database/sql"
	"fmt"
	"os"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

const (
	DBusername = "root"
	DBpassword = ""
	DBname     = "dbblog"
)

var (
	DBUSER = os.Getenv("DB_USER")
	DBPASS = os.Getenv("DB_PASS")
	DBHOST = os.Getenv("DB_HOST")
	DBNAME = os.Getenv("DB_NAME")
)

func Connection() (*sql.DB, error) {
	var url string

	if DBUSER != "" {
		url = fmt.Sprintf("%s:%s@%s/%s?charset=utf8mb4&parseTime=True&loc=Local", DBUSER, DBPASS, DBHOST, DBNAME)
	} else {
		url = fmt.Sprintf("%s:%s@tcp(localhost:3306)/%s?charset=utf8mb4&parseTime=True&loc=Local", DBusername, DBpassword, DBname)
	}
	fmt.Println(url)
	db, err := sql.Open("mysql", url)

	if err != nil {
		return nil, err
	}
	// Limit the number of connection used by the application
	db.SetMaxIdleConns(10)
	// Maximum limit the number of connection used by the application
	db.SetMaxOpenConns(100)
	// Maximum amount of time a connection may be reused.
	db.SetConnMaxLifetime(time.Hour)

	return db, nil
}
