package model

import (
	"go-echo-mongodb/types/dto"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	Id         primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Username   string             `json:"username" bson:"username"`
	FullName   string             `json:"full_name" bson:"full_name"`
	Sex        string             `json:"sex" bson:"sex"`
	Birthday   int64              `json:"birthday" bson:"birthday"`
	Phone      string             `json:"phone" bson:"phone"`
	Avatar     string             `json:"avatar" bson:"avatar"`
	LastLogin  int64              `json:"last_login" bson:"last_login"`
	OS         string             `json:"os" bson:"os"`
	Email      string             `json:"email" bson:"email"`
	IsVerified bool               `json:"is_verified" bson:"is_verified"`
	IsActive   bool               `json:"is_active" bson:"is_active"`
	IsAdmin    bool               `json:"is_admin" bson:"is_admin"`
	Password   string             `json:"password" bson:"password"`
}

func (u User) ConvertFromUserModelToDto() dto.User {
	var uDto dto.User
	uDto.Username = u.Username
	uDto.FullName = u.FullName
	uDto.Avatar = u.Avatar
	uDto.Sex = u.Sex
	uDto.Birthday = u.Birthday
	uDto.Phone = u.Phone
	uDto.OS = u.OS
	uDto.Email = u.Email
	return uDto
}

func CreateUserRequest(uDto *dto.User) (u User) {
	u.Username = uDto.Username
	u.FullName = uDto.FullName
	u.Avatar = uDto.Avatar
	u.Sex = uDto.Sex
	u.Birthday = uDto.Birthday
	u.Phone = uDto.Phone
	u.LastLogin = time.Now().Unix()
	u.OS = uDto.OS
	u.Email = uDto.Email
	u.IsActive = true
	u.IsVerified = false
	u.IsAdmin = false
	return u
}
