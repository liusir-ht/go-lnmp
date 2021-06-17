package Mysql

import (
	"fmt"
	"os"
	"os/exec"
)

var (
	workdir     ="/home/github.com/go-lnmp/mysql" //MySQL的包的目录
	Mysqlrpmc  = "MySQL-client-5.6.40-1.el7.x86_64"                            //mysql-client的包名称
	Mysqlrpms  ="MySQL-server-5.6.40-1.el7.x86_64"                           //mysql-server的包名称
    VersionChan =make(chan string,100)                   //传输MySQL 版本的channel
    //version string //定义一个接收channel数据的变量
	cmd,cc3,cc4,cc5 *exec.Cmd    //
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
	c3:=exec.Command("yum","remove","-y","mysql-community-"+"*")
	err03:=c3.Run()
	if err03 != nil {
		fmt.Printf("删除 Mysql 5.7 依赖 失败\n")
	}else {
		fmt.Printf("删除 Mysql  5.7 依赖成功\n")
	}
	cmd=exec.Command("rpm","-qa","MySQL-server")
	d1,err:=cmd.CombinedOutput()
	if err ==nil{
		if  string(d1) == ""{
			fmt.Printf("没有MySQL-server，无残留 \n")
		}else {
			fmt.Printf("MySQL-server 有残留 %v\n",string(d1))
			c3=exec.Command("rpm","-e","MySQL-server")
			out03,err03:=c3.CombinedOutput()
			if err03 != nil {
				fmt.Printf("删除 MySQL-server 旧包失败 %v\n",string(out03))
				return
			}else {
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
			c3=exec.Command("rpm","-e","MySQL-client")
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
func (m *Mysql) Install() {
	version := <-VersionChan //接收channel data
	switch version {         //判断channel data
	case "5.6":
		cmd = exec.Command("rpm", "-ivh", "MySQL-*"+".rpm") //安装 rpm包
		cc4 = exec.Command("cat", "/root/.mysql_secret")  //查看临时密码
		cc5=exec.Command("systemctl","start","mysql")
	case "5.7":
		cmd = exec.Command("rpm", "-ivh", "mysql-community-*"+".rpm")  //安装 rpm包
		cc3= exec.Command("systemctl","start","mysqld")  //查看临时密码
		cc4 = exec.Command("cat", "/var/log/mysqld.log")
		cc5 = exec.Command("grep", "temporary password")  //查看临时密码
		cc5.Stdin, _ = cc4.StdoutPipe()  //把cc4这个的结果 通过 cc5命令再次操作
		cc5.Stdout = os.Stdout   //输出到系统输出
	}
	cmd.Dir = workdir + "/"  //指定工作路径
	d1, err := cmd.CombinedOutput()   //执行 安装rpm的步骤
	if err != nil {
		fmt.Printf("Rpm Mysql 安装失败 err:%v\n%v\n", err, string(d1))
		return
	} else {
		fmt.Printf("Rpm Mysql 安装成功 %v\n", string(d1))
		c3 := exec.Command("find", "/", "-name", "mysql")  //搜索mysql的目录
		out03, _ := c3.CombinedOutput()
		fmt.Printf("MySQL dir:\n%v\n", string(out03))
		switch version {    //根据version数值的不同输出不同的结果
		case "5.6":
			out,errcc4:=cc4.CombinedOutput()
			if errcc4 !=nil {
				fmt.Printf("MySQL Password Inital err：%v\n:", errcc4)
				fmt.Printf("MySQL Password Inital err：%v\n:", string(out))
				return
			}else {
				fmt.Printf("MySQL Password Inital success：%v\n:", string(out))
				_=cc5.Run()
			}
		case "5.7":
			fmt.Println("MySQL Password Inital :")
			_=cc3.Run()
			_ = cc5.Start()
			_ = cc4.Run()
			_ = cc5.Wait()
		}
	}
}


// Start 启动mysql
func (m *Mysql) Start(){
	version:=<- VersionChan   //接收channel data
	switch version {
	case "5.6":
		cmd = exec.Command("systemctl", "start", "mysql")
	case "5.7":
		cmd = exec.Command("systemctl", "start", "mysqld")
	}
	d1,err:=cmd.CombinedOutput()
	if err != nil {
		fmt.Printf("mysql 启动失败 %v\n%v\n",err,string(d1))
			return
	}else {
		fmt.Printf("mysql 启动成功 %v\n",string(d1))
		c2:=exec.Command("mysql","-V")  //查看mysql版本
		d2,err02:=c2.CombinedOutput()
		if err02 != nil {
			fmt.Printf("查看mysql 版本 err:%v\n%v\n",err02,string(d2))
			return
		}else {
			fmt.Printf("查看mysql版本，成功 \n %v\n",string(d2))
			}
		}
	}


// Stop 停止mysql
func (m *Mysql) Stop() {
	version:=<- VersionChan//接收channel data
	switch version {
	case "5.6":
		cmd = exec.Command("systemctl", "stop", "mysql")
	case "5.7":
		cmd = exec.Command("systemctl", "stop", "mysqld")
	}
	d1, err := cmd.CombinedOutput()
	if err != nil {
			fmt.Printf("mysql 停止失败 %v\n %v\n", err, string(d1))
			return
	} else {
			fmt.Printf("mysql 停止成功 %v\n", string(d1))
		}
}

// Remove 删除mysql
func (m *Mysql) Remove(){
	version:=<- VersionChan //接收channel data
	switch version {  //判断channel的data
	case "5.6":  //当 data = 5.6
		c1:=exec.Command("systemctl","stop","mysql")    //停止 mysql服务
		cmd = exec.Command("rpm", "-e", Mysqlrpmc)  //删除mysql client  rpm包
		c2:=exec.Command("rpm","-e",Mysqlrpms)   //删除 mysql server  rpm 包
		_=c1.Run()   // 执行停止服务
		_=c2.Run()   //执行删除mysql server  rpm包
		//删除残留目录和文件
		_=os.RemoveAll("/usr/lib64/mysql")
		_=os.RemoveAll("/var/lib/mysql")
		_=os.Remove("/root/.mysql_secret")


	case "5.7": //当 data = 5.7
		cmd = exec.Command("yum", "remove", "-y", "mysql-community-*")  //删除 mysql 5.7 的rpm包
		//删除残留目录和文件
		removeerr:=os.RemoveAll("/var/lib/mysql")
		if removeerr == nil {
			fmt.Println("remove /var/lib/mysql success")
		}
		removeerr02:=os.RemoveAll("/usr/share/mysql")
		if removeerr02 == nil {
			fmt.Println("remove /usr/share/mysql success")
		}
		_=os.Remove("/var/log/mysqld.log")
	}
	d1,err:=cmd.CombinedOutput()
	if err != nil {
		fmt.Printf("mysql 删除失败 %v\n%v\n",err,string(d1))
		return
	}else {
		fmt.Printf("mysql 删除成功 %v\n",string(d1))
		}
	}


