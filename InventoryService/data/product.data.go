package data

import (
	"encoding/json"
	"fmt"
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

func GetProduct(productID int) *models.Product {
	productMap.RLock()
	defer productMap.RUnlock()

	product, ok := productMap.prodMap[productID]

	if ok {
		return &product
	}

	return nil
}

func removeProduct(productID int) {
	productMap.Lock()
	defer productMap.Unlock()

	delete(productMap.prodMap, productID)
}

func GetProductList() []models.Product {
	productMap.RLock()
	defer productMap.RUnlock()

	products := make([]models.Product, 0, len(productMap.prodMap))

	for _, prod := range productMap.prodMap {
		products = append(products, prod)
	}
	return products
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
		oldProduct := getProduct(product.ProductID)

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
