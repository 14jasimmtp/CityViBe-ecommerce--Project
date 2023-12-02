package usecase

import (
	"errors"
	"os"

	"github.com/razorpay/razorpay-go"
	"main.go/models"
	"main.go/repository"
	"main.go/utils"
)

func MakePaymentRazorPay(orderID int) (models.Payment, string, error) {
	PaymentDetails, err := repository.GetPaymentDetails(orderID)
	if err != nil {
		return models.Payment{}, "", err
	}

	client := razorpay.NewClient(os.Getenv("KEY_ID_PAY"), os.Getenv("KEY_SECRET_PAY"))

	data := map[string]interface{}{
		"amount":   PaymentDetails.Final_price * 100,
		"currency": "INR",
		"receipt":  "some_receipt_id",
	}
	body, err := client.Order.Create(data, nil)
	if err != nil {

		return models.Payment{}, "", err
	}

	razorPayOrderID := body["id"].(string)

	err = repository.AddRazorPayDetails(orderID, razorPayOrderID)
	if err != nil {
		return models.Payment{}, "", err
	}

	return PaymentDetails, razorPayOrderID, nil

}

func PaymentMethodID(orderID int) (int, error) {
	PaymethodID, err := repository.PayMethod(orderID)
	if err != nil {
		return 0, err
	}
	return PaymethodID, nil
}

func PaymentAlreadyPaid(orderID int) (bool, error) {
	AlreadyPayed, err := repository.PaymentAlreadyPaid(orderID)
	if err != nil {
		return false, err
	}
	return AlreadyPayed, nil
}

func PaymentVerification(details models.PaymentVerify) (models.Invoice, error) {
	result := utils.VerifyPayment(details.OrderID, details.PaymentID, details.Signature,os.Getenv(`KEY_SECRET_PAY`))
	if !result {
		return models.Invoice{}, errors.New(`payment not verified`)
	}

	orders, err := repository.UpdatePaymentStatus(details.OrderID)
	if err != nil {
		return models.Invoice{}, err
	}
	return orders, nil
}
