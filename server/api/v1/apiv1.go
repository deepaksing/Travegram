package v1

import (
	"github.com/deepaksing/Travegram/store"
	"github.com/labstack/echo/v4"
)

type ApiV1Service struct {
	store *store.Store
}

func NewApiv1Service(store *store.Store) *ApiV1Service {
	return &ApiV1Service{
		store: store,
	}
}

func (a *ApiV1Service) Register(rootGroup *echo.Group) {
	apiv1Group := rootGroup.Group("/api/v1")

	// apiv1Group.Use(func(next echo.HandlerFunc) echo.HandlerFunc {
	// 	return JWTMiddleware(a, next, "travegram")
	// })

	//call other services from apiv1
	a.RegisterUserRoute(apiv1Group)

}
