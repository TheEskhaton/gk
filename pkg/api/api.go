package api

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

func FetchOffices() []Office {
	res, err := http.Get("https://oss.uredjenazemlja.hr/oss/public/search-cad-parcels/offices")
	if err != nil {
		log.Fatal(err)
	}
	body, err := io.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
	}
	var offices []Office
	err = json.Unmarshal(body, &offices)
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
	return offices
}

func FetchDepartments() []Department {
	res, err := http.Get("https://oss.uredjenazemlja.hr/oss/public/search-cad-parcels/departments?doFilter=true")
	if err != nil {
		log.Fatal(err)
	}
	body, err := io.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
	}
	var departments []Department
	err = json.Unmarshal(body, &departments)
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
	return departments
}

func FetchMunicipalities(officeId int, departmentId int) []Municipality {
	uri := fmt.Sprintf("https://oss.uredjenazemlja.hr/oss/public/search-cad-parcels/municipalities?search=&officeId=%d&departmentId=%d", officeId, departmentId)
	if departmentId == 0 {
		uri = fmt.Sprintf("https://oss.uredjenazemlja.hr/oss/public/search-cad-parcels/municipalities?search=&officeId=%d", officeId)
	}
	res, err := http.Get(uri)
	if err != nil {
		log.Fatal(err)
	}
	body, err := io.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
	}
	var municipalities []Municipality
	err = json.Unmarshal(body, &municipalities)
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
	return municipalities
}

func SearchParcelIdentifiers(parcelNumber string, municipalityRegNumber int) []ParcelIdentifier {
	res, err := http.Get(fmt.Sprintf("https://oss.uredjenazemlja.hr/oss/public/search-cad-parcels/parcel-numbers?search=%s&municipalityRegNum=%d", parcelNumber, municipalityRegNumber))
	if err != nil {
		log.Fatal(err)
	}
	body, err := io.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
	}
	var parcelIdentifiers []ParcelIdentifier
	err = json.Unmarshal(body, &parcelIdentifiers)
	if err != nil {
		log.Fatal(err)
	}
	return parcelIdentifiers
}

func FetchParcelDetails(parcelId int) []ParcelDetail {
	res, err := http.Post("https://oss.uredjenazemlja.hr/oss/public/cad/search-parcels", "application/json", bytes.NewBufferString(fmt.Sprintf("{\"parcelId\":\"%d\"}", parcelId)))
	if err != nil {
		log.Fatal(err)
	}
	body, err := io.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
	}
	var parcelDetails []ParcelDetail
	err = json.Unmarshal(body, &parcelDetails)
	if err != nil {
		log.Fatal(err)
	}
	return parcelDetails
}

type ExtractType string

const (
	LR            ExtractType = "LR"
	PropertyDeed  ExtractType = "PropertyDeed"
	CadastralPlan ExtractType = "CadastralPlan"
)

func (e *ExtractType) String() string {
	return string(*e)
}

func (e *ExtractType) Set(v string) error {
	switch v {
	case "LR", "PropertyDeed", "CadastralPlan":
		*e = ExtractType(v)
		return nil
	default:
		return errors.New(`must be one of "LR", "PropertyDeed", "CadastralPlan"`)
	}
}

func (e *ExtractType) Type() string {
	return "ExtractType"
}

func FetchPossesionSheetExtract(parcelId int, possessionSheetId int, extractType ExtractType) []PossesionSheetExtract {
	uri := "https://oss.uredjenazemlja.hr/oss/public/reports/get-possessionsheet-extract?parcelId=%d"
	id := parcelId
	switch extractType {
	case PropertyDeed:
		uri = "https://oss.uredjenazemlja.hr/oss/public/reports/get-possessionsheet-extract?possessionSheetId=%d"
		id = possessionSheetId
	case CadastralPlan:
		uri = "https://oss.uredjenazemlja.hr/oss/public/reports/get-kp-extract/%d"
		id = parcelId
	}
	res, err := http.Get(fmt.Sprintf(uri, id))
	if err != nil {
		log.Fatal(err)
	}
	body, err := io.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
	}
	if extractType == CadastralPlan {

		var possesionSheetExtract PossesionSheetExtract
		err = json.Unmarshal(body, &possesionSheetExtract)
		if err != nil {
			log.Fatal(err)
		}
		return []PossesionSheetExtract{possesionSheetExtract}
	} else {

		var possesionSheetExtracts []PossesionSheetExtract
		err = json.Unmarshal(body, &possesionSheetExtracts)
		if err != nil {
			log.Fatal(err)
		}
		return possesionSheetExtracts
	}
}
