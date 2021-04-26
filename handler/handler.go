package handler

import (
	"fmt"
	"github.com/gin-contrib/location"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"time"
	"tinyurl/bf"
	"tinyurl/cache"
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
	shortUrl := c.Param("link")
	log.Println("short link : ", shortUrl)

	if bf.Instance().ExistString(shortUrl) {
		var urlMap *model.TinyUrlMap
		cacheUrl := cache.Instance().Get(cachekey(shortUrl))
		log.Println("cache origin : ", cacheUrl)

		if cacheUrl != "" {
			urlMap = &model.TinyUrlMap{
				OriginUrl: cacheUrl,
			}
		} else {
			urlMap = db.GetLinkByShort(shortUrl)
		}

		if urlMap == nil {
			c.String(http.StatusNotFound, "404 page not found")
			return
		}

		if cacheUrl == "" {
			cache.Instance().Set(cachekey(shortUrl), urlMap.OriginUrl, 24*time.Hour)
		}
		c.Redirect(http.StatusFound, urlMap.OriginUrl)
	} else {
		c.String(http.StatusNotFound, "404 page not found")
	}

}

func cachekey(token string) string {
	return "tinyurl_token_" + token
}
