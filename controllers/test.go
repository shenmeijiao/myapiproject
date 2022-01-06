package controllers

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"myapiproject/lib"

	"github.com/astaxie/beego"
)

type TestController struct {
	DefaultController
	// beego.Controller
}
type LoginAdmField struct {
	LoginName string `json:"login_name"`
	Passwd    string `json:"passwd"`
	CId       string `json:"cid"`
	Code      string `json:"code"`
}

// @Title 加密
// @Description 获取AES加密串
// @Param	str		query	string	true	需要加密的字符串
// @Success 200 {"code":"200","data":{"newStr":""},"msg":"OK"}
// @router /getAESstr [get]
func (t *TestController) GetAESstr() {
	str := t.GetString("str")
	if str == "" {
		t.Outputs(400, nil, "参数不能为空")
	} else {
		key := beego.AppConfig.String("aesKey")
		newStr := lib.EncryptAES([]byte(str), []byte(key))
		// t.Data["json"] = newStr
		// t.ServeJSON()
		t.Outputs(200, map[string]string{
			"newStr": base64.StdEncoding.EncodeToString(newStr),
		}, "OK")
	}
}

// @Title 解密
// @Description Logs out current logged in user session
// @Param	login_name		body string true "登陆名"
// @Param	passwd		body string true "aes加密串"
// @Param	cid			body string true "验证码id"
// @Param   code		body string true "验证码"
// @Success 200 {"code":"200","data":{"pass":""},"msg":"操作成功！"}
// @router /getpwd [post]
func (t *TestController) Getpwd() {
	// str := "KdcYgUzAnz+L24rj5SKtVg=="
	var v LoginAdmField
	if err := json.Unmarshal(t.Ctx.Input.RequestBody, &v); err == nil {
		passwdBytes, _ := base64.StdEncoding.DecodeString(v.Passwd)
		fmt.Printf("pass=%s\n", v.Passwd)
		key := beego.AppConfig.String("aesKey")
		newStr := lib.DecryptAES(passwdBytes, []byte(key))
		// fmt.Println(newStr)
		fmt.Printf("解密后: %s\n", string(newStr))
		t.Outputs(200, map[string]string{"pass": string(newStr)}, "操作成功！")
		// main()
	} else {
		t.Outputs(400, nil, "参数错误！")
	}
}
func main() {
	x := []byte("ZRZXkjfz703")
	key := []byte("hgfedcbahfewyr82")
	x1 := lib.EncryptAES(x, key)
	fmt.Println(x1)
	fmt.Printf("加密后: %s\n", base64.StdEncoding.EncodeToString(x1))
	x2 := lib.DecryptAES(x1, key)
	fmt.Printf("解密后明文: %s\n", x2)
}
