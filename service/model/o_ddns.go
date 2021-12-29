package model

import "time"

func (p *DDNSUpdateDBModel) TableName() string {
	return "o_ddns"
}

type DDNSUpdateDBModel struct {
	ID       uint      `gorm:"column:id;primary_key" json:"id"`
	Ipv4     string    `gorm:"-"`
	Ipv6     string    `gorm:"-"`
	Type     uint      `json:"type" form:"type"`
	Domain   string    `json:"domain" form:"domain"`
	Host     string    `json:"host" form:"host"`
	Key      string    `json:"key" form:"key"`
	Secret   string    `json:"secret" form:"secret"`
	UserName string    `json:"user_name" form:"user_name"`
	Password string    `json:"password" form:"password"`
	CreateAt time.Time `gorm:"<-:create" json:"created_at"`
	UpdateAt time.Time `gorm:"<-:create;<-:update" json:"updated_at"`
}

const DDNSLISTTABLENAME = "o_ddns"

// 返回给前台使用
type DDNSList struct {
	Id        uint      `gorm:"column:id;primary_key" json:"id"`
	Name      string    `json:"name"`
	Domain    string    `json:"domain" form:"domain"`
	Host      string    `json:"host" form:"host"`
	IPV4      string    `json:"ipv_4" gorm:"-"`
	IPV6      string    `json:"ipv_6" gorm:"-"`
	Message   string    `json:"message"`
	State     bool      `json:"state"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
