package users

import (
	"fmt"
	"time"

	"github.com/godesde0/modelos"
)

func AltaUsuario() {
	u := new(modelos.User)
	u.AddUser(10, "Pablo", time.Now(), true)
	v := new(modelos.User)
	v.AddUser(11, "Jose", time.Now(), true)
	w := modelos.AddUser2(12, "Pablo Tilotta", time.Now(), true)
	fmt.Println(u)
	fmt.Println(v)
	fmt.Println(w)
}
