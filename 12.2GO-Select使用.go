/*
select是Go中的一个控制结构, 类似于用于通讯的swift语句. 每个case必须是一个通信操作,要么是发送要么是接收.
select [[随机]]执行一个可运行的case . 如果没有case可运行, 他将阻塞,直到有case可运行.
一个默认的子句应该总是可运行的.

select {
    case communication clause  :
       statement(s);
    case communication clause  :
       statement(s);
    // 你可以定义任意数量的 case
    default : // 可选
	statement(s);
}

*/

package main

import "fmt"
import "time"

// 1.一般使用
func testSelect() {

	var c1, c2, c3 chan int
	var i1, i2 int

	select {
	case i1 = <-c1:
		fmt.Println("received ", i1, "from c1") // 接收
	case c2 <- i2:
		fmt.Println("send", i2, "to c2") // 发送
	case i3, ok := (<-c3): // 或者 i3,ok := <-c3
		if ok {
			fmt.Println("received", i3, "from c3")
		} else {
			fmt.Println("c3 is closed")
		}
	default:
		fmt.Println("no communication")
	}

	// 由于c1,c2,c3都未创建,是空的,直接执行default
}

// 2.根据default特性,我们可以判断channel是否满了!
func testChannel() {

	c := make(chan int, 1)
	c <- 1
	select {
	case c <- 2:
	default:
		fmt.Println("channel is full !")
	}

	/*
		因为 c 插入 1 的时候已经满了， 当 c 要插入 2 的时候，发现 c 已经满了（case1 阻塞住）， 则 select 执行 default 语句。
		这样就可以实现对 channel 是否已满的检测， 而不是一直等待。
		比如我们有一个服务， 当请求进来的时候我们会生成一个 job 扔进 channel， 由其他协程从 channel 中获取 job 去执行。
		但是我们希望当 channel 瞒了的时候， 将该 job 抛弃并回复 【服务繁忙，请稍微再试。】 就可以用 select 实现该需求
	*/
}

// 3.使用select判断是否超时
func testTimeout() {

	c := make(chan bool)
	select {
	case v := <-c:
		fmt.Println("received ", v)
	case <-time.After(3 * time.Second): // 由于没有收到写入channel信息,所以进入超时
		fmt.Println("timeout")
		close(c)
	}

}

// 4.两个goroutine之间传值
var ch chan string

func testPing() {
	i := 0
	for {
		fmt.Printf("ping print:#%v\n", <-ch)
		ch <- fmt.Sprintf("from ping : hi, #%d", i)
		i++
	}
}

func testPang() {
	ch = make(chan string)
	go testPing()
	for i := 0; i < 5; i++ {
		ch <- fmt.Sprintf("from pang: Hello, #%d", i)
		fmt.Printf("pang print: #%v\n", <-ch)

	}

	/*
		ping print:#from pang: Hello, #0
		pang print: #from ping : hi, #0
		ping print:#from pang: Hello, #1
		pang print: #from ping : hi, #1
		ping print:#from pang: Hello, #2
		pang print: #from ping : hi, #2
		ping print:#from pang: Hello, #3
		pang print: #from ping : hi, #3
		ping print:#from pang: Hello, #4
		pang print: #from ping : hi, #
	*/
}
