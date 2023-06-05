package api

import (
	"encoding/json"
	"go-auth/pkg-v1/cache"
	"go-auth/pkg-v1/model"
	"go-auth/pkg-v1/service"
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

func (p *PostAPI) FindAll() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		posts, err := p.PostService.All()
		if err != nil {
			RespondWithError(w, http.StatusNotFound, err.Error())
			return
		}

		RespondWithJSON(w, http.StatusOK, posts)
	}
}

func (p *PostAPI) CreatePost() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var postDTO model.PostDTO

		decoder := json.NewDecoder(r.Body)
		if err := decoder.Decode(&postDTO); err != nil {
			RespondWithError(w, http.StatusBadRequest, err.Error())
			return
		}
		defer r.Body.Close()

		createdPost, err := p.PostService.Save(model.ToPost(&postDTO))
		if err != nil {
			RespondWithError(w, http.StatusInternalServerError, err.Error())
			return
		}

		RespondWithJSON(w, http.StatusOK, model.ToPostDTO(createdPost))
	}
}

func (p *PostAPI) Migrate() {
	err := p.PostService.Migrate()
	if err != nil {
		log.Println(err)
	}
}
