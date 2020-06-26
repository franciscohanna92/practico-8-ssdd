package sample

import (
	user "../../pkg"
)

var Users = map[string]*user.User{
	"01D3XZ3ZHCP3KG9VT4FGAD8KDR": &user.User{
		Id:    "01D3XZ3ZHCP3KG9VT4FGAD8KDR",
		Email: "pedro@example.com",
		Name:  "Pedro",
		Age:   18,
	},
	"01D3XZ89NFJZ9QT2DHVD462AC2": &user.User{
		Id:    "01D3XZ89NFJZ9QT2DHVD462AC2",
		Email: "maria@example.com",
		Name:  "María",
		Age:   48,
	},
	"01D3XZ8JXHTDA6XY05EVJVE9Z2": &user.User{
		Id:    "01D3XZ8JXHTDA6XY05EVJVE9Z2",
		Email: "juan@example.com",
		Name:  "Juan",
		Age:   32,
	},
}
