package repository

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

// func InitMysqlDb() {
	
// 	db, err := sql.Open("mysql", "root:admi@/new_erp")
// 	if err != nil {
// 		fmt.Println(err, "Error")
// 	}

// 	fmt.Println(db, "Database")
// }

func DbConn() (db *sql.DB) {
    dbDriver := "mysql"
    dbUser := "root"
    dbPass := "admin"
    dbName := "new_erp"
    db, err := sql.Open(dbDriver, dbUser+":"+dbPass+"@/"+dbName)
    if err != nil {
        panic(err.Error())
    }
    fmt.Println(db, "Database")
    return db
}

func Select() {
    db := DbConn()
    selDB, err := db.Query("SELECT * FROM adjustment_details ORDER BY id DESC")
    // fmt.Println(selDB, "RESULT INIT")
    if err != nil {
        panic(err)
    }

    // Get column names
	columns, err := selDB.Columns()
	if err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
    }
    
    // Make a slice for the values
	values := make([]sql.RawBytes, len(columns))

	// rows.Scan wants '[]interface{}' as an argument, so we must copy the
	// references into such a slice
	// See http://code.google.com/p/go-wiki/wiki/InterfaceSlice for details
	scanArgs := make([]interface{}, len(values))
	for i := range values {
		scanArgs[i] = &values[i]
	}

    var result []map[string]interface{}
    for selDB.Next() {

        // var id, adjustment_id, employee_id, payroll_item_id int
        // var amount float32

        err = selDB.Scan(scanArgs...)
        if err != nil {
            fmt.Println(err, "RESULT SET")
        }

        var value string
        for i, col := range values {
			// Here we can check if the value is nil (NULL value)
			if col == nil {
				value = "NULL"
			} else {
				value = string(col)
			}
            fmt.Println(columns[i], ": ", value)
		}
        fmt.Println("-----------------------------------")
    }

    defer db.Close()

    fmt.Println(result)
}
