package model

import (
	"time"
)

// CREATE TABLE `products` (
// 	`id` int(10) unsigned NOT NULL AUTO_INCREMENT,
// 	`name` varchar(200) NOT NULL COMMENT '商品名称',
// 	`floral_language` varchar(255) NOT NULL COMMENT '花语',
// 	`price` decimal(10,2) NOT NULL DEFAULT '0.00' COMMENT '商品价格',
// 	`market_price` decimal(10,2) NOT NULL COMMENT '市场价格',
// 	`sale_count` int(11) NOT NULL COMMENT '销售数量',
// 	`stock_count` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '库存数量',
// 	`category_id` int(10) unsigned NOT NULL COMMENT '分类ID',
// 	`sub_category_id` int(11) NOT NULL COMMENT '二级分类ID',
// 	`image_url` varchar(255) NOT NULL COMMENT '商品图片',
// 	`material` varchar(255) NOT NULL COMMENT '原料',
// 	`packing` varchar(255) NOT NULL COMMENT '包装',
// 	`status` tinyint(1) NOT NULL DEFAULT '1' COMMENT '状态：1上架，0下架',
// 	`recommend` tinyint(1) NOT NULL DEFAULT '0' COMMENT '是否推荐',
// 	`sort_order` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '排序',
// 	`apply_user` varchar(255) NOT NULL COMMENT '赠送对象',
// 	`created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
// 	`updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
// 	PRIMARY KEY (`id`),
// 	KEY `idx_category_id` (`category_id`),
// 	KEY `idx_recommend` (`recommend`)
//   ) ENGINE=InnoDB AUTO_INCREMENT=13 DEFAULT CHARSET=utf8mb4 COMMENT='商品表';

// Product represents a product
type Product struct {
	ID             uint      `json:"id" gorm:"column:id;primaryKey"`
	Name           string    `json:"name" gorm:"column:name;not null"`
	FloralLanguage string    `json:"floralLanguage" gorm:"column:floral_language"`
	Price          float64   `json:"price" gorm:"column:price;type:decimal(10,2)"`
	MarketPrice    float64   `json:"marketPrice" gorm:"column:market_price;type:decimal(10,2)"`
	SaleCount      int       `json:"saleCount" gorm:"column:sale_count;default:0"`
	StockCount     int       `json:"stockCount" gorm:"column:stock_count;default:0"`
	CategoryID     uint      `json:"categoryId" gorm:"column:category_id;index"`
	SubCategoryID  uint      `json:"subCategoryId" gorm:"column:sub_category_id;index"`
	Material       string    `json:"material" gorm:"column:material"`
	Packing        string    `json:"packing" gorm:"column:packing"`
	ImageUrl       string    `json:"imageUrl" gorm:"column:image_url"`
	Status         int       `json:"status" gorm:"column:status;default:1"` // 1: on sale, 0: off sale
	Recommend      bool      `json:"recommend" gorm:"column:recommend;default:false"`
	SortOrder      int       `json:"sortOrder" gorm:"column:sort_order;default:0"`
	ApplyUser      string    `json:"applyUser" gorm:"column:apply_user"`
	CreatedAt      time.Time `json:"createdAt" gorm:"column:created_at"`
	UpdatedAt      time.Time `json:"updatedAt" gorm:"column:updated_at"`
}
