package user

import (
	"fmt"
	"time"

	"com.test.users_api_test/database"
	"com.test.users_api_test/models"
)

func CreateUserQuery(user *models.UserTable) (*models.UserTable, error) {
	user.CreatedAt = time.Now()
	if err := database.DB.Create(user).Error; err != nil {
		return nil, err
	}
	return user, nil
}

func SingInQuery(email string, password string) (*models.UserTable, error) {
	loginUserData := models.UserTable{}
	result := database.DB.Where("email = ? AND password = ?", email, password).Find(&loginUserData)
	if result.Error != nil {
		return nil, result.Error
	}

	if result.RowsAffected == 0 {
		return nil, fmt.Errorf("User not found")
	}

	return &loginUserData, nil
}

func SearchByNameQuery(searchedName string) []models.UserTable {
	users := []models.UserTable{}
	likeValue := fmt.Sprintf("%%%s%%", searchedName)
	database.DB.Where("name LIKE ?", likeValue).Find(&users)
	return users
}

func SearchByEmailQuery(searchedEmail string) []models.UserTable {
	users := []models.UserTable{}
	database.DB.Where("email LIKE ?", searchedEmail).Find(&users)
	return users
}

func GetAllUsersQuery() []models.UserTable {
	users := []models.UserTable{}
	database.DB.Find(&users)
	return users
}

func GetUserByIdQuery(id uint64) (*models.UserTable, error) {
	user := models.UserTable{}
	result := database.DB.Where("id = ?", id).Find(&user)
	if result.Error != nil {
		return nil, result.Error
	}
	if result.RowsAffected == 0 {
		return nil, fmt.Errorf("User not found")
	}
	return &user, nil
}

func UpdateUserQuery(user *models.UserTable) {
	database.DB.Save(&user)
}
