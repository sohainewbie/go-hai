.PHONY : format getgo-osx getgo-linux create vendor

format:
	find . -name "*.go" -not -path "./vendor/*" -not -path ".git/*" | xargs gofmt -s -d -w

getgo-osx: main.go
	GOOS=darwin GOARCH=amd64 go build -ldflags '-s -w' -o $@

getgo-linux: main.go
	GOOS=linux GOARCH=amd64 go build -ldflags '-s -w' -o $@

create:
	echo "package controller">>controller/$(filter-out $@,$(MAKECMDGOALS)).go
	echo "package query">>controller/$(filter-out $@,$(MAKECMDGOALS)).go
install:
    go get -u github.com/jinzhu/gorm
    go get -u github.com/jinzhu/gorm/dialects/postgres
    go get -u github.com/go-redis/redis
    go get -u github.com/labstack/echo
    go get -u github.com/dgrijalva/jwt-go
vendor:
	@dep ensure -v