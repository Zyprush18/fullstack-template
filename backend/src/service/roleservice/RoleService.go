package roleservice

import (
	"github.com/Zyprush18/fullstack-template/backend/src/model/response"
	"github.com/Zyprush18/fullstack-template/backend/src/repository/rolerepo"
)

type RoleService interface {
	GetAll() ([]response.Role, error)
}

type claimRepoRole struct {
	repo rolerepo.RolesRepo
}

func NewConnectRepo(r rolerepo.RolesRepo) RoleService  {
	return &claimRepoRole{repo: r}
}

func (c *claimRepoRole) GetAll() ([]response.Role, error) {
	return c.repo.GetAll()
}