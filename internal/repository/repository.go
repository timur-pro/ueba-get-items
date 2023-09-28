package repository

import (
	"encoding/csv"
	"fmt"
	"github.com/timur-pro/ueba-get-items/internal/models"
	"io"
	"log"
	"os"
)

// Repository структура репозитория
// Использую map для поиска записи по id за приблизительно О(1)
type Repository struct {
	idColIndex int
	ueba       map[string]models.Record
}

// New Создание экземпляра репозитория
func New(filePath string, idColIndex int) *Repository {
	f, err := os.Open(filePath)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	reader := csv.NewReader(f)

	ueba := make(map[string]models.Record)
	header, err := reader.Read()
	if err != nil {
		log.Fatal("cannot parse header", err)
	}

	for {
		csvRecord, err := reader.Read()
		if err != nil {
			if err != io.EOF {
				log.Fatal(err)
			}
			break
		}
		record := make(map[string]string)
		for i, v := range csvRecord {
			record[header[i]] = v
		}
		ueba[record[header[idColIndex]]] = record
	}
	return &Repository{
		idColIndex: idColIndex,
		ueba:       ueba,
	}
}

func (r *Repository) GetItem(id string) (models.Record, error) {
	v, ok := r.ueba[id]
	if !ok {
		return nil, fmt.Errorf("item not found")
	}
	return v, nil
}
