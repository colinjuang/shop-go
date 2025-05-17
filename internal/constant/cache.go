package constant

// 缓存键前缀
const (
	// 首页相关缓存
	HomePrefix = ""
	// 商品相关缓存
	ProductPrefix = "product:"
	// 分类相关缓存
	CategoryPrefix = "category:"
	// 用户相关缓存
	UserPrefix = "user:"
	// 订单相关缓存
	OrderPrefix = "order:"
)

// 首页相关缓存键
const (
	// 首页轮播图
	HomeBanners = HomePrefix + "banners"
	// 首页分类
	HomeCategories = HomePrefix + "categories"
	// 首页促销信息
	HomePromotions = HomePrefix + "promotions"
	// 首页推荐商品
	HomeRecommendProducts = HomePrefix + "recommend_products"
	// 首页热门商品
	HomeHotProducts = HomePrefix + "hot_products"
)

// 商品相关缓存键
const (
	// 商品前缀
	ProductPage = ProductPrefix + "page:"
	// 商品详情
	ProductDetail = ProductPrefix + "detail:"
	// 商品列表
	ProductList = ProductPrefix + "list:"
)

// 分类相关缓存键
const (
	// 分类详情
	CategoryDetail = CategoryPrefix + "detail:"
	// 分类列表
	CategoryList = CategoryPrefix + "list"
	// 子分类列表
	CategorySubs = CategoryPrefix + "subs:"
	// 分类按父ID列表
	CategoryParentID = CategoryPrefix + "parent:"
	// 分类树
	CategoryTree = CategoryPrefix + "tree"
)

// 用户相关缓存键
const (
	// 用户信息
	UserInfo = UserPrefix + "info:"
	// 用户地址列表
	UserAddressList = UserPrefix + "address_list:"
)

// 订单相关缓存键
const (
	// 订单前缀
	OrderUserPage = OrderPrefix + "user:"
	// 订单详情
	OrderDetail = OrderPrefix + "detail:"
	// 订单列表
	OrderList = OrderPrefix + "list:"
	// 订单按订单号
	OrderNo = OrderPrefix + "order_no:"
	// 订单状态
	OrderStatus = OrderPrefix + "status:"
)

// 生成带ID的缓存键
func WithID(key string, id interface{}) string {
	return key + "%v"
}

// 生成带分页的缓存键
func WithPage(key string, page, size int) string {
	return key + "page:%d:size:%d"
}
