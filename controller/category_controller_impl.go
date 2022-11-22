package controller

import (
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
	"pustaka-api7.1/helper"
	"pustaka-api7.1/model/web"
	"pustaka-api7.1/service"
)

type CategoryControllerImpl struct {
	CategoryService service.CategoryService

}

func NewCategoryController (categoryService service.CategoryService) CategoryController {
	return &CategoryControllerImpl{
		CategoryService: categoryService,
	}
}

func (controller *CategoryControllerImpl) Save (writer http.ResponseWriter, request *http.Request, params httprouter.Params) {

	
	categoryCreateRequest := web.CategoryCreateRequest{}
	helper.ReadFromBody(request , &categoryCreateRequest)

	categoryResponse :=	controller.CategoryService.Save(request.Context(),categoryCreateRequest )

	webResponse := web.WebResponse{
		Code: 200,
		Status: "OK",
		Data: categoryResponse,

	}


	helper.WriteFromResponse(writer , webResponse)
}

func (controller *CategoryControllerImpl) Update (writer http.ResponseWriter, request *http.Request, params httprouter.Params) {

	
	categoryUpdateRequest := web.CategoryUpdateRequest{}
	helper.ReadFromBody(request , &categoryUpdateRequest)

	categoryId := params.ByName("categoryId")
	id , err := strconv.Atoi(categoryId)
	helper.PanicIfError(err)

	categoryUpdateRequest.Id = id

	UpdateResponse := controller.CategoryService.Update(request.Context() , categoryUpdateRequest )

	webResponse := web.WebResponse {
		Code: 200,
		Status: "OK",
		Data: UpdateResponse,
	}

helper.WriteFromResponse(writer , webResponse)



}

func (controller *CategoryControllerImpl) Delete (writer http.ResponseWriter, request *http.Request, params httprouter.Params) {

	categoryId := params.ByName("categoryId")
	id , err := strconv.Atoi(categoryId)
	helper.PanicIfError(err)

	controller.CategoryService.Delete(request.Context() , id)

	webResponse := web.WebResponse {
		Code: 200,
		Status: "OK",
	}
	helper.WriteFromResponse(writer , webResponse)
}


func (controller *CategoryControllerImpl) FindById (writer http.ResponseWriter, request *http.Request, params httprouter.Params) {

	categoryId := params.ByName("categoryId")
	id , err := strconv.Atoi(categoryId)
	helper.PanicIfError(err)

	categoryResponse :=	controller.CategoryService.FindById(request.Context() , id)

	webResponse := web.WebResponse {
		Code: 200,
		Status: "OK",
		Data: categoryResponse,
	}

	helper.WriteFromResponse(writer , webResponse)
}

func (controller *CategoryControllerImpl) FindAll (writer http.ResponseWriter, request *http.Request, params httprouter.Params) {


	categoryResponse :=	controller.CategoryService.FindAll(request.Context())


	webResponse := web.WebResponse {
		Code: 200,
		Status: "OK",
		Data: categoryResponse,
	}

	helper.WriteFromResponse(writer , webResponse)
}