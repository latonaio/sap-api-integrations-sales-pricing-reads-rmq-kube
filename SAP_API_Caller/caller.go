package sap_api_caller

import (
	"fmt"
	"io/ioutil"
	"net/http"
	sap_api_output_formatter "sap-api-integrations-sales-pricing-reads-rmq-kube/SAP_API_Output_Formatter"
	"strings"
	"sync"

	"github.com/latonaio/golang-logging-library/logger"
	"golang.org/x/xerrors"
)

type RMQOutputter interface {
	Send(sendQueue string, payload map[string]interface{}) error
}

type SAPAPICaller struct {
	baseURL      string
	apiKey       string
	outputQueues []string
	outputter    RMQOutputter
	log          *logger.Logger
}

func NewSAPAPICaller(baseUrl string, outputQueueTo []string, outputter RMQOutputter, l *logger.Logger) *SAPAPICaller {
	return &SAPAPICaller{
		baseURL:      baseUrl,
		apiKey:       GetApiKey(),
		outputQueues: outputQueueTo,
		outputter:    outputter,
		log:          l,
	}
}

func (c *SAPAPICaller) AsyncGetSalesPricingCondition(material, distributionChannel, customer, salesOrganization string, accepter []string) {
	wg := &sync.WaitGroup{}
	wg.Add(len(accepter))
	for _, fn := range accepter {
		switch fn {
		case "MaterialDistChannel":
			func() {
				c.MaterialDistChannel(material, distributionChannel)
				wg.Done()
			}()
		case "MaterialDistChannelCustomer":
			func() {
				c.MaterialDistChannelCustomer(material, distributionChannel, customer)
				wg.Done()
			}()
		case "MaterialSalesOrgDistChannel":
			func() {
				c.MaterialSalesOrgDistChannel(material, salesOrganization, distributionChannel)
				wg.Done()
			}()
		case "MaterialSalesOrgDistChannelCustomer":
			func() {
				c.MaterialSalesOrgDistChannelCustomer(material, salesOrganization, distributionChannel, customer)
				wg.Done()
			}()
		default:
			wg.Done()
		}
	}

	wg.Wait()
}

func (c *SAPAPICaller) MaterialDistChannel(material, distributionChannel string) {
	pricingConditionValidityData, err := c.callSalesPricingConditionSrvAPIRequirementMaterialDistChannel("A_SlsPrcgCndnRecdValidity", material, distributionChannel)
	if err != nil {
		c.log.Error(err)
		return
	}
	err = c.outputter.Send(c.outputQueues[0], map[string]interface{}{"message": pricingConditionValidityData, "function": "SalesPricingConditionValidity"})
	if err != nil {
		c.log.Error(err)
		return
	}
	c.log.Info(pricingConditionValidityData)

	conditionRecordData, err := c.callToConditionRecord(pricingConditionValidityData[0].ToConditionRecord)
	if err != nil {
		c.log.Error(err)
		return
	}
	err = c.outputter.Send(c.outputQueues[0], map[string]interface{}{"message": conditionRecordData, "function": "SalesPricingToConditionRecord"})
	if err != nil {
		c.log.Error(err)
		return
	}
	c.log.Info(conditionRecordData)

}

func (c *SAPAPICaller) callSalesPricingConditionSrvAPIRequirementMaterialDistChannel(api, material, distributionChannel string) ([]sap_api_output_formatter.PricingConditionValidity, error) {
	url := strings.Join([]string{c.baseURL, "API_SLSPRICINGCONDITIONRECORD_SRV", api}, "/")
	req, _ := http.NewRequest("GET", url, nil)

	c.setHeaderAPIKeyAccept(req)
	c.getQueryWithMaterialDistChannel(req, material, distributionChannel)

	resp, err := new(http.Client).Do(req)
	if err != nil {
		return nil, xerrors.Errorf("API request error: %w", err)
	}
	defer resp.Body.Close()

	byteArray, _ := ioutil.ReadAll(resp.Body)
	data, err := sap_api_output_formatter.ConvertToPricingConditionValidity(byteArray, c.log)
	if err != nil {
		return nil, xerrors.Errorf("convert error: %w", err)
	}
	return data, nil
}

