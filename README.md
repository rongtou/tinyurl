# tinyurl

Todo

- refactor
- Logging
- benchmark
- profile
- simple attack defense

How to use

- create database table using `db.sql`
- cp config.sample.toml config.toml
- config the setting with your own environment
- go run .
- `curl -X POST -d 'url=TargetHost' http://localhost:8080/link/create`, and get the shorturl
- visit the short url and will redirect to the TargetHost
