package responses

type ToProductOtherDescriptions struct {
	D struct {
		Results []struct {
			Metadata struct {
				URI  string `json:"uri"`
				Type string `json:"type"`
				Etag string `json:"etag"`
			} `json:"__metadata"`
			ObjectID       string `json:"ObjectID"`
			ParentObjectID string `json:"ParentObjectID"`
			ProductID      string `json:"ProductID"`
			Language       string `json:"Language"`
			LanguageText   string `json:"LanguageText"`
			Description    string `json:"Description"`
			ETag           string `json:"ETag"`
		} `json:"results"`
	} `json:"d"`
}
