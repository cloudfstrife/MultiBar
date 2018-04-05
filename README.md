# MultiBar

## 简介（About）

go语言编写的彩色命令行进度条

show colorful progress bar on the command line,Write in  go programming language.

## 环境（Environment）

Mac Linux 或者 win10

Mac Linux or win10

## 使用（Using）

获取代码（get source code）

```
go get -u github.com/cloudfstrife/bar
```

### 单进度条使用示例（single progress bar demo）

```
package main

import (
	"github.com/cloudfstrife/bar"
	"os"
	"time"
)

func main() {

	process1 := bar.Default()
	process1.Title = "Test"
	process1.Percent = 10
	process1.Show(os.Stdout, 0, true)
	time.Sleep(1 * time.Second)

	process1.Percent = 20
	process1.Show(os.Stdout, 0, true)
	time.Sleep(1 * time.Second)

	process1.Percent = 30
	process1.Show(os.Stdout, 0, true)
	time.Sleep(1 * time.Second)

	process1.Percent = 40
	process1.Show(os.Stdout, 0, true)
	time.Sleep(1 * time.Second)

	process1.Percent = 50
	process1.Show(os.Stdout, 0, true)
	time.Sleep(1 * time.Second)

	process1.Percent = 100
	process1.Show(os.Stdout, 0, true)
}
```

### 单进度条使用示例（multi progress bar demo）

```
package main

import (
	"github.com/cloudfstrife/bar"
	"os"
	"time"
)

func main() {
	multiBar := bar.MultiBar{}

	process1 := bar.Default()
	process1.Title = "default"
	process1.Percent = 5

	process2 := bar.Default()
	process2.Title = "Test"
	process2.Percent = 0

	multiBar.Append("default", process1)
	multiBar.Append("Test", process2)

	multiBar.Show(os.Stdout)
	time.Sleep(1 * time.Second)

	process1.Percent = 15
	process2.Percent = 10
	multiBar.Show(os.Stdout)
	time.Sleep(1 * time.Second)

	process1.Percent = 25
	process2.Percent = 20
	multiBar.Show(os.Stdout)
	time.Sleep(1 * time.Second)

	process1.Percent = 35
	process2.Percent = 30
	multiBar.Show(os.Stdout)
	time.Sleep(1 * time.Second)
}
```

## 代码说明（Code description）

```
.
├── Bar.go				进度条定义（single progress bar definition）
└── MultiBar.go			多进度条定义（multi progress bar definition）
```

### Bar.go

```
func Default() *Bar 
```

创建一个默认的进度条，可以修改此方法中的初始化值构建个性化的默认进度条

create a default progress bar , you can modify the value in this function to build personalized progress bar.

```
func (process *Bar) Show(w io.Writer, maxTitle int, clean bool) 
```

输出进度条

print progress bar

参数说明：

* w			输出目地址

* maxTitle	最长的title，用于多进度条输出时对齐输出内容

* clean		是否清除上一次输出，Bar结构体内部有一个showed，表示是否进行过输出，如果是第一次输出，即使clean为true也不会清理

parameter description

* w			output target 

* maxTitle	use it for align multi progress bar output content

* clean		clean or don't clean the last time output , variable `showed` in `Bar` struct means is already do first out，if current out is the first invoked，this clean parameter is invalid

### MultiBar.go

```
func (multiBar *MultiBar) Append(index string, process *Bar)
```

添加一个进度条到多进度条输出中

append progress bar into multe progress bar

参数说明 

* index		为进度条增加一个键，方便在其它地方获取

* process 	进度条指针

parameter description

* index		key for current progress bar ,you can get this progress bar use this key 

* process		point to progress bar 


```
func (multiBar *MultiBar) Show(w io.Writer) {
```

输出进度条

print progress bar

参数说明 

* w			输出目地址

parameter description

* w			output target 

## Reference(参考资料)

[震惊！原来命令行还可以这么玩？！ ](http://kissyu.org/2017/11/25/%E9%9C%87%E6%83%8A%EF%BC%81%E5%8E%9F%E6%9D%A5%E5%91%BD%E4%BB%A4%E8%A1%8C%E8%BF%98%E5%8F%AF%E4%BB%A5%E8%BF%99%E4%B9%88%E7%8E%A9%EF%BC%9F%EF%BC%81/)

[Build your own Command Line with ANSI escape codes](http://www.lihaoyi.com/post/BuildyourownCommandLinewithANSIescapecodes.html)