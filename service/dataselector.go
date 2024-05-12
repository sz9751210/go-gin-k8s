package service

import (
	"sort"
	"strings"
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

func (d *dataSelector) Len() int {
	return len(d.GenericDataList)
}

func (d *dataSelector) Swap(i, j int) {
	d.GenericDataList[i], d.GenericDataList[j] = d.GenericDataList[j], d.GenericDataList[i]
}

func (d *dataSelector) Less(i, j int) bool {
	a := d.GenericDataList[i].GetCreation()
	b := d.GenericDataList[j].GetCreation()
	return b.Before(a)
}

func (d *dataSelector) Sort() *dataSelector {
	sort.Sort(d)
	return d
}

func (d *dataSelector) Filter() *dataSelector {
	if d.DataSelect.Filter.Name == "" {
		return d
	}
	filtered := []DataCell{}
	for _, value := range d.GenericDataList {
		objName := value.GetName()
		if strings.Contains(objName, d.DataSelect.Filter.Name) {
			filtered = append(filtered, value)
		}
	}
	d.GenericDataList = filtered
	return d
}
