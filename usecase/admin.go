package usecase

import (
	"errors"
	"fmt"
	"strconv"

	"golang.org/x/crypto/bcrypt"
	"main.go/models"
	"main.go/repository"
	"main.go/utils"
)

func AdminLogin(admin models.Admin) (models.Admin, error) {
	AdminDetails, err := repository.AdminLogin(admin)
	fmt.Println(err)
	if err != nil {
		fmt.Println("Admin doesn't exist")
		return models.Admin{}, errors.New("admin not found")
	}

	if bcrypt.CompareHashAndPassword([]byte(AdminDetails.Password), []byte(admin.Password)) != nil {
		fmt.Println("wrong password")
		return models.Admin{}, errors.New("wrong password")
	}

	tokenString, err := utils.AdminTokenGenerate(AdminDetails, "admin")
	if err != nil {
		fmt.Println("error generating token")
		return models.Admin{}, errors.New("error generating token")
	}

	return models.Admin{
		Firstname:   AdminDetails.Firstname,
		Lastname:    admin.Lastname,
		TokenString: tokenString,
	}, nil

}

func GetAllUsers() ([]models.UserDetailsResponse, error) {
	Users, err := repository.GetAllUsers()
	if err != nil {
		return []models.UserDetailsResponse{}, err
	}
	return Users, nil
}

func BlockUser(idStr string) error {
	id, _ := strconv.Atoi(idStr)
	user, err := repository.GetUserById(id)
	if err != nil {
		return err
	}
	if user.Blocked {
		return errors.New("already blocked")
	} else {
		user.Blocked = true
	}
	err = repository.BlockUserByID(user)
	if err != nil {
		return err
	}
	return nil

}

func UnBlockUser(idStr string) error {
	id, _ := strconv.Atoi(idStr)
	user, err := repository.GetUserById(id)
	if err != nil {
		return err
	}
	if !user.Blocked {
		return errors.New("already unblocked")
	} else {
		user.Blocked = false
	}
	err = repository.UnBlockUserByID(user)
	if err != nil {
		return err
	}
	return nil

}

func GetAllOrderDetailsForAdmin() ([]models.ViewAdminOrderDetails, error) {
	orderDetail, err := repository.GetAllOrderDetailsBrief()
	if err != nil {
		return []models.ViewAdminOrderDetails{}, err
	}
	return orderDetail, nil
}

func GetOrderDetails(orderID string) ([]models.OrderProductDetails, error) {
	orderDetails, err := repository.GetSingleOrderDetails(orderID)
	if err != nil {
		return []models.OrderProductDetails{}, err
	}

	return orderDetails, nil
}

func ExecuteGetOffers() (*[]models.Offer, error) {
	offers, err := repository.GetAllOffers()
	if err != nil {
		return nil, err
	}
	avialableoffers := []models.Offer{}
	for _, offers := range offers {
		if offers.UsageLimit != offers.UsedCount {
			avialableoffers = append(avialableoffers, offers)
		}
	}
	return &avialableoffers, nil
}

func ExecuteAddProductOffer(productid, offer int) (*models.Product, error) {

	product, err := repository.GetProductById(productid)
	if err != nil {
		return nil, err
	}
	if offer < 0 || offer > 100 {
		return nil, errors.New("invalid offer percentage")
	}

	amount := float64(offer) / 100.0 * float64(product.Price)
	product.OfferPrize = product.Price - amount
	err1 := repository.UpdateProduct(product)
	if err1 != nil {
		return nil, err
	}
	return product, nil
}

func ExecuteCategoryOffer(catid, offer int) ([]models.Product, error) {

	productlist, err := repository.GetProductsByCategoryoffer(catid)
	if err != nil {
		return nil, err
	}
	if offer < 0 || offer > 100 {
		return nil, errors.New("invalid offer percentage")
	}
	for i := range productlist {
		product := &(productlist)[i]

		amount := float64(offer) / 100.0 * float64(product.Price)
		product.OfferPrize = product.Price - amount
		err := repository.UpdateProduct(product)
		if err != nil {
			return nil, err
		}
	}
	return productlist, nil

}

func DashBoard() (models.CompleteAdminDashboard, error) {
	userDetails, err := repository.DashBoardUserDetails()
	if err != nil {
		return models.CompleteAdminDashboard{}, err
	}
	productDetails, err := repository.DashBoardProductDetails()
	if err != nil {
		return models.CompleteAdminDashboard{}, err
	}
	orderDetails, err := repository.DashBoardOrder()
	if err != nil {
		return models.CompleteAdminDashboard{}, err
	}
	totalRevenue, err := repository.TotalRevenue()
	if err != nil {
		return models.CompleteAdminDashboard{}, err
	}
	amountDetails, err := repository.AmountDetails()
	if err != nil {
		return models.CompleteAdminDashboard{}, err
	}
	return models.CompleteAdminDashboard{
		DashboardUser:    userDetails,
		DashboardProduct: productDetails,
		DashboardOrder:   orderDetails,
		DashboardRevenue: totalRevenue,
		DashboardAmount:  amountDetails,
	}, nil
}


