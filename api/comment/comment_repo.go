package comment

import (
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
