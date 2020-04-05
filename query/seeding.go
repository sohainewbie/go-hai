package query

import (
	"github.com/sohainewbie/go-hai/schemas"
	"github.com/sohainewbie/go-hai/utils"
)

func SeedDataRoles() error {

	var roleID uint64
	roles := []schemas.RolesSchema{}
	params := []schemas.RolesSchema{
		{
			Name:        "user",
			Type:        "user",
			Description: "this role type for access level admin",
		},
		{
			Name:        "admin",
			Type:        "admin",
			Description: "this role type for access level admin",
		},
	}

	for _, val := range params {
		roleID += 1
		listRoleType = append(listRoleType, val.Type)
		dataSetRole[val.Type] = roleID
	}

	utils.GetInstanceMysqlDb().Model("roles").Where("type IN (?)", listRoleType).First(&roles)
	if len(roles) == 0 {
		for _, val := range params {
			if err := utils.GetInstanceMysqlDb().Model("roles").Create(&val).Error; err != nil {
				return err
			}
		}
	}

	return nil
}

func SeedDataUsers() {
	//create demo account
	listUser := []schemas.UsersSchema{
		{
			Name:        "Kucing Liar",
			Email:       "kucing@liar.com",
			Password:    "password",
			PhoneNumber: "+6281288831111",
			Address:     "Jl. Pertenakan Kucing, B.88/no 90",
		},
		{
			Name:        "Kucing Garong",
			Email:       "kucing.garong@gmail.com",
			Password:    "password",
			PhoneNumber: "+628128883222",
			Address:     "Jl. Pertenakan Kucing, B.88/no 90",
		},
		{
			Name:        "Kucing Oren",
			Email:       "kucing.oren@gmail.com",
			Password:    "password",
			PhoneNumber: "+628128883333",
			Address:     "Jl. Pertenakan Kucing, B.88/no 90",
		},
	}

	for _, user := range listUser {
		dataUser := schemas.UsersSchema{}
		dataUser = user
		dataUser.Password, _ = utils.HashPassword(user.Password)

		tx := utils.GetInstanceMysqlDb().Begin()
		tx = tx.Table(`users`).Where("(phone_number = ?  and email = ? ) and deleted_at is null", dataUser.PhoneNumber, dataUser.Email)
		if err := tx.First(&dataUser).Error; err != nil {
			if err.Error() == "record not found" {
				tx.Table(`users`).Save(&dataUser)
				// create relation with table role_user
				usersRoleSchema := schemas.UsersRoleSchema{
					RoleID: 1,
					UserID: dataUser.ID,
				}
				tx.Table("users_role").Create(&usersRoleSchema)
				if err := tx.Commit(); err != nil {
					tx.Rollback()
				}
			}
		}
	}
}
