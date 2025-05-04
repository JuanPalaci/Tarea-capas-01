package csv

import (
	"encoding/csv"
	"errors"
	"layersapi/entities"
	"os"
	"time"
)

type UserRepository struct{}

func NewUserRepository() *UserRepository {
	return &UserRepository{}
}

func (u UserRepository) GetAll() ([]entities.User, error) {
	file, err := os.Open("C:\\Users\\juanc\\OneDrive\\Documentos\\Capas\\Modificado\\layers-backend\\data\\data.csv") //poner la ruta largota
	if err != nil {
		return []entities.User{}, err
	}
	defer file.Close()

	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		return []entities.User{}, err
	}

	var result []entities.User

	for i, record := range records {
		if i == 0 {
			continue
		}

		createdAt, _ := time.Parse(time.RFC3339, record[3])
		updatedAt, _ := time.Parse(time.RFC3339, record[4])
		meta := entities.Metadata{
			CreatedAt: createdAt.String(),
			UpdatedAt: updatedAt.String(),
			CreatedBy: record[5],
			UpdatedBy: record[6],
		}
		result = append(result, entities.NewUser(record[0], record[1], record[2], meta))
	}

	return result, nil
}

func (u UserRepository) GetById(id string) (entities.User, error) {
	file, err := os.Open("C:\\Users\\juanc\\OneDrive\\Documentos\\Capas\\Modificado\\layers-backend\\data\\data.csv")
	if err != nil {
		return entities.User{}, err
	}
	defer file.Close()

	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		return entities.User{}, err
	}

	for i, record := range records {
		if i == 0 {
			continue
		} else if record[0] == id {

			createdAt, _ := time.Parse(time.RFC3339, record[3])
			updatedAt, _ := time.Parse(time.RFC3339, record[4])
			meta := entities.Metadata{
				CreatedAt: createdAt.String(),
				UpdatedAt: updatedAt.String(),
				CreatedBy: record[5],
				UpdatedBy: record[6],
			}
			return entities.NewUser(record[0], record[1], record[2], meta), nil
		}

	}

	return entities.User{}, errors.New("user not found")
}

func (u UserRepository) Create(user entities.User) error {
	file, err := os.OpenFile("C:\\Users\\juanc\\OneDrive\\Documentos\\Capas\\Modificado\\layers-backend\\data\\data.csv", os.O_APPEND|os.O_WRONLY, os.ModeAppend)
	if err != nil {
		return err
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	newUser := []string{
		user.Id,
		user.Name,
		user.Email,
		user.Metadata.CreatedAt,
		user.Metadata.UpdatedAt,
		"webapp",
		"webapp",
	}

	if err := writer.Write(newUser); err != nil {
		return err
	}

	return nil
}

func (u UserRepository) Update(id, name, email string) error {
	const path = "C:\\Users\\juanc\\OneDrive\\Documentos\\Capas\\Modificado\\layers-backend\\data\\data.csv"
	//el path pa el archivo csv
	// c lee
	records, err := readAllUsersFromCSV(path)
	if err != nil {
		return err
	}
	//aqui buscamos el usuario
	found := false
	for i, record := range records {
		if i == 0 {
			continue
		}
		// Comparamos el id
		// record[0] es el id
		if record[0] == id {
			// Actualizamos el nombre, correo y fecha
			records[i][1] = name
			records[i][2] = email
			records[i][4] = time.Now().Format(time.RFC3339) // updatedAt
			records[i][6] = "webapp"                        // updatedBy
			found = true
			break
		}
	}
	//esto es por si no esta el usuario
	if !found {
		return errors.New("user not found")
	}
	// escribimos el archivo de  nuevo
	return writeAllUsersToCSV(path, records)
}

// pa eliminar el usuario
func (u UserRepository) Delete(id string) error {
	const path = "C:\\Users\\juanc\\OneDrive\\Documentos\\Capas\\Modificado\\layers-backend\\data\\data.csv"
	//leemos el archivo
	records, err := readAllUsersFromCSV(path)
	if err != nil {
		return err
	}
	// buscamos el usuario
	var newRecords [][]string
	found := false
	// c recorre igual al del update pero guardando los que no son el usuario
	for i, record := range records {
		if i == 0 {
			newRecords = append(newRecords, record)
			continue
		}
		if record[0] == id {
			found = true
			continue // omitimos este usuario
		}
		newRecords = append(newRecords, record)
	}

	if !found {
		return errors.New("user not found")
	}
	// escribimos el archivo de nuevo
	return writeAllUsersToCSV(path, newRecords)
}

// funciones para leer y escribir el archivo
func readAllUsersFromCSV(path string) ([][]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	reader := csv.NewReader(file)
	return reader.ReadAll()
}

func writeAllUsersToCSV(path string, records [][]string) error {
	file, err := os.Create(path)
	if err != nil {
		return err
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	return writer.WriteAll(records)
}
