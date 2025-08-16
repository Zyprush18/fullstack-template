package response

import (
	"github.com/Zyprush18/fullstack-template/backend/src/database/postgre/entity"

)

type Role struct {
	Role string `json:"role"`
	entity.Model
}