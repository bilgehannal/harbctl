package config

import (
	"encoding/base64"
	"github.com/bilgehannal/harbctl/internal/pkg/errors"
	"time"
)

type Context struct {
	Url      string `json:"url"`
	User     string `json:"user"`
	Password string `json:"password"`
}

func (c Context) Deadline() (deadline time.Time, ok bool) {
	//TODO implement me
	panic("implement me")
}

func (c Context) Done() <-chan struct{} {
	//TODO implement me
	panic("implement me")
}

func (c Context) Err() error {
	//TODO implement me
	panic("implement me")
}

func (c Context) Value(key any) any {
	//TODO implement me
	panic("implement me")
}

type ContextBuilder struct{}

func (ContextBuilder) CreateContext(url string, user string, password string) Context {
	return Context{
		Url:      url,
		User:     user,
		Password: base64.StdEncoding.EncodeToString([]byte(password)),
	}
}

func (c Context) GetApiUrl() string {
	return c.Url + "/api"
}

func (c Context) DecodePassword() string {
	pass, err := base64.StdEncoding.DecodeString(c.Password)
	if err != nil {
		errors.FatalPanic("Password decode error", err)
	}
	return string(pass)
}
