/*
go并发编程中, 使用多核
*/

package main

import (
	"fmt"
	"runtime"
	"sync"
)

func testMultiCore() {

	// 1.控制核数+channel
	runtime.GOMAXPROCS(runtime.NumCPU())

	c := make(chan bool, 10) // 设置10个缓存

	// 这里表示在主线程中执行十个go例程, 这十个例程是异步的,但是整个channel是阻塞当前线程的
	for i := 0; i < 10; i++ {
		go goTask(c, i) // 发起一个goroutine
	}

	for i := 0; i < 10; i++ {
		<-c
	}

	fmt.Println("例程执行完成")

	/*
		goroutine 一个go例程 9 49995001
		goroutine 一个go例程 0 49995001
		goroutine 一个go例程 1 49995001
		goroutine 一个go例程 5 49995001
		goroutine 一个go例程 8 49995001
		goroutine 一个go例程 3 49995001
		goroutine 一个go例程 4 49995001
		goroutine 一个go例程 2 49995001
		goroutine 一个go例程 6 49995001
		goroutine 一个go例程 7 49995001
		例程执行完成
	*/

	// 2.使用waitGroup
	testMultiCore1()
	/*
		goroutine 一个go例程 0 49995001
		goroutine 一个go例程 1 49995001
		goroutine 一个go例程 9 49995001
		goroutine 一个go例程 5 49995001
		goroutine 一个go例程 7 49995001
		goroutine 一个go例程 6 49995001
		goroutine 一个go例程 8 49995001
		goroutine 一个go例程 3 49995001
		goroutine 一个go例程 4 49995001
		goroutine 一个go例程 2 49995001
		例程执行完成

	*/

}

func goTask(c chan bool, index int) {

	a := 1
	for i := 0; i < 10000; i++ {
		a += i
	}
	fmt.Println("goroutine 一个go例程", index, a)

	c <- true
}

func testMultiCore1() {

	// 1.控制核数+WaitGroup
	runtime.GOMAXPROCS(runtime.NumCPU())

	wg := sync.WaitGroup{} // 设置10个缓存
	wg.Add(10)
	// 这里表示在主线程中执行十个go例程, 这十个例程是异步的,但是整个channel是阻塞当前线程的
	for i := 0; i < 10; i++ {
		go goTask1(&wg, i) // 发起一个goroutine
	}

	wg.Wait()

	fmt.Println("例程执行完成")

	/*

	 */

}

func goTask1(wg *sync.WaitGroup, index int) {

	a := 1
	for i := 0; i < 10000; i++ {
		a += i
	}
	fmt.Println("goroutine 一个go例程", index, a)

	wg.Done()
}
