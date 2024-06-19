package service

import (
	appsv1 "k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"
	"sort"
	"strings"
	"time"
)

// dataSelector 結構用於操作和變換一組 DataCell。
type dataSelector struct {
	GenericDataList []DataCell       // DataCell 切片，存放數據元素
	DataSelect      *DataSelectQuery // 指定過濾和分頁的查詢參數
}

// DataCell 接口定義了所有數據單元必須實現的基本方法。
type DataCell interface {
	GetCreation() time.Time // 獲取數據創建時間
	GetName() string        // 獲取數據名稱
}

// DataSelectQuery 結構包含過濾和分頁的具體參數。
type DataSelectQuery struct {
	Filter   *FilterQuery   // 過濾條件
	Paginate *PaginateQuery // 分頁條件
}

type FilterQuery struct {
	Name string // 用於過濾的名稱條件
}

type PaginateQuery struct {
	Limit int // 每頁顯示的項目數
	Page  int // 當前頁碼
}

// Len 返回 GenericDataList 中的元素數量。
func (d *dataSelector) Len() int {
	return len(d.GenericDataList)
}

// Swap 交換 GenericDataList 中的兩個元素。
func (d *dataSelector) Swap(i, j int) {
	d.GenericDataList[i], d.GenericDataList[j] = d.GenericDataList[j], d.GenericDataList[i]
}

// Less 比較兩個元素的創建時間以決定它們的排序。
func (d *dataSelector) Less(i, j int) bool {
	a := d.GenericDataList[i].GetCreation()
	b := d.GenericDataList[j].GetCreation()
	return b.Before(a) // 比較時間，如果 b 在 a 之前則返回 true
}

// Sort 對 GenericDataList 進行排序（基於創建時間）。
func (d *dataSelector) Sort() *dataSelector {
	sort.Sort(d) // 調用 sort.Sort 方法實現排序
	return d
}

// Filter 根據 Name 過濾 GenericDataList 中的元素。
func (d *dataSelector) Filter() *dataSelector {
	if d.DataSelect.Filter.Name == "" {
		return d // 如果過濾名稱為空，則不進行任何操作
	}
	filtered := []DataCell{}
	for _, value := range d.GenericDataList {
		if strings.Contains(value.GetName(), d.DataSelect.Filter.Name) {
			filtered = append(filtered, value) // 符合過濾條件的元素被添加到新切片中
		}
	}
	d.GenericDataList = filtered // 更新列表為過濾後的結果
	return d
}

// Paginate 根據 Limit 和 Page 參數對 GenericDataList 進行分頁。
func (d *dataSelector) Paginate() *dataSelector {
	limit := d.DataSelect.Paginate.Limit
	page := d.DataSelect.Paginate.Page
	if limit <= 0 || page <= 0 {
		return d // 如果參數無效則不進行分頁
	}
	startIndex := (page - 1) * limit
	endIndex := page * limit
	if startIndex >= len(d.GenericDataList) {
		d.GenericDataList = []DataCell{} // 超出範圍，返回空切片
		return d
	}
	if endIndex > len(d.GenericDataList) {
		endIndex = len(d.GenericDataList) // 調整結束索引以避免越界
	}
	d.GenericDataList = d.GenericDataList[startIndex:endIndex] // 獲取分頁切片
	return d
}

// podCell 是 corev1.Pod 的包裝，以滿足 DataCell 接口。
type podCell corev1.Pod

// GetCreation 返回 Pod 的創建時間戳。
func (p podCell) GetCreation() time.Time {
	return p.CreationTimestamp.Time
}

// GetName 返回 Pod 的名稱。
func (p podCell) GetName() string {
	return p.Name
}

type deploymentCell appsv1.Deployment

func (d deploymentCell) GetCreation() time.Time {
	return d.CreationTimestamp.Time
}

func (d deploymentCell) GetName() string {
	return d.Name
}

type namespaceCell corev1.Namespace

func (ns namespaceCell) GetCreation() time.Time {
	return ns.CreationTimestamp.Time
}

func (ns namespaceCell) GetName() string{
	return ns.Name
}
