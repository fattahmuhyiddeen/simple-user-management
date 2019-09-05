package model

import (
	"database/sql"
	"log"
	"strconv"
	"strings"
	"time"

	config "github.com/fattahmuhyiddeen/simple-user-management/config"

	//blank import to init database here
	_ "github.com/lib/pq"
)

//TestConnection is the function to check connection with databases. if fail to connect, it will crash the system. need to restart
func TestConnection() bool {
	connectDB()
	defer disconnectDB()

	query := `SELECT * FROM pg_catalog.pg_tables`
	_, err := db.Query(query)
	log.Println(err)

	return err == nil
}

var db *sql.DB

func connectDB() {
	db, _ = sql.Open("postgres", config.DatabaseURL)
}

func disconnectDB() {
	db.Close()
}

//DateTimeNow return timestamp used for created_at and updated_at
func DateTimeNow() string {
	return time.Now().Format(config.DateTimeFormat)
}

func tablePlaceholder(columns string) (result string) {
	for index := range strings.Split(columns, ",") {
		result += "$" + strconv.Itoa(index+1) + ","
	}
	return strings.TrimSuffix(result, ",")
}
