package dictionary

type Senses struct {
	VariantForms []VariantForms `json:"variantForms"`
	CrossReferenceMarkers []string `json:"crossReferenceMarkers"`
	Domains []string `json:"domains"`
	Examples []Examples `json:"examples"`
	Id string `json:"id"`
	Regions []string `json:"regions"`
	Subsenses []Subsenses `json:"subsenses"`
	Translations []Translations `json:"translations"`
	CrossReferences []CrossReferences `json:"crossReferences"`
	Definitions []string `json:"definitions"`
	Notes []Notes `json:"notes"`
	Pronunciations []Pronunciations `json:"pronunciations"`
	Registers []string `json:"registers"`
}

type Examples struct {
	Notes []Notes `json:"notes"`
	Regions []string `json:"regions"`
	Registers []string `json:"registers"`
	SenseIds []string `json:"senseIds"`
	Text string `json:"text"`
	Translations []Translations `json:"translations"`
	Definitions []string `json:"definitions"`
	Domains []string `json:"domains"`
}

type Translations struct {
	Domains []string `json:"domains"`
	GrammaticalFeatures []GrammaticalFeatures `json:"grammaticalFeatures"`
	Language string `json:"language"`
	Notes []Notes `json:"notes"`
	Regions []string `json:"regions"`
	Registers []string `json:"registers"`
	Text string `json:"text"`
}

type Subsenses struct {
}

type VariantForms struct {
	Regions []string `json:"regions"`
	Text string `json:"text"`
}

type CrossReferences struct {
	Text string `json:"text"`
	Type string `json:"type"`
	Id string `json:"id"`
}

type OxfordDefinitionsApi struct {
	Results []Results `json:"results"`
	Metadata Metadata `json:"metadata"`
}

type Results struct {
	LexicalEntries []LexicalEntries `json:"lexicalEntries"`
	Pronunciations []Pronunciations `json:"pronunciations"`
	Type string `json:"type"`
	Word string `json:"word"`
	Id string `json:"id"`
	Language string `json:"language"`
}

type LexicalEntries struct {
	Entries []Entries `json:"entries"`
	Language string `json:"language"`
	Text string `json:"text"`
	VariantForms []VariantForms `json:"variantForms"`
	DerivativeOf []DerivativeOf `json:"derivativeOf"`
	GrammaticalFeatures []GrammaticalFeatures `json:"grammaticalFeatures"`
	LexicalCategory string `json:"lexicalCategory"`
	Notes []Notes `json:"notes"`
	Pronunciations []Pronunciations `json:"pronunciations"`
}

type Entries struct {
	GrammaticalFeatures []GrammaticalFeatures `json:"grammaticalFeatures"`
	HomographNumber string `json:"homographNumber"`
	Notes []Notes `json:"notes"`
	Pronunciations []Pronunciations `json:"pronunciations"`
	Senses []Senses `json:"senses"`
	VariantForms []VariantForms `json:"variantForms"`
	Etymologies []string `json:"etymologies"`
}

type GrammaticalFeatures struct {
	Text string `json:"text"`
	Type string `json:"type"`
}

type Notes struct {
	Id string `json:"id"`
	Text string `json:"text"`
	Type string `json:"type"`
}

type Pronunciations struct {
	Dialects []string `json:"dialects"`
	PhoneticNotation string `json:"phoneticNotation"`
	PhoneticSpelling string `json:"phoneticSpelling"`
	Regions []string `json:"regions"`
	AudioFile string `json:"audioFile"`
}

type DerivativeOf struct {
	Id string `json:"id"`
	Language string `json:"language"`
	Regions []string `json:"regions"`
	Registers []string `json:"registers"`
	Text string `json:"text"`
	Domains []string `json:"domains"`
}

type Metadata struct {
}

