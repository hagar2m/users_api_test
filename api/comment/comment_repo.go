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
func GetListOfCommentByPostIdQuery(id int) ([]models.Comment, error) {
	//1- Get the post comments
	var postCommentsList []models.Comment
	result := db.DB.Where("post_id = ?", id).Where("parent_id IS NULL").Find(&postCommentsList)
	if result.Error != nil {
		return nil, result.Error
	}
	var newCommentsList []models.Comment
	for _, postComment := range postCommentsList {
		// Get sub-comments for the current post comment
		var subComments []models.Comment
		result2 := db.DB.Where("parent_id = ?", postComment.ID).Find(&subComments)
		if result2.Error == nil {
			// Assign sub-comments to a new copy of postComment
			newComment := postComment
			newComment.Comments = make([]models.Comment, len(subComments))
			copy(newComment.Comments, subComments)

			newCommentsList = append(newCommentsList, newComment)
		}
	}
	return newCommentsList, nil
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
