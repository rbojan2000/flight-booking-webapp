package model

type UserType int64

const (
	Guest         UserType = 0
	Authenticated UserType = 1
	Admin         UserType = 2
)
