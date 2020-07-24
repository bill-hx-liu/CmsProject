package controller

import (
	"CmsProject/service"
	"context"
)

type OrderController struct {
	Ctx context.Context
	Service service.OrderService
}
