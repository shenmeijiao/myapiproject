package controllers

import (
	"encoding/json"
	"fmt"
	"myapiproject/lib"
	"strconv"

	"github.com/astaxie/beego"
)

type DefaultController struct {
	beego.Controller
}

func (c *DefaultController) Outputs(codes int, v interface{}, msg string) {
	result := make(map[string]interface{})
	result["code"] = strconv.Itoa(codes)
	if v == nil {
		result["data"] = struct{}{}
	} else {
		result["data"] = v
	}

	result["msg"] = msg
	// pc1, _, _, _ := runtime.Caller(1)
	lib.FL.Info(fmt.Sprintf("OUTEPUT:%s,%s,%+v", result["code"], result["msg"], result["data"]))
	resultJson, _ := json.Marshal(result)
	lib.FL.Info(fmt.Sprintf("OUTEPUT:%s", string(resultJson)))
	c.Data["json"] = result

	c.ServeJSON()
	c.StopRun()

}
