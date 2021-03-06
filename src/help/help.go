package help

import "fmt"

func OutPut(){   //定义一个函数 输出帮助信息 内容
	fmt.Println(`
Option:
  --web	start,stop,install,remove,deploy,help
  --db	start,stop,install,remove,help
  --php	start,stop,install,remove,help
  --dir	create,delete

Parameter:
  start 用于启动各个服务

  stop 用于停止各个服务

  remove 用于删除各个服务

  install 用于安装各个服务

  deploy 用于nginx 安装好之后,修改配置文件使用

  create 用于第一次初始化目录以及包文件使用

  delete 用于删除使用过的工作目录

  help   显示帮助信息

Example:
  ./main  Option  Parameter`)

}