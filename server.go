package main

import (
	"github.com/captain-blue210/youtube-sample/youtube-manager-go/routes"
	"github.com/joho/godotenv"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/sirupsen/logrus"
)

func init() {
	// 環境変数を読み込む
	err := godotenv.Load()
	if err != nil {
		logrus.Fatal("Error loading .env file")
	}
	logrus.SetLevel(logrus.DebugLevel)
	logrus.SetFormatter(&logrus.JSONFormatter{})
}

func main() {
	// Echoのインスタンス作成
	e := echo.New()

	// ログ取得用MiddleWare
	e.Use(middleware.Logger())
	e.Use(middleware.CORS())

	// Routes
	routes.Init(e)

	// 8080番ポートを指定してサーバーを起動
	e.Logger.Fatal(e.Start(":8080"))
}
