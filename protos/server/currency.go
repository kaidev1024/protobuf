package server

import (
	"context"
	"log"

	protos "github.com/kaidev1024/protobuf/protos/currency"
)

type Currency struct {
	logger *log.Logger
}

func NewCurrency(logger *log.Logger) *Currency {
	return &Currency{logger}
}

func (c *Currency) GetRate(ctx context.Context, rr *protos.RateRequest) (*protos.RateResponse, error) {
	c.logger.Println("Handle GetRate: ", "base: ", rr.GetBase())
	c.logger.Println("Handle GetRate: ", "destination: ", rr.GetDestination())

	return &protos.RateResponse{Rate: 0.5}, nil
}
