package main

import (
	"fmt"
	"os"

	"github.com/go-resty/resty/v2"
	"github.com/rs/zerolog"

	"github.com/sendlovebox/watu-sdk/api"
	"github.com/sendlovebox/watu-sdk/model"
)

func main() {
	fmt.Println("Running Watu SDK")
	_ = os.Setenv("TZ", "Africa/Lagos")
	logger := zerolog.New(os.Stderr).With().Timestamp().Logger()
	client := resty.New()
	logger.Info().Msg("app starting")
	defer logger.Info().Msg("stopped")
	api := api.New(&logger, client, model.BaseURL)
	api.RunInSandboxMode() // to ensure it is running in sandbox mode

	// we need to authenticate the application
	_ = api.Auth("", "", "", "")

	////=================================>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>
	//// Get Bill Groups
	//billGroups, err := api.GetBillGroups(context.Background())
	//if err != nil {
	//	fmt.Printf("err ===> %+v", err)
	//}
	//fmt.Printf("resp ===> %+v", billGroups)

	////=================================>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>
	//// Get Bill By Groups
	//billByGroups, err := api.GetBillsByGroup(context.Background(), "airtime")
	//if err != nil {
	//	fmt.Printf("err ===> %+v", err)
	//}
	//fmt.Printf("resp ===> %+v", billByGroups)

	////=================================>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>
	//// Get Bill Information
	//billInfo, err := api.GetBillInformation(context.Background(), "bill-18")
	//if err != nil {
	//	fmt.Printf("err ===> %+v", err)
	//}
	//fmt.Printf("resp ===> %+v", billInfo)

	////=================================>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>
	//// Get Bill Types
	//billTypes, err := api.GetBillTypes(context.Background(), "bill-18")
	//if err != nil {
	//	fmt.Printf("err ===> %+v", err)
	//}
	//fmt.Printf("resp ===> %+v", billTypes)

	////=================================>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>
	//// Validate Bill
	//body := struct {
	//	Channel         string `json:"channel"`
	//	SmartCardNumber string `json:"smart_card_number"`
	//}{
	//	Channel:         "bill-20",
	//	SmartCardNumber: "10210333125",
	//}

	//body2 := struct {
	//	Channel           string `json:"channel"`
	//	CustomerAccountID string `json:"customer_account_id"`
	//}{
	//	Channel:           "bill-22",
	//	CustomerAccountID: "2206005375",
	//}
	//
	//resp, err := api.ValidateBill(context.Background(), body2)
	//if err != nil {
	//	fmt.Printf("err ===> %+v", err)
	//}
	//fmt.Printf("resp ===> %+v", resp)
	//
	////=================================>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>
	//// Vend Bill
	////this is for airtime
	//body := struct {
	//	Channel           string `json:"channel"`
	//	PhoneNumber       string `json:"phone_number"`
	//	Amount            int    `json:"amount"`
	//	BusinessSignature string `json:"business_signature"`
	//}{
	//	Channel:           "bill-25",
	//	PhoneNumber:       "08094656624",
	//	Amount:            200,
	//	BusinessSignature: "smh",
	//}
	//
	//resp, err := api.VendBill(context.Background(), body)
	//if err != nil {
	//	fmt.Printf("err ===> %+v", err)
	//}
	//fmt.Printf("resp ===> %+v", resp)

}
