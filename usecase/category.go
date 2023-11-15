package usecase

import (
	"errors"

	"main.go/domain"
	"main.go/models"
	"main.go/repository"
)

func GetCategory() ([]domain.Category, error) {
	category, err := repository.GetCategory()
	if err != nil {
		return []domain.Category{}, err
	}
	return category, nil

}
func AddCategory(category models.Category) (domain.Category, error) {
	categories, err := repository.AddCategory(category)
	if err != nil {
		return domain.Category{}, err
	}
	return categories, nil
}
func UpdateCategory(current string, new string) (domain.Category, error) {
	categries, err := repository.CheckCategory(current)
	if err != nil {
		return domain.Category{}, err
	}
	if !categries {
		return domain.Category{}, errors.New("category doesn't exist")
	}
	newCate, err := repository.UpdateCategory(current, new)
	if err != nil {
		return domain.Category{}, err
	}
	return newCate, nil
}
func DeleteCategory(id string) error {
	err := repository.DeleteCategory(id)
	if err != nil {
		return err
	}
	return nil
}
