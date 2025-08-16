package authservice

import (
	"context"
	"errors"
	"fmt"
	"log"

	"github.com/Zyprush18/fullstack-template/backend/src/helper"
	"github.com/Zyprush18/fullstack-template/backend/src/model/request"
	"github.com/Zyprush18/fullstack-template/backend/src/repository/authrepo"
	"github.com/redis/go-redis/v9"
)

type AuthService interface {
	Regist(req *request.User) error
	Login(req *request.Login) (string, error)
	Logout(e string) error
}

type claimRepoAuth struct {
	repo authrepo.AuthRepo
	rdb *redis.Client
	jwtkey string
}

var ctx = context.Background()

func NewConnectRepo(ar authrepo.AuthRepo,r *redis.Client,jkey string) AuthService  {
	return &claimRepoAuth{repo: ar, rdb: r,jwtkey: jkey}
}

func (c *claimRepoAuth) Regist(req *request.User) error {
	pwhash,err:= helper.HashingPass(req.Password)
	if err != nil {
		log.Println(err.Error())
		return err
	}
	req.Password = string(pwhash)
	req.RoleId = 1
	return c.repo.Registration(req)
}


func (c *claimRepoAuth) Login(req *request.Login) (string, error) {

	data, err:= c.repo.Login(req)
	if err != nil {
		return "", err
	}
	// cek pass
	if err:= helper.DecryptPass(data.Password, req.Password);err != nil {
		return "", errors.New("invalid password")
	}

	// generate token
	token, err := helper.GenerateJwtToken(c.jwtkey, data)
	if err != nil {
		return "",err
	}

	// save in redis
	if err:= c.rdb.Set(ctx, data.Email, token, 0).Err();err != nil {
		fmt.Println(err.Error())
		return "",err
	}

	return token,nil
}

func (c *claimRepoAuth) Logout(e string) error {
	mail, err := c.repo.Logout(e)
	
	if err != nil {
		return err
	}
	// delete di token di redis
	if err:=c.rdb.Del(context.Background(), mail).Err();err != nil{
		return err
	}

	return nil
}