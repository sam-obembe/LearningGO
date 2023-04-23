package handlers

import (
	"encoding/json"
	"inventoryService/data"
	"inventoryService/models"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"strings"
)

var productList []models.Product

func init() {
	productsJson := `[
		{
		  "productId": 1,
		  "manufacturer": "Johns-Jenkins",
		  "sku": "p5z343vdS",
		  "upc": "939581000000",
		  "pricePerUnit": "497.45",
		  "quantityOnHand": 9703,
		  "productName": "sticky note"
		},
		{
		  "productId": 2,
		  "manufacturer": "Hessel, Schimmel and Feeney",
		  "sku": "i7v300kmx",
		  "upc": "740979000000",
		  "pricePerUnit": "282.29",
		  "quantityOnHand": 9217,
		  "productName": "leg warmers"
		},
		{
		  "productId": 3,
		  "manufacturer": "Swaniawski, Bartoletti and Bruen",
		  "sku": "q0L657ys7",
		  "upc": "111730000000",
		  "pricePerUnit": "436.26",
		  "quantityOnHand": 5905,
		  "productName": "lamp shade"
		},
		{
		  "productId": 4,
		  "manufacturer": "Runolfsdottir, Littel and Dicki",
		  "sku": "x78426lq1",
		  "upc": "93986215015",
		  "pricePerUnit": "537.90",
		  "quantityOnHand": 2642,
		  "productName": "flowers"
		},
		{
		  "productId": 5,
		  "manufacturer": "Kuhn, Cronin and Spencer",
		  "sku": "r4X793mdR",
		  "upc": "260149000000",
		  "pricePerUnit": "112.10",
		  "quantityOnHand": 6144,
		  "productName": "clamp"
		}]`

	err := json.Unmarshal([]byte(productsJson), &productList)

	if err != nil {
		log.Fatal(err)
	}
}

func getNextId() int {
	highestId := -1
	for _, val := range productList {
		if highestId < val.ProductID {
			highestId = val.ProductID
		}
	}

	return highestId + 1
}

func findProductById(productId int) (*models.Product, int) {
	for i, product := range productList {
		if product.ProductID == productId {
			return &product, i
		}
	}
	return nil, 0
}

func ProductsHandler(writer http.ResponseWriter, req *http.Request) {
	switch req.Method {
	case http.MethodGet:
		productList := data.GetProductList()
		productJson, err := json.Marshal(productList)
		if err != nil {
			writer.WriteHeader(http.StatusInternalServerError)
			return
		}
		writer.Header().Set("Content-Type", "application/json")
		writer.Write(productJson)
	case http.MethodPost:
		var newProduct models.Product
		bodyContent, err := ioutil.ReadAll(req.Body)
		if err != nil {
			writer.WriteHeader(http.StatusBadRequest)
			return
		}

		err = json.Unmarshal(bodyContent, &newProduct)
		if err != nil {
			writer.WriteHeader(http.StatusBadRequest)
			return
		}

		if newProduct.ProductID != 0 {
			writer.WriteHeader(http.StatusBadRequest)
			return
		}

		_, err := data.AddOrUpdateProduct(newProduct)
		newProduct.ProductID = data.GetNextId()
		//productList = append(productList, newProduct)

		if err != nil {
			writer.WriteHeader(http.StatusInternalServerError)
		}

		writer.WriteHeader(http.StatusCreated)
		return

	}
}

func ProductHandler(writer http.ResponseWriter, req *http.Request) {
	urlPathSegments := strings.Split(req.URL.Path, "products/")
	pathIndex := len(urlPathSegments) - 1
	productID, err := strconv.Atoi(urlPathSegments[pathIndex])

	if err != nil {
		writer.WriteHeader(http.StatusNotFound)
		return
	}

	product := data.GetProduct(productID)

	if product == nil {
		writer.WriteHeader(http.StatusNotFound)
		return
	}

	switch req.Method {
	case http.MethodGet:
		productJson, err := json.Marshal(product)
		if err != nil {
			writer.WriteHeader(http.StatusInternalServerError)
			return
		}
		writer.Header().Set("Content-Type", "application/json")
		writer.Write(productJson)
	case http.MethodPut:
		var updatedProduct models.Product
		bodyBytes, err := ioutil.ReadAll(req.Body)
		if err != nil {
			writer.WriteHeader(http.StatusBadRequest)
			return
		}
		err = json.Unmarshal(bodyBytes, &updatedProduct)

		if err != nil {
			writer.WriteHeader(http.StatusBadRequest)
			return
		}

		if updatedProduct.ProductID != productID {
			writer.WriteHeader(http.StatusBadRequest)
			return
		}

		data.AddOrUpdateProduct(updatedProduct)
		writer.WriteHeader(http.StatusOK)
	default:
		writer.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
}
