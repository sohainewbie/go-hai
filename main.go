package main

import (
	"github.com/sohainewbie/go-hai/query"
	"github.com/sohainewbie/go-hai/schemas"
	"github.com/sohainewbie/go-hai/utils"
	"sync"
)

func main() {
	migrate() //migrate script
	seeding()

	wg := sync.WaitGroup{}
	wg.Add(1)
	go func() {
		HTTPServeMain() // http handler
		wg.Done()
	}()

	wg.Wait()
}

func migrate() {
	utils.GetInstanceMysqlDb().AutoMigrate(&schemas.UsersSchema{})
	utils.GetInstanceMysqlDb().AutoMigrate(&schemas.RolesSchema{})
	utils.GetInstanceMysqlDb().AutoMigrate(&schemas.UsersRoleSchema{})
}

func seeding() {
	query.SeedDataUsers()
	query.SeedDataRoles()
}
