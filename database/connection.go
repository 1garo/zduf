package database

import (
	"context"
	"fmt"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
)

var Conn *pgxpool.Pool

func NewConnection(url string) (*pgxpool.Pool, error) {
	var err error

	ctx, cancel := context.WithTimeout(context.Background(), 10 * time.Second)
	defer cancel()

	Conn, err = pgxpool.New(ctx, url)
	if err != nil {
		return nil, fmt.Errorf("Unable to create database pool: %w\n", err)
	}

	return Conn, nil
}



