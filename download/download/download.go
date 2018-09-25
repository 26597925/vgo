package download

import (
	"fmt"
	"sync"
	"time"
	"strconv"
)

type Urls struct {
	Urls []string
	Wg   sync.WaitGroup
	Chs  chan int  // 默认下载量
	Ans  chan bool // 每个进程的下载状态
}

// 初始化下载地址  根据项目确认使用配置文件的方式还是其他方式，此处使用爬虫处理没公开
func (u *Urls) InitUrl(end chan bool) {
	for i := 0; i < 20; i++ {
		u.Urls = append(u.Urls, strconv.Itoa(i))
	}
	end <- true
}

// 实际的下载操作
func downloadHandle(url string) string {
	//需要根据下载内容作存储等处理
	fmt.Println(url)
	time.Sleep(10*time.Second)
	return ""
}

/**
每个线程的操作
url 下载地址
chs 默认下载量
ans 每个线程的下载状态
*/
func (u *Urls) Work(url string) {
	defer func() {
		<-u.Chs  // 某个任务下载完成，让出
		u.Wg.Done()
		fmt.Println("=============================")
	}()
    downloadHandle(url)
	u.Ans <- true // 告知下载完成
}