# 这个程序的初衷
  为了让运维人员,脱离重复而繁琐的部署,提高运维效率,节省时间,此程序因此而生
  
  
  
## 正式开始
>环境介绍：

> Nginx 1.10/1.12/1.14/1.16/1.18/1.20

> MySQL 5.6.40/5.7.33

> Php   5.5.38/7.2.33

> 使用到的目录： /home/github.com/

> GitHub地址： https://github.com/liusir-ht/go-lnmp.git

> GitHub国内地址： https://github.com.cnpmjs.org/liusir-ht/go-lnmp.git
## 选项参数说明
 选项名称     |  参数名称 |  
 ---     |  --- |
 web   | start,stop,install,remove,deploy,help
 db    | start,stop,install,remove,help
 php   | start,stop,install,remove,help
 dir   | create 
 version | nginx:1.16/1.18/1.20 mysql:5.6.40/5.7.33 php:5.5.38

###选项

  `web`   Nginx operation
  
  `db`    MySQL operation
  
  `dir`   First init,directory/pkg
  
  `php`   Php operation
  
  `version` 指定软件的version
  
###  参数 

  `start` 参数用于启动各个服务
  
  `stop`  参数用于停止各个服务 
  
  `remove` 参数用于删除各个服务
  
  `install` 参数用于安装各个服务
  
  `deploy` 参数用于nginx 安装好之后,修改配置文件使用
  
  `create` 参数用于第一次初始化目录以及包文件使用
  
  `help`   参数用于显示命令帮助信息
  
  
  
  
## Example

####  Dir

```
./main  --dir create
```
进行初始化


####  Nginx
  
  
```
./main  --web  start/stop/install/remove/deploy/help
```


####  MySQL
  
  
```
./main  --db  start/stop/install/remove/help
```

####  Php
  
  
```
./main  --php  start/stop/install/remove/help
```

#### Version
需要注意的是 软件操作的时候必须指定 **版本号**
```
./main  --version  1.16  --web   install 
```