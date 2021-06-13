package php

import (
	"fmt"
	"os"
	"os/exec"
)

const (
	Phppackage7 = "php-7.2.33"
	Phppackage5= "php-5.5.38"
)
var (
	workdir ="/home/github.com/go-lnmp/php"  //Php的包的目录
	compiledir7 =workdir+"/php-7.2.33"  //Php编译的目录
	compiledir5 =workdir+"/php-5.5.38"  //Php编译的目录
	installdir="/usr/local/php"
)
type Phproom interface {   //定义Php接口
	Check()
	Install()
	Start()
	Stop()
	Remove()
}
type Php struct {

}
func (p *Php) Check()  {
	cmd:=exec.Command("yum","install","-y","libxml2-devel","libxml2",
		"gd ","zlib-devel","libjpeg" ,"libjpeg-devel","libpng-devel",
		"libXpm-devel"," xz-devel" )
	out,err:=cmd.CombinedOutput()
	if err != nil {
	    fmt.Printf("Php Check err:%v\n",err)
	    fmt.Printf("error:%v\n",string(out))
	} else {
		fmt.Printf("Php Check success:%v\n",string(out))
	}
}
func (p *Php) Install(){
	//解压
	cmd:=exec.Command("tar","-zxf",Phppackage5+".tar.gz")
	cmd.Dir=workdir
	out,err:=cmd.CombinedOutput()
	if err != nil {
		fmt.Printf("Php tar err:%v\n",err)
		fmt.Printf("error:%v\n",string(out))
		return
	} else {
		fmt.Printf("Php tar success:%v\n",string(out))
	}
	//切换目录
	c2:=exec.Command("cd",Phppackage5)
	c2.Dir=workdir
	err02:=c2.Run()
	if err02 != nil {
		fmt.Printf("cd  err:%v\n",err02)
		return
	}
	//编译安装
	c3:=exec.Command("./configure","--prefix=/usr/local/php",
		"--with-gd","--with-zlib","--with-mysql=mysqlnd",
		"--with-mysqli=mysqlnd", "--with-config-file-path=/usr/local/php",
		 "--enable-fpm" , "--enable-mbstring","--with-jpeg-dir=/usr/lib")
	c3.Dir=compiledir5
	fmt.Printf("当前工作目录：%v\n 请耐心等待.......\n目录编译开始：... \n",c3.Dir)
	out03,err03:=c3.CombinedOutput()
	if err03 != nil {
		fmt.Printf("compile PHP  err:%v\n",err03)
		fmt.Printf("compile PHP  err:%v\n",string(out03))
		return
	}else {
		fmt.Printf("compile PHP  success :%v\n",string(out03))
		c4:=exec.Command("make")
		c4.Dir=compiledir5
		fmt.Printf("当前工作目录：%v\n 请耐心等待.......\nmake 开始\n",c4.Dir)
		out04,err04:=c4.CombinedOutput()
		if err04 != nil {
			fmt.Printf("make  PHP  err:%v\n", err04)
			fmt.Printf("make  PHP  err:%v\n",string(out04))
			return
		} else {
			fmt.Printf("make  PHP  success:%v\n",string(out04))

		}

	}
	//make install
	c5:=exec.Command("make","install")
	c5.Dir=compiledir5
	fmt.Printf("当前工作目录：%v\n 请耐心等待.......\nmake install 开始\n",c5.Dir)
	out05,err05:=c5.CombinedOutput()
	if err05 != nil {
		fmt.Printf("make  install PHP  err:%v\n", err05)
		fmt.Printf("make install PHP  err:%v\n",string(out05))
		return
	} else {
		fmt.Printf("make  install PHP  success:%v\n",string(out05))
	}
	//copy ini
	c8:=exec.Command("cp",workdir+"/"+"php.ini","/usr/local/php/php.ini")
	out08,err08:=c8.CombinedOutput()
	if err08 != nil {
		fmt.Printf("copy php.ini err :%v\n",err08)
		fmt.Printf("copy php.ini err :%v\n",string(out08))
		return
	} else {
		fmt.Printf("copy php.ini success :%v\n",string(out08))
	}
	//copy php-fpm.conf
	c9:=exec.Command("cp",workdir+"/"+"php-fpm.conf","/usr/local/php/etc/")
	out09,err09:=c9.CombinedOutput()
	if err09 != nil {
		fmt.Printf("copy php-fpm.conf err :%v\n",err09)
		fmt.Printf("copy php-fpm.conf err :%v\n",string(out09))
		return
	} else {
		fmt.Printf("copy php-fpm.conf success :%v\n",string(out09))
	}
	//copy php-fpm 命令
	c10:=exec.Command("cp","-f",workdir+"/php-5.5.38/sapi/fpm/init.d.php-fpm",
		"/etc/rc.d/init.d/php-fpm")
	out10,err10:=c10.CombinedOutput()
	if err10 != nil {
		fmt.Printf("copy php-fpm  命令 err :%v\n",err10)
		fmt.Printf("copy php-fpm  命令 err :%v\n",string(out10))
		return
	} else {
		fmt.Printf("copy php-fpm 命令 success :%v\n",string(out10))
	}
	errchmod:=os.Chmod("/etc/rc.d/init.d/php-fpm",777)
	if errchmod != nil {
		fmt.Printf("赋予权限失败 err:%v\n",errchmod)
		return
	} else {
		fmt.Printf("赋予权限成功 err:%v\n",errchmod)
	}
	//把php-fpm添加到系统命令
	c11:=exec.Command("chkconfig","--add","/etc/init.d/php-fpm")
	out11,err11:=c11.CombinedOutput()
	if err11 != nil {
		fmt.Printf("添加系统命令  命令 err :%v\n",err11)
		fmt.Printf("添加系统命令   命令 err :%v\n",string(out11))
		return
	} else {
		fmt.Printf("添加系统命令  命令 success :%v\n",string(out11))
		c12:=exec.Command("find","/","-name","php")
		out12,_:=c12.CombinedOutput()
		fmt.Printf("Php dir:\n%v\n",string(out12))
	}
	c4:=exec.Command("cp",workdir+"/index.php","/usr/share/nginx/html/")
	out04,_:=c4.CombinedOutput()
	fmt.Printf("cp Index.PHP :%v\n",string(out04))
/*	//create user php
	c10:=exec.Command("useradd","-M","-s","/sbin/nologin","php")
	out10,err10:=c10.CombinedOutput()
	if err10 != nil {
		fmt.Printf("create user Php err :%v\n",err10)
		fmt.Printf("create user Php  err :%v\n",string(out10))
		return
	} else {
		fmt.Printf("create user Php  :%v\n",string(out10))
	}*/

}
func (p *Php) Start(){
	cmd:=exec.Command("systemctl","start","php-fpm")
	out,err:=cmd.CombinedOutput()
	if err !=nil{
		fmt.Printf("Php 启动失败 :%v\n",err)
		fmt.Printf("Php 启动失败 :%v\n",string(out))
		return
	} else {
		fmt.Printf("Php 启动成功 :%v\n",string(out))
		c2:=exec.Command("netstat","-anput")   //第一个命令
		c3:=exec.Command("grep","9001") //第二个命令
		c3.Stdin,_=c2.StdoutPipe()  //把第一个命令的结果 输出给 第二个命令
		c3.Stdout=os.Stdout //把系统输出 给 第二个命令
		_=c3.Start() //执行第二个命令
		_=c2.Run()  //执行第一个命令
		_=c3.Wait()  //等待第二个命令执行完毕
	}
}
func (p *Php) Stop(){
	c1:=exec.Command("systemctl","stop","php-fpm")
	out,err:=c1.CombinedOutput()
	if err !=nil{
		fmt.Printf("Php 停止失败 :%v\n",err)
		fmt.Printf("Php 停止失败 :%v\n",string(out))
		return
	} else {
		fmt.Printf("Php 停止成功 :%v\n", string(out))
	}
}
func (p *Php) Remove(){
	err:=os.RemoveAll("/usr/local/php")
	if err != nil {
		fmt.Printf("删除 /usr/local/php 目录失败:%v\n",err)
		return
	}else {
		fmt.Printf("删除 /usr/local/php 目录成功:%v\n",err)
	}

	err02:=os.Remove("/etc/rc.d/init.d/php-fpm")
	if err02 != nil {
		fmt.Printf("清除系统命令 失败:%v\n",err02)
		return
	}else {
		fmt.Printf("清除系统命令 成功:%v\n",err02)
	}
}
