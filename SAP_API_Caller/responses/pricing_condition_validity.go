package responses

type PricingConditionValidity struct {
	D struct {
		Results []struct {
			Metadata struct {
				ID   string `json:"id"`
				URI  string `json:"uri"`
				Type string `json:"type"`
				Etag string `json:"etag"`
			} `json:"__metadata"`
			ConditionRecord                string      `json:"ConditionRecord"`
			ConditionValidityEndDate       string      `json:"ConditionValidityEndDate"`
			ConditionValidityStartDate     string      `json:"ConditionValidityStartDate"`
			ConditionApplication           string      `json:"ConditionApplication"`
			ConditionType                  string      `json:"ConditionType"`
			ConditionReleaseStatus         string      `json:"ConditionReleaseStatus"`
			SalesDocument                  string      `json:"SalesDocument"`
			SalesDocumentItem              string      `json:"SalesDocumentItem"`
			ConditionContract              string      `json:"ConditionContract"`
			CustomerGroup                  string      `json:"CustomerGroup"`
			CustomerPriceGroup             string      `json:"CustomerPriceGroup"`
			MaterialPricingGroup           string      `json:"MaterialPricingGroup"`
			SoldToParty                    string      `json:"SoldToParty"`
			BPForSoldToParty               string      `json:"BPForSoldToParty"`
			Customer                       string      `json:"Customer"`
			BPForCustomer                  string      `json:"BPForCustomer"`
			PayerParty                     string      `json:"PayerParty"`
			BPForPayerParty                string      `json:"BPForPayerParty"`
			ShipToParty                    string      `json:"ShipToParty"`
			BPForShipToParty               string      `json:"BPForShipToParty"`
			Supplier                       string      `json:"Supplier"`
			BPForSupplier                  string      `json:"BPForSupplier"`
			MaterialGroup                  string      `json:"MaterialGroup"`
			Material                       string      `json:"Material"`
			PriceListType                  string      `json:"PriceListType"`
			CustomerTaxClassification1     string      `json:"CustomerTaxClassification1"`
			ProductTaxClassification1      string      `json:"ProductTaxClassification1"`
			SDDocument                     string      `json:"SDDocument"`
			ReferenceSDDocument            string      `json:"ReferenceSDDocument"`
			ReferenceSDDocumentItem        string      `json:"ReferenceSDDocumentItem"`
			SalesOffice                    string      `json:"SalesOffice"`
			SalesGroup                     string      `json:"SalesGroup"`
			SalesOrganization              string      `json:"SalesOrganization"`
			DistributionChannel            string      `json:"DistributionChannel"`
			TransactionCurrency            string      `json:"TransactionCurrency"`
			ConditionProcessingStatus      string      `json:"ConditionProcessingStatus"`
			PricingDate                    string      `json:"PricingDate"`
			ConditionScaleBasisValue       string      `json:"ConditionScaleBasisValue"`
			TaxCode                        string      `json:"TaxCode"`
			ServiceDocument                string      `json:"ServiceDocument"`
			ServiceDocumentItem            string      `json:"ServiceDocumentItem"`
			CustomerConditionGroup         string      `json:"CustomerConditionGroup"`
			ToConditionRecord struct {
				Deferred struct {
					URI string `json:"uri"`
				} `json:"__deferred"`
			} `json:"to_SlsPrcgConditionRecord"`
		} `json:"results"`
	} `json:"d"`
}
