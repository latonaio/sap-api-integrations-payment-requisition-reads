package sap_api_caller

import (
	"fmt"
	"io/ioutil"
	"net/http"
	sap_api_output_formatter "sap-api-integrations-payment-requisition-reads/SAP_API_Output_Formatter"
	"strings"
	"sync"

	"github.com/latonaio/golang-logging-library-for-sap/logger"
	"golang.org/x/xerrors"
)

type SAPAPICaller struct {
	baseURL string
	apiKey  string
	log     *logger.Logger
}

func NewSAPAPICaller(baseUrl string, l *logger.Logger) *SAPAPICaller {
	return &SAPAPICaller{
		baseURL: baseUrl,
		apiKey:  GetApiKey(),
		log:     l,
	}
}

func (c *SAPAPICaller) AsyncGetPaymentRequisition(paymentRequisitionUUID, paymentRequisitionStrategyUUID, paymentRequisitionItemUUID string, accepter []string) {
	wg := &sync.WaitGroup{}
	wg.Add(len(accepter))
	for _, fn := range accepter {
		switch fn {
		case "Strategies":
			func() {
				c.Strategies(paymentRequisitionUUID, paymentRequisitionStrategyUUID)
				wg.Done()
			}()
		case "Item":
			func() {
				c.Item(paymentRequisitionUUID, paymentRequisitionItemUUID)
				wg.Done()
			}()
		case "Requisitions":
			func() {
				c.Requisitions(paymentRequisitionUUID)
				wg.Done()
			}()
		default:
			wg.Done()
		}
	}
	wg.Wait()
}

func (c *SAPAPICaller) Strategies(paymentRequisitionUUID, paymentRequisitionStrategyUUID string) {
	data, err := c.callPaymentRequisitionSrvAPIRequirementStrategies("Strategies", paymentRequisitionUUID, paymentRequisitionStrategyUUID)
	if err != nil {
		c.log.Error(err)
		return
	}
	c.log.Info(data)

}

func (c *SAPAPICaller) callPaymentRequisitionSrvAPIRequirementStrategies(api, paymentRequisitionUUID, paymentRequisitionStrategyUUID string) ([]sap_api_output_formatter.Strategies, error) {
	url := strings.Join([]string{c.baseURL, "c4codataapi", api}, "/")
	req, _ := http.NewRequest("GET", url, nil)

	c.setHeaderAPIKeyAccept(req)
	c.getQueryWithStrategies(req, paymentRequisitionUUID, paymentRequisitionStrategyUUID)

	resp, err := new(http.Client).Do(req)
	if err != nil {
		return nil, xerrors.Errorf("API request error: %w", err)
	}
	defer resp.Body.Close()

	byteArray, _ := ioutil.ReadAll(resp.Body)
	data, err := sap_api_output_formatter.ConvertToStrategies(byteArray, c.log)
	if err != nil {
		return nil, xerrors.Errorf("convert error: %w", err)
	}
	return data, nil
}

func (c *SAPAPICaller) Item(paymentRequisitionUUID, paymentRequisitionItemUUID string) {
	data, err := c.callPaymentRequisitionSrvAPIRequirementItem("Item", paymentRequisitionUUID, paymentRequisitionItemUUID)
	if err != nil {
		c.log.Error(err)
		return
	}
	c.log.Info(data)

}

func (c *SAPAPICaller) callPaymentRequisitionSrvAPIRequirementItem(api, paymentRequisitionUUID, paymentRequisitionItemUUID string) ([]sap_api_output_formatter.Item, error) {
	url := strings.Join([]string{c.baseURL, "c4codataapi", api}, "/")
	req, _ := http.NewRequest("GET", url, nil)

	c.setHeaderAPIKeyAccept(req)
	c.getQueryWithItem(req, paymentRequisitionUUID, paymentRequisitionItemUUID)

	resp, err := new(http.Client).Do(req)
	if err != nil {
		return nil, xerrors.Errorf("API request error: %w", err)
	}
	defer resp.Body.Close()

	byteArray, _ := ioutil.ReadAll(resp.Body)
	data, err := sap_api_output_formatter.ConvertToItem(byteArray, c.log)
	if err != nil {
		return nil, xerrors.Errorf("convert error: %w", err)
	}
	return data, nil
}

func (c *SAPAPICaller) Requisitions(paymentRequisitionUUID string) {
	data, err := c.callPaymentRequisitionSrvAPIRequirementRequisitions("Requisitions", paymentRequisitionUUID)
	if err != nil {
		c.log.Error(err)
		return
	}
	c.log.Info(data)

}

func (c *SAPAPICaller) callPaymentRequisitionSrvAPIRequirementRequisitions(api, paymentRequisitionUUID string) ([]sap_api_output_formatter.Requisitions, error) {
	url := strings.Join([]string{c.baseURL, "c4codataapi", api}, "/")
	req, _ := http.NewRequest("GET", url, nil)

	c.setHeaderAPIKeyAccept(req)
	c.getQueryWithRequisitions(req, paymentRequisitionUUID)

	resp, err := new(http.Client).Do(req)
	if err != nil {
		return nil, xerrors.Errorf("API request error: %w", err)
	}
	defer resp.Body.Close()

	byteArray, _ := ioutil.ReadAll(resp.Body)
	data, err := sap_api_output_formatter.ConvertToRequisitions(byteArray, c.log)
	if err != nil {
		return nil, xerrors.Errorf("convert error: %w", err)
	}
	return data, nil
}

func (c *SAPAPICaller) setHeaderAPIKeyAccept(req *http.Request) {
	req.Header.Set("APIKey", c.apiKey)
	req.Header.Set("Accept", "application/json")
}

func (c *SAPAPICaller) getQueryWithStrategies(req *http.Request, paymentRequisitionUUID, paymentRequisitionStrategyUUID string) {
	params := req.URL.Query()
	params.Add("$filter", fmt.Sprintf("PaymentRequisitionUUID eq '%s' and paymentRequisitionStrategyUUID eq '%s'", paymentRequisitionUUID, paymentRequisitionStrategyUUID))
	req.URL.RawQuery = params.Encode()
}

func (c *SAPAPICaller) getQueryWithItem(req *http.Request, paymentRequisitionUUID, paymentRequisitionItemUUID string) {
	params := req.URL.Query()
	params.Add("$filter", fmt.Sprintf("paymentRequisitionUUID eq '%s' and paymentRequisitionItemUUID eq '%s'", paymentRequisitionUUID, paymentRequisitionItemUUID))
	req.URL.RawQuery = params.Encode()
}

func (c *SAPAPICaller) getQueryWithRequisitions(req *http.Request, paymentRequisitionUUID string) {
	params := req.URL.Query()
	params.Add("$filter", fmt.Sprintf("paymentRequisitionUUID eq '%s'", paymentRequisitionUUID))
	req.URL.RawQuery = params.Encode()
}
