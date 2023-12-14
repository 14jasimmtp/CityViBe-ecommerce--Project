package repository

import (
	"errors"
	"fmt"
	"strconv"
	"time"

	"gorm.io/gorm"
	initialisers "main.go/Initialisers"
	"main.go/domain"
	"main.go/models"
)

func AddProduct(product models.AddProduct) (models.UpdateProduct, error) {
	var dproduct models.UpdateProduct
	var p domain.Product
	result := initialisers.DB.Raw("INSERT INTO products(name,description,category_id,size_id,stock,price,color,image_url) values(?,?,?,?,?,?,?,?)", product.Name, product.Description, product.CategoryID, product.Size, product.Stock, product.Price, product.Color, product.ImageURL).Scan(&p)
	fmt.Println(p)
	if result.Error != nil {
		return models.UpdateProduct{}, result.Error
	}
	query := initialisers.DB.Raw(`SELECT products.id,name,description,categories.category,sizes.size,stock,price,color FROM products INNER JOIN categories ON categories.id = products.category_id INNER JOIN sizes ON sizes.id=products.size_id WHERE name = ?`, product.Name).Scan(&dproduct)
	if query.Error != nil {
		return models.UpdateProduct{}, query.Error
	}
	fmt.Println(dproduct)
	return dproduct, nil
}

func EditProductDetails(id string, product models.AddProduct) (models.UpdateProduct, error) {
	var updatedProduct models.UpdateProduct

	result := initialisers.DB.Raw("UPDATE products SET name=?,description=?,category_id=?,size_id=?,stock=?,price=?,color=? WHERE id=?", product.Name, product.Description, product.CategoryID, product.Size, product.Stock, product.Price, product.Color, id).Scan(&updatedProduct)
	if result.Error != nil {
		return models.UpdateProduct{}, result.Error
	}
	query := initialisers.DB.Raw(`SELECT products.id,name,description,categories.category,sizes.size,stock,color,price FROM products INNER JOIN categories ON categories.id = products.category_id INNER JOIN sizes ON sizes.id=products.size_id WHERE products.id = ?`, id).Scan(&updatedProduct)
	if query.Error != nil {
		return models.UpdateProduct{}, query.Error

	}
	return updatedProduct, nil
}

func DeleteProduct(id int) error {
	query := initialisers.DB.Exec(`UPDATE products SET deleted = true WHERE id = ?`, id)
	if query.Error != nil {
		return errors.New("no product found to delete")
	}
	return nil
}

func GetAllProducts() ([]models.Product, error) {
	var products []models.Product
	query := initialisers.DB.Raw(`SELECT products.id,name,description,categories.category,sizes.size,stock,color,price,offer_prize FROM products INNER JOIN categories ON categories.id = products.category_id INNER JOIN sizes ON sizes.id=products.size_id WHERE deleted = false ORDER BY id ASC`).Scan(&products)
	if query.Error != nil {
		return []models.Product{}, query.Error
	}
	return products, nil
}

func SeeAllProducts() ([]domain.Product, error) {
	var products []domain.Product
	err := initialisers.DB.Raw("SELECT * FROM products ORDER BY id ASC").Scan(&products).Error
	if err != nil {
		return nil, err
	}
	return products, nil
}

func GetSingleProduct(id string) (models.Product, error) {
	var product models.Product
	idint, err := strconv.Atoi(id)
	if err != nil {
		return models.Product{}, errors.New("error while converting id to int")
	}

	query := initialisers.DB.Raw("SELECT name,description,categories.category,sizes.size,stock,color,price FROM products INNER JOIN categories ON categories.id = products.category_id INNER JOIN sizes ON sizes.id=products.size_id WHERE products.id = ?", idint).Scan(&product)
	if product.Name == "" {
		return models.Product{}, errors.New("no products found with this id")
	}

	if query.Error != nil {
		return models.Product{}, errors.New("something went wrong")
	}

	return product, nil
}

func FilterProductCategoryWise(category string) ([]models.Product, error) {
	var product []models.Product
	initialisers.DB.Raw(`SELECT name,description,categories.category,sizes.size,stock,color,price FROM products INNER JOIN categories ON categories.id = products.category_id INNER JOIN sizes ON sizes.id=products.size_id WHERE categories.category = ? `, category).Scan(&product)
	return product, nil
}

