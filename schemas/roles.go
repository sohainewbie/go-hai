package schemas

import (
	"github.com/sohainewbie/go-hai/utils"
)

type RolesSchema struct {
	utils.BaseModel
	Name        string `json:"name"`
	Type        string `json:"type"`
	Description string `json:"description"`
}

func (RolesSchema) TableName() string {
	return "roles"
}
