package nginx

import (
	"fmt"
	"os"
	"os/exec"
)

var (
	Ngxrpm string  //定义安装Nginx的包名*/
	workdir ="/home/github.com/go-lnmp/pkg/nginx"  //Nginx的包的目录
	confdir="/home/github.com/go-lnmp/conf"
	ngxconf ="/etc/nginx/conf.d/default.conf"
)

type  Ngxroom interface {   //定义nginx方法的集合
	Check()
	Install()
	Start()
	Remove()
	Stop()
	Reload()
	Deploy()
}
type Ngx struct {}  //定义一个Ngx的空结构体

// Check 检查nginx环境信息
func (n *Ngx)  Check(){
	cmd:=exec.Command("rpm","-qa","nginx")   //检查是否已经安装过nginx
	d1,err:=cmd.Output()
	if err ==nil{
		if  string(d1) == ""{  //如果没有则不操作
			fmt.Printf("Nginx 无残留\n")
		}else {  //如果有则输出已经存在
			fmt.Printf("nginx 有残留 %v\n",string(d1))
			fmt.Println(string(d1))
			c2:=exec.Command("rpm","-e",string(d1))   //删除已经存在的nginx包
			d2,err02:=c2.CombinedOutput()
			if err02 != nil {
				fmt.Printf("删除旧包失败 %v\n %v\n",err02,string(d2) )
				return
			}else {
				fmt.Printf("删除旧包成功 %v\n", string(d2))
			}
		}
	}
	return
}

// Install 安装nginx
func (n *Ngx) Install(){
	cmd:=exec.Command("yum","localinstall","-y",Ngxrpm+".rpm")  //安装nginx Rpm包
	cmd.Dir=workdir  //指定包的目录
	d1,err:=cmd.CombinedOutput()
	if err !=nil{
		fmt.Printf("Rpm Nginx 安装失败 err:%v\n %v\n",err,string(d1))
		return
	}else {
		fmt.Printf("Rpm Nginx 安装成功 %v\n",string(d1))
		n.Start()  //启动nginx
		c2:=exec.Command("nginx","-v")  //验证是否安装完毕
		d2,err02:=c2.Output()
		if err02 != nil {
			fmt.Printf("验证rpm Nginx 包失败 err:%v\n",err02)
			return
		}else {
			fmt.Printf("验证rpm Nginx 包成功 \n %v\n",string(d2))
			c3:=exec.Command("find","/","-name","nginx")
			out03,_:=c3.CombinedOutput()
			fmt.Printf("Nginx dir:\n%v\n",string(out03))
			err03 := os.Chmod("/usr/share/nginx/html/", 777)
			if err03 != nil {
				fmt.Printf("Chmod dir /usr/share/nginx/html  err:%v\n", err02)
				return
			}else {
				fmt.Printf("Chmod dir /usr/share/nginx/html  success:%v\n", err02)
			}
		}
	}
}

// Start 启动nginx
func (n *Ngx) Start(){
	cmd:=exec.Command("nginx")  //启动nginx
	d1,err:=cmd.CombinedOutput()
	if err != nil {
		fmt.Printf("Nginx  启动失败 %v\n %v\n",err,string(d1))
		return
	}else {
		fmt.Printf("Nginx 启动成功 %v\n",string(d1))
	}

}

// Reload 重载nginx
func (n *Ngx) Reload(){
	cmd:=exec.Command("nginx","-s","reload")  //重载nginx
	d1,err:=cmd.CombinedOutput()
	if err != nil {
		fmt.Printf("Nginx  重载失败 %v\n %v\n",err,string(d1))
		return
	}else {
		fmt.Printf("Nginx 重载成功 %v\n",string(d1))
	}

}

// Stop 停止nginx
func (n *Ngx) Stop(){
	cmd:=exec.Command("nginx","-s","stop")  //停止nginx
	d1,err:=cmd.CombinedOutput()
	if err != nil {
		fmt.Printf("Nginx 停止失败 %v\n %v\n",err,string(d1))
		return
	}else {
		fmt.Printf("Nginx 停止成功 %v\n",string(d1))
	}

}

// Remove 删除nginx
func (n *Ngx) Remove(){
	cmd:=exec.Command("yum","remove","-y",Ngxrpm)
	d1,err:=cmd.CombinedOutput()
	if err != nil {
		fmt.Printf("Nginx 删除失败 %v\n",string(d1))
		fmt.Printf("Nginx 删除失败 %v\n",err)
		return
	}else {
		fmt.Printf("Nginx 删除成功 %v\n",string(d1))
	}
}

// Deploy 配置nginx
func (n *Ngx) Deploy(){
	err:=os.Remove(ngxconf)
	if err != nil {
		fmt.Printf("删除 nginx conf errr:%v\n",err)
		return
	}else {
		fmt.Printf("删除 nginx conf 成功\n")
	}
	cmd:=exec.Command("cp","-rp",confdir+"/default.conf","/etc/nginx/conf.d/")
	out02,err02:=cmd.CombinedOutput()
	if err02 != nil {
		fmt.Printf("copy  ngxconf  err:%v\n",err02)
		fmt.Printf("copy  ngxconf  err:%v\n",string(out02))
		return
	} else {
		fmt.Printf("copy  ngxconf  success:%v\n",string(out02))
	}
	var r Ngxroom
	r=&Ngx{}
	r.Reload() //调用Reload 方法

}


