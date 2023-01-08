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

type UserInfoService struct {
}

func (l UserInfoService) GetUserInfo(req *request.Userinfo) (*response.UserInfoResp, error) {
	log.Printf("11--->%+v", req)
	db := global.CONFIG.DB

	var userinfo entity.UserInfo
	err := db.Where("email = ?", req.Email).First(&userinfo).Error

	if err != nil {
		return nil, errors.New("查找不到用户")
	}

	log.Printf("22--->%+v", userinfo)

	var resp response.UserInfoResp
	jsonconv.DeepCopyByJson(userinfo, &resp)
	return &resp, nil

	//return nil, errors.New("你什么都没有做")
}
