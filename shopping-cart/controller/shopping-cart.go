package controller

import "shopping-cart/persistent/dao"

// Movie manages Shopping-cart CRUD
type Movie struct {
	shoppingCartDao dao.ShoppingCart
}