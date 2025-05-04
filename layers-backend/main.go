package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"layersapi/controllers"
	"layersapi/repositories/memory"
	"layersapi/services"
	"net/http"
)

func main() {
	r := mux.NewRouter()

	userRepository := memory.NewUserRepository() //para el csv poner el csv modifcando el nombre del paquete osea solo poner csv
	userService := services.NewUserService(userRepository)
	userController := controllers.NewUserController(*userService)

	//para obtener todos los usuarios
	r.HandleFunc("/users", userController.GetAllUsersHandler).Methods(http.MethodGet)
	//para obtener un usuario por id
	r.HandleFunc("/users/{id}", userController.GetUserByIdHandler).Methods(http.MethodGet)
	//para crear un usuario
	r.HandleFunc("/users", userController.CreateUserHandler).Methods(http.MethodPost)
	//para actualizar el usuario
	r.HandleFunc("/users/{id}", userController.UpdateUserHandler).Methods(http.MethodPut)
	//para eliminar el usuario
	r.HandleFunc("/users/{id}", userController.DeleteUserHandler).Methods(http.MethodDelete)

	fmt.Println("Listening on port 8080")
	http.ListenAndServe(":8080", r)

}
