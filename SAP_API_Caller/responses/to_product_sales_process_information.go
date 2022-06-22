package responses

type ToProductSalesProcessInformation struct {
	D struct {
		Results []struct {
			Metadata struct {
				URI  string `json:"uri"`
				Type string `json:"type"`
				Etag string `json:"etag"`
			} `json:"__metadata"`
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
		} `json:"results"`
	} `json:"d"`
}
