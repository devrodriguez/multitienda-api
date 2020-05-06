package models

type GoogleApiAuto struct {
	Predictions []struct {
		Description       string `json:"description,omitempty"`
		ID                string `json:"id,omitempty"`
		MatchedSubstrings []struct {
			Length int `json:"length"`
			Offset int `json:"offset"`
		} `json:"matched_substrings,omitempty"`
		PlaceID              string `json:"place_id,omitempty"`
		Reference            string `json:"reference,omitempty"`
		StructuredFormatting struct {
			MainText                  string `json:"main_text"`
			MainTextMatchedSubstrings []struct {
				Length int `json:"length"`
				Offset int `json:"offset"`
			} `json:"main_text_matched_substrings"`
			SecondaryText string `json:"secondary_text"`
		} `json:"structured_formatting,omitempty"`
		Terms []struct {
			Offset int    `json:"offset"`
			Value  string `json:"value"`
		} `json:"terms,omitempty"`
		Types []string `json:"types,omitempty"`
	} `json:"predictions"`
	Status string `json:"status"`
}
