package utils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRemoveDuplicateUsers(t *testing.T) {
	users := []int{123456, 51256, 123456, 98741, 9837982, 83651, 123456, 123456 , 129751, 999999, 123456, 123456, 51256, 51256, 83651, 83651}
	filtered := RemoveDuplicateUsers(users)

	assert.Equal(t, 7, len(filtered))
}