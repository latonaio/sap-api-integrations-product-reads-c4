package sap_api_caller

import (
	"fmt"
	"io/ioutil"
	"net/http"
	sap_api_output_formatter "sap-api-integrations-product-reads-c4/SAP_API_Output_Formatter"
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

func (c *SAPAPICaller) AsyncGetProduct(productID string, accepter []string) {
	wg := &sync.WaitGroup{}
	wg.Add(len(accepter))
	for _, fn := range accepter {
		switch fn {
		case "ProductCollection":
			func() {
				c.ProductCollection(productID)
				wg.Done()
			}()
		default:
			wg.Done()
		}
	}

	wg.Wait()
}

func (c *SAPAPICaller) ProductCollection(productID string) {
	productCollectionData, err := c.callProductSrvAPIRequirementProductCollection("ProductCollection", productID)
	if err != nil {
		c.log.Error(err)
		return
	}
	c.log.Info(productCollectionData)

	productOtherDescriptionsData, err := c.callProductOtherDescriptions(productCollectionData[0].ToProductOtherDescriptions)
	if err != nil {
		c.log.Error(err)
		return
	}
	c.log.Info(productOtherDescriptionsData)

	productSalesProcessInformationData, err := c.callProductSalesProcessInformation(productCollectionData[0].ToProductSalesProcessInformation)
	if err != nil {
		c.log.Error(err)
		return
	}
	c.log.Info(productSalesProcessInformationData)

}

func (c *SAPAPICaller) callProductSrvAPIRequirementProductCollection(api, productID string) ([]sap_api_output_formatter.ProductCollection, error) {
	url := strings.Join([]string{c.baseURL, "c4codataapi", api}, "/")
	req, _ := http.NewRequest("GET", url, nil)

	c.setHeaderAPIKeyAccept(req)
	c.getQueryWithProductCollection(req, productID)

	resp, err := new(http.Client).Do(req)
	if err != nil {
		return nil, xerrors.Errorf("API request error: %w", err)
	}
	defer resp.Body.Close()

	byteArray, _ := ioutil.ReadAll(resp.Body)
	data, err := sap_api_output_formatter.ConvertToProductCollection(byteArray, c.log)
	if err != nil {
		return nil, xerrors.Errorf("convert error: %w", err)
	}
	return data, nil
}

func (c *SAPAPICaller) callProductOtherDescriptions(url string) ([]sap_api_output_formatter.ProductOtherDescriptions, error) {
	req, _ := http.NewRequest("GET", url, nil)
	c.setHeaderAPIKeyAccept(req)

	resp, err := new(http.Client).Do(req)
	if err != nil {
		return nil, xerrors.Errorf("API request error: %w", err)
	}
	defer resp.Body.Close()

	byteArray, _ := ioutil.ReadAll(resp.Body)
	data, err := sap_api_output_formatter.ConvertToProductOtherDescriptions(byteArray, c.log)
	if err != nil {
		return nil, xerrors.Errorf("convert error: %w", err)
	}
	return data, nil
}

func (c *SAPAPICaller) callProductSalesProcessInformation(url string) ([]sap_api_output_formatter.ProductSalesProcessInformation, error) {
	req, _ := http.NewRequest("GET", url, nil)
	c.setHeaderAPIKeyAccept(req)

	resp, err := new(http.Client).Do(req)
	if err != nil {
		return nil, xerrors.Errorf("API request error: %w", err)
	}
	defer resp.Body.Close()

	byteArray, _ := ioutil.ReadAll(resp.Body)
	data, err := sap_api_output_formatter.ConvertToProductSalesProcessInformation(byteArray, c.log)
	if err != nil {
		return nil, xerrors.Errorf("convert error: %w", err)
	}
	return data, nil
}

func (c *SAPAPICaller) setHeaderAPIKeyAccept(req *http.Request) {
	req.Header.Set("APIKey", c.apiKey)
	req.Header.Set("Accept", "application/json")
}

func (c *SAPAPICaller) getQueryWithProductCollection(req *http.Request, productID string) {
	params := req.URL.Query()
	params.Add("$filter", fmt.Sprintf("ProductID eq '%s'", productID))
	req.URL.RawQuery = params.Encode()
}
