# go-hai
Boilerplate golang with MVC structure with Echo framework.

###HOW TO START
1. rename config.json.example into config.json
2. go run *.go
3. if you want to build just make build

### A typical top-level directory layout
	get-getgo/
    .
    ├── ...
    ├── service/*.go          # you can put another service in here
    │   └── email.go          # this is example    
    ├── controller/*.go       # you can put all validation from router before you processing data into database
    │   └── index.go          # this is example
    └── query/*.go            # this folder for execute the query from controller
    │   └── user.go           # this is example    
    └── schemas/*.go          # this folder for define all your models
    │   └── users.go          # this is example        
    └── utils/*.go            # this folder all function like helper for your backend, it's handle like for http req/res
    └── test/*.py             # this folder for testing your api.
    └── vendor/               # this folder contains all library and make sure it should be added at .gitignore
    └── config.json.example   # this is config for the application you can rename .example into config.json
    └── main_http.go          # the router path you can added another path in here


*fell free if you want to contribute
