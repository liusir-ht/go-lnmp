package Mysql

import (
	"fmt"
	"os/exec"
)

const (
	mysqlrpmc="MySQL-client-5.6.40-1.el7.x86_64"  //mysql-client的包名称
	mysqlrpms="MySQL-server-5.6.40-1.el7.x86_64"  //mysql-server的包名称

)
var (
	workdir ="/home/github.com/go-lnmp/mysql"  //MySQL的包的目录
)
type Mysqlroom interface {   //定义Mysql接口
	Check()
	Install()
	Start()
	Remove()
	Stop()
}
type  Mysql struct {

}
func (m *Mysql)  Check(){
		c2 := exec.Command("yum", "remove","-y","mariadb"+"*")
		out02, err02 := c2.CombinedOutput()
		if err02 != nil {
			fmt.Printf("删除mariadb libs 依赖失败 :%v\n", err02)
			fmt.Printf("命令结果错误：%v\n", string(out02))
			panic("删除mariadb libs 依赖失败")
			return
		} else {
			fmt.Printf("删除mariadb 依赖成功 :%v\n", string(out02))
		}
	cmd:=exec.Command("rpm","-qa","MySQL-server")
	d1,err:=cmd.CombinedOutput()
	if err ==nil{
		if  string(d1) == ""{
			fmt.Printf("没有MySQL-server，无残留 \n")
		}else {
			fmt.Printf("MySQL-server 有残留 %v\n",string(d1))
			c3:=exec.Command("rpm","-e","MySQL-server")
			out03,err03:=c3.CombinedOutput()
			if err03 != nil {
				fmt.Printf("删除 MySQL-server 旧包失败 %v\n",err03 )
				fmt.Printf("删除 MySQL-server 旧包失败 %v\n",string(out03))
				return
			}else {
				fmt.Printf("删除 MySQL-server 旧包成功\n")
				fmt.Printf("删除 MySQL-server 旧包成功 %v\n",string(out03))
			}
		}
	}
	d3:=exec.Command("rpm","-qa","MySQL-client")
	out03,err03:=d3.CombinedOutput()
	if err03 ==nil{
		if  string(out03) == ""{
			fmt.Printf("没有MySQL-client，无残留 \n")
		}else {
			fmt.Printf("MySQL-client 有残留 %v\n",string(out03))
			//os.Setenv("MySQL-client",string(out03))
			c3:=exec.Command("rpm","-e","MySQL-client")
			out04,err04:=c3.CombinedOutput()
			if err04 != nil {
				fmt.Printf("删除 MySQL-client 旧包失败 %v\n",err04 )
				fmt.Printf("删除 MySQL-client 旧包失败 %v\n",string(out04))
				return
			}else {
				fmt.Printf("删除 MySQL-client 旧包成功\n")
				fmt.Printf("删除 MySQL-client 旧包成功 %v\n",string(out04))
			}
		}
	}
}

// Install 安装mysql
func (m *Mysql) Install(){
	cmd:=exec.Command("rpm","-ivh","*"+".rpm")
	cmd.Dir=workdir
	d1,err:=cmd.CombinedOutput()
	if err !=nil{
		fmt.Printf("Rpm Mysql 安装失败 err:%v\n%v\n",err,string(d1))
		return
	}else {
		fmt.Printf("Rpm Mysql 安装成功 %v\n",string(d1))
		c3:=exec.Command("find","/","-name","mysql")
		out03,_:=c3.CombinedOutput()
		fmt.Printf("MySQL dir:\n%v\n",string(out03))
		c4:=exec.Command("cat","/root/.mysql_secret")
		out04,_:=c4.CombinedOutput()
		fmt.Printf("MySQL Password Inital :\n%v\n",string(out04))
	}
}

// Start 启动mysql
func (m *Mysql) Start(){
	cmd:=exec.Command("systemctl","start","mysql")
	d1,err:=cmd.CombinedOutput()
	if err != nil {
		fmt.Printf("mysql 启动失败 %v\n%v\n",err,string(d1))
		return
	}else {
		fmt.Printf("mysql 启动成功 %v\n",string(d1))
		c2:=exec.Command("mysql","-V")
		d2,err02:=c2.CombinedOutput()
		if err02 != nil {
			fmt.Printf("验证rpm mysql包失败 err:%v\n%v\n",err02,string(d2))
			return
		}else {
			fmt.Printf("验证rpm mysql包成功 \n %v\n",string(d2))
		}
	}
}

// Stop 停止mysql
func (m *Mysql) Stop(){
	cmd:=exec.Command("systemctl","stop","mysql")
	d1,err:=cmd.CombinedOutput()
	if err != nil {
		fmt.Printf("mysql 停止失败 %v\n %v\n",err,string(d1))
		return
	}else {
		fmt.Printf("mysql 停止成功 %v\n",string(d1))
	}

}

// Remove 删除mysql
func (m *Mysql) Remove(){
	cmd:=exec.Command("rpm","-e",mysqlrpmc)
	cmd.Dir=workdir
	d1,err:=cmd.CombinedOutput()
	if err != nil {
		fmt.Printf("mysql client 删除失败 %v\n%v\n",err,string(d1))
		return
	}else {
		fmt.Printf("mysql client 删除成功 %v\n",string(d1))
		c2:=exec.Command("rpm","-e",mysqlrpms)
		c2.Dir=workdir
		err02:=c2.Run()
		if err02 != nil {
			fmt.Printf("mysql server 删除失败 %v\n",err02)
			return
		}else {
			fmt.Printf("mysql server 删除成功\n")
		}
	}
}
//备份数据库