func (c *SAPAPICaller) callToConditionRecord(url string) (*sap_api_output_formatter.ToConditionRecord, error) {
	req, _ := http.NewRequest("GET", url, nil)
	c.setHeaderAPIKeyAccept(req)

	resp, err := new(http.Client).Do(req)
	if err != nil {
		return nil, xerrors.Errorf("API request error: %w", err)
	}
	defer resp.Body.Close()

	byteArray, _ := ioutil.ReadAll(resp.Body)
	data, err := sap_api_output_formatter.ConvertToToConditionRecord(byteArray, c.log)
	if err != nil {
		return nil, xerrors.Errorf("convert error: %w", err)
	}
	return data, nil
}

func (c *SAPAPICaller) MaterialDistChannelCustomer(material, distributionChannel, customer string) {
	pricingConditionValidityData, err := c.callSalesPricingConditionSrvAPIRequirementMaterialDistChannelCustomer("A_SlsPrcgCndnRecdValidity", material, distributionChannel, customer)
	if err != nil {
		c.log.Error(err)
		return
	}
	err = c.outputter.Send(c.outputQueues[0], map[string]interface{}{"message": pricingConditionValidityData, "function": "SalesPricingConditionValidity"})
	if err != nil {
		c.log.Error(err)
		return
	}
	c.log.Info(pricingConditionValidityData)

	conditionRecordData, err := c.callToConditionRecord(pricingConditionValidityData[0].ToConditionRecord)
	if err != nil {
		c.log.Error(err)
		return
	}
	err = c.outputter.Send(c.outputQueues[0], map[string]interface{}{"message": conditionRecordData, "function": "SalesPricingToConditionRecord"})
	if err != nil {
		c.log.Error(err)
		return
	}
	c.log.Info(conditionRecordData)
}

func (c *SAPAPICaller) callSalesPricingConditionSrvAPIRequirementMaterialDistChannelCustomer(api, material, distributionChannel, customer string) ([]sap_api_output_formatter.PricingConditionValidity, error) {
	url := strings.Join([]string{c.baseURL, "API_SLSPRICINGCONDITIONRECORD_SRV", api}, "/")
	req, _ := http.NewRequest("GET", url, nil)

	c.setHeaderAPIKeyAccept(req)
	c.getQueryWithMaterialDistChannelCustomer(req, material, distributionChannel, customer)

	resp, err := new(http.Client).Do(req)
	if err != nil {
		return nil, xerrors.Errorf("API request error: %w", err)
	}
	defer resp.Body.Close()

	byteArray, _ := ioutil.ReadAll(resp.Body)
	data, err := sap_api_output_formatter.ConvertToPricingConditionValidity(byteArray, c.log)
	if err != nil {
		return nil, xerrors.Errorf("convert error: %w", err)
	}
	return data, nil
}

func (c *SAPAPICaller) MaterialSalesOrgDistChannel(material, salesOrganization, distributionChannel string) {
	pricingConditionValidityData, err := c.callSalesPricingConditionSrvAPIRequirementMaterialSalesOrgDistChannel("A_SlsPrcgCndnRecdValidity", material, salesOrganization, distributionChannel)
	if err != nil {
		c.log.Error(err)
		return
	}
	err = c.outputter.Send(c.outputQueues[0], map[string]interface{}{"message": pricingConditionValidityData, "function": "SalesPricingConditionValidity"})
	if err != nil {
		c.log.Error(err)
		return
	}
	c.log.Info(pricingConditionValidityData)

	conditionRecordData, err := c.callToConditionRecord(pricingConditionValidityData[0].ToConditionRecord)
	if err != nil {
		c.log.Error(err)
		return
	}
	err = c.outputter.Send(c.outputQueues[0], map[string]interface{}{"message": conditionRecordData, "function": "SalesPricingToConditionRecord"})
	if err != nil {
		c.log.Error(err)
		return
	}
	c.log.Info(conditionRecordData)

}

