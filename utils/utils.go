package utils

import (
	"denounce-abandoned-items/domain"
)

func RemoveDuplicateUsers(users []int) []int {
	allKeys := make(map[int]bool)
	var list []int
	for _, usr := range users {
		if _, value := allKeys[usr]; !value {
			allKeys[usr] = true
			list = append(list, usr)
		}
	}
	return list
}

func BuildEmail(userID int) domain.Email {
	ctx := domain.Context{}
	email := domain.Email{
		UserID:   userID,
		Template: "MOD_PAUSE_TEST",
		Context:  ctx,
	}
	return email
}
