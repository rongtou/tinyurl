package handler

import (
	"fmt"
	"github.com/gin-contrib/location"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"time"
	"tinyurl/db"
	"tinyurl/model"
	"tinyurl/service"
)

func CreateLink(c *gin.Context) {
	link := c.PostForm("url")
	log.Println("link : ", link)

	urlMap := &model.TinyUrlMap{}
	urlMap.OriginUrl = link
	urlMap.ShortUrl = service.GenToken()
	urlMap.CreatedTime = time.Now()

	err := db.CreateLink(urlMap)

	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code":    1001,
			"message": "create link fail",
		})
		return
	}

	url := location.Get(c)
	c.JSON(http.StatusOK, gin.H{
		"code":    0,
		"message": "succ",
		"data": gin.H{
			"short_url": fmt.Sprintf("%s://%s/l/%s", url.Scheme, url.Host, urlMap.ShortUrl),
		},
	})
}

func RedirectLink(c *gin.Context) {
	//读取cache
	//有则跳转
	//没有则查询数据库
	//还没有就设置null的过期，并返回404
	shortUrl := c.Param("link")
	log.Println("short link : ", shortUrl)
	urlMap := db.GetLinkByShort(shortUrl)

	if urlMap == nil {
		c.String(http.StatusNotFound, "404 page not found")
		return
	}

	c.Redirect(http.StatusFound, urlMap.OriginUrl)
}
