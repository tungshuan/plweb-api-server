package db

import (
	"bytes"
	"database/sql"
	"github.com/Yuniii/plweb-api-server/config"
	_ "github.com/go-sql-driver/mysql"
)

var (
	DB *sql.DB
)

func init() {
	var err error
	DB, err = sql.Open("mysql", buildDataSource())
	if err != nil {
		panic(err)
	}
}

//[username[:password]@][protocol[(address)]]/dbname
func buildDataSource() string {
	var buffer bytes.Buffer
	buffer.WriteString(config.DB_CONFIG.User)
	buffer.WriteString(":")
	buffer.WriteString(config.DB_CONFIG.Password)
	buffer.WriteString("@tcp(")
	buffer.WriteString(config.DB_CONFIG.Host)
	buffer.WriteString(":")
	buffer.WriteString(config.DB_CONFIG.Port)
	buffer.WriteString(")/")
	buffer.WriteString(config.DB_CONFIG.DBName)
	return buffer.String()
}