func CheckStock(pid int) error {
	var stock int
	initialisers.DB.Raw(`SELECT stock from products WHERE id = ?`, pid).Scan(&stock)
	if stock < 1 {
		return errors.New("product out of stock")
	}
	return nil
}

func GetProductAmountFromID(pid string) (float64, error) {
	var price struct {
		Price      float64
		OfferPrize float64
	}

	if err := initialisers.DB.Raw("select price,offer_prize from products where id = ?", pid).Scan(&price).Error; err != nil {
		return 0.0, err
	}
	if price.OfferPrize != 0 {
		return price.OfferPrize, nil
	}
	return price.Price, nil
}

func SearchProduct(search string) ([]models.Product, error) {
	var products []models.Product

	query := initialisers.DB.Raw(
		`SELECT id,name,description,Categories.category,Sizes.size,stock,price,color
	 	 FROM products INNER JOIN categories ON categories.id=products.category_id
		 INNER JOIN sizes ON sizes.id = products.size_id
		 WHERE name ILIKE $1 OR description ILIKE $1 OR sizes.size ILIKE $1 OR categories.category ILIKE $1`, "%"+search+"%",
	).Scan(&products)
	if query.Error != nil {
		return []models.Product{}, errors.New(`something went wrong`)
	}
	if query.RowsAffected < 1 {
		return []models.Product{}, errors.New(`no products found`)
	}

	return products, nil
}

func FilterProducts(category, size string, minPrice, maxPrice float64) ([]models.UpdateProduct, error) {
	var Products []models.UpdateProduct
	query := initialisers.DB.Raw(
		`SELECT products.id,products.name,products.description,categories.category,sizes.size,products.stock,products.price,products.color,products.offer_prize as offer_price FROM products
		 INNER JOIN categories ON products.category_id=categories.id
		 INNER JOIN sizes ON products.size_id=sizes.id
		 WHERE (categories.category = ? OR ? = '')
		 AND (sizes.size = ? OR ? = '')
		 AND ((products.price >= ? AND products.price <= ?) OR ? = 0.0)`,
		category, category, size, size, minPrice, maxPrice, minPrice).Scan(&Products)
	if query.Error != nil {
		return []models.UpdateProduct{}, errors.New(`something went wrong`)
	}
	if query.RowsAffected < 1 {
		return []models.UpdateProduct{}, errors.New(`no products found`)
	}
	return Products, nil
}

func CreateOffer(offer *models.Offer) error {
	if err := initialisers.DB.Create(offer).Error; err != nil {
		return err
	}
	return nil
}

func GetAllOffers() ([]models.Offer, error) {
	var offer []models.Offer
	currenttime := time.Now()
	err := initialisers.DB.Where("valid_until > ?", currenttime).Find(&offer).Error
	if err != nil {
		return nil, errors.New("record not found")
	}
	return offer, nil

}

func GetProductsByCategoryoffer(id int) ([]models.Product, error) {
	var product []models.Product

	err := initialisers.DB.Raw(
		`SELECT products.id,products.name,products.description,categories.category,sizes.size,products.stock,price,offer_prize as offer_price,color
		from products
		inner join categories on categories.id=products.category_id
		inner join sizes on sizes.id=products.size_id 
		Where category_id= ? AND deleted = ?`,
		id, false).Find(&product).Error
	if err != nil {
		return nil, errors.New("record not found")
	}
	return product, nil
}

func GetProductById(id int) (*models.Product, error) {
	var product models.Product
	result := initialisers.DB.Raw(
		`SELECT products.id,products.name,products.description,categories.category,sizes.size,products.stock,price,offer_prize as offer_price,color
		from products
		inner join categories on categories.id=products.category_id
		inner join sizes on sizes.id=products.size_id 
		Where products.id = ? AND deleted = ?`,
		id, false).Find(&product)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, result.Error
		}
		return nil, result.Error
	}
	return &product, nil
}

func UpdateProduct(product *models.Product) error {
	query := initialisers.DB.Exec(`UPDATE products SET offer_prize = ? WHERE id = ? `, product.OfferPrize, product.ID)
	if query.Error != nil {
		return errors.New(`something went wrong`)
	}
	if query.RowsAffected == 0 {
		return errors.New(`no products found with this id`)
	}
	return nil
}
