package responses

type ToConditionRecord struct {
	D struct {
		Metadata struct {
			ID   string `json:"id"`
			URI  string `json:"uri"`
			Type string `json:"type"`
			Etag string `json:"etag"`
		} `json:"__metadata"`
		ConditionRecord              string      `json:"ConditionRecord"`
		ConditionSequentialNumber    string      `json:"ConditionSequentialNumber"`
		ConditionTable               string      `json:"ConditionTable"`
		ConditionApplication         string      `json:"ConditionApplication"`
		ConditionType                string      `json:"ConditionType"`
		ConditionValidityEndDate     string      `json:"ConditionValidityEndDate"`
		ConditionValidityStartDate   string      `json:"ConditionValidityStartDate"`
		CreationDate                 string      `json:"CreationDate"`
		PricingScaleType             string      `json:"PricingScaleType"`
		PricingScaleBasis            string      `json:"PricingScaleBasis"`
		ConditionScaleQuantity       string      `json:"ConditionScaleQuantity"`
		ConditionScaleQuantityUnit   string      `json:"ConditionScaleQuantityUnit"`
		ConditionScaleAmount         string      `json:"ConditionScaleAmount"`
		ConditionScaleAmountCurrency string      `json:"ConditionScaleAmountCurrency"`
		ConditionCalculationType     string      `json:"ConditionCalculationType"`
		ConditionRateValue           string      `json:"ConditionRateValue"`
		ConditionRateValueUnit       string      `json:"ConditionRateValueUnit"`
		ConditionQuantity            string      `json:"ConditionQuantity"`
		ConditionQuantityUnit        string      `json:"ConditionQuantityUnit"`
		BaseUnit                     string      `json:"BaseUnit"`
		ConditionIsDeleted           bool        `json:"ConditionIsDeleted"`
		PaymentTerms                 string      `json:"PaymentTerms"`
		IncrementalScale             string      `json:"IncrementalScale"`
		PricingScaleLine             string      `json:"PricingScaleLine"`
		ConditionReleaseStatus       string      `json:"ConditionReleaseStatus"`
	} `json:"d"`
}
