package server

import (
	"context"
	"log"
)

type Currency struct {
	logger *log.Logger
}

func (c *Currency) GetRate(context.Context, *RateRequest) (*RateResponse, error) {

}
