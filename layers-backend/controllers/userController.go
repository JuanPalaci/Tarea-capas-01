package controllers

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"layersapi/entities/dto"
	"layersapi/services"
	"net/http"
)

type UserController struct {
	userService services.UserService
}

func NewUserController(userService services.UserService) *UserController {
	return &UserController{userService: userService}
}

func (u UserController) GetAllUsersHandler(w http.ResponseWriter, r *http.Request) {
	resData, err := u.userService.GetAll()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
	}
	res, err := json.Marshal(resData)
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	w.Write(res)
}

func (u UserController) GetUserByIdHandler(w http.ResponseWriter, r *http.Request) {
	//obtener el id de la url
	vars := mux.Vars(r)
	id := vars["id"]
	// Llamar al servicio
	resData, err := u.userService.GetById(id)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	res, err := json.Marshal(resData)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(res)
}

func (u UserController) UpdateUserHandler(w http.ResponseWriter, r *http.Request) {
	//obtener el id de la url
	vars := mux.Vars(r)
	id := vars["id"]
	// Leer el cuerpo del request
	var user dto.UpdateUser
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}
	// Llamar al servicio
	err = u.userService.Update(id, user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	// Devolver una respuesta de éxito
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(`{"message": "user updated successfully"}`))
}

// creando usuario
func (u UserController) CreateUserHandler(w http.ResponseWriter, r *http.Request) {
	// Leer el cuerpo del request
	var user dto.CreateUser
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}
	// Llamar al servicio
	err = u.userService.Create(user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	// Devolver una respuesta de éxito
	w.WriteHeader(http.StatusCreated)
	w.Write([]byte(`{"message": "user created successfully"}`))
}

func (u UserController) DeleteUserHandler(w http.ResponseWriter, r *http.Request) {
	//obtener el id de la url
	vars := mux.Vars(r)
	id := vars["id"]
	// Llamar al servicio
	err := u.userService.Delete(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	// Devolver una respuesta de éxito
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(`{"message": "user deleted successfully"}`))
}
