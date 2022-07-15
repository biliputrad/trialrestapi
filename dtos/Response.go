package dtos

import "RestAPI-GETNPOST/Entity"

type Data interface {
	Entity.User
}
type Response struct {
	Success bool        `json:"success"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}
