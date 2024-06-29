package service

import (
	"context"
	"errors"

	"github.com/wonderivan/logger"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var Namespace namespace

type namespace struct{}

type NamespaceResp struct {
	Items []corev1.Namespace `json:"items"`
	Total int                `json:"total"`
}

func (ns *namespace) GetNamespaces(filterName string, limit, page int) (namespaceResp *NamespaceResp, err error) {
	namespaceList, err := K8s.ClientSet.CoreV1().Namespaces().List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		logger.Error("get namespace list error: ", err.Error())
		return nil, errors.New("get namespace list error" + err.Error())
	}

	selectableData := &dataSelector{
		GenericDataList: ns.toCells(namespaceList.Items),
		DataSelect: &DataSelectQuery{
			Filter: &FilterQuery{
				Name: filterName,
			},
			Paginate: &PaginateQuery{
				Limit: limit,
				Page:  page,
			},
		},
	}
	filtered := selectableData.Filter()
	total := len(filtered.GenericDataList)
	data := filtered.Sort().Paginate()

	namespaces := ns.fromCells(data.GenericDataList)
	namespaceResp = &NamespaceResp{
		Items: namespaces,
		Total: total,
	}
	return namespaceResp, nil
}

func (ns *namespace) GetNamespaceDetail(namespaceName string) (namespace *corev1.Namespace, err error) {
	namespace, err = K8s.ClientSet.CoreV1().Namespaces().Get(context.TODO(), namespaceName, metav1.GetOptions{})
	if err != nil {
		logger.Error("get namespace detail error: ", err.Error())
		return nil, errors.New("get namespace detail error" + err.Error())
	}
	return namespace, nil
}

func (ns *namespace) DeleteNamespace(namespaceName string) (err error) {
	err = K8s.ClientSet.CoreV1().Namespaces().Delete(context.TODO(), namespaceName, metav1.DeleteOptions{})
	if err != nil {
		logger.Error("delete namespace error: ", err.Error())
		return errors.New("delete namespace error" + err.Error())
	}

	return nil
}



func (ns *namespace) toCells(std []corev1.Namespace) []DataCell {
	cells := make([]DataCell, len(std))
	for i := range std {
		cells[i] = namespaceCell(std[i])
	}
	return cells
}

func (ns *namespace) fromCells(cells []DataCell) []corev1.Namespace {
	// 製作空的儲存原始 Namespace 結構體的空切片
	namespaces := make([]corev1.Namespace, len(cells))
	// 將 cells 切片中的元素轉換為 namespaceCell 並存入 namespaces 切片中
	for i := range cells {
		namespaces[i] = corev1.Namespace(cells[i].(namespaceCell))
	}
	return namespaces
}
