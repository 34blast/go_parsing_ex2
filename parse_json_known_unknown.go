package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	//	"strings"
)

type Product struct {
	DocType      string                 `json:"docType"`
	ProductId    string                 `json:"productId"`
	Gtin         string                 `json:"gtin"`
	SerialNumber string                 `json:"serialNumber"`
	LotNumber    string                 `json:"lotNumber"`
	ExpiryDate   string                 `json:"expiryDate"`
	FingerPrint  string                 `json:"fingerPrint"`
	Data         map[string]interface{} `json:"-"` // Unknown fields should go here.
}

// use one long string for comparisions below
const JSON_PRODUCT1 = `{"docType":"product","productId":"productId1","gtin":"gtin1","serialNumber":"serialNumber1","lotNumber":"lotNumber1","expiryDate":"2019-01-07T17:31:00.106Z", "fingerPrint":"fingerPrint1", "drugName":"drugName","mfgName":"mfgName","mfgDateTime":"2019-01-07T17:31:00.106Z","gtinSource":"glnSource", "extraAttribute1":"extraAttribute1","extraAttribute2":"extraAttribute2"}`
const JSON_PRODUCT_NO_EXTRAS = `{"docType":"product","productId":"productId1","gtin":"gtin1","serialNumber":"serialNumber1","lotNumber":"lotNumber1","expiryDate":"2019-01-07T17:31:00.106Z","fingerPrint":"fingerPrint1"}`
const JSON_PRODUCT_NO_FINGERPRINT = `{"docType":"product","productId":"productId1","gtin":"gtin1","serialNumber":"serialNumber1","lotNumber":"lotNumber1","expiryDate":"2019-01-07T17:31:00.106Z"}`

/*
const JSON_PRODUCT1 = `{ 
	"docType": "product", 
	"productId": "productId1", 
	"gtin": "gtin1", 
	"serialNumber": "serialNumber1", 
	"lotNumber": "lotNumber1",  
	"expiryDate": "2019-01-07T17:31:00.106Z", 
	"fingerPrint": "fingerPrint1", 
	"drugName": "drugName", "mfgName": 
	"mfgName", "mfgDateTime": "2019-01-07T17:31:00.106Z", 
	"gtinSource": "glnSource", 
	"extraAttribute1": "extraAttribute1",
	"extraAttribute2": "extraAttribute2"
	}`
*/
// This shows how to parse JSON with known attributes and unknown attributes.
// The unknown go into a MAP.  I parse everythign into the map,  then manually set the known fields from the map
// then remove them from the map
func main() {

	var product Product
	if err2 := json.Unmarshal([]byte(JSON_PRODUCT1), &product.Data); err2 != nil {
		log.Fatal(err2)
		os.Exit(-1)
	}

	product.DocType = product.Data["docType"].(string)
	product.ProductId = product.Data["productId"].(string)
	product.Gtin = product.Data["gtin"].(string)
	product.SerialNumber = product.Data["serialNumber"].(string)
	product.LotNumber = product.Data["lotNumber"].(string)
	product.ExpiryDate = product.Data["expiryDate"].(string)
	if val, ok := product.Data["fingerPrint"]; ok {
		product.FingerPrint = val.(string)
		delete(product.Data, "fingerPrint")
	} 

	delete(product.Data, "docType")
	delete(product.Data, "productId")
	delete(product.Data, "gtin")
	delete(product.Data, "serialNumber")
	delete(product.Data, "lotNumber")
	delete(product.Data, "expiryDate")

	fmt.Println()
	fmt.Println("JSON_PRODUCT1")
	fmt.Printf("%s\n", product)
	fmt.Println()

	// Put it back in the original String will have to Marshall it twice
	var combinedBytes []byte
	bytesOuter, errO := json.Marshal(product)
	if errO != nil {
		log.Fatal(errO)
		os.Exit(0)
	}

	if len(product.Data) > 0 {
		bytesInner, errI := json.Marshal(product.Data)
		if errI != nil {
			log.Fatal(errI)
			os.Exit(0)
		}
		byteSpaceSlice := []byte(" ")
		byteCommaSlice := []byte(",")
		bytesOuter[len(bytesOuter)-1] = byteSpaceSlice[0]
		bytesInner[0] = byteCommaSlice[0]
		combinedBytes = append(bytesOuter, bytesInner...)
	} else {
		combinedBytes = bytesOuter
	}

	fmt.Println()
	fmt.Println("    ------ Original JSON ------")
	fmt.Println(JSON_PRODUCT1)
	fmt.Println("    ------ Reconstructed JSON ------")
	fmt.Println(string(combinedBytes))
	fmt.Println()
}
