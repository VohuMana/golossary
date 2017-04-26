package dictionary

type PearsonSensesContract struct {
	Definition []string `json:"definition"`
}

type PearsonAudioContract struct {
	Lang string `json:"lang"`
	Type string `json:"type"`
	URL string `json:"url"`
}

type PearsonPronunciationsContract struct {
	Audio []PearsonAudioContract `json:"audio"`
}

type PearsonDefinitionContract struct {
	Datasets []string `json:"datasets"`
	Headword string `json:"headword"`
	Id string `json:"id"`
	PartOfSpeech string `json:"part_of_speech"`
	Pronunciations []PearsonPronunciationsContract `json:"pronunciations"`
	Senses []PearsonSensesContract `json:"senses"`
	URL string `json:"url"`
}

type PearsonDefineWordContract struct {
	Status uint `json:"status"`
	Offset uint `json:"offset"`
	Limit uint `json:"limit"`
	Count uint `json:"count"`
	Total uint `json:"total"`
	URL string `json:"url"`
	Results []PearsonDefinitionContract `json:"results"`
}