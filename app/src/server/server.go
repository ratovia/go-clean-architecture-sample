package server

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type apiHandler func(*gin.Context) (int, interface{}, error)

func RunServer(env string, port string) {
	fmt.Printf("Running server in %s environment...\n", env)

	// DBの初期化などの処理
	db := InitializeDB()

	// コントローラの初期化
	controller := NewController(db, env)

	// ルータの作成
	router := gin.Default()

	// ルーティングの設定
	handleRootGroup(controller, router.Group("/"))

	// サーバーの起動
	log.Fatal(http.ListenAndServe(":" + port, router))
}
