package bootstrap

import (
	"github.com/joyleewei/goblog/routes"

	"github.com/gorilla/mux"
)

// SetupRoute 路由初始化
func SetupRoute() *mux.Router {
	router := mux.NewRouter()
	routes.RegisterWebRoutes(router)
	return router
}