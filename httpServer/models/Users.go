package models

import (
	"github.com/globalsign/mgo"
	"github.com/globalsign/mgo/bson"
)

// UserModel 用户数据库
type UserModel struct {
	DB *mgo.Collection
}

// Users 用户
type Users struct {
	ID    bson.ObjectId `bson:"_id"`   // 用户ID
	Name  string        `bson:"name"`  // 用户唯一名字
	Email string        `bson:"email"` // 邮箱
	Info  UserInfo      `bson:"info"`  // 用户个性信息
	Token string        `bson:"token"` // Violet 访问令牌
	Level int64         `bson:"level"` // 用户等级
}

// 性别
const (
	GenderMan int = iota
	GenderWoman
	GenderUnknown
)

// UserInfo 用户个性信息
type UserInfo struct {
	Avatar   string `bson:"avatar"`   // 头像URL
	Gender   int    `bson:"gender"`   // 性别
	NikeName string `bson:"nikeName"` // 昵称
}

// AddUser 添加用户
func (m *UserModel) AddUser() (bson.ObjectId, error) {
	newUser := bson.NewObjectId()
	err := m.DB.Insert(&Users{
		ID:   newUser,
		Name: "user_" + string(newUser),
	})
	if err != nil {
		return "", err
	}
	return newUser, nil
}

// SetUserInfo 设置用户信息
func (m *UserModel) SetUserInfo(id string, info UserInfo) (err error) {
	_, err = m.DB.UpsertId(bson.ObjectIdHex(id), bson.M{"$set": info})
	return
}

// SetUserName 设置用户名
func (m *UserModel) SetUserName(id, name string) (err error) {
	_, err = m.DB.UpsertId(bson.ObjectIdHex(id), bson.M{"$set": bson.M{"name": name}})
	return
}

// GetUserByID 根据ID查询用户
func (m *UserModel) GetUserByID(id string) (*Users, error) {
	user := new(Users)
	err := m.DB.FindId(id).One(&user)
	if err != nil {
		return nil, err
	}
	return user, nil
}

// GetUserByName
func (m *UserModel) GetUserByName(name string) (*Users, error) {
	user := new(Users)
	err := m.DB.Find(bson.M{}) (id).One(&user)
	if err != nil {
		return nil, err
	}
	return user, nil
}

// GetUserByEmail
func (m *UserModel) GetUserByEmail(email string) (*Users, error) {
	user := new(Users)
	err := m.DB.FindId(id).One(&user)
	if err != nil {
		return nil, err
	}
	return user, nil
}