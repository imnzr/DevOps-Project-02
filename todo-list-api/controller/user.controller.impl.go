package controller

import (
	"encoding/json"
	"net/http"
	"strconv"
	"todo-list-api/helper"
	"todo-list-api/models/web"
	"todo-list-api/service"

	"github.com/julienschmidt/httprouter"
)

type UserControllerImpl struct {
	UserService service.UserService
}

func NewUserController(userService service.UserService) UserController {
	return &UserControllerImpl{
		UserService: userService,
	}
}

// Create implements UserController.
func (controller UserControllerImpl) Create(writter http.ResponseWriter, request *http.Request, params httprouter.Params) {
	decoder := json.NewDecoder(request.Body)
	userCreateRequest := web.UserCreateRequest{}
	err := decoder.Decode(&userCreateRequest)
	helper.PanicIfError(err)

	userResponse := controller.UserService.Create(request.Context(), userCreateRequest)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   userResponse,
	}

	writter.Header().Add("Content-Type", "application/json")

	encoder := json.NewEncoder(writter)
	err = encoder.Encode(webResponse)
	helper.PanicIfError(err)

}

// Update implements UserController.
func (controller UserControllerImpl) Update(writter http.ResponseWriter, request *http.Request, params httprouter.Params) {
	decoder := json.NewDecoder(request.Body)
	userUpdateRequest := web.UserUpdateRequest{}
	err := decoder.Decode(&userUpdateRequest)
	helper.PanicIfError(err)

	userId := params.ByName("userId")
	id, err := strconv.Atoi(userId)
	helper.PanicIfError(err)

	userUpdateRequest.Id = id

	userResponse := controller.UserService.Update(request.Context(), userUpdateRequest)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   userResponse,
	}

	writter.Header().Add("Content-Type", "application/json")

	encoder := json.NewEncoder(writter)
	err = encoder.Encode(webResponse)
	helper.PanicIfError(err)

}

// Delete implements UserController.
func (controller UserControllerImpl) Delete(writter http.ResponseWriter, request *http.Request, params httprouter.Params) {
	userId := params.ByName("userId")
	id, err := strconv.Atoi(userId)
	helper.PanicIfError(err)

	controller.UserService.Delete(request.Context(), id)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
	}

	writter.Header().Add("Content-Type", "application/json")

	encoder := json.NewEncoder(writter)
	err = encoder.Encode(webResponse)
	helper.PanicIfError(err)
}

// FindById implements UserController.
func (controller UserControllerImpl) FindById(writter http.ResponseWriter, request *http.Request, params httprouter.Params) {
	userId := params.ByName("userId")
	id, err := strconv.Atoi(userId)
	helper.PanicIfError(err)

	userResponse := controller.UserService.FindById(request.Context(), id)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   userResponse,
	}

	writter.Header().Add("Content-Type", "application/json")

	encoder := json.NewEncoder(writter)
	err = encoder.Encode(webResponse)
	helper.PanicIfError(err)
}

// FindByAll implements UserController.
func (controller UserControllerImpl) FindByAll(writter http.ResponseWriter, request *http.Request, params httprouter.Params) {
	userResponses := controller.UserService.FindByAll(request.Context())
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   userResponses,
	}

	writter.Header().Add("Content-Type", "application/json")

	encoder := json.NewEncoder(writter)
	err := encoder.Encode(webResponse)
	helper.PanicIfError(err)
}

// Login implements UserController.
func (controller *UserControllerImpl) Login(writter http.ResponseWriter, request *http.Request, params httprouter.Params) {
	decoder := json.NewDecoder(request.Body)
	loginRequest := web.LoginRequest{}
	err := decoder.Decode(&loginRequest)
	helper.PanicIfError(err)

	userResponse := controller.UserService.Login(request.Context(), loginRequest)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   userResponse,
	}

	encoder := json.NewEncoder(writter)
	err = encoder.Encode(webResponse)
	helper.PanicIfError(err)
}

// UpdateUsername implements UserController.
func (controller *UserControllerImpl) UpdateUsername(writter http.ResponseWriter, request *http.Request, params httprouter.Params) {
	decoder := json.NewDecoder(request.Body)
	userUpdateRequest := web.UserUpdateRequestUsername{}
	err := decoder.Decode(&userUpdateRequest)
	helper.PanicIfError(err)

	userId := params.ByName("userId")
	id, err := strconv.Atoi(userId)
	helper.PanicIfError(err)

	userUpdateRequest.Id = id

	updateUser, err := controller.UserService.UpdateUsername(request.Context(), userUpdateRequest)
	helper.PanicIfError(err)

	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   updateUser,
	}

	// Beri response ke client
	writter.Header().Set("Content-Type", "application/json")
	writter.WriteHeader(http.StatusOK)
	json.NewEncoder(writter).Encode(webResponse)

}
