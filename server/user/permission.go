package user

type Permission uint

var CurrentIndex uint = 0

const (
	CreateUser Permission = iota
	ReadUser
)
