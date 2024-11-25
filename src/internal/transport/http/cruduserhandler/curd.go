package cruduserhandler

import (
	"crud/model"
	"encoding/json"
	"net/http"
	"strconv"
)

func (h *handlerusercrud) Create(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	req := &model.Usermodel{}
	err := json.NewDecoder(r.Body).Decode(req)
	if err != nil {

		json.NewEncoder(w).Encode(err.Error())
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	if req.Email == "" || req.Name == "" || req.Pass == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	err = h.service.Create(*req)
	if err != nil {
		json.NewEncoder(w).Encode(err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)

}

func (h *handlerusercrud) Get(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	id := r.URL.Query().Get("id")
	if id == "" {
		json.NewEncoder(w).Encode("id is required")
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	if _, err := strconv.Atoi(id); err != nil {
		json.NewEncoder(w).Encode(err.Error())
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	user, err := h.service.Get(id)
	if err != nil {
		json.NewEncoder(w).Encode(err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(user)
	w.WriteHeader(http.StatusOK)
}

func (h *handlerusercrud) List(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	users, err := h.service.List()
	if err != nil {
		json.NewEncoder(w).Encode(err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(users)
	w.WriteHeader(http.StatusOK)

}

func (h *handlerusercrud) Update(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	req := &[]model.Usermodel{}
	err := json.NewDecoder(r.Body).Decode(req)

	if err != nil {
		json.NewEncoder(w).Encode(err.Error())
		w.WriteHeader(http.StatusNotFound)
		return
	}
	if len(*req) != 2 {
		json.NewEncoder(w).Encode("too many or less data")
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	old := (*req)[0]
	new := (*req)[1]

	if old.ID != new.ID && new.Email != "" && new.Name != "" && new.Pass != "" {
		json.NewEncoder(w).Encode("invalid data or id")
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	err = h.service.Update(old, new)
	if err != nil {

		json.NewEncoder(w).Encode(err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}

func (h *handlerusercrud) Delete(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	id := r.URL.Query().Get("id")
	if id == "" {
		json.NewEncoder(w).Encode("id is required")
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	if _, err := strconv.Atoi(id); err != nil {
		json.NewEncoder(w).Encode(err.Error())
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	req, err := h.service.Get(id)
	if err != nil {
		json.NewEncoder(w).Encode(err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	err = h.service.Delete(req)
	if err != nil {
		json.NewEncoder(w).Encode(err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)

}
