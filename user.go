package main

type User struct {
	id       int64  `json:"-"`
	login    string `json:"login"`
	password string `json:"password"`
}
