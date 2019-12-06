package main

// https://icyapril.com/go/programming/2017/12/17/object-orientation-in-go.html

import (
	"fmt"
	"strings"
)

type permissions struct {
	admin bool
	suspended bool
}

type user struct {
	name  string
	access permissions
}

func (u user) FirstName() string {
	names := strings.Fields(u.name)
	if len(names) > 0 {
		return names[0]
	}
	return "Unnamed"
}

func (u user) IsAdmin() bool {
	return u.access.admin
}

func (u user) IsSuspended() bool {
	return u.access.suspended
}

func Init() {
	newAdminStatus := permissions{admin: true, suspended: false}
	paul := user{name: "Paul Smithson", access: newAdminStatus}
	fmt.Printf("First Name: %s\n", paul.FirstName())
	fmt.Printf("Admin? %t\n", paul.IsAdmin())
	fmt.Printf("Suspended? %t\n", paul.IsSuspended())
}