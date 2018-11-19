package main

import (
	"github.com/jinzhu/gorm"
	"time"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

var DB *gorm.DB


func init() {

	url := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local", "username",
		"yourpassword",
		"localhost", "scrago_test")
	gormDB, err := gorm.Open("mysql", url)

	if err != nil {
		panic(err)
	}
	DB = gormDB
}

type Model struct {
	ID		int		`json:"id" sql:"id"`
	CreatedAt	time.Time	`sql:"default:CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt	time.Time	`sql:"default:CURRENT_TIMESTAMP" json:"updated_at"`
}

type User struct {
	Model
	UserName string `gorm:"column:user_name"`
	Avatar string
	Intro string
	FollowCount int
	Type int8
	FansCount int
	UserType int8
	IsVip int8
	PubshareCount int
	HotUk int64
	AlbumCount int
}

func (User) TableName() string {
	return "user"
}