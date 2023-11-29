package usecase

import (
	"errors"
	"strconv"

	"main.go/domain"
	"main.go/models"
	"main.go/repository"
)

func CreateCoupon(coupon models.Coupon) (domain.Coupon, error) {
	CheckCouponExist, err := repository.CheckCouponExist(coupon.Coupon)
	if CheckCouponExist {
		return domain.Coupon{}, errors.New(`coupon already exist`)
	}
	if err != nil {
		return domain.Coupon{}, err
	}

	Coupon, err := repository.CreateCoupon(coupon)
	if err != nil {
		return domain.Coupon{}, err
	}
	return Coupon, nil
}

func DisableCoupon(coupon uint) error {
	err := repository.DisableCoupon(coupon)
	if err != nil {
		return err
	}
	return nil
}

func EnableCoupon(coupon uint) error {
	err := repository.EnableCoupon(coupon)
	if err != nil {
		return err
	}
	return nil
}

func GetCouponsForAdmin() ([]domain.Coupon, error) {
	Coupons, err := repository.GetCouponsForAdmin()
	if err != nil {
		return []domain.Coupon{}, err
	}
	return Coupons, nil
}

func UpdateCoupon(coupon models.Coupon, coupon_id string) (domain.Coupon, error) {
	cid, err := strconv.Atoi(coupon_id)
	if err != nil {
		return domain.Coupon{}, err
	}
	CheckCoupon, err := repository.CheckCouponExistWithID(cid)

	if !CheckCoupon {
		return domain.Coupon{}, errors.New(`no coupon found with this id`)
	}
	if err != nil {
		return domain.Coupon{}, err
	}
	Coupon, err := repository.UpdateCoupon(coupon, coupon_id)
	if err != nil {
		return domain.Coupon{}, err
	}
	return Coupon, nil
}
