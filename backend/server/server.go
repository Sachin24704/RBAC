package server

import (
	"backend/ent"
	"backend/handler"
	"backend/services"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

func StartServer(client *ent.Client) {
	router := mux.NewRouter()
	router.HandleFunc("/health",func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "server is running")
	}).Methods("GET")
	modelService := services.NewModelService(client)
	modelHandler := handler.NewModelHandler(modelService)
	router.HandleFunc("/models",modelHandler.ListModelsWithOperationsHandler).Methods("GET")
	roleHandler := handler.NewRoleHandler(services.NewRoleService(client))
	// add permission for role.
	router.HandleFunc("/roles/{role_id}/operations", roleHandler.AddOperationsToRoleHandler).Methods("POST")
	// delete permissions for role.
    router.HandleFunc("/roles/{role_id}/operations", roleHandler.RemoveOperationsFromRoleHandler).Methods("DELETE")
	permissionHandler := handler.NewPermissionHandler(services.NewPermissionService(client))
	router.HandleFunc("/users/{id}/permissions", permissionHandler.GetUserPermissions)
	log.Println("server started at 8000")
	c := cors.New(cors.Options{
        AllowedOrigins:   []string{"http://localhost:3000"},
        AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
        AllowedHeaders:   []string{"Content-Type", "Authorization"},
        AllowCredentials: true,
    })
    handlerWithCORS := c.Handler(router)
    log.Println("server started at 8000")
    log.Fatal(http.ListenAndServe(":8000", handlerWithCORS))
}
