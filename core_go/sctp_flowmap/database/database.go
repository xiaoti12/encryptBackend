package database

import (
	"database/sql"
	"fmt"
	"github.com/ClickHouse/clickhouse-go"
	"log"
)

var (
	Connect *sql.DB
)

var task string

// Init connect before main function
func init() {
	var err error

	// connect clickhouse
	Connect, err = sql.Open("clickhouse", "tcp://10.3.242.84:9000?&compress=true&debug=false&password=password")
	CheckErr(err)
	if err := Connect.Ping(); err != nil {
		if exception, ok := err.(*clickhouse.Exception); ok {
			fmt.Printf("[%d] %s \n%s\n", exception.Code, exception.Message, exception.StackTrace)
		} else {
			fmt.Println(err)
		}
		return
	}

	// 创建 SCTP 数据库
	_, err = Connect.Exec("CREATE DATABASE IF NOT EXISTS SCTP_1;")
	CheckErr(err)
}

func CheckErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
