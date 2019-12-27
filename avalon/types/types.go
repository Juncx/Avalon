package types

import "time"

type ItemInfo struct {
	Id       int     `json:"id" xorm:"not null pk autoincr"`
	Name     string  `json:"name" xorm:"not null"`
	Title    string  `json:"title" xorm:"not null"`
	Describe string  `json:"describe"`
	Price    float32 `json:"price" xorm:"not null"`
	Images   string  `json:"images"`
	Externs  string  `json:"externs"`

	CreateTime time.Time `xorm:"created"`
	UpdateTime time.Time `xorm:"updated"`
}

type ItemShelf struct {
	ItemID     int     `json:"itemid" xorm:"not null pk"`
	Itemname   string  `json:"itemname" xorm:"not null"`
	ItemPrice  float32 `json:"itemprice"`
	ItemRemain int     `json:"itemremain" xorm:"not null"`
	ItemSold   int     `json:"itemsold" xorm:"not null"`

	CreateTime time.Time `xorm:"created"`
	UpdateTime time.Time `xorm:"updated"`
}

type UserInfo struct {
	Id     int    `json:"id" xorm:"not null pk autoincr"`
	Name   string `json:"name" xorm:"not null"`
	Intro  string `json:"intro"`
	Avatar string `json:"avatar"`
	Brith  string `json:"brith"`
	Addr   string `json:"address"`

	CreateTime time.Time `xorm:"created"`
	UpdateTime time.Time `xorm:"updated"`
}

type OrderInfo struct {
	Id         int     `json:"id" xorm:"not null pk autoincr"`
	UserID     int     `json:"userid" xorm:"not null"`
	UserName   string  `json:"username" xorm:"not null"`
	ItemID     int     `json:"itemid" xorm:"not null"`
	Itemname   string  `json:"itemname" xorm:"not null"`
	ItemPrice  float32 `json:"itemprice"`
	Quantity   int     `json:"quantity" xorm:"not null"`
	Describe   string  `json:"describe"`
	Settlement bool    `json:"settlement"`

	CreateTime time.Time `xorm:"created"`
	UpdateTime time.Time `xorm:"updated"`
}
