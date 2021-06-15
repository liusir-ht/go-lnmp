package directory

import (
	"context"
	"fmt"
	"os"
	"os/exec"
	"time"
)

var (
	Workdir       ="/home/github.com"                         //dir目录
	GithubAddress ="https://github.com.cnpmjs.org/liusir-ht/go-lnmp.git" //仓库地址
)

func CreateDir(){
	_,err:=os.Stat(Workdir)  //检查目录info
	if err !=nil{
		if  os.IsNotExist(err) {   //判断文件或者目录是否存在
			err=os.MkdirAll(Workdir,os.ModePerm)  //如果不存在 创建目录
			if err != nil {
				fmt.Println("创建目录失败\n",err)
				return
			}else {
				fmt.Println("创建GitHub目录成功\n","开始DownLoadpkg......")
				DownLoadPkg()  //下载GitHub 所需要的包
			}
		}
	}else {
		fmt.Println("GitHub目录已存在\n"," 开始DownLoadpkg......")
		DownLoadPkg()  //下载GitHub 所需要的包
	}
}
func DownLoadPkg()  {
	c3:=exec.Command("yum","install","-y","git")  //安装 Git 客户端工具
	c3.Dir=Workdir
	err03:=c3.Run()  //阻塞运行 命令
	if err03 != nil {
		fmt.Printf("git  install err:%v\n",err03)
		return
	} else {
		fmt.Printf("git  install success\n")
	}
	cmd:=exec.Command("git","init")  //初始化 Git
	cmd.Dir=Workdir
	out,err:=cmd.CombinedOutput()
	if err != nil {
		fmt.Printf("初始化git err:%v\n",err)
		fmt.Printf("初始化git err:%v\n",string(out))
		return
	} else {
		fmt.Printf("初始化git success:%v\n 正在克隆 Github仓库 请耐心等待....\n",string(out))
		//提高 GitHub的HTTP 缓存
		c4:=exec.Command("git ","config"," --global"," http.postBuffer 100M")
		c4.Dir=Workdir
		_=c4.Run()
	}
	start:=time.Now()
	ctx,cancel:=context.WithTimeout(context.Background(),time.Minute * 5 )   //定一个 context超时控制
	defer cancel()
	c2:=exec.CommandContext(ctx,"git","clone", GithubAddress) //克隆 GitHub 到本地
	c2.Dir=Workdir
	out01,err01:=c2.CombinedOutput()
	if err01 != nil {
		fmt.Printf("git  clone err:%v\n",err01)
		fmt.Printf("git clone  err:%v\n",string(out01))
		stop:=time.Since(start)
		fmt.Printf("请查看网络后 重试\n")
		fmt.Printf("耗时 %v\n",stop)
		ctx.Done()

	} else {
		fmt.Printf("git  clone  success:%v\n",string(out01))
		fmt.Printf("git  clone success:%v\n",err01)
		stop:=time.Since(start)
		fmt.Printf("耗时 %v\n",stop)
		ctx.Done()
	}



}