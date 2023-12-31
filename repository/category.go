package repository

import (
	"errors"
	"strconv"

	initialisers "main.go/Initialisers"
	"main.go/domain"
	"main.go/models"
)

func GetCategory() ([]domain.Category, error) {
	var category []domain.Category
	err := initialisers.DB.Raw("SELECT * FROM categories").Scan(&category).Error
	if err != nil {
		return nil, err
	}
	return category, nil
}
func AddCategory(category models.Category) (domain.Category, error) {
	var categore string
	err := initialisers.DB.Raw("INSERT INTO categories (category) VALUES (?) RETURNING category", category.Category).Scan(&categore).Error
	if err != nil {
		return domain.Category{}, err
	}
	var categoriesResponse domain.Category
	err = initialisers.DB.Raw("SELECT id , category FROM categories  WHERE category = ?", categore).Scan(&categoriesResponse).Error
	if err != nil {
		return domain.Category{}, err
	}
	return categoriesResponse, nil
}
func DeleteCategory(id string) error {
	category_id, err := strconv.Atoi(id)
	if err != nil {
		return err
	}
	var count int
	if err := initialisers.DB.Raw("SELECT COUNT(*) FROM categories WHERE id=?", category_id).Scan(&count).Error; err != nil {
		return err
	}
	if count < 1 {
		return errors.New("category for given id does not exist")
	}

	if err := initialisers.DB.Exec("DELETE FROM categories WHERE id=?", category_id).Error; err != nil {
		return err
	}
	if query := initialisers.DB.Raw(`UPDATE products SET deleted = true WHERE category_id = ?`, category_id).Error; err != nil {
		return query
	}
	return nil
}
func UpdateCategory(current string, new string) (domain.Category, error) {
	if initialisers.DB == nil {
		return domain.Category{}, errors.New("database connection is nil")
	}
	if err := initialisers.DB.Exec("UPDATE categories SET category=? WHERE category = ?", new, current).Error; err != nil {
		return domain.Category{}, err
	}
	var newcat domain.Category
	if err := initialisers.DB.Raw("SELECT id,category FROM categories WHERE category = ?", new).Scan(&newcat).Error; err != nil {
		return domain.Category{}, nil
	}
	return newcat, nil
}
func CheckCategory(current string) (bool, error) {
	var count int
	err := initialisers.DB.Raw("SELECT COUNT(*) FROM categories WHERE category=?", current).Scan(&count).Error
	if err != nil {
		return false, err
	}
	if count == 0 {
		return false, err
	}
	return true, err
}
