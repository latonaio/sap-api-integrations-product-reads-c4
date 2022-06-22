package sap_api_output_formatter

import (
	"encoding/json"
	"sap-api-integrations-product-reads-c4/SAP_API_Caller/responses"

	"github.com/latonaio/golang-logging-library-for-sap/logger"
	"golang.org/x/xerrors"
)

func ConvertToProductCollection(raw []byte, l *logger.Logger) ([]ProductCollection, error) {
	pm := &responses.ProductCollection{}

	err := json.Unmarshal(raw, pm)
	if err != nil {
		return nil, xerrors.Errorf("cannot convert to ProductCollection. unmarshal error: %w", err)
	}
	if len(pm.D.Results) == 0 {
		return nil, xerrors.New("Result data is not exist")
	}
	if len(pm.D.Results) > 10 {
		l.Info("raw data has too many Results. %d Results exist. show the first 10 of Results array", len(pm.D.Results))
	}

	productCollection := make([]ProductCollection, 0, 10)
	for i := 0; i < 10 && i < len(pm.D.Results); i++ {
		data := pm.D.Results[i]
		productCollection = append(productCollection, ProductCollection{
			ObjectID:                         data.ObjectID,
			ProductID:                        data.ProductID,
			UUID:                             data.UUID,
			Language:                         data.Language,
			LanguageText:                     data.LanguageText,
			Description:                      data.Description,
			ProductCategoryID:                data.ProductCategoryID,
			Status:                           data.Status,
			StatusText:                       data.StatusText,
			Usage:                            data.Usage,
			UsageText:                        data.UsageText,
			Division:                         data.Division,
			DivisionText:                     data.DivisionText,
			BaseUOM:                          data.BaseUOM,
			BaseUOMText:                      data.BaseUOMText,
			CreatedBy:                        data.CreatedBy,
			LastChangedBy:                    data.LastChangedBy,
			CreatedOn:                        data.CreatedOn,
			LastChangedOn:                    data.LastChangedOn,
			CreatedByUUID:                    data.CreatedByUUID,
			LastChangedByUUID:                data.LastChangedByUUID,
			ExternalSystem:                   data.ExternalSystem,
			ExternalID:                       data.ExternalID,
			EntityLastChangedOn:              data.EntityLastChangedOn,
			ETag:                             data.ETag,
			ToProductOtherDescriptions:       data.ProductOtherDescriptions.Deferred.URI,
			ToProductSalesProcessInformation: data.ProductSalesProcessInformation.Deferred.URI,
		})
	}

	return productCollection, nil
}

func ConvertToProductOtherDescriptions(raw []byte, l *logger.Logger) ([]ProductOtherDescriptions, error) {
	pm := &responses.ToProductOtherDescriptions{}

	err := json.Unmarshal(raw, pm)
	if err != nil {
		return nil, xerrors.Errorf("cannot convert to ProductOtherDescriptions. unmarshal error: %w", err)
	}
	if len(pm.D.Results) == 0 {
		return nil, xerrors.New("Result data is not exist")
	}
	if len(pm.D.Results) > 10 {
		l.Info("raw data has too many Results. %d Results exist. show the first 10 of Results array", len(pm.D.Results))
	}

	productOtherDescriptions := make([]ProductOtherDescriptions, 0, 10)
	for i := 0; i < 10 && i < len(pm.D.Results); i++ {
		data := pm.D.Results[i]
		productOtherDescriptions = append(productOtherDescriptions, ProductOtherDescriptions{
			ObjectID:       data.ObjectID,
			ParentObjectID: data.ParentObjectID,
			ProductID:      data.ProductID,
			Language:       data.Language,
			LanguageText:   data.LanguageText,
			Description:    data.Description,
			ETag:           data.ETag,
		})
	}

	return productOtherDescriptions, nil
}

func ConvertToProductSalesProcessInformation(raw []byte, l *logger.Logger) ([]ProductSalesProcessInformation, error) {
	pm := &responses.ToProductSalesProcessInformation{}

	err := json.Unmarshal(raw, pm)
	if err != nil {
		return nil, xerrors.Errorf("cannot convert to ProductSalesProcessInformation. unmarshal error: %w", err)
	}
	if len(pm.D.Results) == 0 {
		return nil, xerrors.New("Result data is not exist")
	}
	if len(pm.D.Results) > 10 {
		l.Info("raw data has too many Results. %d Results exist. show the first 10 of Results array", len(pm.D.Results))
	}

	productSalesProcessInformation := make([]ProductSalesProcessInformation, 0, 10)
	for i := 0; i < 10 && i < len(pm.D.Results); i++ {
		data := pm.D.Results[i]
		productSalesProcessInformation = append(productSalesProcessInformation, ProductSalesProcessInformation{
			ObjectID:                data.ObjectID,
			ParentObjectID:          data.ParentObjectID,
			ProductID:               data.ProductID,
			SalesOrganisationID:     data.SalesOrganisationID,
			DistributionChannel:     data.DistributionChannel,
			DistributionChannelText: data.DistributionChannelText,
			Status:                  data.Status,
			StatusText:              data.StatusText,
			SalesUoM:                data.SalesUoM,
			SalesUoMText:            data.SalesUoMText,
			MinimumOrderQuantity:    data.MinimumOrderQuantity,
			PricingProductID:        data.PricingProductID,
			ETag:                    data.ETag,
		})
	}

	return productSalesProcessInformation, nil

}
