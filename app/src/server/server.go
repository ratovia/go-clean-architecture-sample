package server

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func RunServer(env string, port string) {
	// DBの初期化などの処理
	db := InitializeDB()

	// コントローラの初期化
	controller := NewController(db, env)

	// ルータの作成
	router := gin.Default()

	// ルーティングの設定
	handleRootGroup(controller, router.Group("/"))

	// サーバーの起動
	log.Fatal(http.ListenAndServe(":"+port, router))
}
