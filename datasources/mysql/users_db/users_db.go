package users_db

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	// pkg for .env
	_ "github.com/joho/godotenv/autoload"
	// the _ is to import the pkg but not use it, just for initializing
	_"github.com/go-sql-driver/mysql"
)

const (
	mysql_users_db_username = "mysql_users_db_username"
	mysql_users_db_password = "mysql_users_db_password"
	mysql_users_db_host = "mysql_users_db_host"
	mysql_users_db_schema = "mysql_users_db_schema"
)

var (
	Client *sql.DB

	username = os.Getenv(mysql_users_db_username)
	password = os.Getenv(mysql_users_db_password)
	host = os.Getenv(mysql_users_db_host)
	schema = os.Getenv(mysql_users_db_schema)
)

// init() get triggered as soon as package is imported
func init() {
	// log.Println(username, password, host, schema)
	datasourceName := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8",
		username, 
		password, 
		host, 
		schema,
	)
	var err error
	Client, err = sql.Open("mysql", datasourceName)
	if err != nil {
		panic(err)
	}
	if err = Client.Ping(); err != nil {
		panic(err)
	}
	log.Println("database successfully configured")
}