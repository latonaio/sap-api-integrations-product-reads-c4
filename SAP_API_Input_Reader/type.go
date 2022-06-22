package sap_api_input_reader

type EC_MC struct {
	ConnectionKey string `json:"connection_key"`
	Result        bool   `json:"result"`
	RedisKey      string `json:"redis_key"`
	Filepath      string `json:"filepath"`
	Document      struct {
		DocumentNo     string `json:"document_no"`
		DeliverTo      string `json:"deliver_to"`
		Quantity       string `json:"quantity"`
		PickedQuantity string `json:"picked_quantity"`
		Price          string `json:"price"`
		Batch          string `json:"batch"`
	} `json:"document"`
	ProductionOrder struct {
		DocumentNo           string `json:"document_no"`
		Status               string `json:"status"`
		DeliverTo            string `json:"deliver_to"`
		Quantity             string `json:"quantity"`
		CompletedQuantity    string `json:"completed_quantity"`
		PlannedStartDate     string `json:"planned_start_date"`
		PlannedValidatedDate string `json:"planned_validated_date"`
		ActualStartDate      string `json:"actual_start_date"`
		ActualValidatedDate  string `json:"actual_validated_date"`
		Batch                string `json:"batch"`
		Work                 struct {
			WorkNo                   string `json:"work_no"`
			Quantity                 string `json:"quantity"`
			CompletedQuantity        string `json:"completed_quantity"`
			ErroredQuantity          string `json:"errored_quantity"`
			Component                string `json:"component"`
			PlannedComponentQuantity string `json:"planned_component_quantity"`
			PlannedStartDate         string `json:"planned_start_date"`
			PlannedStartTime         string `json:"planned_start_time"`
			PlannedValidatedDate     string `json:"planned_validated_date"`
			PlannedValidatedTime     string `json:"planned_validated_time"`
			ActualStartDate          string `json:"actual_start_date"`
			ActualStartTime          string `json:"actual_start_time"`
			ActualValidatedDate      string `json:"actual_validated_date"`
			ActualValidatedTime      string `json:"actual_validated_time"`
		} `json:"work"`
	} `json:"production_order"`
	APISchema      string `json:"api_schema"`
	MaterialCode   string `json:"material_code"`
	Plant_Supplier string `json:"plant/supplier"`
	Stock          string `json:"stock"`
	DocumentType   string `json:"document_type"`
	DocumentNo     string `json:"document_no"`
	PlannedDate    string `json:"planned_date"`
	ValidatedDate  string `json:"validated_date"`
	Deleted        bool   `json:"deleted"`
}

type SDC struct {
	ConnectionKey      string `json:"connection_key"`
	Result             bool   `json:"result"`
	RedisKey           string `json:"redis_key"`
	Filepath           string `json:"filepath"`
	ProductCollection struct {
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
			ObjectID       string `json:"ObjectID"`
			ParentObjectID string `json:"ParentObjectID"`
			ProductID      string `json:"ProductID"`
			Language       string `json:"Language"`
			LanguageText   string `json:"LanguageText"`
			Description    string `json:"Description"`
			ETag           string `json:"ETag"`
			ProductSalesProcessInformation struct {
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
			} `json:"ProductSalesProcessInformation"`
		} `json:"ProductOtherDescriptions"`
	} `json:"ProductCollection"`
	APISchema    string   `json:"api_schema"`
	Accepter     []string `json:"accepter"`
	ProductCode string   `json:"product_code"`
	Deleted      bool     `json:"deleted"`
}