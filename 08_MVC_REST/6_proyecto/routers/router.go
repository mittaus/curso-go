package routers

import (
	"net/http"

	"github.com/gin-gonic/gin"

	_ "github.com/EDDYCJY/go-gin-example/docs"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"

	"github.com/EDDYCJY/go-gin-example/middleware/jwt"
	"github.com/EDDYCJY/go-gin-example/pkg/export"
	"github.com/EDDYCJY/go-gin-example/pkg/qrcode"
	"github.com/EDDYCJY/go-gin-example/pkg/upload"
	"github.com/EDDYCJY/go-gin-example/routers/api"
	v1 "github.com/EDDYCJY/go-gin-example/routers/api/v1"
)

//InitRouter configuración inicial
func InitRouter() *gin.Engine {
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	r.StaticFS("/export", http.Dir(export.GetExcelFullPath()))
	r.StaticFS("/upload/images", http.Dir(upload.GetImageFullPath()))
	r.StaticFS("/qrcode", http.Dir(qrcode.GetQrCodeFullPath()))

	r.POST("/auth", api.GetAuth)
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	r.POST("/upload", api.UploadImage)

	apiv1 := r.Group("/api/v1")
	apiv1.Use(jwt.JWT())
	{
		//Obtenga una lista de etiquetas
		apiv1.GET("/tags", v1.GetTags)
		//Nueva etiqueta
		apiv1.POST("/tags", v1.AddTag)
		//Actualizar etiqueta especificada
		apiv1.PUT("/tags/:id", v1.EditTag)
		//Eliminar etiqueta especificada
		apiv1.DELETE("/tags/:id", v1.DeleteTag)
		//Exportar etiqueta
		r.POST("/tags/export", v1.ExportTag)
		//Importar etiquetas
		r.POST("/tags/import", v1.ImportTag)

		//Obtenga una lista de artículos
		apiv1.GET("/articles", v1.GetArticles)
		//Obtén el artículo especificado
		apiv1.GET("/articles/:id", v1.GetArticle)
		//Articulo nuevo
		apiv1.POST("/articles", v1.AddArticle)
		//Actualizar el artículo especificado
		apiv1.PUT("/articles/:id", v1.EditArticle)
		//Eliminar artículo especificado
		apiv1.DELETE("/articles/:id", v1.DeleteArticle)
		//Generar póster de artículo
		apiv1.POST("/articles/poster/generate", v1.GenerateArticlePoster)

	}

	return r
}
