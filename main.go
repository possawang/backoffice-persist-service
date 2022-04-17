package main

import (
	"github.com/possawang/backoffice-persist-service/entity"
	"github.com/possawang/go-persist-lib-common/startup"
)

func main() {
	component := startup.PersistMainComponent{
		Models: models(),
	}
	startup.StartingPersistService(component)
}

func models() []interface{} {
	var model []interface{}
	model = append(model, entity.Role{}, entity.Allowed{}, entity.User{})
	return model
}
