package memory

import (
	"errors"
	"layersapi/data"
	"layersapi/entities"
)

type UserRepository struct{}

func NewUserRepository() *UserRepository {
	return &UserRepository{}
}

func (u UserRepository) GetAll() ([]entities.User, error) {
	res := data.Data
	return res, nil
}

func (u UserRepository) GetById(id string) (entities.User, error) {
	for _, v := range data.Data {
		if v.Id == id {
			return v, nil
		}
	}

	return entities.User{}, errors.New("user not found")
}

func (u UserRepository) Create(user entities.User) error {
	data.Data = append(data.Data, user)
	return nil
}

func (u UserRepository) Update(id, name, email string) error {
	for i, v := range data.Data {
		if v.Id == id {
			data.Data[i].Name = name
			data.Data[i].Email = email
			return nil
		}
	}

	return errors.New("user not found")
}

// elminando el usuario
func (u UserRepository) Delete(id string) error {
	// recorriendo los datos en memoria
	for i, v := range data.Data {
		// si el id coincide con el id que se quiere eliminar
		if v.Id == id {
			// eliminando el usuario
			data.Data = append(data.Data[:i], data.Data[i+1:]...)
			return nil
		}
	}
	return errors.New("user not found")
}
