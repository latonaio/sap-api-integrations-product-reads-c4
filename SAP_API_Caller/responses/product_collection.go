package responses

type ProductCollection struct {
	D struct {
		Results []struct {
			Metadata struct {
				URI  string `json:"uri"`
				Type string `json:"type"`
				Etag string `json:"etag"`
			} `json:"__metadata"`
			ObjectID                 string `json:"ObjectID"`
			ProductID                string `json:"ProductID"`
			UUID                     string `json:"UUID"`
			Language                 string `json:"Language"`
			LanguageText             string `json:"LanguageText"`
			Description              string `json:"Description"`
			ProductCategoryID        string `json:"ProductCategoryID"`
			Status                   string `json:"Status"`
			StatusText               string `json:"StatusText"`
			Usage                    string `json:"Usage"`
			UsageText                string `json:"UsageText"`
			Division                 string `json:"Division"`
			DivisionText             string `json:"DivisionText"`
			BaseUOM                  string `json:"BaseUOM"`
			BaseUOMText              string `json:"BaseUOMText"`
			CreatedBy                string `json:"CreatedBy"`
			LastChangedBy            string `json:"LastChangedBy"`
			CreatedOn                string `json:"CreatedOn"`
			LastChangedOn            string `json:"LastChangedOn"`
			CreatedByUUID            string `json:"CreatedByUUID"`
			LastChangedByUUID        string `json:"LastChangedByUUID"`
			ExternalSystem           string `json:"ExternalSystem"`
			ExternalID               string `json:"ExternalID"`
			EntityLastChangedOn      string `json:"EntityLastChangedOn"`
			ETag                     string `json:"ETag"`
			ProductOtherDescriptions struct {
				Deferred struct {
					URI string `json:"uri"`
				} `json:"__deferred"`
			} `json:"ProductOtherDescriptions"`
			ProductSalesProcessInformation struct {
				Deferred struct {
					URI string `json:"uri"`
				} `json:"__deferred"`
			} `json:"ProductSalesProcessInformation"`
		} `json:"results"`
	} `json:"d"`
}
