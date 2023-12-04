package usecase

import (
	"errors"
	"fmt"
	"os"

	"github.com/razorpay/razorpay-go"
	"main.go/models"
	"main.go/repository"
)

func MakePaymentRazorPay(orderID int) (models.Payment, string, error) {
	PaymentDetails, err := repository.GetPaymentDetails(orderID)
	if err != nil {
		return models.Payment{}, "", err
	}

	client := razorpay.NewClient(os.Getenv("KEY_ID_PAY"), os.Getenv("KEY_SECRET_PAY"))

	data := map[string]interface{}{
		"amount":   int(PaymentDetails.Final_price * 100),
		"currency": "INR",
		"receipt":  "some_receipt_id",
	}
	body, err := client.Order.Create(data, nil)
	if err != nil {
		fmt.Println("hello")
		fmt.Println(err)
		return models.Payment{}, "", err
	}

	razorPayOrderID := body["id"].(string)

	err = repository.AddRazorPayDetails(orderID, razorPayOrderID)
	if err != nil {
		fmt.Println("hig")
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

func SavePaymentDetails(orderID int, paymentID string) error {
	status, err := repository.CheckPaymentStatus(orderID)
	if err != nil {
		return err
	}
	if status == "not paid" {
		err = repository.UpdatePaymentDetails(orderID, paymentID)
		if err != nil {
			return err
		}
		err := repository.UpdateShipmentAndPaymentByOrderID("processing", "paid", orderID)
		if err != nil {
			return err
		}
		return nil
	}
	return errors.New("already paid")
}
