package dictionary

type MirramWebsterDefinitionContract struct {
	Date string `xml:"date"`
	DefinitionNumbers []string `xml:"sn"`
	Definitions []string `xml:"dt"`
}

type MirramWebsterEntryContract struct {
	Id string `xml:"id,attr"`
	Entryword string `xml:"ew"`
	Headword string `xml:"hw"`
	Pronunciation string `xml:"pr"`
	PartOfSpeech string `xml:"fl"`
}