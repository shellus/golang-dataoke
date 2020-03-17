package dadtaoke

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
	Cid              uint
	OriginalPrice    float64
	ActualPrice      float64
	ShopType         uint
	MonthSales       uint
	SellerId         string
	ShopName         string
	CouponPrice      float64
	CouponConditions string
	CouponTotalNum   uint
	CouponReceiveNum uint
	CouponLink       string
	CommissionRate   float64
	CommissionType   uint
	ItemLink         string
	CouponEndTime    string
	CouponStartTime  string
}

type Items struct {
	PageId   string
	TotalNum int
	List     []*Item
}

type Result struct {
	Time time.Duration
	Code int
	Msg  string
	Data *json.RawMessage
}
