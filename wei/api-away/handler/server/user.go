package server

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/yuhang-jieke/week2/pkg"
	"github.com/yuhang-jieke/week2/wei/api-away/basic/config"
	"github.com/yuhang-jieke/week2/wei/api-away/handler/request"
	__ "github.com/yuhang-jieke/week2/wei/user-server/handler/proto"
)

func Create(c *gin.Context) {
	var form request.Goods
	if err := c.ShouldBind(&form); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code": 400,
			"msg":  "参数不正确",
		})
		return
	}
	_, err := config.UserClient.Goods(c, &__.GoodsReq{
		Name:  form.Name,
		Price: float32(form.Price),
		Num:   int64(form.Num),
		Total: form.Title,
	})

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code": 400,
			"msg":  "添加失败",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "添加成功",
	})
	return
}
func Upload(c *gin.Context) {
	var form request.Upload
	if err := c.ShouldBind(&form); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code": 400,
			"msg":  "参数不正确",
		})
		return
	}
	pkg.Upload(form.Str)
	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "上传成功",
	})
	return
}
