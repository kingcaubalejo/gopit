package repository

import (
	"database/sql"
	"fmt"

	"go-api-jwt/settings"

    _ "github.com/go-sql-driver/mysql"
)

var database *sql.DB

func InitDatabase() {
	
    dbDataSource := fmt.Sprintf("%s:%s@tcp(%s)/%s", settings.Get().DbHost, settings.Get().DbPassword, settings.Get().DbServerAddr, settings.Get().DbName)
    
    db, err := sql.Open(settings.Get().DbDriver, dbDataSource)
    
    if err != nil {
        panic(err.Error())
    }
    
    database = db
    fmt.Sprintf("Database connection to %s is established.", dbDataSource)
}