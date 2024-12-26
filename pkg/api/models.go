package api

type Office struct {
	ID   int    `json:"id,string"`
	Name string `json:"name"`
}

type Department struct {
	ID       int    `json:"id,string"`
	Name     string `json:"name"`
	OfficeID int    `json:"officeId,string"`
}

type Municipality struct {
	ID                 int    `json:"key1,string"`
	Name               string `json:"displayValue1"`
	RegistrationNumber int    `json:"key2,string"`
}

type ParcelIdentifier struct {
	ID     int    `json:"key1,string"`
	Number string `json:"value1"`
}

type ParcelDetail struct {
	ParcelID              int64           `json:"parcelId"`
	ParcelNumber          string          `json:"parcelNumber"`
	CADMunicipalityID     int64           `json:"cadMunicipalityId"`
	CADMunicipalityRegNum string          `json:"cadMunicipalityRegNum"`
	CADMunicipalityName   string          `json:"cadMunicipalityName"`
	InstitutionID         int64           `json:"institutionId"`
	Address               string          `json:"address"`
	Area                  string          `json:"area"`
	BuildingRemark        int64           `json:"buildingRemark"`
	DetailSheetNumber     string          `json:"detailSheetNumber"`
	HasBuildingRight      bool            `json:"hasBuildingRight"`
	LastChangeLog         string          `json:"lastChangeLog"`
	LastElaborateNumber   string          `json:"lastElaborateNumber"`
	LegalRegimeTypeNames  []string        `json:"legalRegimeTypeNames"`
	ParcelParts           []ParcelPart    `json:"parcelParts"`
	PossessionSheet       PossessionSheet `json:"possessionSheet"`
	ParcelLinks           []ParcelLink    `json:"parcelLinks"`
	IsAdditionalDataSet   bool            `json:"isAdditionalDataSet"`
	LrUnit                LrUnit          `json:"lrUnit"`
	LegalRegime           bool            `json:"legalRegime"`
	Graphic               bool            `json:"graphic"`
	AlphaNumeric          bool            `json:"alphaNumeric"`
	IsHarmonized          bool            `json:"isHarmonized"`
}

type ParcelLink struct {
	ParcelID     int64         `json:"parcelId"`
	ParcelNumber string        `json:"parcelNumber"`
	Address      string        `json:"address"`
	LrUnit       LrUnit        `json:"lrUnit"`
	ParcelParts  []interface{} `json:"parcelParts"`
}

type LrUnit struct {
	LrUnitID                 int64           `json:"lrUnitId"`
	LrUnitNumber             string          `json:"lrUnitNumber"`
	MainBookID               int64           `json:"mainBookId"`
	MainBookName             string          `json:"mainBookName"`
	InstitutionID            int64           `json:"institutionId"`
	Status                   string          `json:"status"`
	Verificated              bool            `json:"verificated"`
	Condominiums             bool            `json:"condominiums"`
	CadastreMunicipalityID   int64           `json:"cadastreMunicipalityId"`
	InstitutionName          string          `json:"institutionName"`
	StatusName               string          `json:"statusName"`
	LrUnitTypeID             int64           `json:"lrUnitTypeId"`
	LrUnitTypeName           string          `json:"lrUnitTypeName"`
	OwnershipSheetB          OwnershipSheetB `json:"ownershipSheetB"`
	OSSDataRestrictionStatus int64           `json:"ossDataRestrictionStatus"`
}

type ParcelPart struct {
	ParcelPartID          int64   `json:"parcelPartId"`
	Name                  string  `json:"name"`
	Area                  string  `json:"area"`
	PossessionSheetID     int64   `json:"possessionSheetId"`
	PossessionSheetNumber string  `json:"possessionSheetNumber"`
	LastChangeLog         string  `json:"lastChangeLog"`
	LastElaborateNumber   string  `json:"lastElaborateNumber"`
	Building              bool    `json:"building"`
	Type                  *string `json:"type,omitempty"`
	BuildingRight         *int64  `json:"buildingRight,omitempty"`
}

type PossessionSheet struct {
	PossessionSheetID     int64       `json:"possessionSheetId"`
	PossessionSheetNumber string      `json:"possessionSheetNumber"`
	CADMunicipalityID     int64       `json:"cadMunicipalityId"`
	CADMunicipalityRegNum string      `json:"cadMunicipalityRegNum"`
	CADMunicipalityName   string      `json:"cadMunicipalityName"`
	Possessors            []Possessor `json:"possessors"`
}

type Possessor struct {
	Name      string `json:"name"`
	Ownership string `json:"ownership"`
	Address   string `json:"address"`
}

type OwnershipSheetB struct {
	LrUnitShares []LrUnitShare `json:"lrUnitShares"`
	LrEntries    []interface{} `json:"lrEntries"`
}

type LrUnitShare struct {
	Description         string        `json:"description"`
	LrOwners            []LrOwner     `json:"lrOwners"`
	LrUnitShareID       int64         `json:"lrUnitShareId"`
	Share               string        `json:"share"`
	CondominiumNumber   string        `json:"condominiumNumber"`
	Condominiums        []string      `json:"condominiums"`
	SubSharesAndEntries []interface{} `json:"subSharesAndEntries"`
	Status              int64         `json:"status"`
	OrderNumber         string        `json:"orderNumber"`
}

type LrOwner struct {
	LrOwnerID int64  `json:"lrOwnerId"`
	Name      string `json:"name"`
	Address   string `json:"address"`
	TaxNumber string `json:"taxNumber"`
}
