package handlers

import (
	"encoding/json"
	"github.com/timur-pro/ueba-get-items/internal/models"
	"log"
	"net/http"
)

type GetItemHandler struct {
	repo UebaRepository
}

// New создание экземпляра репозитория
func New(repo UebaRepository) *GetItemHandler {
	return &GetItemHandler{
		repo: repo,
	}
}

// Get обработчик запрос ручки get-items
func (h *GetItemHandler) Get(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	// Так как нет определенного контракта на эндоинт,
	// то обработаем запрос вида GET /get-items?id=1&id=3
	err := r.ParseForm()
	if err != nil {
		log.Printf("cannot parse request params: %s\n", err)
		http.Error(
			w,
			"cannot parse request params",
			http.StatusBadRequest,
		)
		return
	}

	ids := r.Form["id"]
	if len(ids) == 0 {
		log.Printf("nothing to do")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	result := []models.Record{}
	for _, id := range ids {
		record, err := h.repo.GetItem(id)
		if err != nil {
			log.Printf("id %s not found", id)
			continue
		}
		result = append(result, record)
	}

	err = json.NewEncoder(w).Encode(result)
	if err != nil {
		log.Printf("error encoding result: %s\n", err)
		w.WriteHeader(http.StatusInternalServerError)
	}
}
