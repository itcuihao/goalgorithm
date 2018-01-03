package main

import (
	"fmt"
	"sync"
)

func main() {
	// 互斥锁
	// 在main函数所在线程中执行两次mu.Lock()，当第二次加锁时会因为锁已经被占用（不是递归锁）而阻塞，main函数的阻塞状态驱动后台线程继续向前执行。当后台线程执行到mu.Unock()时解锁，解锁会导致main函数中第二个mu.Lock()阻塞状态取消，但是这时已经确保打印工作完成了。但是解锁后，后台线程和主线程再没有其它的同步事件参考，它们何时退出的事件将是并发的：在main函数退出导致程序退出时，后台线程可能已经退出了，也可能没有退出。虽然，无法确定两个线程退出的时间，但是打印工作是可以正确完成的。
	var mx sync.Mutex
	mx.Lock()
	go func() {
		fmt.Println("你好")
		mx.Unlock()
	}()
	mx.Lock()

	// 无缓冲通道
	// 对于从无缓冲信道进行的接收，发生在对该信道进行的发送完成之前。因此，后台线程<-done接收操作完成之后，main线程的done <- 1发生操作才可能完成（从而退出main、退出程序），而此时打印工作已经完成了。
	done := make(chan int)
	go func() {
		fmt.Println("你好")
		fmt.Println("a")
		i, ok := <-done
		fmt.Println(i, ok)
	}()
	fmt.Println("b")
	done <- 1
	fmt.Println("c")

	// 带缓冲通道
	// 对于带缓冲的Channel，对于Channel的第K个接收完成操作发生在第K+C个发送操作完成之前，其中C是Channel的缓存大小。虽然管道是带缓存的，main线程接收完成是在后台线程发送开始但还未完成的时刻，此时打印工作也是已经完成的。
	done1 := make(chan int, 1)
	go func() {
		fmt.Println("你好")
		fmt.Println("a")
		done1 <- 1
	}()
	fmt.Println("b")
	<-done1
	fmt.Println("c")

	// 将打印线程扩展到10个

	done10 := make(chan int, 10)
	for i := 0; i < cap(done10); i++ {
		go func() {
			fmt.Println("你好")
			done10 <- 1
		}()
	}

	for i := 0; i < cap(done10); i++ {
		<-done10
	}

	// 等待N个线程完成后再进行下一步的同步操作有一个简单的做法，就是使用sync.WaitGroup

	fmt.Println("waitgroup")
	var wg sync.WaitGroup

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			fmt.Println(i)
		}(i)
	}
	wg.Wait()
	// 其中wg.Add(1)用于增加等待事件的个数，必须确保在后台线程启动之前执行（如果放到后台线程之中执行这不能保证被正常执行到）。当后台现在完成打印工作之后，调用wg.Done()表示完成一个事件。main函数的wg.Wait()是等待全部的事件完成。

}
