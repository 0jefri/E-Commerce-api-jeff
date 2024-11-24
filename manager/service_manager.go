package manager

import (
	"github.com/e-commerce-api/api/address"
	"github.com/e-commerce-api/api/banner"
	"github.com/e-commerce-api/api/cart"
	"github.com/e-commerce-api/api/category"
	"github.com/e-commerce-api/api/orders"
	"github.com/e-commerce-api/api/product"
	"github.com/e-commerce-api/api/users"
	"github.com/e-commerce-api/api/wishlist"
)

type ServiceManager interface {
	UserService() users.Service
	BannerService() banner.BannerService
	CategoryService() category.CategoryService
	ProductService() product.ProductService
	BestSellingProductsService() product.ProductService
	WishlistService() wishlist.WishlistService
	CartService() cart.CartService
	OrderService() orders.OrderService
	AddresService() address.AddressService
}

type serviceManager struct {
	repoManager RepoManager
}

func NewServiceManager(repo RepoManager) ServiceManager {
	return &serviceManager{
		repoManager: repo,
	}
}

func (m *serviceManager) AddresService() address.AddressService {
	return address.NewAddressService(m.repoManager.AddresRepo())
}
func (m *serviceManager) OrderService() orders.OrderService {
	return orders.NewOrderService(m.repoManager.OrderRepo())
}
func (m *serviceManager) UserService() users.Service {
	return users.NewService(m.repoManager.UserRepo())
}

func (m *serviceManager) BannerService() banner.BannerService {
	return banner.NewBannerService(m.repoManager.BannerRepo())
}

func (m *serviceManager) CategoryService() category.CategoryService {
	return category.NewService(m.repoManager.CategoryRepo())
}

func (m *serviceManager) ProductService() product.ProductService {
	return product.NewProductService(m.repoManager.ProductRepo())
}

func (m *serviceManager) BestSellingProductsService() product.ProductService {
	return product.NewProductService(m.repoManager.BestSellingProductsRepo())
}

func (m *serviceManager) WishlistService() wishlist.WishlistService {
	return wishlist.NewWishlistService(m.repoManager.WishlistRepo())
}

func (m *serviceManager) CartService() cart.CartService {
	return cart.NewCartService(
		m.repoManager.CartRepo(),
		m.ProductService(), // Tambahkan ProductService di sini
	)
}
