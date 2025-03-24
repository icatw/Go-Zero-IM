package main

type User struct {
	Id    string
	Name  string
	Phone string
}

var users = map[string]*User{
	"1": &User{"1", "Joe", "123456789"},
	"2": &User{"2", "Jane", "987654321"},
}
