package v1

import (
	"github.com/dagou8/blog-service/pkg/app"
	"github.com/dagou8/blog-service/pkg/errcode"
	"github.com/gin-gonic/gin"
)

type Article struct{}

func NewArticle() Article {
	return Article{}
}

// @summary 获取单篇文章
// @Produce json
// @Param id path int true "文章ID"
// @Success 200 {object} model.Article "成功"
// @Failure 400 {object} errcode.Error "请求错误"
// @Failure 500 {object} errcode.Error "内部错误"
// @Router /api/v1/articles/{id} [get]
func (a Article) Get(c *gin.Context) {
	app.NewResponse(c).ToErrorResponse(errcode.ServerError)
	return
}

// @Summary 获取文章列表
// @Produce json
// @Param name query string false "文章标题"
// @Param tag_id query int false "标签ID"
// @Param state query int false "状态"
// @Param page query int false "页码"
// @Param page_size query int false "每页数量"
// @Success 200 {object} model.ArticleSwagger "成功"
// @Failure 400 {object} errcode.Error "请求错误"
// @Failure 500 {object} errcode.Error "内部错误"
// @Router /api/v1/articles [get]
func (a Article) List(c *gin.Context) {}

// @Summary 创建文章
// @Produce json
// @Param tag_id body int true "标签ID"
// @Param title body string true "文章标题"
// @Param desc body string false "文章描述"
// @Param content body string true "文章内容"
// @Param cover_image_url body string true "封面图片地址"
// @Param created_by body string true "创建者"
// @Param state body int false "状态"
// @Success 200 {object} model.Article "成功"
// @Failure 400 {object} errcode.Error "内部错误"
// @Failure 500 {object} errcode.Error "内部错误"
// @Router /api/v1/articles [post]
func (a Article) Create(c *gin.Context) {}

// @Summary 更新文章
// @Produce json
// @Param id path int true "文章ID"
// @Param tag_id body int false "标签ID"
// @Param title body string false "文章标题"
// @Param decs body string false "文章描述"
// @Param cover_image_url body string false "封面图地址"
// @Param content body string false "文章内容"
// @Param modified_by body string true "修改者"
// @Success 200 {object} model.Article "成功"
// @Failure 400 {object} errcode.Error "请求错误"
// @Failure 500 {object} errcode.Error "内部错误"
// @Router /api/v1/articles/{id} [put]
func (a Article) Update(c *gin.Context) {}

// @Summary 删除文章
// @Produce json
// @Param id path int true "文章ID"
// @Success 200 {string} string "成功"
// @Failure 400 {object} errcode.Error "请求错误"
// @Failure 500 {object} errcode.Error "内部错误"
// @Router /api/v1/articles/{id} [delete]
func (a Article) Delete(c *gin.Context) {}