func (c *SAPAPICaller) callSalesPricingConditionSrvAPIRequirementMaterialSalesOrgDistChannel(api, material, salesOrganization, distributionChannel string) ([]sap_api_output_formatter.PricingConditionValidity, error) {
	url := strings.Join([]string{c.baseURL, "API_SLSPRICINGCONDITIONRECORD_SRV", api}, "/")
	req, _ := http.NewRequest("GET", url, nil)

	c.setHeaderAPIKeyAccept(req)
	c.getQueryWithMaterialSalesOrgDistChannel(req, material, salesOrganization, distributionChannel)

	resp, err := new(http.Client).Do(req)
	if err != nil {
		return nil, xerrors.Errorf("API request error: %w", err)
	}
	defer resp.Body.Close()

	byteArray, _ := ioutil.ReadAll(resp.Body)
	data, err := sap_api_output_formatter.ConvertToPricingConditionValidity(byteArray, c.log)
	if err != nil {
		return nil, xerrors.Errorf("convert error: %w", err)
	}
	return data, nil
}

func (c *SAPAPICaller) MaterialSalesOrgDistChannelCustomer(material, salesOrganization, distributionChannel, customer string) {
	pricingConditionValidityData, err := c.callSalesPricingConditionSrvAPIRequirementMaterialSalesOrgDistChannelCustomer("A_SlsPrcgCndnRecdValidity", material, salesOrganization, distributionChannel, customer)
	if err != nil {
		c.log.Error(err)
		return
	}
	err = c.outputter.Send(c.outputQueues[0], map[string]interface{}{"message": pricingConditionValidityData, "function": "SalesPricingConditionValidity"})
	if err != nil {
		c.log.Error(err)
		return
	}
	c.log.Info(pricingConditionValidityData)

	conditionRecordData, err := c.callToConditionRecord(pricingConditionValidityData[0].ToConditionRecord)
	if err != nil {
		c.log.Error(err)
		return
	}
	err = c.outputter.Send(c.outputQueues[0], map[string]interface{}{"message": conditionRecordData, "function": "SalesPricingToConditionRecord"})
	if err != nil {
		c.log.Error(err)
		return
	}
	c.log.Info(conditionRecordData)

}

func (c *SAPAPICaller) callSalesPricingConditionSrvAPIRequirementMaterialSalesOrgDistChannelCustomer(api, material, salesOrganization, distributionChannel, customer string) ([]sap_api_output_formatter.PricingConditionValidity, error) {
	url := strings.Join([]string{c.baseURL, "API_SLSPRICINGCONDITIONRECORD_SRV", api}, "/")
	req, _ := http.NewRequest("GET", url, nil)

	c.setHeaderAPIKeyAccept(req)
	c.getQueryWithMaterialSalesOrgDistChannelCustomer(req, material, salesOrganization, distributionChannel, customer)

	resp, err := new(http.Client).Do(req)
	if err != nil {
		return nil, xerrors.Errorf("API request error: %w", err)
	}
	defer resp.Body.Close()

	byteArray, _ := ioutil.ReadAll(resp.Body)
	data, err := sap_api_output_formatter.ConvertToPricingConditionValidity(byteArray, c.log)
	if err != nil {
		return nil, xerrors.Errorf("convert error: %w", err)
	}
	return data, nil
}

func (c *SAPAPICaller) setHeaderAPIKeyAccept(req *http.Request) {
	req.Header.Set("APIKey", c.apiKey)
	req.Header.Set("Accept", "application/json")
}

func (c *SAPAPICaller) getQueryWithMaterialDistChannel(req *http.Request, material, distributionChannel string) {
	params := req.URL.Query()
	params.Add("$filter", fmt.Sprintf("Material eq '%s' and DistributionChannel eq '%s'", material, distributionChannel))
	req.URL.RawQuery = params.Encode()
}

