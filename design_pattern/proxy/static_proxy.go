package main

import (
	"fmt"
	"time"
)

func main(){
	proxyUser := NewUserProxy()
	err := proxyUser.Login("user", "pwd")
	if err != nil {
		fmt.Println(err)
	}
}

type IUser interface {
	Login(username, password string) error
}

type User struct {

}

func (u *User) Login(user, pwd string) error {
	time.Sleep(1*time.Second)
	return nil
}

type UserProxy struct {
	user *User
}

func NewUserProxy() *UserProxy {
	return &UserProxy{
		user: &User{},
	}
}

func (u *UserProxy) Login(user, pwd string) error {
	s := time.Now()

	if err := u.user.Login(user, pwd); err != nil {
		return err
	}

	fmt.Printf("time spend on logining: %v\n", time.Now().Sub(s))
	return nil
}
