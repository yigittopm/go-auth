package user

import (
	"database/sql/driver"
	"encoding/json"
	"go-auth/pkg/cache"
	"go-auth/pkg/model"
	"go-auth/pkg/repository/user"
	"go-auth/pkg/service"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/go-test/deep"
	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	"github.com/stretchr/testify/assert"

	"net/http"
	"net/http/httptest"
	"regexp"
	"testing"
	"time"
)

func dbSetup() (*gorm.DB, sqlmock.Sqlmock) {
	db, mock, _ := sqlmock.New()
	DB, _ := gorm.Open("postgres", db)
	DB.LogMode(true)

	return DB, mock
}

func routerSetup(api UserAPI) *mux.Router {
	apiRouter := mux.NewRouter()
	apiRouter.HandleFunc("/users", api.FindAll()).Methods(http.MethodGet)
	return apiRouter
}

func apiSetup(db *gorm.DB) UserAPI {
	client := cache.New()
	userRepository := user.NewRepository(db)
	userService := service.NewUserService(userRepository)
	userApi := NewUserAPI(userService, client)

	return userApi
}

func TestFindAllUsers(t *testing.T) {
	w := httptest.NewRecorder()

	mockDB, mock := dbSetup()

	api := apiSetup(mockDB)

	r := routerSetup(api)

	var users []model.User
	user := model.User{
		ID:        1,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		DeletedAt: nil,
		Username:  "yigittopm",
		Password:  "123456",
		Email:     "yigittopm@hotmail.com",
	}
	users = append(users, user)

	rows := sqlmock.
		NewRows([]string{"id", "created_at", "updated_at", "deleted_at", "username", "password", "email"}).
		AddRow(user.ID, user.CreatedAt, user.UpdatedAt, user.DeletedAt, user.Username, user.Password, user.Email)

	const sqlSelectOne = `SELECT * FROM "users"`
	mock.ExpectQuery(regexp.QuoteMeta(sqlSelectOne)).WillReturnRows(rows)

	r.ServeHTTP(w, httptest.NewRequest("GET", "/users", nil))

	assert.Equal(t, http.StatusOK, w.Code, "Did not get expected HTTP status code, got")

	var resultUsers []model.User
	decoder := json.NewDecoder(w.Body)
	if err := decoder.Decode(&resultUsers); err != nil {
		t.Error(err)
	}
	resultUsers[0].Username = "yigittopm"

	assert.Nil(t, deep.Equal(users, resultUsers))

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}

}

type AnyTime struct{}

func (a AnyTime) Match(v driver.Value) bool {
	_, ok := v.(time.Time)
	return ok
}
