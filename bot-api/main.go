package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"sync"
	"time"
)

type OrderPayload struct {
	Token    string `json:"token"`
	Code     string `json:"code"`
	Comments struct {
		CustomerComment string `json:"customerComment"`
		VendorComment   string `json:"vendorComment"`
	} `json:"comments"`
	CreatedAt time.Time `json:"createdAt"`
	Customer  struct {
		Email                  string `json:"email"`
		FirstName              string `json:"firstName"`
		LastName               string `json:"lastName"`
		MobilePhone            string `json:"mobilePhone"`
		Code                   string `json:"code"`
		ID                     string `json:"id"`
		MobilePhoneCountryCode string `json:"mobilePhoneCountryCode"`
	} `json:"customer"`
	Delivery struct {
		Address struct {
			Postcode int    `json:"postcode"`
			City     string `json:"city"`
			Street   string `json:"street"`
			Number   int    `json:"number"`
		} `json:"address"`
		ExpectedDeliveryTime time.Time `json:"expectedDeliveryTime"`
		ExpressDelivery      bool      `json:"expressDelivery"`
		RiderPickupTime      time.Time `json:"riderPickupTime"`
	} `json:"delivery"`
	Discounts []struct {
		Name   string `json:"name"`
		Amount string `json:"amount"`
		Type   string `json:"type"`
	} `json:"discounts"`
	ExpeditionType  string    `json:"expeditionType"`
	ExpiryDate      time.Time `json:"expiryDate"`
	ExtraParameters struct {
		Property1 string `json:"property1"`
		Property2 string `json:"property2"`
	} `json:"extraParameters"`
	LocalInfo struct {
		CountryCode            string `json:"countryCode"`
		CurrencySymbol         string `json:"currencySymbol"`
		Platform               string `json:"platform"`
		PlatformKey            string `json:"platformKey"`
		CurrencySymbolPosition string `json:"currencySymbolPosition"`
		CurrencySymbolSpaces   string `json:"currencySymbolSpaces"`
		DecimalDigits          string `json:"decimalDigits"`
		DecimalSeparator       string `json:"decimalSeparator"`
		Email                  string `json:"email"`
		Phone                  string `json:"phone"`
		ThousandsSeparator     string `json:"thousandsSeparator"`
		Website                string `json:"website"`
	} `json:"localInfo"`
	Payment struct {
		Status              string `json:"status"`
		Type                string `json:"type"`
		RemoteCode          string `json:"remoteCode"`
		RequiredMoneyChange string `json:"requiredMoneyChange"`
		VatID               string `json:"vatId"`
		VatName             string `json:"vatName"`
	} `json:"payment"`
	Test               bool        `json:"test"`
	ShortCode          string      `json:"shortCode"`
	PreOrder           bool        `json:"preOrder"`
	Pickup             interface{} `json:"pickup"`
	PlatformRestaurant struct {
		ID string `json:"id"`
	} `json:"platformRestaurant"`
	Price struct {
		DeliveryFees []struct {
			Name  string  `json:"name"`
			Value float64 `json:"value"`
		} `json:"deliveryFees"`
		GrandTotal                       string `json:"grandTotal"`
		MinimumDeliveryValue             string `json:"minimumDeliveryValue"`
		PayRestaurant                    string `json:"payRestaurant"`
		RiderTip                         string `json:"riderTip"`
		SubTotal                         string `json:"subTotal"`
		VatTotal                         string `json:"vatTotal"`
		Comission                        string `json:"comission"`
		ContainerCharge                  string `json:"containerCharge"`
		DeliveryFee                      string `json:"deliveryFee"`
		CollectFromCustomer              string `json:"collectFromCustomer"`
		DiscountAmountTotal              string `json:"discountAmountTotal"`
		DeliveryFeeDiscount              string `json:"deliveryFeeDiscount"`
		ServiceFeePercent                string `json:"serviceFeePercent"`
		ServiceFeeTotal                  string `json:"serviceFeeTotal"`
		ServiceTax                       int    `json:"serviceTax"`
		ServiceTaxValue                  int    `json:"serviceTaxValue"`
		DifferenceToMinimumDeliveryValue string `json:"differenceToMinimumDeliveryValue"`
		VatVisible                       bool   `json:"vatVisible"`
		VatPercent                       string `json:"vatPercent"`
	} `json:"price"`
	Products []struct {
		CategoryName     string `json:"categoryName"`
		Name             string `json:"name"`
		PaidPrice        string `json:"paidPrice"`
		Quantity         string `json:"quantity"`
		RemoteCode       string `json:"remoteCode"`
		SelectedToppings []struct {
			Children   []interface{} `json:"children"`
			Name       string        `json:"name"`
			Price      string        `json:"price"`
			Quantity   int           `json:"quantity"`
			ID         string        `json:"id"`
			RemoteCode interface{}   `json:"remoteCode"`
			Type       string        `json:"type"`
		} `json:"selectedToppings"`
		UnitPrice       string        `json:"unitPrice"`
		Comment         string        `json:"comment"`
		Description     string        `json:"description"`
		DiscountAmount  string        `json:"discountAmount"`
		HalfHalf        bool          `json:"halfHalf"`
		ID              string        `json:"id"`
		SelectedChoices []interface{} `json:"selectedChoices"`
		Variation       struct {
			Name string `json:"name"`
		} `json:"variation"`
		VatPercentage string `json:"vatPercentage"`
	} `json:"products"`
	CorporateOrder  bool   `json:"corporateOrder"`
	CorporateTaxID  string `json:"corporateTaxId"`
	IntegrationInfo struct {
	} `json:"integrationInfo"`
	MobileOrder bool          `json:"mobileOrder"`
	WebOrder    bool          `json:"webOrder"`
	Vouchers    []interface{} `json:"vouchers"`
}

