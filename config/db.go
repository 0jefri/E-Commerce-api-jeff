package config

import (
	"fmt"
	"sync"
	"time"

	"github.com/e-commerce-api/api/address"
	"github.com/e-commerce-api/api/banner"
	"github.com/e-commerce-api/api/cart"
	"github.com/e-commerce-api/api/category"
	"github.com/e-commerce-api/api/order"
	"github.com/e-commerce-api/api/orders"
	"github.com/e-commerce-api/api/product"
	"github.com/e-commerce-api/api/users"
	"github.com/e-commerce-api/api/wishlist"
	_ "github.com/lib/pq"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB
var once sync.Once

func InitDB() {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=disable TimeZone=Asia/Shanghai", Cfg.Database.Host, Cfg.Database.Username, Cfg.Database.Password, Cfg.Database.Dbname, Cfg.Database.Port)
	Db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		NowFunc: func() time.Time {
			return time.Now().Local()
		},
		PrepareStmt: true,
	})

	if err != nil {
		panic(err)
	}

	once.Do(func() {
		DB = Db
		fmt.Println("Successfully Connected To Database!")
	})
}

func SyncDB() {
	if err := DB.AutoMigrate(&users.Users{}, &banner.Banner{}, &category.Category{}, &product.Product{}, &order.Order{}, &wishlist.Wishlist{}, &cart.Cart{}, &orders.Orderes{}, &address.Address{}); err != nil {
		fmt.Printf("AutoMigrate error: %s\n", err)
		panic(err)
	} else {
		fmt.Println("Database migrated successfully!")
	}

	// if err := seeder.SeedBanners(DB); err != nil {
	// 	panic("Failed to seed data")
	// }
}

// func SeedBanners(db *gorm.DB) error {
// 	// Data awal untuk tabel banners
// 	banners := []banner.Banner{
// 		{
// 			ID:       uuid.New().String(),
// 			Photo:    "banner1.jpg",
// 			Title:    "Sale",
// 			Subtitle: "50% off on all items",
// 			PathPage: "/sale",
// 		},
// 		{
// 			ID:       uuid.New().String(),
// 			Photo:    "banner2.jpg",
// 			Title:    "New Arrivals",
// 			Subtitle: "Check out our latest products",
// 			PathPage: "/new-arrivals",
// 		},
// 		{
// 			ID:       uuid.New().String(),
// 			Photo:    "banner3.jpg",
// 			Title:    "Exclusive Deals",
// 			Subtitle: "Special prices for members",
// 			PathPage: "/deals",
// 		},
// 	}

// Insert data ke tabel banners
// 	for _, banner := range banners {
// 		if err := db.Create(&banner).Error; err != nil {
// 			fmt.Println("error:", err)
// 			return err
// 		}
// 	}
// 	return nil
// }
