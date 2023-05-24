package main

import (
	_ "github.com/sakirsensoy/genv/dotenv/autoload"

	"go-app/lib/rbac"
)

func main() {
	rbac.New().AddRoleForUser("废墟", "superadmin")
	// var user model.User

	// user.ID = 1
	// user.Password = hash.HmacSha256Hex([]byte(config.HASH.HmacSha256Key), "admin123")
	// db.Conn.Save(&user)
}
