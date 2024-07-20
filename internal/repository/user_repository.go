package repository

import (
    "myapp/internal/db"
    "myapp/internal/models"
)

func GetUsers() ([]models.User, error) {
    rows, err := db.GetDB().Query("SELECT id, name, email FROM users")
    if err != nil {
        return nil, err
    }
    defer rows.Close()

    var users []models.User
    for rows.Next() {
        var user models.User
        if err := rows.Scan(&user.ID, &user.Name, &user.Email); err != nil {
            return nil, err
        }
        users = append(users, user)
    }
    return users, nil
}

func CreateUser(user models.User) error {
    _, err := db.GetDB().Exec("INSERT INTO users (name, email) VALUES (?, ?)", user.Name, user.Email)
    return err
}

func UpdateUser(user models.User) error {
    _, err := db.GetDB().Exec("UPDATE users SET name = ?, email = ? WHERE id = ?", user.Name, user.Email, user.ID)
    return err
}

func DeleteUser(id int) error {
    _, err := db.GetDB().Exec("DELETE FROM users WHERE id = ?", id)
    return err
}
