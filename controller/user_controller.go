package controller

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/imnzr/DevOps-Project-02/models/web"
	"github.com/imnzr/DevOps-Project-02/service"
	"github.com/julienschmidt/httprouter"
)

type UserControllerImplementation struct {
	UserService service.UserService
}

func NewUserController(userService service.UserService) UserController {
	return &UserControllerImplementation{
		UserService: userService,
	}
}

// Login implements UserController.
func (controller *UserControllerImplementation) Login(writter http.ResponseWriter, request *http.Request, params httprouter.Params) {
	decoder := json.NewDecoder(request.Body)
	loginRequest := web.UserLoginRequest{}
	err := decoder.Decode(&loginRequest)
	if err != nil {
		http.Error(writter, err.Error(), http.StatusBadRequest)
		return
	}

	userResponse := controller.UserService.Login(request.Context(), loginRequest)
	webResponse := web.WebResponse{
		Code:   http.StatusOK,
		Status: "OK",
		Data:   userResponse,
	}
	writter.Header().Add("Content-Type", "application/json")
	writter.WriteHeader(http.StatusOK)
	err = json.NewEncoder(writter).Encode(webResponse)
	if err != nil {
		http.Error(writter, err.Error(), http.StatusBadRequest)
		return
	}
}

// Create implements UserController.
func (controller *UserControllerImplementation) Create(writter http.ResponseWriter, request *http.Request, params httprouter.Params) {
	decoder := json.NewDecoder(request.Body)
	userCrateRequest := web.UserCreateRequest{}
	err := decoder.Decode(&userCrateRequest)
	if err != nil {
		http.Error(writter, err.Error(), http.StatusBadRequest)
	}
	userResponse := controller.UserService.Create(request.Context(), userCrateRequest)
	webResponse := web.WebResponse{
		Code:   http.StatusOK,
		Status: "OK",
		Data:   userResponse,
	}

	writter.Header().Add("Content-Type", "application/json")
	writter.WriteHeader(http.StatusOK)
	json.NewEncoder(writter).Encode(webResponse)
}

// Delete implements UserController.
func (controller *UserControllerImplementation) Delete(writter http.ResponseWriter, request *http.Request, params httprouter.Params) {
	userId := params.ByName("userId")
	id, err := strconv.Atoi(userId)
	if err != nil {
		http.Error(writter, err.Error(), http.StatusBadRequest)
		return
	}
	controller.UserService.Delete(request.Context(), id)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
	}
	writter.Header().Add("Content-Type", "application/json")
	writter.WriteHeader(http.StatusOK)
	err = json.NewEncoder(writter).Encode(webResponse)
	if err != nil {
		http.Error(writter, err.Error(), http.StatusBadRequest)
	}
}

// FindAll implements UserController.
func (controller *UserControllerImplementation) FindAll(writter http.ResponseWriter, request *http.Request, params httprouter.Params) {
	user := controller.UserService.FindAll(request.Context())
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   user,
	}

	writter.Header().Add("Content-Type", "application/json")
	writter.WriteHeader(http.StatusOK)
	err := json.NewEncoder(writter).Encode(webResponse)
	if err != nil {
		http.Error(writter, err.Error(), http.StatusBadRequest)
	}
}

// FindById implements UserController.
func (controller *UserControllerImplementation) FindById(writter http.ResponseWriter, request *http.Request, params httprouter.Params) {
	userId := params.ByName("userId")
	id, err := strconv.Atoi(userId)
	if err != nil {
		http.Error(writter, err.Error(), http.StatusBadRequest)
		return
	}

	user := controller.UserService.FindById(request.Context(), id)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   user,
	}

	writter.Header().Add("Content-Type", "application/json")
	writter.WriteHeader(http.StatusOK)
	err = json.NewEncoder(writter).Encode(webResponse)
	if err != nil {
		http.Error(writter, err.Error(), http.StatusBadRequest)
	}
}

// Update implements UserController.
func (controller *UserControllerImplementation) Update(writter http.ResponseWriter, request *http.Request, params httprouter.Params) {
	decoder := json.NewDecoder(request.Body)
	userUpdateRequest := web.UserUpdateRequest{}
	err := decoder.Decode(&userUpdateRequest)
	if err != nil {
		http.Error(writter, err.Error(), http.StatusBadRequest)
		return
	}

	userId := params.ByName("userId")
	id, err := strconv.Atoi(userId)
	if err != nil {
		http.Error(writter, err.Error(), http.StatusBadRequest)
		return
	}

	userUpdateRequest.Id = id

	userResponse := controller.UserService.Update(request.Context(), userUpdateRequest)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   userResponse,
	}

	writter.Header().Add("Content-Type", "application/json")
	writter.WriteHeader(http.StatusOK)
	err = json.NewEncoder(writter).Encode(webResponse)
	if err != nil {
		http.Error(writter, err.Error(), http.StatusBadRequest)
	}

}
