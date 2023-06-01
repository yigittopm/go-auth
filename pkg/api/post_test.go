package api

import (
	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	"github.com/stretchr/testify/assert"
	"go-auth/pkg/cache"
	"go-auth/pkg/repository"
	"go-auth/pkg/service"
	"net/http"
	"net/http/httptest"
	"testing"
)

func postRouterSetup(api PostAPI) *mux.Router {
	apiRouter := mux.NewRouter()
	apiRouter.HandleFunc("/posts", api.FindAll()).Methods(http.MethodGet)
	return apiRouter
}

func postApiSetup(db *gorm.DB) PostAPI {
	client := cache.New()
	postRepository := repository.NewPostRepository(db)
	postService := service.NewPostService(postRepository)
	postApi := NewPostAPI(postService, client)

	return postApi
}

func TestPostAPI_FindAll(t *testing.T) {
	w := httptest.NewRecorder()
	mockDB, _ := dbSetup()
	api := postApiSetup(mockDB)
	r := postRouterSetup(api)

	r.ServeHTTP(w, httptest.NewRequest("GET", "/users", nil))
	assert.Equal(t, http.StatusOK, w.Code, "Did not get expected HTTP status code, got")
}
