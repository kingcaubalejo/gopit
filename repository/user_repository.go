package repository

import (
	 "database/sql"
    _ "fmt"
    "errors"
    
    "go-api-jwt-v2/services/models"
    "go-api-jwt-v2/interfaces"

    _ "github.com/go-sql-driver/mysql"
)

type DbUserRepo struct {
    Repository interfaces.User
}

func (dbUserRepo *DbUserRepo) DisplayList() ([]models.Users, error) {
    rows, err := database.Query("SELECT uuid, username, password FROM users ORDER BY uuid DESC")

    if err != nil {
        return []models.Users{}, err
    }

    var users []models.Users
    for rows.Next() {
        user := models.Users{}
        err = rows.Scan(&user.UUID, &user.Username, &user.Password)

        if err != nil {
            return []models.Users{}, err
        }

        users = append(users, user)
    }

    err = rows.Err()
    if err != nil {
        return []models.Users{}, err
    }

    return users, err
}

func (dbUserRepo *DbUserRepo) DisplayListById(uuid int) (models.Users, error) {
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

func (dbUserRepo *DbUserRepo) Save(u models.Users) error {
    createUser, err := database.Prepare("INSERT INTO users (username, password) VALUES(?,?)")
    if err != nil {
        return err
    }

    createUser.Exec(u.Username, u.Password)
    
    return nil
}

func (dbUserRepo *DbUserRepo) Update(u models.Users) error {
    updateUser, err := database.Prepare("UPDATE users SET username=? WHERE uuid=?")
    if err != nil {
        return err
    }

    updateUser.Exec(u.Username, u.UUID)

    return nil
}

func (dbUserRepo *DbUserRepo) Delete(uuid int) error {
    deleteUser, err := database.Prepare("DELETE FROM users WHERE uuid=?")
    if err != nil {
        return err
    }

    deleteUser.Exec(uuid)

    return nil
}