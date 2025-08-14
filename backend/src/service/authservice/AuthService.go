package authservice

import (
	"log"

	"github.com/Zyprush18/fullstack-template/backend/src/helper"
	"github.com/Zyprush18/fullstack-template/backend/src/model/request"
	"github.com/Zyprush18/fullstack-template/backend/src/repository/authrepo"
)

type AuthService interface {
	Regist(req *request.User) error
}

type claimRepoAuth struct {
	repo authrepo.AuthRepo
}

func NewConnectRepo(r authrepo.AuthRepo) AuthService  {
	return &claimRepoAuth{repo: r}
}

func (c *claimRepoAuth) Regist(req *request.User) error {
	pwhash,err:= helper.HashingPass(req.Password)
	if err != nil {
		log.Println(err.Error())
		return err
	}
	req.Password = string(pwhash)
	return c.repo.Registration(req)
}