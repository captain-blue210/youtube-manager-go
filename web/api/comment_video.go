package api

import (
	"firebase.google.com/go/auth"
	"github.com/captain-blue210/youtube-sample/youtube-manager-go/middlewares"
	"github.com/captain-blue210/youtube-sample/youtube-manager-go/models"
	"github.com/labstack/echo"
	"github.com/valyala/fasthttp"
)

type CommentVideoResponse struct {
	VideoID string `json:"video_id"`
	Comment string `json:"comment"`
}

func CommentVideo() echo.HandlerFunc {
	return func(c echo.Context) error {
		dbs := c.Get("dbs").(*middlewares.DatabaseClient)
		token := c.Get("auth").(*auth.Token)

		user := models.User{}
		if dbs.DB.Table("users").
			Where(models.User{UID: token.UID}).First(&user).RecordNotFound() {
			user = models.User{UID: token.UID}
			dbs.DB.Create(&user)
		}

		cm := models.Comment{}
		videoId := c.Param("id")
		comment := c.Param("comment")
		if dbs.DB.Table("comments").
			Where(models.Comment{UserID: user.ID, VideoID: videoId}).
			First(&cm).RecordNotFound() {
			cm = models.Comment{UserID: user.ID, VideoID: videoId, Comment: comment}
			dbs.DB.Create(&cm)
		}

		res := CommentVideoResponse{
			VideoID: videoId,
			Comment: comment,
		}

		return c.JSON(fasthttp.StatusOK, res)
	}
}
