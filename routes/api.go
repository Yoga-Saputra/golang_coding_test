package routes

import (
	"net/http"
	"transfer-pinnacle/app/handler"

	"github.com/gin-gonic/gin"
	"github.com/jpillora/overseer"
)

func InitApi(state overseer.State) {
	router := gin.Default()
	router.SetTrustedProxies(nil)
	router.Use(gin.Logger())

	router.POST("/", func(ctx *gin.Context) {
		ctx.JSON(http.StatusMethodNotAllowed, "Method not allowed")
	})

	router.GET("/", handler.Hello)
	router.GET("/language", handler.GetLanguage)
	router.GET("palindrom", handler.GetPalindrom)

	router.POST("/language", handler.PostLanguage)
	router.GET("/languages", handler.GetLanguages)
	router.GET("/language/:id", handler.GetLanguageById)
	router.PATCH("/language/:id", handler.UpdateLanguage)
	router.DELETE("/language/:id", handler.DeleteLanguage)

	router.Run(":3003")

	http.Serve(state.Listener, router)
}
