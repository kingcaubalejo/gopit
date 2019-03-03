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
    rows, err := db.Query("SELECT * FROM adjustment_details ORDER BY id DESC")
    
    type AdjustmentDetails struct {
        Id                  int
        AdjustmentId        int
        EmployeeId          int
        PayrollItemId       int
        Amount              float32
    }


    if err != nil {
        panic(err)
    }

    defer rows.Close()

    for rows.Next() {
        adjustmentDetails := AdjustmentDetails{}
        err = rows.Scan(&adjustmentDetails.Id, &adjustmentDetails.AdjustmentId, &adjustmentDetails.EmployeeId, &adjustmentDetails.PayrollItemId, &adjustmentDetails.Amount)

        if err != nil {
            panic(err)
        }
        fmt.Println(adjustmentDetails)
    }

    err = rows.Err()
    if err != nil {
        panic(err)
    }
}