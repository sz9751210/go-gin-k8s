package service

import (
	"time"
)

type dataSelector struct {
	GenericDataList []DataCell
	DataSelect      *DataSelectQuery
}

type DataCell interface {
	GetCreation() time.Time
	GetName() string
}

type DataSelectQuery struct {
	Filter   *FilterQuery
	Paginate *PaginateQuery
}

type FilterQuery struct {
	Name string
}

type PaginateQuery struct {
	Limit int
	Page  int
}
