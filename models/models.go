package models

import (
	"time"

	"github.com/edgedb/edgedb-go"
)

type Book struct {
	ID           edgedb.UUID           `edgedb:"id" json:"Id"`
	Title        string                `edgedb:"title" json:"Title"`
	Author       string                `edgedb:"author" json:"Author"`
	Genre        string                `edgedb:"genre" json:"Genre"`
	Release_date time.Time             `edgedb:"release_date" json:"Release_date"`
	Quantity     int16                 `edgedb:"quantity" json:"Quantity"`
	Sales        edgedb.OptionalBigInt `edgedb:"sales" json:"Sales"`
}

type Groups struct {
	Key struct {
		Genre  string `edgedb:"genre" json:"Genre"`
		Author string `edgedb:"author" json:"Author"`
	} `edgedb:"key" json:"Key"`
	Grouping []string `edgedb:"grouping" json:"grouping"`
	Elements []Book   `edgedb:"elements" json:"elements"`
}
