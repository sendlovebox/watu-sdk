package model

type (
	// BillGroup schema
	BillGroup struct {
		ID                         string `json:"id"`
		Name                       string `json:"name"`
		Description                string `json:"description"`
		Status                     bool   `json:"status"`
		Icon                       string `json:"icon"`
		Icon2                      string `json:"icon_2"`
		FixedCharge                string `json:"fixed_charge"`
		PercentCharge              string `json:"percent_charge"`
		CurrencyCode               string `json:"currency_code"`
		InternationalPercentCharge string `json:"international_percent_charge"`
		InternationalFixedCharge   string `json:"international_fixed_charge"`
	}

	// Bill schema
	Bill struct {
		ID             string `json:"id"`
		Name           string `json:"name"`
		Description    string `json:"description"`
		Icon           string `json:"icon"`
		HasType        bool   `json:"has_type"`
		CanValidate    bool   `json:"can_validate"`
		HasPhoneNumber bool   `json:"has_phone_number"`
		RequestDetails []struct {
			Name          string      `json:"name"`
			Parameter     string      `json:"parameter"`
			Description   interface{} `json:"description"`
			IsRequired    bool        `json:"is_required"`
			IsType        bool        `json:"is_type"`
			InType        bool        `json:"in_type"`
			IsAccount     int         `json:"is_account"`
			IsAmount      int         `json:"is_amount"`
			InValidate    bool        `json:"in_validate"`
			Validator     string      `json:"validator"`
			ConcatParams  interface{} `json:"concat_params"`
			ParameterType string      `json:"parameter_type"`
			Options       []struct {
				Label string `json:"label"`
				Value string `json:"value"`
			} `json:"options"`
		} `json:"request_details"`
		ValidateParameters []struct {
			Name          string      `json:"name"`
			Parameter     string      `json:"parameter"`
			Description   interface{} `json:"description"`
			IsRequired    bool        `json:"is_required"`
			IsType        bool        `json:"is_type"`
			InType        bool        `json:"in_type"`
			IsAccount     int         `json:"is_account"`
			IsAmount      int         `json:"is_amount"`
			InValidate    bool        `json:"in_validate"`
			Validator     string      `json:"validator"`
			ConcatParams  interface{} `json:"concat_params"`
			ParameterType string      `json:"parameter_type"`
			Options       []struct {
				Label string `json:"label"`
				Value string `json:"value"`
			} `json:"options"`
		} `json:"validate_parameters"`
		TypeParameters []interface{} `json:"type_parameters"`
		Route          struct {
			PercentCharge              string      `json:"percent_charge"`
			InternationalPercentCharge string      `json:"international_percent_charge"`
			FixedCharge                string      `json:"fixed_charge"`
			InternationalFixedCharge   string      `json:"international_fixed_charge"`
			LocalCap                   interface{} `json:"local_cap"`
			InternationalCap           interface{} `json:"international_cap"`
			FixedCommission            string      `json:"fixed_commission"`
			PercentCommission          string      `json:"percent_commission"`
			CurrencyCode               string      `json:"currency_code"`
			IsFree                     bool        `json:"is_free"`
		} `json:"route"`
		NotificationChannel []string `json:"notification_channel"`
		Partner             struct {
			SquareIcon    string `json:"square_icon"`
			LandscapeIcon string `json:"landscape_icon"`
			Name          string `json:"name"`
			Slug          string `json:"slug"`
		} `json:"partner"`
		Subtitle     string      `json:"subtitle"`
		UssdCode     interface{} `json:"ussd_code"`
		AccountField struct {
			Name          string      `json:"name"`
			Parameter     string      `json:"parameter"`
			Description   interface{} `json:"description"`
			IsRequired    bool        `json:"is_required"`
			IsType        bool        `json:"is_type"`
			InType        bool        `json:"in_type"`
			IsAccount     int         `json:"is_account"`
			IsAmount      int         `json:"is_amount"`
			InValidate    bool        `json:"in_validate"`
			Validator     string      `json:"validator"`
			ConcatParams  interface{} `json:"concat_params"`
			ParameterType string      `json:"parameter_type"`
			Options       []struct {
				Label string `json:"label"`
				Value string `json:"value"`
			} `json:"options"`
		} `json:"account_field"`
	}

	// BillType schema
	BillType struct {
		Name          string `json:"name"`
		ProductCode   string `json:"product_code"`
		Description   string `json:"description"`
		Amount        int    `json:"amount"`
		Validity      int    `json:"validity"`
		InvoicePeriod int    `json:"invoice_period"`
	}

	// VendBillResponse schema
	VendBillResponse struct {
		Feature              string `json:"feature"`
		FeatureId            string `json:"feature_id"`
		FeatureType          string `json:"feature_type"`
		Status               string `json:"status"`
		Message              string `json:"message"`
		TransactionReference string `json:"transaction_reference"`
		DateCreated          string `json:"date_created"`
		MerchantReference    string `json:"merchant_reference"`
		Channel              string `json:"channel"`
		WalletStatement      []struct {
			Amount      string `json:"amount"`
			ValueBefore string `json:"value_before"`
			ValueAfter  string `json:"value_after"`
			Description string `json:"description"`
			Type        string `json:"type"`
			Currency    string `json:"currency"`
		} `json:"wallet_statement"`
		OtherData        []interface{} `json:"other_data"`
		Description      string        `json:"description"`
		TransactionGroup string        `json:"transaction_group"`
		PaymentData      struct {
			Currency      string `json:"currency"`
			Amount        string `json:"amount"`
			Fees          string `json:"fees"`
			FixedCharge   string `json:"fixed_charge"`
			PercentCharge string `json:"percent_charge"`
			TotalAmount   string `json:"total_amount"`
			Account       struct {
				AccountId   string      `json:"account_id"`
				Name        interface{} `json:"name"`
				Email       interface{} `json:"email"`
				PhoneNumber interface{} `json:"phone_number"`
			} `json:"account"`
		} `json:"payment_data"`
		TransactionParameters struct {
			Message    string `json:"message"`
			StatusCode int    `json:"status_code"`
		} `json:"transaction_parameters"`
	}
)
