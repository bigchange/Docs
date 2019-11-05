# advanced-go-programming-book

[图书pdf地址](https://github.com/chai2010/advanced-go-programming-book)

注：Intellij IDE 设置go classpath： 安装插件成功后 —> 进入File —> Other Settings —> Default Project Structure…添加设置 classpath为当前工程的目录，不然项目内部包的导入会出现错误

学习开源项目的实现

查看处理器和系统架构
uname -a;
uname -m;
dpkg --print-architecture；
getconf LONG_BIT;
file /sbin/init;

## GOLANG跨平台编译可运行的二进制文件

"env GOOS=linux GOARCH=amd64 go build main.go"
GOOS：目标平台的操作系统（darwin、freebsd、linux、windows)
GOARCH：目标平台的体系架构（386、amd64、arm）

## 一、go基础知识

1.1  导入包
import(     . "fmt” )
这个点操作的含义就是这个包导入之后在你调用这个包的函数时，你可以省略前缀的包名，也就是前面你调用的fmt.Println("hello world")可以省略的写成Println("hello world")
别名操作别名操作顾名思义我们可以把包命名成另一个我们用起来容易记忆的名字 
import(     f "fmt" )别名操作的话调用包函数时前缀变成了我们的前缀，即f.Println("hello world")
_ 操作这个操作经常是让很多人费解的一个操作符，请看下面这个import
        import (
            "database/sql"
            _ "github.com/ziutek/mymysql/godrv"
        )
_ 操作其实是引入该包，而不直接使用包里面的函数，而是调用了该包里面的init函数。
只要在包中声明了init函数，引用这个包的时候该init函数就会自动被调用，像包中的常量和变量一样被初始化执行 
1.2  基础  
    使用反引号`作为原始字符串符号:
    s :=`Starting part
            Ending part`
1.3  函数
        recover仅在延迟函数中有效。
1.4  进阶
    （1）内建函数 new 本质上说跟其他语言中的同名函数功能一样:new(T) 分配了零值填充的 T 类型的内存空间，并且返回其地址，一个 *T 类型的值。用 Go 的术语说，它返回了一个指针，指向新分配的类型 T 的零值。
    （2） 内建函数 make(T, args) 与 new(T) 有着不同的功能。它只能创建slice，map 和 channel，并且返回一个有初始值(非零)的 T 类型，而不是 *T。本质来讲，导致这三个类型有所不同的原因是指向数据结构的引用在使用前必须被初始化。
    （3） Go 不是面向对象语言，因此并没有继承。但是有时又会需要从已经实现的类型中“继承”并修改一些方法。在 Go 中可以用嵌入一个类型的方式来实现
               组合： type PrintableMutex struct {Mutex } —  PrintableMutex 已经从 Mutex 继承了方法集合
1.5 接口
   (1） 对于接口 I，S 是合法的实现，因为它定义了 I 所需的两个方法。注意，即便是没有明确定义 S 实现了 I，这也是正确的
   (2） var _ io.Writer = (*myWriter)(nil)， 编译器自动检测类型是否实现接口
   (3) 例如某类型有 m 个方法，某接口有 n 个方法，则很容易知道这种判定的时间复杂度为 O(mn)，Go 会对方法集的函数按照函数名的字典序进行排序，所以实际的时间复杂度为 O(m+n)。
1.6 并发
    （1） 有一些丑陋的东西;不得不从 channel 中读取两次(第 14 和 15 行)。在这个例子中没问题，但是如果不知道有启动了多少个 goroutine 怎么办呢?这里有另一个Go 内建的关键字:select。通过 select(和其他东西)可以监听 channel 上输入的数据。
    （2）channel作为变量传入的形式： func dup3(in <−chan int) (<−chan int, <−chan int, <−chan int) 
               解释： 传入一个channel，返回三个channel
1.7 goroutine, context - 并发编程
     Context 可以用来通知后台goroutine的退出，防止泄露
1.8 切片操纵
  切片操作会导致整个底层的数据被锁定，得不到释放。 - 解决办法: 重新克隆一份得到切片内容
1.9 常见问题
  闭包错误引用同一个变量 —   解决办法 ：生成一个局部变量赋值
  在循环内部使用defer —  解决办法：z哎局部函数内部执行defer
  不同goroutine之前不满足顺序一致性内存模型  —  解决办法：显示同步，使用channel
  死循环会独占cpu导致其他goroutine饿死 —  解决办法：加入runtime.Gosched()调度函数
  goroutine是协作式调度，goroutine本身不会主动放弃cpu
  recover必须在defer函数中运行， recover捕获的是祖父级调用异常，直接调用无效
10.0 goroutine竞争检测
 golang在1.1之后引入了竞争检测的概念。我们可以使用go run -race 或者 go build -race 来进行竞争检测。
11. 测试类
12.唤醒锁： sync.Cond
    p.cond = sync.NewCond(&p.lock)
13.方法的定义
   (1)不管方法的接收者是什么类型，该类型的值和指针都可以调用；
   (2) 实现了接收者是值类型的方法，相当于自动实现了接收者是指针类型的方法；而实现了接收者是指针类型的方法，不会自动生成对应接收者是值类型的方法。
   (3) 类型 T 只有接受者是 T 的方法；而类型 *T 拥有接受者是 T 和 *T 的方法。语法上 T 能直接调 *T 的方法仅仅是 Go 的语法糖
