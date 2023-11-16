package conventer

import (
	"encoding/json"
	"fmt"
	"strconv"

	"github.com/gin-gonic/gin"
)

func ConvertStringToUint(value string) (uint, error) {
	valueUint, err := strconv.Atoi(value)
	if err != nil {
		return 0, err
	}

	return uint(valueUint), nil
}

func ParseRequestBody(ctx *gin.Context, v interface{}) error {
	if err := json.NewDecoder(ctx.Request.Body).Decode(v); err != nil {
		return fmt.Errorf("error parsing JSON: %v", err)
	}
	return nil
}