package data

import (
	"encoding/json"
	"fmt"
	"inventoryService/database"
	"inventoryService/models"
	"io/ioutil"
	"log"
	"os"
	"sort"
	"sync"
)

var productMap = struct {
	sync.RWMutex
	prodMap map[int]models.Product
}{prodMap: make(map[int]models.Product)}

func init() {
	fmt.Println("loading products ...")
	prodMap, err := loadProductMap()

	productMap.prodMap = prodMap

	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%d products loaded ...\n", len(productMap.prodMap))
}

func loadProductMap() (map[int]models.Product, error) {
	fileName := "products.json"
	_, err := os.Stat(fileName)

	if os.IsNotExist(err) {
		return nil, fmt.Errorf("file [%s] does not exist", fileName)
	}

	file, _ := ioutil.ReadFile(fileName)
	productList := make([]models.Product, 0)

	err = json.Unmarshal([]byte(file), &productList)

	if err != nil {
		log.Fatal(err)
	}

	prodMap := make(map[int]models.Product)

	for i := 0; i < len(productList); i++ {
		prodMap[productList[i].ProductID] = productList[i]
	}

	return prodMap, err
}

func GetProduct(productID int) (*models.Product, error) {
	row := database.DbConn.QueryRow(`SELECT productId, 
	manufacturer, 
	sku, 
	upc, 
	pricePerUnit, 
	quantityOnHand, 
	productName 
	FROM products WHER productId =?`, productID)

	product := &models.Product{}

	row.Scan(&product.Manufacturer, &product.Sku, &product.Upc, &product.PricePerUnit, &product.QuantityOnHand, &product.ProductName)
	productMap.RLock()
	defer productMap.RUnlock()

	product, ok := productMap.prodMap[productID]

	if ok {
		return &product
	}

	return nil
}

func RemoveProduct(productID int) {
	productMap.Lock()
	defer productMap.Unlock()

	delete(productMap.prodMap, productID)
}

func GetProductList() ([]models.Product, error) {
	results, err := database.DbConn.Query(`SELECT productId, 
	manufacturer, 
	sku, 
	upc, 
	pricePerUnit, 
	quantityOnHand, 
	productName 
	FROM products`)

	if err != nil {
		return nil, err
	}
	defer results.Close()
	products := make([]models.Product, 0)

	for results.Next() {
		var product models.Product

		//fields have to be in the same order as query above
		results.Scan(&product.Manufacturer, &product.Sku, &product.Upc, &product.PricePerUnit, &product.QuantityOnHand, &product.ProductName)
		products = append(products, product)
	}

	return products, nil
}

func getProductIds() []int {
	productMap.RLock()
	defer productMap.RUnlock()

	prodIds := make([]int, 0, len(productMap.prodMap))

	for _, prod := range productMap.prodMap {
		prodIds = append(prodIds, prod.ProductID)
	}

	sort.Ints(prodIds)
	return prodIds
}

func GetNextId() int {
	ids := getProductIds()

	maxId := ids[len(ids)-1]

	return maxId + 1
}

func AddOrUpdateProduct(product models.Product) (int, error) {
	addOrUpdateId := -1

	if product.ProductID > 0 {
		oldProduct := GetProduct(product.ProductID)

		if oldProduct == nil {
			return 0, fmt.Errorf("Product id [%d] does not exist", product.ProductID)
		}
		addOrUpdateId = product.ProductID
	} else {
		addOrUpdateId = GetNextId()
		product.ProductID = addOrUpdateId
	}
	productMap.Lock()
	productMap.prodMap[addOrUpdateId] = product
	productMap.Unlock()
	return addOrUpdateId, nil
}
