package directory

import (
	"context"
	"fmt"
	"os"
	"os/exec"
	"sync"
	"time"
)

var (
	Workdir       ="/home/github.com"                         //dir目录
	GithubAddress ="https://github.com.cnpmjs.org/liusir-ht/go-lnmp.git" //仓库地址
	wg sync.WaitGroup
	GitlfsAddress="https://github.com/git-lfs/git-lfs/releases/download/v2.13.3/git-lfs-linux-amd64-v2.13.3.tar.gz"
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
				fmt.Println("创建GitHub目录成功","开始DownLoadpkg......")
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
		c5:=exec.Command("wget",GitlfsAddress)  //下载 Gitlfs
		c5.Dir="/home"  // 指定工作目录
		c6:=exec.Command("tar","-zxf","git-lfs-linux-amd64-v2.13.3.tar.gz")
		c7:=exec.Command("/home/git-lfs","install") //安装Gitlfs
		c6.Dir=c5.Dir
		out05,err05:=c5.CombinedOutput()
		if err05!=nil {
			fmt.Printf("获取git-lfs err:%v\n",string(out05))
			fmt.Printf("获取git-lfs err:%v\n",err05)
			return
		}else {
			fmt.Println("获取git-lfs  success")
		}
		time.Sleep(time.Second *3)
		_=c6.Run()
		out07,err07:=c7.CombinedOutput()
		if err07!=nil {
			fmt.Printf("安装git-lfs err:%v\n",string(out07))
			fmt.Printf("安装git-lfs err:%v\n",err07)
			return
		}else {
			fmt.Println("安装git-lfs  success")
		}
	}
	ctx,cancel:=context.WithTimeout(context.Background(),time.Minute*5 )   //定一个 context超时控制
	wg.Add(1)  //设置goroutine的个数 1
	go GitClone(ctx)  //把这个带有超时的 context传入进去
	wg.Wait()  //阻塞等待 goroutine 执行结束
	select { //select 多路复用
		case <-ctx.Done(): //等待上级信号,一旦从这个channel获取到内容 就 输出命令超时，否则一直堵塞
			fmt.Println("执行命令超时")
			defer cancel()  //执行cancel函数
			return
		default:
		}
	defer cancel()
}
func  GitClone(ctx context.Context){    //定义一个GitClone函数
	   start:=time.Now()  //定义程序开始时间
	   fmt.Printf("开始时间 ： %v\n",start)
		c2:=exec.CommandContext(ctx,"git","clone", GithubAddress) //克隆 GitHub 到本地
		c2.Dir=Workdir //指定工作目录
		out01,err01:=c2.CombinedOutput()
		if err01 != nil {
			fmt.Printf("git clone  err:%v\n",string(out01))
			fmt.Printf("请查看网络后 重试\n")
			stop:=time.Since(start) //计算程序运行的时间
			fmt.Printf("耗时 %v\n",stop)
			_=os.RemoveAll(Workdir)  //如果克隆失败，就删除原来的目录
			ctx.Done()  //发送 ctx.Done 信号 给 上级
			return
		} else {
			fmt.Printf("git  clone  success:%v\n",string(out01))
			stop:=time.Since(start)  //计算程序运行的时间
			fmt.Printf("耗时 %v\n",stop)
			ctx.Done()  //发送 ctx.Done 信号 给 上级
		}
		wg.Done()  //让设置 goroutine的个数  -1
}