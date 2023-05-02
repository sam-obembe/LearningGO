package handlers

import (
	"encoding/json"
	"fmt"
	"inventoryService/data"
	"inventoryService/middleware"
	"inventoryService/models"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
)

const productsBasePath = "products"

func checkErr(err error, writer http.ResponseWriter) {
	if err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		return
	}

}

func SetupRoutes(apiBasePath string) {
	handleProducts := http.HandlerFunc(ProductsHandler)
	handleProduct := http.HandlerFunc(ProductHandler)

	productsPath := fmt.Sprintf("%s/%s", apiBasePath, productsBasePath)
	productPath := fmt.Sprintf("%s/%s/", apiBasePath, productsBasePath)

	http.Handle(productsPath, middleware.CorsMiddleware(handleProducts))
	http.Handle(productPath, middleware.CorsMiddleware(handleProduct))
}

func ProductsHandler(writer http.ResponseWriter, req *http.Request) {
	switch req.Method {
	case http.MethodGet:
		productList, err := data.GetProductList()
		if err != nil {
			writer.WriteHeader(http.StatusInternalServerError)
			return
		}

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

		_, err = data.AddOrUpdateProduct(newProduct)
		newProduct.ProductID = data.GetNextId()
		//productList = append(productList, newProduct)

		if err != nil {
			writer.WriteHeader(http.StatusInternalServerError)
		}

		writer.WriteHeader(http.StatusCreated)
		return
	case http.MethodOptions:
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
	case http.MethodDelete:
		data.RemoveProduct(productID)
	default:
		writer.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
}
