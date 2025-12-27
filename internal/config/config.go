package config

import (
	"net/http"
	"time"
)

// user type enums
type userRole string

const (
	USER   userRole = "user"
	EXPERT userRole = "expert"
	ADMIN  userRole = "admin"
)

// server type enums
type serverRole string

const (
	CORE  serverRole = "core"
	PROXY serverRole = "proxy"
)

type SMuser struct {
	Name     string
	Token    string
	UserRole userRole
}

type SMserver struct {
	Address    string
	ServerRole serverRole
}

// config to be passed around in functions, user.token and core.address used in making http requests
type Config struct {
	User       SMuser
	Core       SMserver
	HttpClient *http.Client
}

func (c *Config) SetAdmin(name, token string) {
	c.User = SMuser{
		Name:     name,
		Token:    token,
		UserRole: ADMIN,
	}
}

func (c *Config) SetCore(address string) {
	c.Core = SMserver{
		Address:    address,
		ServerRole: CORE,
	}
}

func (c *Config) NewClient(timeoutSeconds int) {
	c.HttpClient = &http.Client{
		Timeout: time.Duration(timeoutSeconds) * time.Second,
	}
}
