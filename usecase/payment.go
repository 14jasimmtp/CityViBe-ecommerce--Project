package usecase

import (
	"errors"
	"fmt"
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



func VerifyPayment(details models.PaymentVerify, order_id int) (models.OrderDetails, error) {
	paid, err := repository.CheckVerifiedPayment(order_id)
	if err != nil {
		return models.OrderDetails{}, err
	}
	if paid {
		return models.OrderDetails{}, errors.New(`already payment verified`)
	}

	result := utils.VerifyPayment(details.OrderID, details.PaymentID, details.Signature, os.Getenv("KEY_SECRET_PAY"))
	if !result {
		return models.OrderDetails{}, errors.New("payment is unsuccessful")
	}

	orders, err := repository.UpdateShipmentAndPaymentByOrderID("processing", "paid", order_id)
	if err != nil {
		return models.OrderDetails{}, err
	}
	return orders, nil
}
