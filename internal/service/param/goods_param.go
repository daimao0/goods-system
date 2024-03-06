package param

// GoodsParam 商品参数
type GoodsParam struct {
	//商品名称
	Name string
	//商品编号
	GoodsNumber string
	//商品类型
	GoodsType string
}

// GoodsReduceParam 减少商品库存
type GoodsReduceParam struct {
	//id
	Id uint64
	//购买数量
	Num int
}

// GoodsPageParam 商品分页参数
type GoodsPageParam struct {
	//商品名称
	Name string
	//商品类型
	GoodsType string
	//当前页
	Page int
	//页面大小
	Size int
}

// GoodsSaveParam 商品保存
type GoodsSaveParam struct {
	//商品名称
	Name string `json:"name"`
	//商品类型
	GoodsType string `json:"goodsType"`
	//商品编号
	GoodsNumber string `json:"goodsNumber"`
	//商品价格
	Price float64 `json:"price"`
	//数量
	Count int64 `json:"count"`
	//描述
	Desc string `json:"desc"`
}

// GoodsUpdateParam 商品保存
type GoodsUpdateParam struct {
	//主键
	Id uint64 `gorm:"column:id;not null;autoIncrement;primaryKey;comment:'主键id'"`
	//商品名称
	Name string `gorm:"column:name;default:'';not null;type:varchar(255);comment:'商品名称';index"`
	//商品编号
	GoodsNumber string `gorm:"column:goods_number;default:'';not null;type:varchar(255);comment:'商品编号';index"`
	//商品类型
	GoodsType string `gorm:"column:goods_type;default:'';not null;type:varchar(255);comment:'商品类型'"`
	//商品价格
	Price float64 `gorm:"column:price;default:0;not null;comment:'商品价格'"`
	//数量
	Count int64 `gorm:"column:count;default:0;not null;type:int(10);comment:'商品数量'"`
	//描述
	Desc string `gorm:"column:desc;default:'';not null;type:text;comment:'商品描述'"`
}
