package post

import (
	"encoding/json"
	"go-auth/pkg/api"
	"go-auth/pkg/app"
	"go-auth/pkg/cache"
	"go-auth/pkg/model"
	"go-auth/pkg/repository/post"
	"go-auth/pkg/service"
	"log"
	"net/http"
)

type PostAPI struct {
	PostService service.PostService
	client      *cache.Client
}

func NewPostAPI(p service.PostService, c *cache.Client) PostAPI {
	return PostAPI{
		PostService: p,
		client:      c,
	}
}

func InitPostAPI(app *app.App) {
	postRepository := post.NewRepository(app.DB)
	postService := service.NewPostService(postRepository)
	postAPI := NewPostAPI(postService, app.Cache)
	postAPI.Migrate()

	app.Router.HandleFunc("/posts", postAPI.FindAll()).Methods(http.MethodGet)
	app.Router.HandleFunc("/posts", postAPI.CreatePost()).Methods(http.MethodPost)
}

func (p *PostAPI) FindAll() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		posts, err := p.PostService.All()
		if err != nil {
			api.RespondWithError(w, http.StatusNotFound, err.Error())
			return
		}

		api.RespondWithJSON(w, http.StatusOK, posts)
	}
}

func (p *PostAPI) CreatePost() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var postDTO model.PostDTO

		decoder := json.NewDecoder(r.Body)
		if err := decoder.Decode(&postDTO); err != nil {
			api.RespondWithError(w, http.StatusBadRequest, err.Error())
			return
		}
		defer r.Body.Close()

		createdPost, err := p.PostService.Save(model.ToPost(&postDTO))
		if err != nil {
			api.RespondWithError(w, http.StatusInternalServerError, err.Error())
			return
		}

		api.RespondWithJSON(w, http.StatusOK, model.ToPostDTO(createdPost))
	}
}

func (p *PostAPI) Migrate() {
	err := p.PostService.Migrate()
	if err != nil {
		log.Println(err)
	}
}
