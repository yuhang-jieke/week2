package request

type Goods struct {
	Name  string  `form:"name"   binding:"required"`
	Price float64 `form:"price"  binding:"required"`
	Num   int     `form:"num"  binding:"required"`
	Title string  `form:"title"  binding:"required"`
}
type Upload struct {
	Str string `form:"str"   binding:"required"`
}
