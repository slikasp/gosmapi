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

type user struct {
	name     string
	token    string
	userRole userRole
}

type server struct {
	address    string
	serverRole serverRole
}

// config to be passed around in functions, user.token and core.address used in making http requests
type config struct {
	user user
	core server
}

func (c *config) SetAdmin(name, token string) config {
	c.user = user{
		name:     name,
		token:    token,
		userRole: ADMIN,
	}
	return *c
}

func (c *config) SetCore(address string) config {
	c.core = server{
		address:    address,
		serverRole: CORE,
	}
	return *c
}
