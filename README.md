# 这个程序的初衷
  为了让运维人员,脱离重复而繁琐的部署,提高运维效率,节省时间,此程序因此而生
  
  
  
## 正式开始
>环境介绍：

> Nginx 1.20.X

> MySQL 5.6.X

> Php   5.5.X

> 使用到的目录： /home/github.com/

> GitHub地址： https://github.com/liusir-ht/go-lnmp.git

> GitHub国内地址： https://github.com.cnpmjs.org/liusir-ht/go-lnmp.git
## 选项参数说明
 选项名称     |  参数名称 |  
 ---     |  --- |
 web   | start,stop,install,remove,deploy
 db    | start,stop,install,remove
 php   | start,stop,install,remove
 dir   | create 
 
 
  `start` 参数用于启动各个服务
  
  `stop`  参数用于停止各个服务 
  
  `remove` 参数用于删除各个服务
  
  `install` 参数用于安装各个服务
  
  `deploy` 参数用于nginx 安装好之后,修改配置文件使用
  
  `create` 参数用于第一次初始化目录以及包文件使用
 
  
## Example

####  Dir

```
./main  --dir create
```
进行初始化


####  Nginx
  
  
```
./main  --web  start/stop/install/remove/deploy
```


####  MySQL
  
  
```
./main  --db  start/stop/install/remove
```

####  Php
  
  
```
./main  --php  start/stop/install/remove
```