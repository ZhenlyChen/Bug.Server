package controllers

import (
	"github.com/ZhenlyChen/BugServer/httpServer/models"
	"github.com/ZhenlyChen/BugServer/httpServer/services"
	"github.com/kataras/iris"
	"github.com/kataras/iris/sessions"
)

// GameController 用户控制
type GameController struct {
	Ctx     iris.Context
	Service services.GameService
	Session *sessions.Session
}

// GameRes 游戏版本内容
type GameRes struct {
	Status string      `json:"status"`
	Data   models.Game `json:"data"`
}

// GetNew GET /game/new 获取最新版本号
func (c GameController) GetNew() (res GameRes) {
	res.Status = StatusSuccess
	res.Data = c.Service.GetNewestVersion()
	return
}
