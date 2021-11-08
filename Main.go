package main

import (
	"github.com/urfave/cli"
	"os"
	"ps/common"
)
var data map[string]interface{}

func initialize() {
	filePath := "./myPassword.csv"
	ifFileNotExistGenerate(filePath)
	data = common.File2Map(filePath)
}

func ifFileNotExistGenerate(path string) {
	if _,err := os.Stat(path); err != nil {
		if os.IsNotExist(err) {
			f,_ := os.Create(path)
			f.Write([]byte("{}"))
		}
	}
}

func main() {

	initialize()

	//实例化cli
	app := cli.NewApp()
	//Name可以设定应用的名字
	app.Name = "passwordHelper"
	// Version可以设定应用的版本号
	app.Version = "1.0.0"

	// Commands用于创建命令
	app.Commands = []cli.Command{
		// 显示数据库中存储的所有密码
		{
			// 命令的名字
			Name: "showAll",
			// 命令的缩写，就是不输入language只输入lang也可以调用命令
			Aliases: []string{"all"},
			// 命令的用法注释，这里会在输入 程序名 -help的时候显示命令的使用方法
			Usage:   "使用all或者showAll即可获取密码库中所有数据",
			// 命令的处理函数
			Action: getAllInfoAction,
		},
		{
			Name: "get",
			Aliases: []string{"g"},
			Usage: "get后跟平台，可以得到指定平台保存的所有用户名和密码\n get后跟平台 用户名，可以得到对应的密码",
			Action: getAction,
		},
		{
			Name: "add",
			Aliases: []string{"a"},
			Usage: "增加密码。需传三个参数（平台 用户名 密码）",
			Action: addPasswordAction,
		},
		{
			Name: "set",
			Aliases: []string{"s"},
			Usage: "修改密码。需传三个参数（平台 用户名 密码）",
			Action: setPasswordAction,
		},
	}
	// 接受os.Args启动程序
	app.Run(os.Args)
}

