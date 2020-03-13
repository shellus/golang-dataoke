package fetcher

import (
	"fmt"
	"math"
	"sync"
)

var c chan int
var o = make(chan *Items, 5)

//var itemChan = make(chan Item, 1)

var wg = new(sync.WaitGroup)

func getTotal() int {
	items, err := Paginator(1)
	if err != nil {
		panic(err.Error())
	}
	totalNum := items.TotalNum
	total := int(math.Ceil(float64(totalNum / 100)))
	return total
}

func Run() {

	// 获取页数，全部塞进channel中
	total := getTotal()

	fmt.Println(total)
	//total := 10
	c = make(chan int, total)
	for p := 1; p < total; p++ {
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
