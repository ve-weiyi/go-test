package api

import (
	"go-test/model/entity"
	"log"
	"testing"
)

func TestName(t *testing.T) {
	var user entity.UserAuth

	user.Username = "abc"
	user.Password = "123"

	log.Printf("1111--->%v", user)
	log.Printf("1111--->%+v", user)
}
