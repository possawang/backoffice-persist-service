package main

import (
	"github.com/possawang/backoffice-persist-service/endpoints"
	"github.com/possawang/backoffice-persist-service/entity"
	"github.com/possawang/backoffice-persist-service/mdw"
	"github.com/possawang/go-persist-lib-common/startup"
	"github.com/possawang/go-service-lib-common/routerutils"
)

func main() {
	component := startup.PersistMainComponent{
		Models:    models(),
		Endpoints: router(),
	}
	startup.StartingPersistService(component)
}

func router() map[string]routerutils.Endpoint {
	route := make(map[string]routerutils.Endpoint)
	route["/getuserdata"] = routerutils.Endpoint{
		Method:    "POST",
		Mdw:       mdw.GetUserData,
		Execution: endpoints.GetUserData,
	}
	return route
}

func models() []interface{} {
	var model []interface{}
	model = append(model, entity.Role{}, entity.Allowed{}, entity.User{})
	return model
}
