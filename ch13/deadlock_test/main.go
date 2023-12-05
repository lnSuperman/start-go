package main

func main() {
	var msg chan int
	msg = make(chan int)
	//如果是没有缓冲的channel 在没有启动一个消费者之前 放数据就会报错
	msg <- 1
	//data := <-msg
	//fmt.Println(data)
}
