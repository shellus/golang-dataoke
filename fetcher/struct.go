package fetcher

import (
	"encoding/json"
	"time"
)

type Item struct {
	GoodsId          string
	Dtitle           string
	Title            string
	Desc             string
	MainPic          string
	Cid              int
	OriginalPrice    float64
	ActualPrice      float64
	ShopType         int
	MonthSales       int
	SellerId         string
	ShopName         string
	CouponPrice      float64
	CouponConditions string
	CouponTotalNum   int
	CouponReceiveNum int
	CouponLink       string
	CommissionRate   interface{}
	CommissionType   int
	ItemLink         string
	CouponEndTime    string
	CouponStartTime  string
}

type Items struct {
	PageId   string
	TotalNum int
	List     []Item
}

type Result struct {
	Time time.Duration
	Code int
	Msg  string
	Data *json.RawMessage
}
