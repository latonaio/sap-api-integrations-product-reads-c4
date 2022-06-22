package sap_api_output_formatter

type ProductCollection struct {
	ObjectID                         string `json:"ObjectID"`
	ProductID                        string `json:"ProductID"`
	UUID                             string `json:"UUID"`
	Language                         string `json:"Language"`
	LanguageText                     string `json:"LanguageText"`
	Description                      string `json:"Description"`
	ProductCategoryID                string `json:"ProductCategoryID"`
	Status                           string `json:"Status"`
	StatusText                       string `json:"StatusText"`
	Usage                            string `json:"Usage"`
	UsageText                        string `json:"UsageText"`
	Division                         string `json:"Division"`
	DivisionText                     string `json:"DivisionText"`
	BaseUOM                          string `json:"BaseUOM"`
	BaseUOMText                      string `json:"BaseUOMText"`
	CreatedBy                        string `json:"CreatedBy"`
	LastChangedBy                    string `json:"LastChangedBy"`
	CreatedOn                        string `json:"CreatedOn"`
	LastChangedOn                    string `json:"LastChangedOn"`
	CreatedByUUID                    string `json:"CreatedByUUID"`
	LastChangedByUUID                string `json:"LastChangedByUUID"`
	ExternalSystem                   string `json:"ExternalSystem"`
	ExternalID                       string `json:"ExternalID"`
	EntityLastChangedOn              string `json:"EntityLastChangedOn"`
	ETag                             string `json:"ETag"`
	ToProductOtherDescriptions       string `json:"ProductOtherDescriptions"`
	ToProductSalesProcessInformation string `json:"ProductSalesProcessInformation"`
}

type ProductOtherDescriptions struct {
	ObjectID       string `json:"ObjectID"`
	ParentObjectID string `json:"ParentObjectID"`
	ProductID      string `json:"ProductID"`
	Language       string `json:"Language"`
	LanguageText   string `json:"LanguageText"`
	Description    string `json:"Description"`
	ETag           string `json:"ETag"`
}

type ProductSalesProcessInformation struct {
	ObjectID                string `json:"ObjectID"`
	ParentObjectID          string `json:"ParentObjectID"`
	ProductID               string `json:"ProductID"`
	SalesOrganisationID     string `json:"SalesOrganisationID"`
	DistributionChannel     string `json:"DistributionChannel"`
	DistributionChannelText string `json:"DistributionChannelText"`
	Status                  string `json:"Status"`
	StatusText              string `json:"StatusText"`
	SalesUoM                string `json:"SalesUoM"`
	SalesUoMText            string `json:"SalesUoMText"`
	MinimumOrderQuantity    string `json:"MinimumOrderQuantity"`
	PricingProductID        string `json:"PricingProductID"`
	ETag                    string `json:"ETag"`
}
