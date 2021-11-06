package main

import (
	"encoding/csv"
	"fmt"
	"github.com/urfave/cli"
	"io"
	"io/ioutil"
	"log"
	"os"
	"ps/common"
)

var getAllInfoAction = func (c *cli.Context) error{
	csvFile, err := os.Open("./myPassword.csv")
	common.TestError(err, "can not open the file, err is %+v")
	defer csvFile.Close()
	reade := csv.NewReader(csvFile)
	for {
		row, err := reade.Read()
		if err != nil && err != io.EOF {
			log.Fatalf("can not read, err is %+v", err)
			return err
		}
		if err == io.EOF {
			break
		}
		fmt.Printf("platform=%s, username=%s, password=%s\n", row[0], row[1], row[2])
	}
	return nil
}

var getAction = func(c *cli.Context) error {
	var1 := c.Args().Get(0)
	var2 := c.Args().Get(1)
	platForms := data
	if var1 == "" && var2 == "" {
		println("啥也不获取，你想干啥？")
	} else if var2 == "" {
		println("获取指定平台下所有账号密码")
		m, exist := platForms[var1].(map[string]interface{})
		if exist == false {
			fmt.Println("不存在指定平台")
			return nil
		}
		fmt.Printf("username\tpassword\n")
		for key, value := range m {
			fmt.Printf("%s\t%s\n",key,value)
		}
	} else {
		m := platForms[var1].(map[string]interface{})
		fmt.Println(m[var2])
	}
	return nil
}

var addPasswordAction = func(c *cli.Context) error {
	platform := c.Args().Get(0)
	username := c.Args().Get(1)
	password := c.Args().Get(2)
	usernames, exitPlatForm := data[platform].(map[string]interface{})
	if exitPlatForm == false {
		data[platform] = map[string]interface{}{username: password}
	} else {
		_, exitUsername := usernames[username]
		if exitUsername == false {
			usernames[username] = password
			data[platform] = usernames
		} else {
			fmt.Printf("平台用户名已存在，如要修改请使用set进行设置")
		}
	}
	// 将数据传入文件
 	dataJsonString := []byte(common.Map2jsonString(data))
	ioutil.WriteFile("myPassword.csv", dataJsonString, 0666)
	return nil
}

var setPasswordAction = func(c *cli.Context) error {
	platform := c.Args().Get(0)
	username := c.Args().Get(1)
	password := c.Args().Get(2)
	usernames, exitPlatForm := data[platform].(map[string]interface{})
	if exitPlatForm == false {
		fmt.Printf("找不到<%s>平台\n" ,platform)
	} else {
		_, exitUsername := usernames[username]
		if exitUsername == false {
			fmt.Printf("在<%s>平台找不到<%s>用户名\n" ,platform ,username)
		} else {
			usernames[username] = password
			data[platform] = usernames
		}
	}
	// 将数据传入文件
	dataJsonString := []byte(common.Map2jsonString(data))
	ioutil.WriteFile("myPassword.csv", dataJsonString, 0666)
	return nil

}