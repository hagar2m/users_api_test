package post

import (
	"fmt"
	"time"

	"com.test.users_api_test/api/comment"
	"com.test.users_api_test/api/models"
	"com.test.users_api_test/db"
)

func CreatePostQuery(post *models.Post) (*models.Post, error) {
	post.CreatedAt = time.Now()
	result := db.DB.Create(&post)
	if result.Error != nil {
		return nil, result.Error
	}

	return post, nil
}
func GetPostCommentsByIdQuery(id int) (*models.Post, error) {
	//1-  Get The post
	post := models.Post{}
	result := db.DB.Where("id = ?", id).Find(&post)
	if result.Error != nil {
		return nil, result.Error
	}
	if result.RowsAffected == 0 {
		return nil, fmt.Errorf("user not found")
	}
	//2- get Comments of this Post
	commentsList, err := comment.GetListOfCommentByPostIdQuery(id)
	if err != nil {
		return nil, err
	}
	//3- append the sub comment to the parent comment
	// for _, parentComment := range commentsList {
	// 	subComment, err := comment.GetCommentByIdQuery(parentComment.ID)
	// 	fmt.Println("subComment: ", subComment)
	// 	if err != nil {
	// 		// post.Comments = append(post.Comments, *subComment)
	// 	} else {
	// 	fmt.Println("=====err: ", err)

	// 		return nil, err
	// 	}
	// }
	post.Comments = commentsList
	return &post, nil
}
