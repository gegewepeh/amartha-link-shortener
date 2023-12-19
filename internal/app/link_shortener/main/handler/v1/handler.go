package handler

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"

	v1 "link-shortener/internal/app/link_shortener/main/services/v1"
	models "link-shortener/internal/app/link_shortener/models"
	network "link-shortener/internal/pkg/network"
	utils "link-shortener/internal/pkg/utils"

	"github.com/gorilla/mux"
	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin:     func(r *http.Request) bool { return true },
}

func getAllLinks(w http.ResponseWriter, r *http.Request) (*network.Response, *utils.AppError) {
	ctx := createContext(r)
	vars := mux.Vars(r)
	id := vars["id"]
	username := r.URL.Query().Get("username")
	password := r.URL.Query().Get("password")

	if username == "" || password == "" {
		return nil, models.WrapError("User", "InvalidParameters", errors.New("Invalid Username or Password"), nil)
	}

	pool := v1.GetPool()
	user, err := pool.FindUser(ctx, username, password)
	if err != nil {
		return nil, err
	}
	data, _ := pool.GetLinks(ctx, id, user.ID)

	return &network.Response{
		Code: 200,
		Data: data,
		Meta: nil,
	}, nil
}

func getBySlugId(w http.ResponseWriter, r *http.Request) (*network.Response, *utils.AppError) {
	ctx := createContext(r)
	vars := mux.Vars(r)
	id := vars["id"]

	pool := v1.GetPool()
	data, err := pool.GetBySlugId(ctx, id)
	if err != nil {
		return nil, err
	}

	return &network.Response{
		Code: 301,
		Data: data,
		Meta: nil,
	}, nil
}

func createSlug(w http.ResponseWriter, r *http.Request) (*network.Response, *utils.AppError) {
	var request models.CreateSlugRequest
	ctx := createContext(r)

	decodeErr := json.NewDecoder(r.Body).Decode(&request)
	if decodeErr != nil {
		return nil, models.WrapError("General", "JSONError", decodeErr, nil)
	}

	if utils.IsValidURL(request.FullLink) == false {
		return nil, models.WrapError("Link", "InvalidParameters", errors.New("Invalid URL Format"), nil)
	}

	if request.Username == "" || request.Password == "" {
		return nil, models.WrapError("User", "InvalidParameters", errors.New("Invalid Username or Password"), nil)
	}

	generatedSlugId := utils.RandString(6, false)
	pool := v1.GetPool()
	user, err := pool.FindUser(ctx, request.Username, request.Password)
	if err != nil {
		return nil, err
	}

	data, err := pool.CreateSlug(ctx, request.FullLink, generatedSlugId, user.ID)
	if err != nil {
		return nil, err
	}

	return &network.Response{
		Code: 201,
		Data: data,
		Meta: nil,
	}, nil
}

func updateSlug(w http.ResponseWriter, r *http.Request) (*network.Response, *utils.AppError) {
	var request models.UpdateSlugRequest
	ctx := createContext(r)
	vars := mux.Vars(r)
	id := vars["id"]

	decodeErr := json.NewDecoder(r.Body).Decode(&request)
	if decodeErr != nil {
		return nil, models.WrapError("General", "JSONError", decodeErr, nil)
	}
	request.Slug = id

	if request.Username == "" || request.Password == "" {
		return nil, models.WrapError("User", "InvalidParameters", errors.New("Invalid Username or Password"), nil)
	}

	generatedSlugId := utils.RandString(6, false)
	pool := v1.GetPool()
	user, err := pool.FindUser(ctx, request.Username, request.Password)
	if err != nil {
		return nil, err
	}

	data, err := pool.UpdateSlugId(ctx, request.Slug, generatedSlugId, user.ID)
	if err != nil {
		return nil, err
	}

	return &network.Response{
		Code: 201,
		Data: data,
		Meta: nil,
	}, nil
}

func createUser(w http.ResponseWriter, r *http.Request) (*network.Response, *utils.AppError) {
	var request models.CreateUserRequest
	ctx := createContext(r)

	decodeErr := json.NewDecoder(r.Body).Decode(&request)
	if decodeErr != nil {
		return nil, models.WrapError("General", "JSONError", decodeErr, nil)
	}

	pool := v1.GetPool()
	data, err := pool.CreateUser(ctx, request.Username, request.Password)
	if err != nil {
		return nil, err
	}

	return &network.Response{
		Code: 201,
		Data: data,
		Meta: nil,
	}, nil
}

func createContext(r *http.Request) context.Context {
	ctx := r.Context()
	reqID := r.Header.Get("X-Request-Id")
	if reqID == "" {
		reqID = utils.RandString(8, true)
	}
	ctx = utils.SetRequestID(ctx, reqID)
	return ctx
}
