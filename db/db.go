package db

import (
	"go-basic-server/models"
)

var (
	Userdb = make(map[string]models.User)
)
