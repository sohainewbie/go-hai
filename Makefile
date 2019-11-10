.PHONY : format gohai-osx gohai-linux create vendor

format:
	find . -name "*.go" -not -path "./vendor/*" -not -path ".git/*" | xargs gofmt -s -d -w

gohai-osx: main.go
	GOOS=darwin GOARCH=amd64 go build -ldflags '-s -w' -o $@

gohai-linux: main.go
	GOOS=linux GOARCH=amd64 go build -ldflags '-s -w' -o $@

create:
	echo "package controller">>controller/$(filter-out $@,$(MAKECMDGOALS)).go
	echo "package query">>controller/$(filter-out $@,$(MAKECMDGOALS)).go
install:
	go get -v github.com/jinzhu/gorm
	go get -v github.com/jinzhu/gorm/dialects/postgres
	go get -v github.com/jinzhu/gorm/dialects/mysql
	go get -v github.com/skip2/go-qrcode
	go get -v github.com/go-redis/redis
	go get -v github.com/labstack/echo
	go get -v github.com/dgrijalva/jwt-go	
vendor:
	@dep ensure -v