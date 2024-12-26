package api

import (
	"bytes"
	"encoding/json"
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
	fmt.Println(string(body))
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
	res, err := http.Get(fmt.Sprintf("https://oss.uredjenazemlja.hr/oss/public/search-cad-parcels/municipalities?search=&officeId=%d&departmentId=%d", officeId, departmentId))
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