func (c *SAPAPICaller) getQueryWithMaterialDistChannelCustomer(req *http.Request, material, distributionChannel, customer string) {
	params := req.URL.Query()
	params.Add("$filter", fmt.Sprintf("Material eq '%s' and DistributionChannel eq '%s' and Customer eq '%s'", material, distributionChannel, customer))
	req.URL.RawQuery = params.Encode()
}

func (c *SAPAPICaller) getQueryWithMaterialSalesOrgDistChannel(req *http.Request, material, salesOrganization, distributionChannel string) {
	params := req.URL.Query()
	params.Add("$filter", fmt.Sprintf("Material eq '%s' and SalesOrganization eq '%s' and DistributionChannel eq '%s'", material, salesOrganization, distributionChannel))
	req.URL.RawQuery = params.Encode()
}

func (c *SAPAPICaller) getQueryWithMaterialSalesOrgDistChannelCustomer(req *http.Request, material, salesOrganization, distributionChannel, customer string) {
	params := req.URL.Query()
	params.Add("$filter", fmt.Sprintf("Material eq '%s' and SalesOrganization eq '%s' and DistributionChannel eq '%s' and Customer eq '%s'", material, salesOrganization, distributionChannel, customer))
	req.URL.RawQuery = params.Encode()
}

//func (c *SAPAPICaller) getQueryWithMaterial(req *http.Request, purchasingInfoRecord, purchasingInfoRecordCategory, supplier, material, purchasingOrganization, plant string) {
//	params := req.URL.Query()
//	params.Add("$filter", fmt.Sprintf("PurchasingInfoRecord ne '' and PurchasingInfoRecordCategory ne '' and Supplier eq '%s' and Material eq '%s' and PurchasingOrganization eq '%s' and Plant eq '%s'", supplier, material, purchasingOrganization, plant))
//	req.URL.RawQuery = params.Encode()
//}

//func (c *SAPAPICaller) getQueryWithMaterialGroup(req *http.Request, purchasingInfoRecord, purchasingInfoRecordCategory, supplier, materialGroup, purchasingOrganization, plant string) {
//	params := req.URL.Query()
//	params.Add("$filter", fmt.Sprintf("PurchasingInfoRecord ne '' and PurchasingInfoRecordCategory ne '' and Supplier eq '%s' and MaterialGroup eq '%s' and PurchasingOrganization eq '%s' and Plant eq '%s'", supplier, materialGroup, purchasingOrganization, plant))
//	req.URL.RawQuery = params.Encode()
//}

// func (c *SAPAPICaller) getQueryWithPricingConditionMaterial(req *http.Request, purchasingInfoRecord, PurchasingInfoRecordCategory, supplier, material, purchasingOrganization, plant, conditionType string) {
//	params := req.URL.Query()
//	params.Add("$filter", fmt.Sprintf("PurchasingInfoRecord ne '' and PurchasingInfoRecordCategory ne '' and Supplier eq '%s' and Material eq '%s' and PurchasingOrganization='%s' and Plant eq '%s' and conditionType eq '%s'", purchasingInfoRecord, PurchasingInfoRecordCategory, supplier, material, purchasingOrganization, plant, conditionType))
//	req.URL.RawQuery = params.Encode()
// }

// func (c *SAPAPICaller) getQueryWithPricingConditionMaterialGroup(req *http.Request, purchasingInfoRecord, PurchasingInfoRecordCategory, supplier, materialGroup, purchasingOrganization, plant, conditionType string) {
//	params := req.URL.Query()
//	params.Add("$filter", fmt.Sprintf("PurchasingInfoRecord ne '' and PurchasingInfoRecordCategory ne '' and Supplier eq '%s' and MaterialGroup eq '%s' and PurchasingOrganization='%s' and Plant eq '%s' and conditionType eq '%s'", purchasingInfoRecord, PurchasingInfoRecordCategory, supplier, material, purchasingOrganization, plant, conditionType))
//	req.URL.RawQuery = params.Encode()
// }
