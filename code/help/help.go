package help

import "fmt"

func OutPut(){
	fmt.Println(`
Option:
  --web	start,stop,install,remove,deploy
  --db	start,stop,install,remove
  --php	start,stop,install,remove
  --dir	create

Parameter:
  start 用于启动各个服务

  stop 用于停止各个服务

  remove 用于删除各个服务

  install 用于安装各个服务

  deploy 用于nginx 安装好之后,修改配置文件使用

  create 用于第一次初始化目录以及包文件使用

Example:
  ./main  Option  Parameter`)

}