package usecases

import "github.com/fiatfour/go-clean/entities"

type OrderRepository interface {
	Save(order entities.Order) error
}
