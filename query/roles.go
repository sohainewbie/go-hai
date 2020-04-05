package query

import (
	"github.com/sohainewbie/go-hai/schemas"
	"github.com/sohainewbie/go-hai/utils"
)

var (
	listRoleType = []string{}
	dataSetRole  = make(map[string]uint64)
)

func GetRoleByType(roleType string) (role schemas.RolesSchema) {
	utils.GetInstanceMysqlDb().Model("roles").Where("(type = ? ) and deleted_at is null", roleType).First(&role)
	return
}

func SetRole(payload schemas.UsersRoleSchema) error {
	return utils.GetInstanceMysqlDb().Table("users_role").Where("user_id = ?  and deleted_at is null", payload.UserID).Update(&payload).Error
}
