package main

import (
	"fmt"
	"strings"
	"time"
	"os"
	"bufio"
	"io"
)

type Reader interface { //用接口实现
	read(rc chan []byte)
}

type Write interface {
	write(pw chan string)
}

type LogProcess struct { //通过go的结构体 相当于创建了一个类
	rc chan []byte //创建channel
	pw chan string
	read Reader
	write Write

}

type LogRead struct {
	path string //读取文件的路径
}

type LogWrite struct {
	influxdbdsn string //data soure
}


//用引用,结构体很大就不需要copy
func (r *LogRead) read(rc chan []byte){
	//读取模块
	//打开文件
	f, err := os.Open(r.path)
	if err != nil{
		panic(fmt.Sprintf("open file error:%s",err.Error()))
	}
	//从文件末尾开始逐行读取
	f.Seek(0,2)
	fd := bufio.NewReader(f)

	for {
		line,err :=fd.ReadBytes('\n')
		if err == io.EOF{ //文件末尾抛出的异常
			time.Sleep(500 *time.Millisecond)
			continue
		} else if err != nil{
			panic(fmt.Sprintf("out of file:%s",err.Error()))
		}
		//line := "message"
		rc <- line[:len(line)-1]  //line是slice
	}
}
func (l *LogProcess) LogPaser(){
	//解析模块
	for v := range l.rc {
		l.pw <- strings.ToUpper(string(v))
	}

}

func (w *LogWrite) write(pw chan string){
	//写入模块
	for v := range pw {
		fmt.Println(v)
	}

}

func main() {

	r := &LogRead{  //读取器
		path : "/tmp/access.log",
	}
	w := &LogWrite{ //写入器
		influxdbdsn : "username&password",
	}
	lp := &LogProcess{
		rc: make(chan []byte),
		pw: make(chan string),
		read:r,
		write:w,
	}

	go lp.read.read(lp.rc)
	go lp.LogPaser()
	go lp.write.write(lp.pw)

	time.Sleep(100*time.Second)
}
