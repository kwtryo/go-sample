package main

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kwtryo/go-sample/config"
	"github.com/kwtryo/go-sample/handler"
)

// TODO: routerパッケージ
func setupRouter(ctx context.Context, cfg *config.Config) *gin.Engine {
	router := gin.Default()
	// ヘルスチェック
	router.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"status": "ok"})
	})

	// clocker := clock.RealClocker{}
	// r := store.Repository{Clocker: clocker}
	// db, cleanup, err := store.New(ctx, cfg)
	// if err != nil {
	// 	return nil, cleanup, err
	// }

	// テスト用
	router.GET("/users", handler.GetAllUser)

	// ru := &handler.RegisterUser{
	// 	DB:   db,
	// 	Repo: &r,
	// }
	// router.POST("/register", ru.ServeHTTP)

	return router
}
