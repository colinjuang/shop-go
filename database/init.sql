-- 添加系统管理员用户（如果有需要）
INSERT INTO `users` (`open_id`, `nickname`, `avatar`, `gender`, `city`, `province`, `country`) VALUES
('admin_openid', '系统管理员', 'avatars/admin.jpg', 1, '上海', '上海', '中国');

-- 添加更多产品分类
INSERT INTO `categories` (`name`, `parent_id`, `image`, `sort_order`) VALUES 
('生日鲜花', 1, 'category/birthday_flowers.jpg', 5),
('纪念日鲜花', 1, 'category/anniversary_flowers.jpg', 6),
('办公室绿植', 2, 'category/office_plants.jpg', 1),
('小型绿植', 2, 'category/small_plants.jpg', 2),
('多肉植物', 2, 'category/succulents.jpg', 3);

-- 添加更多产品
INSERT INTO `products` (`name`, `description`, `price`, `stock`, `category_id`, `images`, `main_image`, `status`, `hot`, `recommend`, `sort_order`) VALUES
('混合花束', '精选多种鲜花组合而成的混合花束，色彩丰富', 238.00, 50, 1, 'products/mixed_bouquet_1.jpg,products/mixed_bouquet_2.jpg', 'products/mixed_bouquet_1.jpg', 1, 1, 1, 5),
('粉玫瑰花束', '19枝精选粉玫瑰，寓意甜蜜爱情', 268.00, 100, 5, 'products/pink_roses_1.jpg,products/pink_roses_2.jpg', 'products/pink_roses_1.jpg', 1, 0, 1, 6),
('白百合花束', '5枝白百合，寓意纯洁与祝福', 219.00, 60, 6, 'products/white_lilies_1.jpg,products/white_lilies_2.jpg', 'products/white_lilies_1.jpg', 1, 0, 0, 7),
('康乃馨花束', '12枝粉色康乃馨，表达对母亲的爱', 188.00, 80, 7, 'products/carnations_1.jpg,products/carnations_2.jpg', 'products/carnations_1.jpg', 1, 1, 0, 8),
('红掌绿植', '红掌盆栽，美观且易打理', 168.00, 50, 12, 'products/anthurium_1.jpg,products/anthurium_2.jpg', 'products/anthurium_1.jpg', 1, 0, 1, 9),
('多肉组合', '多种精选多肉植物组合盆栽', 99.00, 100, 14, 'products/succulents_1.jpg,products/succulents_2.jpg', 'products/succulents_1.jpg', 1, 1, 1, 10),
('商务花篮', '高档商务花篮，适合开业、庆典等场合', 399.00, 30, 3, 'products/business_basket_1.jpg,products/business_basket_2.jpg', 'products/business_basket_1.jpg', 1, 0, 1, 11),
('生日礼盒', '生日主题鲜花礼盒，含蛋糕装饰', 299.00, 40, 4, 'products/birthday_box_1.jpg,products/birthday_box_2.jpg', 'products/birthday_box_1.jpg', 1, 1, 1, 12);

-- 创建测试用户订单
-- 首先创建一个测试用户
INSERT INTO `users` (`open_id`, `nickname`, `avatar`, `gender`, `city`, `province`, `country`) VALUES
('test_user_openid', '测试用户', 'avatars/test_user.jpg', 1, '北京', '北京', '中国');

-- 添加测试用户地址
INSERT INTO `addresses` (`user_id`, `name`, `phone`, `province`, `city`, `district`, `detail_addr`, `is_default`) VALUES
(2, '张三', '13800138000', '北京', '北京市', '海淀区', '中关村科技园区1号楼', 1);

-- 创建一个测试订单
INSERT INTO `orders` (`user_id`, `order_no`, `total_amount`, `payment_amount`, `status`, `payment_time`, `address_id`, `receiver_name`, `receiver_phone`, `address`) VALUES
(2, 'ORD20220101001', 188.00, 188.00, 1, NOW(), 1, '张三', '13800138000', '北京市海淀区中关村科技园区1号楼');

-- 添加订单商品
INSERT INTO `order_items` (`order_id`, `product_id`, `quantity`, `price`, `name`, `image`) VALUES
(1, 1, 1, 188.00, '红玫瑰花束', 'products/red_roses_1.jpg'); 