package dao

import (
	"context"
	"encoding/json"
	"github.com/go-redis/redis/v8"
	log "github.com/sirupsen/logrus"
	"shopping-cart/entities"
	"shopping-cart/persistent"
	"time"
)

var ctx = context.Background()

type ShoppingCart struct {
	CustomerId string
}

// Insert adds a new Prod into redis'
func (m *ShoppingCart) AddProdInCart(prod entities.AddProd) error {
	val, err := persistent.Database.Client.Get(ctx, m.CustomerId).Result()
	if err == redis.Nil {
		log.Warn(m.CustomerId + " key does not exist")
		array := []entities.AddProd{prod}
		err = m.newCartAdd(array)
		if err != nil {
			return err
		}
	} else if err != nil {
		return err
	}

	data := []byte(val)
	var jsonArr []entities.AddProd
	err = json.Unmarshal(data, &jsonArr)
	if err != nil {
		return err
	}

	log.Debug(jsonArr)

	jsonArr = append(jsonArr, prod)
	err = m.newCartAdd(jsonArr)
	if err != nil {
		return err
	}

	return nil
}

// Insert adds a new Prod into redis'
func (m *ShoppingCart) newCartAdd(array []entities.AddProd) error {
	err := persistent.Database.Client.Set(ctx, m.CustomerId, array, 1*time.Hour).Err()
	if err != nil {
		return err
	}

	return nil
}
