package sap_api_output_formatter

import (
	"encoding/json"
	"sap-api-integrations-sales-pricing-reads-rmq-kube/SAP_API_Caller/responses"

	"github.com/latonaio/golang-logging-library/logger"
	"golang.org/x/xerrors"
)

func ConvertToPricingConditionValidity(raw []byte, l *logger.Logger) ([]PricingConditionValidity, error) {
	pm := &responses.PricingConditionValidity{}

	err := json.Unmarshal(raw, pm)
	if err != nil {
		return nil, xerrors.Errorf("cannot convert to PricingConditionValidity. unmarshal error: %w", err)
	}
	if len(pm.D.Results) == 0 {
		return nil, xerrors.New("Result data is not exist")
	}
	if len(pm.D.Results) > 10 {
		l.Info("raw data has too many Results. %d Results exist. show the first 10 of Results array", len(pm.D.Results))
	}

	pricingConditionValidity := make([]PricingConditionValidity, 0, 10)
	for i := 0; i < 10 && i < len(pm.D.Results); i++ {
		data := pm.D.Results[i]
		pricingConditionValidity = append(pricingConditionValidity, PricingConditionValidity{
			ConditionRecord:            data.ConditionRecord,
			ConditionValidityEndDate:   data.ConditionValidityEndDate,
			ConditionValidityStartDate: data.ConditionValidityStartDate,
			ConditionApplication:       data.ConditionApplication,
			ConditionType:              data.ConditionType,
			ConditionReleaseStatus:     data.ConditionReleaseStatus,
			SalesDocument:              data.SalesDocument,
			SalesDocumentItem:          data.SalesDocumentItem,
			ConditionContract:          data.ConditionContract,
			CustomerGroup:              data.CustomerGroup,
			CustomerPriceGroup:         data.CustomerPriceGroup,
			MaterialPricingGroup:       data.MaterialPricingGroup,
			SoldToParty:                data.SoldToParty,
			BPForSoldToParty:           data.BPForSoldToParty,
			Customer:                   data.Customer,
			BPForCustomer:              data.BPForCustomer,
			PayerParty:                 data.PayerParty,
			BPForPayerParty:            data.BPForPayerParty,
			ShipToParty:                data.ShipToParty,
			BPForShipToParty:           data.BPForShipToParty,
			Supplier:                   data.Supplier,
			BPForSupplier:              data.BPForSupplier,
			MaterialGroup:              data.MaterialGroup,
			Material:                   data.Material,
			PriceListType:              data.PriceListType,
			CustomerTaxClassification1: data.CustomerTaxClassification1,
			ProductTaxClassification1:  data.ProductTaxClassification1,
			SDDocument:                 data.SDDocument,
			ReferenceSDDocument:        data.ReferenceSDDocument,
			ReferenceSDDocumentItem:    data.ReferenceSDDocumentItem,
			SalesOffice:                data.SalesOffice,
			SalesGroup:                 data.SalesGroup,
			SalesOrganization:          data.SalesOrganization,
			DistributionChannel:        data.DistributionChannel,
			TransactionCurrency:        data.TransactionCurrency,
			ConditionProcessingStatus:  data.ConditionProcessingStatus,
			PricingDate:                data.PricingDate,
			ConditionScaleBasisValue:   data.ConditionScaleBasisValue,
			TaxCode:                    data.TaxCode,
			ServiceDocument:            data.ServiceDocument,
			ServiceDocumentItem:        data.ServiceDocumentItem,
			CustomerConditionGroup:     data.CustomerConditionGroup,
			ToConditionRecord:          data.ToConditionRecord.Deferred.URI,
		})
	}

	return pricingConditionValidity, nil
}

func ConvertToToConditionRecord(raw []byte, l *logger.Logger) (*ToConditionRecord, error) {
	pm := &responses.ToConditionRecord{}

	err := json.Unmarshal(raw, &pm)
	if err != nil {
		return nil, xerrors.Errorf("cannot convert to ToConditionRecord. unmarshal error: %w", err)
	}

	return &ToConditionRecord{
		ConditionRecord:              pm.D.ConditionRecord,
		ConditionSequentialNumber:    pm.D.ConditionSequentialNumber,
		ConditionTable:               pm.D.ConditionTable,
		ConditionApplication:         pm.D.ConditionApplication,
		ConditionType:                pm.D.ConditionType,
		ConditionValidityEndDate:     pm.D.ConditionValidityEndDate,
		ConditionValidityStartDate:   pm.D.ConditionValidityStartDate,
		CreationDate:                 pm.D.CreationDate,
		PricingScaleType:             pm.D.PricingScaleType,
		PricingScaleBasis:            pm.D.PricingScaleBasis,
		ConditionScaleQuantity:       pm.D.ConditionScaleQuantity,
		ConditionScaleQuantityUnit:   pm.D.ConditionScaleQuantityUnit,
		ConditionScaleAmount:         pm.D.ConditionScaleAmount,
		ConditionScaleAmountCurrency: pm.D.ConditionScaleAmountCurrency,
		ConditionCalculationType:     pm.D.ConditionCalculationType,
		ConditionRateValue:           pm.D.ConditionRateValue,
		ConditionRateValueUnit:       pm.D.ConditionRateValueUnit,
		ConditionQuantity:            pm.D.ConditionQuantity,
		ConditionQuantityUnit:        pm.D.ConditionQuantityUnit,
		BaseUnit:                     pm.D.BaseUnit,
		ConditionIsDeleted:           pm.D.ConditionIsDeleted,
		PaymentTerms:                 pm.D.PaymentTerms,
		IncrementalScale:             pm.D.IncrementalScale,
		PricingScaleLine:             pm.D.PricingScaleLine,
		ConditionReleaseStatus:       pm.D.ConditionReleaseStatus,
	}, nil
}
