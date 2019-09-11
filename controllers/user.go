package controllers

import (
	"encoding/json"
	"net/http"
	"regexp"
	"strconv"

	"github.com/eliises/gowebservice/models"
)

type userController struct {
	userIDPattern *regexp.Regexp
}

func (uc userController) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path == "/users" {
		switch r.Method {
		case http.MethodGet:
			uc.getAll(w, r)

		case http.MethodPost:
			uc.post(w, r)
		default:
			w.WriteHeader(http.StatusNotImplemented)
		}
	} else {
		matches := uc.userIDPattern.FindStringSubmatch(r.URL.Path)
		if len(matches) == 0 {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		id, err := strconv.Atoi(matches[1])
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		switch r.Method {
		case http.MethodGet:
			uc.get(id, w)
		case http.MethodPut:
			uc.put(id, w, r)
		case http.MethodDelete:
			uc.delete(id, w)
		default:
			w.WriteHeader(http.StatusNotImplemented)
		}
	}
}

func (uc *userController) getAll(w http.ResponseWriter, r *http.Request) {
	encodeResponseAsJson(models.GetUsers(), w)
}

func (uc *userController) get(id int, w http.ResponseWriter) {
	user, err := models.GetUserById(id)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		// w.Write([]byte(err.Error()))
		return
	}
	encodeResponseAsJson(user, w)
}

func (uc *userController) post(w http.ResponseWriter, r *http.Request) {
	user, err := uc.parseRequest(r)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Could not parse user object"))
		return
	}
	user, err = models.AddUser(user)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	encodeResponseAsJson(user, w)
}

func (uc *userController) put(id int, w http.ResponseWriter, r *http.Request) {
	user, err := uc.parseRequest(r)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Could not parse User object"))
		return
	}
	if id != user.ID {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("ID of submitted user must match ID in URL"))
		return
	}
	encodeResponseAsJson(user, w)
}

func (uc *userController) delete(id int, w http.ResponseWriter) {
	err := models.RemoveUserById(id)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	w.WriteHeader(http.StatusOK)
}

func (uc *userController) parseRequest(r *http.Request) (models.User, error) {
	dec := json.NewDecoder(r.Body)

	var user models.User
	err := dec.Decode(&user)
	if err != nil {
		return models.User{}, err
	}
	return user, nil
}

func newUserController() *userController {
	return &userController{
		userIDPattern: regexp.MustCompile(`^/users/(\d+)/?`),
	}
}
