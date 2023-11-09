package user

import (
	"fmt"
	"time"

	"com.test.users_api_test/db"
	"com.test.users_api_test/models"
)

func CreateUserQuery(user *models.UserTable) (*models.UserTable, error) {
	user.CreatedAt = time.Now()
	if err := db.DB.Create(user).Error; err != nil {
		return nil, err
	}
	return user, nil
}

func SingInQuery(email string, password string) (*models.UserTable, error) {
	loginUserData := models.UserTable{}
	result := db.DB.Where("email = ? AND password = ?", email, password).Find(&loginUserData)
	if result.Error != nil {
		return nil, result.Error
	}

	if result.RowsAffected == 0 {
		return nil, fmt.Errorf("user not found")
	}

	return &loginUserData, nil
}

func SearchByNameQuery(searchedName string) []models.UserTable {
	users := []models.UserTable{}
	likeValue := fmt.Sprintf("%%%s%%", searchedName)
	db.DB.Where("name LIKE ?", likeValue).Find(&users)
	return users
}

func SearchByEmailQuery(searchedEmail string) []models.UserTable {
	users := []models.UserTable{}
	db.DB.Where("email LIKE ?", searchedEmail).Find(&users)
	return users
}

func GetAllUsersQuery() []models.UserTable {
	users := []models.UserTable{}
	db.DB.Find(&users)
	return users
}

func GetUserByIdQuery(id uint) (*models.UserTable, error) {
	user := models.UserTable{}
	result := db.DB.Where("id = ?", id).Find(&user)
	if result.Error != nil {
		return nil, result.Error
	}
	if result.RowsAffected == 0 {
		return nil, fmt.Errorf("user not found")
	}
	return &user, nil
}

func UpdateUserQuery(user *models.UserTable) {
	db.DB.Save(&user)
}

func DeletUserQuery(id uint) (*models.UserTable, error) {
	user := models.UserTable{
		ID: id,
	}
	result := db.DB.Where("id = ?", id).Delete(&user)
	if result.Error != nil {
		return nil, fmt.Errorf("failed to delete user: " + result.Error.Error())

	} else if result.RowsAffected == 0 {
		return nil, fmt.Errorf("user not found")
	}
	return &user, nil
}
