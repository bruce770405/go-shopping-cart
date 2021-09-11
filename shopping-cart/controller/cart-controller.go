package controller

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"net/http"
	"shopping-cart/common"
	"shopping-cart/entities"
	"shopping-cart/persistent/dao"
	"shopping-cart/utils"
)

// ShoopingCart manages Shopping-cart CRUD
type Shopping struct {
	shoppingCartDao dao.ShoppingCart
}

type UserClaims struct {
	Name string `json:"name"`
	jwt.StandardClaims
}

// AddProdInCart godoc
// @Summary Add a new prod in cart
// @Description Add a new prod in customer cart
// @Tags movie
// @Accept  json
// @Produce  json
// @Param Authorization header string true "Token"
// @Param user body entities.AddProd true "Add Prod in cart"
// @Failure 401 {object} entities.Error
// @Success 200 {object} entities.Message
// @Router /cart [post]
func (m *Shopping) AddProdInCart(ctx *gin.Context) {
	var addProd entities.AddProd
	if err := ctx.BindJSON(&addProd); err != nil {
		ctx.AbortWithStatus(http.StatusBadRequest)
		return
	}

	tokenString := ctx.GetHeader("Authorization")
	customerId, err := utils.Jwt{Token: tokenString}.GetCustomerIdByJWT()
	if err != nil {
		ctx.JSON(http.StatusForbidden, entities.Error{Code: common.StatusCodeUnknown, Message: err.Error()})
		return
	}

	m.shoppingCartDao.CustomerId = customerId
	err = m.shoppingCartDao.AddProdInCart(addProd)
	if err == nil {
		ctx.JSON(http.StatusOK, entities.Message{Message: "Successfully"})
	} else {
		ctx.JSON(http.StatusForbidden, entities.Error{Code: common.StatusCodeUnknown, Message: err.Error()})
		log.Debug("[ERROR]: ", err)
	}
}
