package controllers

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/zhanglianx111/c2o/models"
	"github.com/zhanglianx111/kompose/pkg/app"
	"os"
)

// Operations about docker-compose file
type ConvertController struct {
	beego.Controller
}

// @Title convert docker-compose file
// @Description convert a docker compose file
// @Param body body models.File true "docker compose file content"
// @Success 200 {string} models.File.Name
// @router / [post]
func (c *ConvertController) Post() {
	// body is writed to a file
	var files []string
	file := models.WriteToFile(c.Ctx.Input.RequestBody)
	if len(file) == 0 {
		err := errors.New("write to file fail")
		panic(err)
	}
	files = append(files, file)
	obj := models.InitConvertOpt(files)
	fmt.Printf("kobject: %v\n", obj)
	// convert cmd
	data, err := app.Convert(obj)
	if err != nil {
		fmt.Println(err)
	}
	if err := os.Remove(files[0]); err != nil {
		fmt.Println(err)
	}
	// do convertion

	// return result of convertion

	//c.Data["json"] = map[string]string{"Name": "name"}
	var v interface{}
	err = json.Unmarshal(data, &v)
	if err != nil {
		fmt.Println(err)
	}
	c.Data["json"] = v
	c.ServeJSON()
}
