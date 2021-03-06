package controllers

import (
	"PrometheusAlert/models"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"strconv"
	"time"
)

//取到tpl路径
//fmt.Println(filepath.Join(beego.AppPath,"tpl"))

type MainController struct {
	beego.Controller
}
//main page
func (c *MainController) Get() {
	c.Data["IsIndex"]=true
	c.TplName = "index.html"
}
//test page
func (c *MainController) Test() {
	c.Data["IsTest"]=true
	c.TplName = "test.html"
}
//template page
func (c *MainController) Template() {
	c.Data["IsTemplate"]=true
	c.TplName = "template.html"
	Template, err := models.GetAllTpl()
	if err != nil {
		logs.Error(err)
	}
	c.Data["Template"] = Template
}
//template add
func (c *MainController) TemplateAdd() {
	c.Data["IsTemplate"]=true
	c.TplName = "template_add.html"
}
func (c *MainController) AddTpl()  {
	//获取表单信息
	tid:=c.Input().Get("id")
	name:=c.Input().Get("name")
	t_tpye:=c.Input().Get("type")
	t_use:=c.Input().Get("use")
	content:=c.Input().Get("content")
	if len(tid)==0 {
		id,_:=strconv.Atoi(tid)
		models.AddTpl(id,name,t_tpye,t_use,content)
	}else {
		id,_:=strconv.Atoi(tid)
		models.UpdateTpl(id,name,t_tpye,t_use,content)
	}
	c.Redirect("/template",302)
}
func (c *MainController) TemplateEdit() {
	c.Data["IsTemplate"]=true
	c.TplName = "template_edit.html"
	s_id,_:=strconv.Atoi(c.Input().Get("id"))
	Template, err := models.GetTpl(s_id)
	if err != nil {
		logs.Error(err)
	}
	c.Data["Template"] = Template
}
func (c *MainController) TemplateTest() {
	c.Data["IsTemplate"]=true
	c.TplName = "template_test.html"
	s_id,_:=strconv.Atoi(c.Input().Get("id"))
	Template, err := models.GetTpl(s_id)
	if err != nil {
		logs.Error(err)
	}
	c.Data["Template"] = Template
}
func (c *MainController) TemplateDel() {
	s_id,_:=strconv.Atoi(c.Input().Get("id"))
	err := models.DelTpl(s_id)
	if err != nil {
		logs.Error(err)
	}
	c.Redirect("/template", 302)
}

func LogsSign()string  {
	return strconv.FormatInt(time.Now().UnixNano(),10)
}

func (c *MainController)AlertTest()  {
	MessageData:=c.Input().Get("mtype")
	logsign:="["+LogsSign()+"]"
	switch MessageData {
	case "wx":
		wxtext:="[PrometheusAlert](https://github.com/feiyu563/PrometheusAlert)\n>**测试告警**\n>`告警级别:`测试\n**PrometheusAlert**"
		ret:=PostToWeiXin(wxtext,beego.AppConfig.String("wxurl"),logsign)
		c.Data["json"]=ret
	case "dd":
		ddtext:="## [PrometheusAlert](https://github.com/feiyu563/PrometheusAlert)\n\n"+"#### 测试告警\n\n"+"###### 告警级别：测试\n\n##### PrometheusAlert\n\n"+"![PrometheusAlert]("+beego.AppConfig.String("logourl")+")"
	    ret:=PostToDingDing("PrometheusAlert",ddtext,beego.AppConfig.String("ddurl"),logsign)
		c.Data["json"]=ret
	case "fs":
		fstext:="[PrometheusAlert](https://github.com/feiyu563/PrometheusAlert)\n\n"+"测试告警\n\n"+"告警级别：测试\n\nPrometheusAlert\n\n"+"![PrometheusAlert]("+beego.AppConfig.String("logourl")+")"
		ret:=PostToFeiShu("PrometheusAlert",fstext,beego.AppConfig.String("fsurl"),logsign)
		c.Data["json"]=ret
	case "txdx":
		MobileMessage:="PrometheusAlertCenter测试告警"
		ret:=PostTXmessage(MobileMessage,beego.AppConfig.String("defaultphone"),logsign)
		c.Data["json"]=ret
	case "txdh":
		ret:=PostTXphonecall("PrometheusAlertCenter测试告警",beego.AppConfig.String("defaultphone"),logsign)
		c.Data["json"]=ret
	case "hwdx":
		MobileMessage:="PrometheusAlertCenter测试告警"
		ret:=PostHWmessage(MobileMessage,beego.AppConfig.String("defaultphone"),logsign)
		c.Data["json"]=ret
	case "alydx":
		MobileMessage:="PrometheusAlertCenter测试告警"
		ret:=PostALYmessage(MobileMessage,beego.AppConfig.String("defaultphone"),logsign)
		c.Data["json"]=ret
	case "alydh":
		MobileMessage:="PrometheusAlertCenter测试告警"
		ret:=PostALYphonecall(MobileMessage,beego.AppConfig.String("defaultphone"),logsign)
		c.Data["json"]=ret
	case "rlydh":
		MobileMessage:="PrometheusAlertCenter测试告警"
		ret:=PostRLYphonecall(MobileMessage,beego.AppConfig.String("defaultphone"),logsign)
		c.Data["json"]=ret
	default:
		c.Data["json"]="hahaha!"
	}
	c.ServeJSON()
}