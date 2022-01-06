package lib

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
)

var FL *logs.BeeLogger

// var XLH string

const (
	LOGS_DOTYPE_LOGIN = "登录"
	// LOGS_DOTYPE_JOBS_ADD = "增加任务"
	// LOGS_DOTYPE_JOBS_UP = "修改任务"
	// LOGS_DOTYPE_VIDEOS_DEL = "编辑视频"
)

func init() {
	FL = logs.NewLogger()
	FL.EnableFuncCallDepth(true)
	FL.SetLogger(logs.AdapterFile, `{"filename":"`+beego.AppConfig.String("logPathFile")+".log"+`","level":7,"daily":false,"maxdays":10,"color":false}`)
}
