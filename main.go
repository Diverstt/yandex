package main

import (
	"strings"
)

type UserCreateRequest struct {
	FirstName string
	Age       int
}

func Validate(req UserCreateRequest) string {
	err := strings.Contains(req.FirstName, " ")
	if err == true || req.FirstName == "" {
		return "invalid request"
	}
	if req.Age <= 0 || req.Age > 150 {
		return "invalid request"
	}
	return ""

}
