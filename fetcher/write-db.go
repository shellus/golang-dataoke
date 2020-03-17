package fetcher

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var db *gorm.DB

func Connect(driver string, dsn string) {
	var err error

	db, err = gorm.Open(driver, dsn)
	if err != nil {
		panic(fmt.Sprintf("failed to connect database: %s\n", err.Error()))
	}

	// Migrate the schema
	db.AutoMigrate(&Product{})
}

func Close() error {
	return db.Close()
}

type Product struct {
	gorm.Model
	UniqueId        string
	ShortTitle      string
	Title           string
	Intro           string
	Tags            string
	PicUrl          string
	CateId          uint
	OriginPrice     float64
	Price           float64
	ShopType        string
	SellNum         uint
	SellerId        string
	SellerNick      string
	CouponPrice     float64
	CouponRate      string
	CouponCondition string
	CouponSurplus   uint
	CouponReceive   uint
	CouponPcUrl     string
	Commission      float64
	CommissionType  uint
	ItemUrl         string
	StopAt          string
	StartAt         string
}

func WriteItem(item *Item) {
	fmt.Println(item.Dtitle)

	var dbItem Product

	// 唯一id
	dbItem.UniqueId = item.GoodsId

	// 短标题
	dbItem.ShortTitle = item.Dtitle

	// 标题
	dbItem.Title = item.Title

	// 介绍词语
	dbItem.Intro = item.Desc

	// tags
	dbItem.Tags = ""
	// 图片url
	dbItem.PicUrl = item.MainPic

	// 分类id
	dbItem.CateId = item.Cid

	// 原价(元)
	dbItem.OriginPrice = item.OriginalPrice
	// 现价(元)
	dbItem.Price = item.ActualPrice

	// 店铺累型，B店天猫，C店淘宝
	if item.ShopType == 1 {
		dbItem.ShopType = "B"
	} else {
		dbItem.ShopType = "C"
	}

	// 商品销量
	dbItem.SellNum = item.MonthSales

	// 卖家id
	dbItem.SellerId = item.SellerId
	// 卖家昵称
	dbItem.SellerNick = item.ShopName

	//优惠券金额(元)
	dbItem.CouponPrice = item.CouponPrice

	// 折扣率， 例如原价5元现价4元就是4/5*10=8折
	dbItem.CouponRate = "0"

	// 优惠券条件
	dbItem.CouponCondition = item.CouponConditions
	// 优惠券库存量
	dbItem.CouponSurplus = item.CouponTotalNum
	// 已经领取数量
	dbItem.CouponReceive = item.CouponReceiveNum
	// 电脑券链接
	dbItem.CouponPcUrl = item.CouponLink
	// 佣金比例
	dbItem.Commission = item.CommissionRate

	// 佣金类型，0-通用，1-定向，2-高佣，3-营销计划
	dbItem.CommissionType = item.CommissionType
	// 商品链接
	dbItem.ItemUrl = item.ItemLink
	// 优惠券结束时间（等同于优惠券结束时间）
	dbItem.StopAt = item.CouponEndTime
	// 优惠券结束时间（等同于优惠券结束时间）
	dbItem.StartAt = item.CouponStartTime

	var user Product

	db.Where("unique_id = ?", dbItem.UniqueId).First(&user)
	if user.ID != 0 {
		fmt.Println("已存在，更新")
		db.Model(user).Updates(&dbItem)
	} else {
		fmt.Println("插入新纪录")
		db.Create(&dbItem)
	}
}
