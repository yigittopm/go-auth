package api

import (
	"encoding/json"
	"fmt"
	"go-auth/pkg/cache"
	"go-auth/pkg/model"
	"go-auth/pkg/service"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
)

type UserAPI struct {
	UserService service.UserService
	client      *cache.Client
}

func NewUserAPI(u service.UserService, c *cache.Client) UserAPI {
	return UserAPI{
		UserService: u,
		client:      c,
	}
}

func (u *UserAPI) FindAll() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		users, err := u.UserService.All()
		if err != nil {
			RespondWithError(w, http.StatusNotFound, err.Error())
			return
		}

		RespondWithJSON(w, http.StatusOK, users)
	}
}

func (u *UserAPI) FindByID() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		id, err := strconv.Atoi(vars["id"])
		if err != nil {
			RespondWithError(w, http.StatusBadRequest, err.Error())
			return
		}

		// Caching
		value, _ := u.client.Get(string(id))
		if value != nil {
			var user *model.User
			data := fmt.Sprintf("CACHE: %v", value)
			err := json.Unmarshal(value, &user)
			if err != nil {
				log.Println(err)
			}
			fmt.Println(data)
			RespondWithJSON(w, http.StatusOK, model.ToUserDTO(user))
			return
		}

		user, err := u.UserService.FindById(uint(id))
		if err != nil {
			RespondWithError(w, http.StatusNotFound, err.Error())
			return
		}

		// Caching set
		data, err := json.Marshal(user)
		if err != nil {
			fmt.Println(err)
		}
		err = u.client.Set(string(id), data, time.Second*10)
		fmt.Println("CACHELENDI")

		if err != nil {
			log.Println(err)
		}
		RespondWithJSON(w, http.StatusOK, model.ToUserDTO(user))
	}
}

func (u *UserAPI) CreateUser() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var userDTO model.UserDTO

		decoder := json.NewDecoder(r.Body)
		if err := decoder.Decode(&userDTO); err != nil {
			RespondWithError(w, http.StatusBadRequest, err.Error())
			return
		}
		defer r.Body.Close()

		createdUser, err := u.UserService.Save(model.ToUser(&userDTO))
		if err != nil {
			RespondWithError(w, http.StatusInternalServerError, err.Error())
			return
		}

		RespondWithJSON(w, http.StatusOK, model.ToUserDTO(createdUser))

	}
}

func (u *UserAPI) UpdateUser() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		id, err := strconv.Atoi(vars["id"])
		if err != nil {
			RespondWithError(w, http.StatusBadRequest, err.Error())
			return
		}

		var userDTO model.User
		decoder := json.NewDecoder(r.Body)
		if err := decoder.Decode(&userDTO); err != nil {
			RespondWithError(w, http.StatusBadRequest, err.Error())
			return
		}
		defer r.Body.Close()

		user, err := u.UserService.FindById(uint(id))
		if err != nil {
			RespondWithError(w, http.StatusNotFound, err.Error())
			return
		}

		user.Username = userDTO.Username
		user.Password = userDTO.Password
		user.Email = userDTO.Email
		updatedUser, err := u.UserService.Save(user)
		if err != nil {
			RespondWithError(w, http.StatusInternalServerError, err.Error())
			return
		}

		RespondWithJSON(w, http.StatusOK, model.ToUserDTO(updatedUser))
	}
}

func (u *UserAPI) DeleteUser() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		id, err := strconv.Atoi(vars["id"])
		if err != nil {
			RespondWithError(w, http.StatusBadRequest, err.Error())
			return
		}

		user, err := u.UserService.FindById(uint(id))
		if err != nil {
			RespondWithError(w, http.StatusNotFound, err.Error())
			return
		}

		err = u.UserService.Delete(user.ID)
		if err != nil {
			RespondWithError(w, http.StatusNotFound, err.Error())
			return
		}
		type Response struct {
			Message string
		}
		response := Response{
			Message: "Post deleted successfully!",
		}

		RespondWithJSON(w, http.StatusOK, response)
	}
}

func (u *UserAPI) Migrate() {
	err := u.UserService.Migrate()
	if err != nil {
		log.Println(err)
	}
}
