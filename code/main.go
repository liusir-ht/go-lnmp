package main

import (
	"flag"
	"fmt"
	"go-tools/Mysql"
	"go-tools/directory"
	"go-tools/help"
	"go-tools/nginx"
	"go-tools/php"
	"os/user"
)


func main() {
	helpflag:=flag.String("help","not","帮助参数")
	ngxflag:=flag.String("web","","nginx")  //配置nginx选项的默认参数
	dir:=flag.String("dir","","workdir")//配置dir选项的默认参数
	dbflag:=flag.String("db","","mysql") //配置db选项的默认参数
	phpflag:=flag.String("php","","php") //配置php选项的默认参数
	flag.Parse() //解析参数
	//用户模块
	u,_:=user.Current()
	fmt.Printf("当前执行程序用户：%v\n",u.Username)
	n:=&nginx.Ngx{}
	var  web nginx.Ngxroom
	web=n  //给接口赋值
	var  ppp php.Phproom
	phpinit:=&php.Php{}
	ppp=phpinit
	//帮助区域
	if *helpflag=="" {
		help.OutPut()  //输出帮助内容
	}
	//初始化区域
	if *dir=="create" {
		directory.CreateDir()
	}

	//Nginx  区域
	switch *ngxflag {  //使用switch 进行多条件判断
	case "install":
		web.Check()
		web.Install()
	case "start":
		web.Start()
	case "stop":
		web.Stop()
	case "remove":
		web.Stop()
		web.Remove()
	case "deploy":
		web.Deploy()
		ppp.Stop()
		ppp.Start()
	case "help":
		help.OutPut()  //输出帮助内容
	}
	//MySQL 区域
	var  db  Mysql.Mysqlroom
	dbinit:=&Mysql.Mysql{}
	db=dbinit //给接口赋值
	switch *dbflag{   //使用switch 进行多条件判断
	case "install":
		db.Check()
		db.Install()
	case "start":
		db.Start()
	case "stop":
		db.Stop()
	case "remove":
		db.Stop()
		db.Remove()
	case "help":
		help.OutPut()  //输出帮助内容
	}

    //PHP区域
	switch *phpflag{ //使用switch 进行多条件判断
	case "install":
		ppp.Check()
		ppp.Install()
	case "start":
		ppp.Start()
	case "stop":
		ppp.Stop()
	case "remove":
		ppp.Remove()
	case "help":
		help.OutPut()  //输出帮助内容
	}
}
