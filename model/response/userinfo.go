package response

import "time"

type UserInfoResp struct {
	ID         int        `gorm:"column:id;type:int;primaryKey;autoIncrement:true" json:"id"`             // 用户ID
	Email      *string    `gorm:"column:email;type:varchar(50)" json:"email"`                             // 邮箱号
	Nickname   string     `gorm:"column:nickname;type:varchar(50);not null" json:"nickname"`              // 用户昵称
	Avatar     string     `gorm:"column:avatar;type:varchar(1024);not null" json:"avatar"`                // 用户头像
	Intro      *string    `gorm:"column:intro;type:varchar(255)" json:"intro"`                            // 用户简介
	WebSite    *string    `gorm:"column:web_site;type:varchar(255)" json:"web_site"`                      // 个人网站
	IsDisable  int        `gorm:"column:is_disable;type:tinyint(1);not null" json:"is_disable"`           // 是否禁用
	CreateTime time.Time  `gorm:"column:create_time;type:int unsigned;autoCreateTime" json:"create_time"` // 创建时间
	UpdateTime *time.Time `gorm:"column:update_time;type:int unsigned;autoUpdateTime" json:"update_time"` // 更新时间
}
