package repository

import (
	 "database/sql"
    "fmt"
    "errors"
    
    "go-api-jwt-v2/services/models"
    "go-api-jwt-v2/interfaces"

    _ "github.com/go-sql-driver/mysql"
)

type DbUserRepo struct {
    Repository interfaces.UserSaver
}

func SelectAll() (map[string]interface{}) {
    rows, err := database.Query("SELECT uuid, username, password FROM users ORDER BY uuid DESC")

    if err != nil {
        panic(err)
    }

    var users []map[string]interface{}
    for rows.Next() {
        user := models.Users{}
        err = rows.Scan(&user.UUID, &user.Username, &user.Password)

        if err != nil {
            panic(err)
        }

        users = append(users, map[string]interface{}{
            "uuid": user.UUID,
            "username": user.Username,
            "password": user.Password,
        })
    }

    err = rows.Err()
    if err != nil {
        panic(err)
    }

    return map[string]interface{}{
        "result": users,
    }
}

func SelectWhere(uuid int) (map[string]interface{}) {
    distinctUsers := models.Users{}
    errUsers := database.QueryRow("SELECT uuid, username, password FROM users WHERE uuid=?", uuid).Scan(&distinctUsers.UUID, &distinctUsers.Username, &distinctUsers.Password)

    if errUsers != nil {
        if errUsers == sql.ErrNoRows {
            fmt.Println("Zero rows found")
        } else {
            panic(errUsers)
        }
    }

    return map[string]interface{}{
        "result": distinctUsers,
    }
}

func CreateUser(user *models.Users) ([]byte) {
    createUser, err := database.Prepare("INSERT INTO users (username, password) VALUES(?,?)")
    if err != nil {
        panic(err.Error())
    }
    createUser.Exec(user.Username, user.Password)
    
    return []byte("User is successfully created")
}

func UpdateUser(user *models.Users) {
    updateUser, err := database.Prepare("UPDATE users SET username=? WHERE uuid=?")
    if err != nil {
        panic(err.Error())
    }
    updateUser.Exec(user.Username, user.UUID)
}

func DeleteUser(user *models.Users) {
    deleteUser, err := database.Prepare("DELETE FROM users WHERE uuid=?")
    if err != nil {
        panic(err.Error())
    }
    deleteUser.Exec(user.UUID)
}

//Implementation of interfaces
func (dbUserRepo *DbUserRepo) DisplayList(uuid int) (models.Users, error) {
    distinctUsers := models.Users{}
    errUsers := database.QueryRow("SELECT uuid, username, password FROM users WHERE uuid=?", uuid).Scan(&distinctUsers.UUID, &distinctUsers.Username, &distinctUsers.Password)

    if errUsers != nil {
        if errUsers != sql.ErrNoRows {
            return distinctUsers, errUsers
        }
        return distinctUsers, errors.New("Zero rows found.")
    }

    return distinctUsers, nil
}