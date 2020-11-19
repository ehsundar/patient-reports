package user

import (
	"github.com/ehsundar/reports/internal/helper"
	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
	"net/http"
)

type Controller struct {
}

func NewController() *Controller {
	return &Controller{}
}

func (ctrl *Controller) Register(mx *mux.Router) {
	mx.HandleFunc("/create", ctrl.getCreate).Methods("GET")
	mx.HandleFunc("/create", ctrl.postCreate).Methods("POST")

	mx.HandleFunc("/edit/{userID:[0-9]+}", ctrl.getEdit).Methods("GET")
	mx.HandleFunc("/edit/{userID:[0-9]+}", ctrl.postEdit).Methods("POST")
}

func (ctrl *Controller) getCreate(w http.ResponseWriter, r *http.Request) {
	data := make(map[string]interface{})
	data["organizations"] = []string{"Pegah Salamat", "Erfan Hospital", "Test Hospital"}
	err := helper.RenderTemplate("./templates/user/create.html", data, w)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
	w.WriteHeader(http.StatusOK)
}

func (ctrl *Controller) postCreate(w http.ResponseWriter, r *http.Request) {
	_, err := w.Write([]byte("test"))
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
	http.Redirect(w, r, "/user/create", http.StatusTemporaryRedirect)
}

func (ctrl *Controller) getEdit(w http.ResponseWriter, r *http.Request) {
	data := make(map[string]interface{})
	data["organizations"] = []string{"Pegah Salamat", "Erfan Hospital", "Test Hospital"}
	data["formData"] = map[string]interface{}{
		"username":     "ehsandar",
		"password":     "***",
		"fullName":     "Amir Ehsandar",
		"isActive":     false,
		"organization": "Erfan Hospital",
	}
	err := helper.RenderTemplate("./templates/user/edit.html", data, w)
	if err != nil {
		log.WithError(err).Error("failed to render template")
		w.WriteHeader(http.StatusInternalServerError)
	}
	w.WriteHeader(http.StatusOK)
}

func (ctrl *Controller) postEdit(w http.ResponseWriter, r *http.Request) {
	_, err := w.Write([]byte("test"))
	if err != nil {
		log.WithError(err).Error("failed to render template")
		w.WriteHeader(http.StatusInternalServerError)
	}
	http.Redirect(w, r, "/user/create", http.StatusTemporaryRedirect)
}
