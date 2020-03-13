package fetcher

import "fmt"

type DBItem struct {
	Unique_id        string
	Short_title      string
	Title            string
	Intro            string
	Tags             string
	Pic_url          string
	Cate_id          int
	Origin_price     float64
	Price            float64
	Shop_type        string
	Sell_num         int
	Seller_id        string
	Seller_nick      string
	Coupon_price     float64
	Coupon_rate      string
	Coupon_condition string
	Coupon_surplus   int
	Coupon_receive   int
	Coupon_pc_url    string
	Commission       interface{}
	Commission_type  int
	Item_url         string
	Stop_at          string
	Start_at         string
}

func WriteItem(item Item) {
	fmt.Println(item.Dtitle)

	var dbItem DBItem

	// 唯一id
	dbItem.Unique_id = item.GoodsId

	// 短标题
	dbItem.Short_title = item.Dtitle

	// 标题
	dbItem.Title = item.Title

	// 介绍词语
	dbItem.Intro = item.Desc

	// tags
	dbItem.Tags = ""
	// 图片url
	dbItem.Pic_url = item.MainPic

	// 分类id
	dbItem.Cate_id = item.Cid

	// 原价(元)
	dbItem.Origin_price = item.OriginalPrice
	// 现价(元)
	dbItem.Price = item.ActualPrice

	// 店铺累型，B店天猫，C店淘宝
	if item.ShopType == 1 {
		dbItem.Shop_type = "B"
	} else {
		dbItem.Shop_type = "C"
	}

	// 商品销量
	dbItem.Sell_num = item.MonthSales

	// 卖家id
	dbItem.Seller_id = item.SellerId
	// 卖家昵称
	dbItem.Seller_nick = item.ShopName

	//优惠券金额(元)
	dbItem.Coupon_price = item.CouponPrice

	// 折扣率， 例如原价5元现价4元就是4/5*10=8折
	dbItem.Coupon_rate = "0"

	// 优惠券条件
	dbItem.Coupon_condition = item.CouponConditions
	// 优惠券库存量
	dbItem.Coupon_surplus = item.CouponTotalNum
	// 已经领取数量
	dbItem.Coupon_receive = item.CouponReceiveNum
	// 电脑券链接
	dbItem.Coupon_pc_url = item.CouponLink
	// 佣金比例
	dbItem.Commission = item.CommissionRate

	// 佣金类型，0-通用，1-定向，2-高佣，3-营销计划
	dbItem.Commission_type = item.CommissionType
	// 商品链接
	dbItem.Item_url = item.ItemLink
	// 优惠券结束时间（等同于优惠券结束时间）
	dbItem.Stop_at = item.CouponEndTime
	// 优惠券结束时间（等同于优惠券结束时间）
	dbItem.Start_at = item.CouponStartTime
}
