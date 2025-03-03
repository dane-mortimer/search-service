package utils

import (
	"search-service/models"
	"strings"
)

func CreateAdminPath(path string, method string) models.AdminPath {
	method = strings.ToUpper(method)

	validMethods := map[string]bool{
		"GET":     true,
		"POST":    true,
		"PUT":     true,
		"DELETE":  true,
		"PATCH":   true,
		"HEAD":    true,
		"OPTIONS": true,
		"*":       true,
	}

	if !validMethods[method] {
		method = "*"
	}

	return models.AdminPath{
		Path:   path,
		Method: method,
	}
}