var DefaultPayload = &OrderPayload{
	Token: "",
	Code:  "",
	Comments: struct {
		CustomerComment string `json:"customerComment"`
		VendorComment   string `json:"vendorComment"`
	}{},
	CreatedAt: time.Time{},
	Customer: struct {
		Email                  string `json:"email"`
		FirstName              string `json:"firstName"`
		LastName               string `json:"lastName"`
		MobilePhone            string `json:"mobilePhone"`
		Code                   string `json:"code"`
		ID                     string `json:"id"`
		MobilePhoneCountryCode string `json:"mobilePhoneCountryCode"`
	}{},
	Delivery: struct {
		Address struct {
			Postcode int    `json:"postcode"`
			City     string `json:"city"`
			Street   string `json:"street"`
			Number   int    `json:"number"`
		} `json:"address"`
		ExpectedDeliveryTime time.Time `json:"expectedDeliveryTime"`
		ExpressDelivery      bool      `json:"expressDelivery"`
		RiderPickupTime      time.Time `json:"riderPickupTime"`
	}{},
	Discounts:      nil,
	ExpeditionType: "",
	ExpiryDate:     time.Time{},
	ExtraParameters: struct {
		Property1 string `json:"property1"`
		Property2 string `json:"property2"`
	}{},
	LocalInfo: struct {
		CountryCode            string `json:"countryCode"`
		CurrencySymbol         string `json:"currencySymbol"`
		Platform               string `json:"platform"`
		PlatformKey            string `json:"platformKey"`
		CurrencySymbolPosition string `json:"currencySymbolPosition"`
		CurrencySymbolSpaces   string `json:"currencySymbolSpaces"`
		DecimalDigits          string `json:"decimalDigits"`
		DecimalSeparator       string `json:"decimalSeparator"`
		Email                  string `json:"email"`
		Phone                  string `json:"phone"`
		ThousandsSeparator     string `json:"thousandsSeparator"`
		Website                string `json:"website"`
	}{},
	Payment: struct {
		Status              string `json:"status"`
		Type                string `json:"type"`
		RemoteCode          string `json:"remoteCode"`
		RequiredMoneyChange string `json:"requiredMoneyChange"`
		VatID               string `json:"vatId"`
		VatName             string `json:"vatName"`
	}{},
	Test:      false,
	ShortCode: "",
	PreOrder:  false,
	Pickup:    nil,
	PlatformRestaurant: struct {
		ID string `json:"id"`
	}{},
	Price: struct {
		DeliveryFees []struct {
			Name  string  `json:"name"`
			Value float64 `json:"value"`
		} `json:"deliveryFees"`
		GrandTotal                       string `json:"grandTotal"`
		MinimumDeliveryValue             string `json:"minimumDeliveryValue"`
		PayRestaurant                    string `json:"payRestaurant"`
		RiderTip                         string `json:"riderTip"`
		SubTotal                         string `json:"subTotal"`
		VatTotal                         string `json:"vatTotal"`
		Comission                        string `json:"comission"`
		ContainerCharge                  string `json:"containerCharge"`
		DeliveryFee                      string `json:"deliveryFee"`
		CollectFromCustomer              string `json:"collectFromCustomer"`
		DiscountAmountTotal              string `json:"discountAmountTotal"`
		DeliveryFeeDiscount              string `json:"deliveryFeeDiscount"`
		ServiceFeePercent                string `json:"serviceFeePercent"`
		ServiceFeeTotal                  string `json:"serviceFeeTotal"`
		ServiceTax                       int    `json:"serviceTax"`
		ServiceTaxValue                  int    `json:"serviceTaxValue"`
		DifferenceToMinimumDeliveryValue string `json:"differenceToMinimumDeliveryValue"`
		VatVisible                       bool   `json:"vatVisible"`
		VatPercent                       string `json:"vatPercent"`
	}{},
	Products:        nil,
	CorporateOrder:  false,
	CorporateTaxID:  "",
	IntegrationInfo: struct{}{},
	MobileOrder:     false,
	WebOrder:        false,
	Vouchers:        nil,
}

func main() {

	var wg sync.WaitGroup
	for i := 0; i < 10000; i++ {
		wg.Add(1)
		go func(index int, wg *sync.WaitGroup) {
			defer wg.Done()
			newPayload := new(OrderPayload)
			*newPayload = *DefaultPayload
			newPayload.Code = fmt.Sprintf("CODE %d", index+1)

			jsonModel, _ := json.Marshal(newPayload)

			resp, err := http.Post("http://localhost:8081/order", "application/json",
				bytes.NewBuffer(jsonModel))

			if err != nil {
				fmt.Println(err.Error())
			}

			fmt.Printf("Response received -> %d: %s\n", index, resp.Status)

		}(i, &wg)
		time.Sleep(20 * time.Millisecond)

	}
	wg.Wait()
}
