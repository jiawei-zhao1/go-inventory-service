package config

type MySQL struct {
	Host, Port, User, Password, Name string
	MaxOpenConn,MaxIdleConn int
}