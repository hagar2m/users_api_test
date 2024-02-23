package comment

import (
	"fmt"
	"time"

	"com.test.users_api_test/api/models"
	"com.test.users_api_test/db"
)

func CreateCommentQuery(comment *models.Comment) (*models.Comment, error) {
	comment.CreatedAt = time.Now()
	result := db.DB.Create(&comment)
	if result.Error != nil {
		return nil, result.Error
	}

	return comment, nil
}

func GetCommentByIdQuery(id int) (*models.Comment, error) {
	//1- Get the parent comments
	var parentComment models.Comment
	db.DB.Where("id = ?", id).Find(&parentComment)

	//2- Get the sub comments
	var subComments []models.Comment
	result2 := db.DB.Where("parent_id = ?", id).Find(&subComments)
	if result2.Error != nil {
		return nil, result2.Error
	}
	//3- append the sub comment to the parent comment
	for _, subComment := range subComments {
		fmt.Println("subComment: ", subComment)
		parentComment.Comments = append(parentComment.Comments, subComment)
	}

	return &parentComment, nil
}
