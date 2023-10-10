package models

type ResponseModel struct {
	Message string      `json:"msg"`
	Data    interface{} `json:"data"`
}
