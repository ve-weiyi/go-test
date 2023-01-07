package service

import (
	"go-test/global"
	"go-test/model/entity"
	"go-test/model/request"
	"log"
)

type LoginService struct {
}

// 参数:user request.UserLoginReq  返回值(response.UserLoginResp,error)
func (l LoginService) Login(user *request.UserLoginReq) (*entity.User, error) {
	log.Printf("打印接收参数--->%+v", user)

	db := global.CONFIG.DB     //开启数据库链接
	var user_table entity.User //定义数据库表 user_table，表明查询这张表
	err := db.Where("username = ? and password =? ", user.Uloginid, user.Upassword).First(&user_table).Error
	log.Printf("user_table--->%+v", user_table)
	if err != nil {

		return nil, err
	} //查询不到

	return &user_table, nil //查询到了
}
