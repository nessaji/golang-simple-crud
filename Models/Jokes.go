package Models

import (
	_ "github.com/go-sql-driver/mysql"
)

type Joke struct {
	ID     int     `json:"id" binding:"required"`
	Likes  int     `json:"likes"`
	Joke   string  `json:"joke" binding:"required"`
  }

