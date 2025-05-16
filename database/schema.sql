-- 创建花店数据库
CREATE DATABASE IF NOT EXISTS `flower_shop` DEFAULT CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci;

USE `flower_shop`;

-- 用户表
CREATE TABLE IF NOT EXISTS `users` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `open_id` varchar(100) NOT NULL COMMENT '微信OpenID',
  `nickname` varchar(100) DEFAULT NULL COMMENT '用户昵称',
  `avatar` varchar(255) DEFAULT NULL COMMENT '头像URL',
  `gender` tinyint(1) DEFAULT 0 COMMENT '性别：0未知，1男，2女',
  `city` varchar(50) DEFAULT NULL COMMENT '城市',
  `province` varchar(50) DEFAULT NULL COMMENT '省份',
  `district` varchar(50) DEFAULT NULL COMMENT '区县',
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  UNIQUE KEY `idx_open_id` (`open_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='用户表';

-- 地址表
CREATE TABLE IF NOT EXISTS `addresses` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `user_id` int(10) unsigned NOT NULL COMMENT '用户ID',
  `name` varchar(50) NOT NULL COMMENT '收货人姓名',
  `phone` varchar(20) NOT NULL COMMENT '联系电话',
  `province` varchar(50) NOT NULL COMMENT '省份',
  `city` varchar(50) NOT NULL COMMENT '城市',
  `district` varchar(50) NOT NULL COMMENT '区县',
  `detail_addr` varchar(255) NOT NULL COMMENT '详细地址',
  `is_default` tinyint(1) NOT NULL DEFAULT 0 COMMENT '是否默认地址',
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  KEY `idx_user_id` (`user_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='收货地址表';

-- 商品分类表
CREATE TABLE IF NOT EXISTS `categories` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `name` varchar(100) NOT NULL COMMENT '分类名称',
  `parent_id` int(10) unsigned DEFAULT 0 COMMENT '父分类ID',
  `image` varchar(255) DEFAULT NULL COMMENT '分类图片',
  `sort_order` int(10) unsigned DEFAULT 0 COMMENT '排序',
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  KEY `idx_parent_id` (`parent_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='商品分类表';

-- 商品表
CREATE TABLE IF NOT EXISTS `products` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `name` varchar(200) NOT NULL COMMENT '商品名称',
  `description` text DEFAULT NULL COMMENT '商品描述',
  `price` decimal(10,2) NOT NULL DEFAULT 0.00 COMMENT '商品价格',
  `stock` int(10) unsigned NOT NULL DEFAULT 0 COMMENT '库存数量',
  `category_id` int(10) unsigned NOT NULL COMMENT '分类ID',
  `images` text DEFAULT NULL COMMENT '商品图片，逗号分隔',
  `main_image` varchar(255) DEFAULT NULL COMMENT '主图',
  `status` tinyint(1) NOT NULL DEFAULT 1 COMMENT '状态：1上架，0下架',
  `hot` tinyint(1) NOT NULL DEFAULT 0 COMMENT '是否热销',
  `recommend` tinyint(1) NOT NULL DEFAULT 0 COMMENT '是否推荐',
  `sort_order` int(10) unsigned DEFAULT 0 COMMENT '排序',
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  KEY `idx_category_id` (`category_id`),
  KEY `idx_hot` (`hot`),
  KEY `idx_recommend` (`recommend`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='商品表';

-- 轮播图表
CREATE TABLE IF NOT EXISTS `banners` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `image` varchar(255) NOT NULL COMMENT '图片URL',
  `link` varchar(255) DEFAULT NULL COMMENT '链接URL',
  `sort_order` int(10) unsigned DEFAULT 0 COMMENT '排序',
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='首页轮播图表';

-- 促销活动表
CREATE TABLE IF NOT EXISTS `promotions` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `title` varchar(100) NOT NULL COMMENT '活动标题',
  `image` varchar(255) NOT NULL COMMENT '活动图片',
  `link` varchar(255) DEFAULT NULL COMMENT '活动链接',
  `sort_order` int(10) unsigned DEFAULT 0 COMMENT '排序',
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='促销活动表';

-- 购物车表
CREATE TABLE IF NOT EXISTS `cart_items` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `user_id` int(10) unsigned NOT NULL COMMENT '用户ID',
  `product_id` int(10) unsigned NOT NULL COMMENT '商品ID',
  `quantity` int(10) unsigned NOT NULL DEFAULT 1 COMMENT '数量',
  `selected` tinyint(1) NOT NULL DEFAULT 1 COMMENT '是否选中',
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  KEY `idx_user_id` (`user_id`),
  KEY `idx_product_id` (`product_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='购物车表';

-- 订单表
CREATE TABLE IF NOT EXISTS `orders` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `user_id` int(10) unsigned NOT NULL COMMENT '用户ID',
  `order_no` varchar(100) NOT NULL COMMENT '订单编号',
  `total_amount` decimal(10,2) NOT NULL COMMENT '订单总金额',
  `payment_amount` decimal(10,2) NOT NULL COMMENT '实付金额',
  `status` tinyint(1) NOT NULL DEFAULT 0 COMMENT '订单状态：0待付款，1已付款，2已发货，3已完成，4已取消',
  `payment_time` timestamp NULL DEFAULT NULL COMMENT '付款时间',
  `address_id` int(10) unsigned DEFAULT NULL COMMENT '地址ID',
  `receiver_name` varchar(50) DEFAULT NULL COMMENT '收货人姓名',
  `receiver_phone` varchar(20) DEFAULT NULL COMMENT '收货人电话',
  `address` varchar(255) DEFAULT NULL COMMENT '收货地址',
  `payment_type` tinyint(1) NOT NULL DEFAULT 1 COMMENT '支付方式：1微信支付',
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '订单创建时间',
  `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  UNIQUE KEY `idx_order_no` (`order_no`),
  KEY `idx_user_id` (`user_id`),
  KEY `idx_status` (`status`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='订单表';

-- 订单商品表
CREATE TABLE IF NOT EXISTS `order_items` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `order_id` int(10) unsigned NOT NULL COMMENT '订单ID',
  `product_id` int(10) unsigned NOT NULL COMMENT '商品ID',
  `quantity` int(10) unsigned NOT NULL COMMENT '数量',
  `price` decimal(10,2) NOT NULL COMMENT '价格',
  `name` varchar(200) NOT NULL COMMENT '商品名称',
  `image` varchar(255) DEFAULT NULL COMMENT '商品图片',
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  KEY `idx_order_id` (`order_id`),
  KEY `idx_product_id` (`product_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='订单商品表';

-- 添加外键约束（如果需要）
-- ALTER TABLE `addresses` ADD CONSTRAINT `fk_address_user` FOREIGN KEY (`user_id`) REFERENCES `users` (`id`) ON DELETE CASCADE;
-- ALTER TABLE `cart_items` ADD CONSTRAINT `fk_cart_user` FOREIGN KEY (`user_id`) REFERENCES `users` (`id`) ON DELETE CASCADE;
-- ALTER TABLE `cart_items` ADD CONSTRAINT `fk_cart_product` FOREIGN KEY (`product_id`) REFERENCES `products` (`id`) ON DELETE CASCADE;
-- ALTER TABLE `orders` ADD CONSTRAINT `fk_order_user` FOREIGN KEY (`user_id`) REFERENCES `users` (`id`) ON DELETE CASCADE;
-- ALTER TABLE `order_items` ADD CONSTRAINT `fk_order_item_order` FOREIGN KEY (`order_id`) REFERENCES `orders` (`id`) ON DELETE CASCADE;
-- ALTER TABLE `order_items` ADD CONSTRAINT `fk_order_item_product` FOREIGN KEY (`product_id`) REFERENCES `products` (`id`) ON DELETE SET NULL;

-- 添加初始数据
INSERT INTO `categories` (`name`, `parent_id`, `image`, `sort_order`) VALUES 
('鲜花', 0, 'category/flowers.jpg', 1),
('绿植', 0, 'category/plants.jpg', 2),
('花篮', 0, 'category/baskets.jpg', 3),
('礼盒', 0, 'category/gift_boxes.jpg', 4);

INSERT INTO `categories` (`name`, `parent_id`, `image`, `sort_order`) VALUES 
('玫瑰', 1, 'category/roses.jpg', 1),
('百合', 1, 'category/lilies.jpg', 2),
('康乃馨', 1, 'category/carnations.jpg', 3),
('向日葵', 1, 'category/sunflowers.jpg', 4);

INSERT INTO `banners` (`image`, `link`, `sort_order`) VALUES
('banners/spring_sale.jpg', '/products?category_id=1', 1),
('banners/new_arrivals.jpg', '/products?recommend=1', 2),
('banners/gift_ideas.jpg', '/products?category_id=4', 3);

INSERT INTO `promotions` (`title`, `image`, `link`, `sort_order`) VALUES
('母亲节特惠', 'promotions/mothers_day.jpg', '/products?category_id=3', 1),
('新品上市', 'promotions/new_collection.jpg', '/products?hot=1', 2),
('节日礼盒', 'promotions/holiday_gifts.jpg', '/products?category_id=4', 3);

-- 添加示例产品
INSERT INTO `products` (`name`, `description`, `price`, `stock`, `category_id`, `images`, `main_image`, `status`, `hot`, `recommend`, `sort_order`) VALUES
('红玫瑰花束', '11枝精选红玫瑰，寓意热烈的爱', 188.00, 100, 5, 'products/red_roses_1.jpg,products/red_roses_2.jpg', 'products/red_roses_1.jpg', 1, 1, 1, 1),
('向日葵花束', '7枝向日葵，寓意积极乐观', 158.00, 80, 8, 'products/sunflowers_1.jpg,products/sunflowers_2.jpg', 'products/sunflowers_1.jpg', 1, 1, 0, 2),
('粉百合花束', '5枝粉色百合，清新淡雅', 199.00, 60, 6, 'products/lilies_1.jpg,products/lilies_2.jpg', 'products/lilies_1.jpg', 1, 0, 1, 3),
('永生花礼盒', '永不凋谢的玫瑰花礼盒，长久保存', 299.00, 50, 4, 'products/preserved_roses_1.jpg,products/preserved_roses_2.jpg', 'products/preserved_roses_1.jpg', 1, 1, 1, 4); 