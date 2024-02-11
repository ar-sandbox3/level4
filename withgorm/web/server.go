package web

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/ar-sandbox3/level4/withgorm/db/models"
	"gorm.io/gorm"
)

func Handler(conn *gorm.DB) http.Handler {
	departments := models.DepartmentStore{DB: conn}

	mux := http.NewServeMux()
	mux.HandleFunc("/departments", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodPost:
			name := r.FormValue("name")
			if name == "" {
				http.Error(w, "missing name", http.StatusBadRequest)
				return
			}
			if err := departments.Create(name); err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			w.WriteHeader(http.StatusCreated)
		case http.MethodGet:
			id := r.FormValue("id")
			if id == "" {
				departments, err := departments.GetAll()
				if err != nil {
					http.Error(w, err.Error(), http.StatusInternalServerError)
					return
				}
				b, err := json.Marshal(departments)
				if err != nil {
					http.Error(w, err.Error(), http.StatusInternalServerError)
					return
				}
				w.Header().Set("Content-Type", "application/json")
				w.Write(b)
				return
			}

			departmentID, err := strconv.Atoi(id)
			if err != nil {
				http.Error(w, err.Error(), http.StatusBadRequest)
				return
			}
			department, err := departments.Get(int64(departmentID))
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			b, err := json.Marshal(department)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			w.Header().Set("Content-Type", "application/json")
			w.Write(b)

		default:
			http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		}
	})
	return mux
}
