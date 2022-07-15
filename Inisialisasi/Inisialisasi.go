package Inisialisasi

import (
	"RestAPI-GETNPOST/Entity"
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

type Model struct {
	//ModelProduct      interface{}
	//ModelCart         interface{}
	//ModelShoppingCart interface{}
	ModelUser interface{}
}

func GetGormConn(host, user, dbname, password string, port int) (*gorm.DB, error) {
	dsn := fmt.Sprintf("host=%s port=%d user=%s dbname=%s password=%s sslmode=disable", host, port, user, dbname, password)
	return gorm.Open(postgres.Open(dsn), &gorm.Config{})
}

func Initialize() (db *gorm.DB, err error) {
	fmt.Println("Connecting to Database please wait..")
	db, err = GetGormConn("localhost", "postgres", "postgres", "32c560b2", 5432)

	if err != nil {
		panic("Failed to Connect Database")
	}
	fmt.Println("Register Models of Product")
	//for _, modelproduct := range RegisterModelsProduct() {
	//	err = db.Debug().AutoMigrate(modelproduct.ModelProduct)
	//	if err != nil {
	//		log.Fatal(err)
	//
	//	}
	//}
	//fmt.Println("Register Models of Cart")
	//for _, modelcart := range RegisterModelsCart() {
	//	err = db.Debug().AutoMigrate(modelcart.ModelCart)
	//	if err != nil {
	//		log.Fatal(err)
	//	}
	fmt.Println("Register Models of User")
	for _, modeluser := range RegisterModelsUser() {
		err = db.Debug().AutoMigrate(modeluser.ModelUser)
		if err != nil {
			log.Fatal(err)
		}
	}
	//fmt.Println("Register Models of Shopping Cart")
	//for _, modelshoppingcart := range RegisterModelsShoppingCart() {
	//	err = db.Debug().AutoMigrate(modelshoppingcart.ModelShoppingCart)
	//	if err != nil {
	//		log.Fatal(err)
	//	}
	//}

	fmt.Println("Success Migrate Database")
	return db, err
}

//func RegisterModelsProduct() []Model {
//	return []Model{
//		{ModelProduct: Entity.Product{}},
//	}
//}
//func RegisterModelsCart() []Model {
//	return []Model{
//		{ModelCart: Entity.Cart{}},
//	}
//}
//func RegisterModelsShoppingCart() []Model {
//	return []Model{
//		{ModelShoppingCart: Entity.ShoppingCart{}},
//	}
//}
func RegisterModelsUser() []Model {
	return []Model{
		{ModelUser: Entity.User{}},
	}
}
