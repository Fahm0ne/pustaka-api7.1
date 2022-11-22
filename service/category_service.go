package service

import (
	"context"

	"pustaka-api7.1/model/web"
)

type CategoryService interface {
	Save(ctx context.Context , request web.CategoryCreateRequest ) web.CategoryResponse
	Update(ctx context.Context, request web.CategoryUpdateRequest ) web.CategoryResponse
	Delete(ctx context.Context, categoryId int)
	FindById(ctx context.Context, categoryId int) web.CategoryResponse
	FindAll(ctx context.Context) []web.CategoryResponse
}