package db

import (
	"database/sql"
	"github.com/didi/gendry/builder"
	"github.com/didi/gendry/manager"
	"github.com/didi/gendry/scanner"
	_ "github.com/go-sql-driver/mysql"
	"github.com/spf13/viper"
	"log"
	"time"
	"tinyurl/model"
)

var db *sql.DB

func init() {
	host := viper.GetString("database.host")
	port := viper.GetInt("database.port")
	user := viper.GetString("database.user")
	password := viper.GetString("database.password")
	dbName := viper.GetString("database.dbname")
	db1, err := manager.
		New(dbName, user, password, host).
		Set(
			manager.SetCharset("utf8"),
			manager.SetAllowCleartextPasswords(true),
			manager.SetInterpolateParams(true),
			manager.SetTimeout(1*time.Second),
			manager.SetReadTimeout(1*time.Second),
			manager.SetParseTime(true),
		).Port(port).Open(true)
	if err != nil {
		log.Fatalln(err)
	}
	db = db1
	db.SetConnMaxLifetime(time.Minute * 3)
	db.SetMaxOpenConns(5)
	db.SetMaxIdleConns(5)
}

func CreateLink(urlMap *model.TinyUrlMap) error {
	table := "tinyurl_maps"
	var data []map[string]interface{}
	data = append(data, map[string]interface{}{
		"origin_url":   urlMap.OriginUrl,
		"short_url":    urlMap.ShortUrl,
		"created_time": urlMap.CreatedTime,
	})

	cond, vals, err := builder.BuildInsert(table, data)
	if err != nil {
		log.Println(err)
		return err
	}

	result, err := db.Exec(cond, vals...)
	log.Println(result)

	if err != nil {
		log.Println(err)
		return err
	}

	return nil
}

func GetLinkByShort(shortUrl string) *model.TinyUrlMap {
	where := map[string]interface{}{
		"short_url": shortUrl,
	}
	table := "tinyurl_maps"

	selectFields := []string{"id", "origin_url", "short_url", "created_time"}
	cond, vals, err := builder.BuildSelect(table, where, selectFields)
	if err != nil {
		log.Println(err)
		return nil
	}

	rows, err := db.Query(cond, vals...)
	if err != nil {
		log.Println("err :", err)
		return nil
	}

	var urlMap model.TinyUrlMap
	err = scanner.ScanClose(rows, &urlMap)
	if err != nil {
		log.Println(err)
		return nil

	}

	return &urlMap
}
