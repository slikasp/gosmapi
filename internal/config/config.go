package config

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
	User SMuser
	Core SMserver
}

func (c *Config) SetAdmin(name, token string) Config {
	c.User = SMuser{
		Name:     name,
		Token:    token,
		UserRole: ADMIN,
	}
	return *c
}

func (c *Config) SetCore(address string) Config {
	c.Core = SMserver{
		Address:    address,
		ServerRole: CORE,
	}
	return *c
}
