package fetcher

import (
	"sync"
)

var c chan int
var o = make(chan *Items, 5)

//var itemChan = make(chan Item, 1)

var wg = new(sync.WaitGroup)

func Run() {

	// 获取页数，全部塞进channel中
	totalPage := GetTotalPage()

	//totalPage := 10
	c = make(chan int, totalPage)
	for p := 1; p < totalPage; p++ {
		c <- p
	}
	close(c)

	// 开启5个消费goroutine 并等待
	go func() {
		for i := 0; i < 10; i++ {
			wg.Add(1)
			go PaginatorIterator(c, o, wg)
		}
		wg.Wait()
		close(o)
	}()

	// 消费抓取到的数据, 知道o chan 结束
	for items := range o {
		for _, item := range items.List {
			//itemChan <- item
			WriteItem(item)
		}
	}

}
