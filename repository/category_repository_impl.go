package repository

import (
	"context"
	"database/sql"
	"errors"

	"pustaka-api7.1/helper"
	"pustaka-api7.1/model/domain"
)

type CategoryRepositoryImpl struct {
}

func NewCategoryRepository () CategoryRepository {
	return &CategoryRepositoryImpl{}
	
}

func (repo *CategoryRepositoryImpl) Save(ctx context.Context , tx *sql.Tx , category domain.Category) domain.Category {
	SQL := "Insert into category (name) values (?)"
	result , err := tx.ExecContext(ctx , SQL , category.Name)
	helper.PanicIfError(err)

   id , err := result.LastInsertId()
   if err != nil {
	panic(err)
   }

   category.Id = int(id)
   return category



}

func (repo *CategoryRepositoryImpl) Update(ctx context.Context , tx *sql.Tx , category domain.Category) domain.Category {

SQL := "update category set name = ? where id = ?"

	_ , err :=	tx.ExecContext(ctx, SQL , category.Id)
	helper.PanicIfError(err)

	return category

}

func (repo *CategoryRepositoryImpl) Delete(ctx context.Context , tx *sql.Tx , category domain.Category) {

	SQL := "delete from category where (id) = ?"

	_, err := tx.ExecContext(ctx , SQL , category.Id)
	helper.PanicIfError(err)


}

func (repo *CategoryRepositoryImpl) FindById(ctx context.Context , tx *sql.Tx , categoryId int) (domain.Category , error) {

	SQL := "select from category where (id) = ?"

	rows , err := tx.QueryContext(ctx , SQL , categoryId)
	helper.PanicIfError(err)

	category := domain.Category{}
	if rows.Next() {
		err := rows.Scan(&category.Id , &category.Name)
		helper.PanicIfError(err)
		return category , nil
	} else {
		return category , errors.New("category is not found")
	}

}
func (repo *CategoryRepositoryImpl) FindAll(ctx context.Context , tx *sql.Tx ) []domain.Category {

	SQL := "select id , name from category"

	rows , err :=	tx.QueryContext(ctx , SQL)
	helper.PanicIfError(err)

	var categories []domain.Category
	for rows.Next() {
		category := domain.Category{}
		err := rows.Scan(&category.Id , &category.Name)
		helper.PanicIfError(err)
		categories = append(categories, category )
	}

	return categories
}