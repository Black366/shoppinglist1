package models1

import "time"

type ShopList struct {
	Id        int
	Name      string
	CreatedAt time.Time
	Items     []Product
}
