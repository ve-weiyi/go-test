package service

import (
	"errors"
	"go-test/global"
	"go-test/model/entity"
	"go-test/model/request"
	"go-test/model/response"
	"go-test/utils/jsonconv"
	"log"
)

type LoginService struct {
}

func (l LoginService) Login(user *request.UserLoginReq) (*response.UserLoginResp, error) {
	log.Printf("11--->%+v", user)
	db := global.CONFIG.DB

	var login entity.UserAuth
	err := db.Where("username = ?", user.Username).First(&login).Error

	if err != nil {
		return nil, errors.New("查找不到用户")
	}

	log.Printf("22--->%+v", login)

	var resp response.UserLoginResp
	jsonconv.DeepCopyByJson(login, &resp)
	return &resp, nil

	//return nil, errors.New("你什么都没有做")
}
