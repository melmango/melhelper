package helpers

import (
	"github.com/astaxie/beego/logs"
)

var LOG *logs.BeeLogger

func init() {
	LOG =logs.NewLogger(10000)
	LOG.EnableFuncCallDepth(true)
	LOG.SetLogFuncCallDepth(3)
	LOG.SetLogger("console","")
}

