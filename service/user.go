package service

import (
	"Learn-CasaOS/pkg/config"
	"io"
	"mime/multipart"
	"os"
)

type UserService interface {
	SetUser(username, pwd, token, email, desc string) error
	UploadFile(file multipart.File, name string) error
}

type user struct{}

func (c *user) SetUser(username, pwd, token, email, desc string) error {
	if len(username) > 0 {
		config.Cfg.Section("user").Key("UserName").SetValue(username)
		config.UserInfo.UserName = username
	}
	if len(pwd) > 0 {
		config.Cfg.Section("user").Key("PWD").SetValue(pwd)
		config.UserInfo.PWD = pwd
	}
	if len(token) > 0 {
		config.Cfg.Section("user").Key("Token").SetValue(token)
		config.UserInfo.Token = token
	}
	if len(email) > 0 {
		config.Cfg.Section("user").Key("Email").SetValue(email)
		config.UserInfo.Email = email
	}
	if len(desc) > 0 {
		config.Cfg.Section("user").Key("Description").SetValue(desc)
		config.UserInfo.Description = desc
	}
	config.Cfg.SaveTo("conf/conf.ini")
	return nil
}

// 上传文件
func (c *user) UploadFile(file multipart.File, url string) error {
	out, _ := os.OpenFile(url, os.O_WRONLY|os.O_TRUNC|os.O_CREATE, 0644)
	defer out.Close()
	io.Copy(out, file)
	return nil
}

func NewUserService() UserService {
	return &user{}
}
