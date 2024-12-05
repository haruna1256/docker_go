// GinでWebサーバにアクセスする
package main

import "github.com/gin-gonic/gin"

func main() {
	router := gin.Default()
	// ルートエンドポイント
	router.GET("/", func(c *gin.Context) {
    c.JSON(200, gin.H{
        "message": "Hello World!!",
    })
	})

	// サーバーを起動
	router.Run(":8000") // ポート8000で接続をリッスン
}