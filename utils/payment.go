package utils

import (
	"fmt"

	"github.com/razorpay/razorpay-go/utils"
)

func VerifyPayment(orderID, paymentID, signature, razopaySecret string) bool {

	params := map[string]interface{}{
		"razorpay_order_id":   orderID,
		"razorpay_payment_id": paymentID,
	}

	// secret := "qvxbhiwTJZLHHE3tNQQv8Mty"
	result := utils.VerifyPaymentSignature(params, signature, razopaySecret)
	fmt.Println("*****", result)
	return result
}
