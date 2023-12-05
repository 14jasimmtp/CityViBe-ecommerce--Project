package repository

import (
	"errors"

	initialisers "main.go/Initialisers"
	"main.go/domain"
	"main.go/models"
)

func CheckCouponExist(coupon string) (bool, error) {
	var couponvalid domain.Coupon
	query := `SELECT * FROM coupons WHERE coupon = ?`

	db := initialisers.DB.Raw(query, coupon).Scan(&couponvalid)
	if db.Error != nil {
		return true, errors.New(`something went wrong`)
	}

	if db.RowsAffected > 0 {
		return true, errors.New(`already exist`)
	}

	return false, nil

}

func CheckCouponExistWithID(coupon int) (bool, error) {
	var couponvalid domain.Coupon
	query := `SELECT * FROM coupons WHERE id = ?`

	db := initialisers.DB.Raw(query, coupon).Scan(&couponvalid)
	if db.Error != nil {
		return true, errors.New(`something went wrong`)
	}

	if db.RowsAffected > 0 {
		return true, nil
	}

	return false, nil
}

func CreateCoupon(coupon models.Coupon) (domain.Coupon, error) {
	var Coupons domain.Coupon
	query := initialisers.DB.Raw(`INSERT INTO coupons (coupon,discount_percentage,usage_limit) VALUES (?,?,?) RETURNING id,coupon,discount_percentage,usage_limit,active`, coupon.Coupon, coupon.DiscoutPercentage, coupon.UsageLimit).Scan(&Coupons)
	if query.Error != nil {
		return domain.Coupon{}, errors.New(`something went wrong`)
	}
	return Coupons, nil
}

func DisableCoupon(coupon uint) error {
	query := initialisers.DB.Exec(`UPDATE coupons SET active = false WHERE id = ?`, coupon)
	if query.RowsAffected < 1 {
		return errors.New(`no coupons found with this id`)
	}
	if query.Error != nil {
		return errors.New(`something went wrong`)
	}
	return nil
}

func EnableCoupon(coupon uint) error {
	query := initialisers.DB.Exec(`UPDATE coupons SET active = true WHERE id = ?`, coupon)
	if query.RowsAffected < 1 {
		return errors.New(`no coupons found with this id`)
	}
	if query.Error != nil {
		return errors.New(`something went wrong`)
	}
	return nil
}

func GetCouponsForAdmin() ([]domain.Coupon, error) {
	var Coupons []domain.Coupon
	query := initialisers.DB.Raw(`SELECT * FROM coupons`).Scan(&Coupons)
	if query.Error != nil {
		return []domain.Coupon{}, errors.New(`something went wrong`)
	}
	if query.RowsAffected < 1 {
		return []domain.Coupon{}, errors.New(`no coupons added.Add a coupon to view`)
	}
	return Coupons, nil
}

func UpdateCoupon(coupon models.Coupon, coupon_id string) (domain.Coupon, error) {
	var coupons domain.Coupon
	query := initialisers.DB.Raw(`UPDATE coupons SET coupon = ? ,discount_percentage = ? RETURNING id,coupon,discount_percentage,valid`, coupon.Coupon, coupon.DiscoutPercentage).Scan(&coupons)
	if query.Error != nil {
		return domain.Coupon{}, errors.New(`something went wrong`)
	}
	return coupons, nil
}

func GetDiscountRate(coupon string) (float64, error) {
	var discountRate float64
	query := initialisers.DB.Raw(`SELECT discount_percentage from coupons WHERE coupon = ? AND active = true`, coupon).Scan(&discountRate)
	if query.Error != nil {
		return 0.0, errors.New(`something went wrong`)
	}
	if query.RowsAffected < 1 {
		return 0.0, errors.New(`no coupons found`)
	}
	return discountRate, nil
}

func UpdateCouponUsage(userID uint, coupon string) error {
	query := initialisers.DB.Exec(`insert into used_coupons (user_id,coupon) values(?,?)`, userID, coupon)
	if query.Error != nil {
		return errors.New(`something went wrong`)
	}
	return nil
}

func UpdateCouponCount(coupon string) error {
	query := initialisers.DB.Exec(`UPDATE coupons SET usage_limit = usage_limit - 1`)
	if query.Error != nil {
		return errors.New(`something went wrong`)
	}
	return nil
}

func CheckCouponUsage(userId uint, coupon string) error {
	var count int
	query := initialisers.DB.Raw(`SELECT count(*) from used_coupons where user_id = ? AND coupon = ?`, userId, coupon).Scan(&count)
	if query.Error != nil {
		return errors.New(`something went wrong`)
	}
	if count >= 1 {
		return errors.New(`coupon already used`)
	}
	return nil
}

func ViewUserCoupons(userID uint) ([]models.Couponlist, error) {
	var coupons []models.Couponlist
	query := initialisers.DB.Raw(`SELECT coupon,discount_percentage FROM coupons `).Scan(&coupons)
	if query.Error != nil {
		return []models.Couponlist{}, errors.New(`something went wrong`)
	}
	if query.RowsAffected == 0 {
		return []models.Couponlist{}, errors.New(`no coupons found`)
	}
	return coupons, nil
}
