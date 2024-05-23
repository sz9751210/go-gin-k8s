package service

import (
	"context"
	"errors"

	"github.com/gin-gonic/gin"
	"github.com/wonderivan/logger"
	appsv1 "k8s.io/api/apps/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var Deployment deployment

type deployment struct{}

// 定義列表的返回內容，Item是 deployment 元素列表，Total 為 deployment 總數。
type DeploymentsResp struct {
	Total int                `json:"total"`
	Items []appsv1.Deployment `json:"items"`
}

func(d *deployment) GetDeployments(filterName, namespace string, limit, page int) (deploymentsResp *DeploymentsResp, err error) {
	deploymentList, err := K8s.ClientSet.AppsV1().Deployments(namespace).List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		logger.Error("get deployment list error: ", err.Error())
		return nil, errors.New("get deployment list error" + err.Error())
	}
	selectableData := &dataSelector{
		GenericDataList: d.toCells(deploymentList.Items),
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

	data  := filtered.Sort().Paginate()
	deployments := d.fromCells(data.GenericDataList)
	deploymentsResp = &DeploymentsResp{
		Total: total,
		Items: deployments,
	}

	return deploymentsResp, nil
}

func(d *deployment) GetDeploymentDetail(deloymentName, namespace string)(deployment *appsv1.Deployment, err error) {
	deployment, err = K8s.ClientSet.AppsV1().Deployments(namespace).Get(context.TODO(), deloymentName, metav1.GetOptions{})
	if err != nil {
		logger.Error("get deployment detail error: ", err.Error())
		return nil, errors.New("get deployment detail error" + err.Error())
	}
	return deployment, nil
}


// toCells 方法將 corev1.Pod 的切片轉換成 DataCell 的切片。
func (d *deployment) toCells(deployments []appsv1.Deployment) []DataCell {
	cells := make([]DataCell, len(deployments)) // 創建與 pods 相同長度的 DataCell 切片
	for i := range deployments {
		cells[i] = deploymentCell(deployments[i]) // 假定 podCell 是 DataCell 的一個實現
	}
	return cells
}

// fromCells 方法將 DataCell 的切片轉換回 corev1.Pod 的切片。
func (d *deployment) fromCells(cells []DataCell) []appsv1.Deployment {
	deployments := make([]appsv1.Deployment, len(cells)) // 創建與 cells 相同長度的 corev1.Pod 切片
	for i, cell := range cells {
		dc, ok := cell.(deploymentCell) // 假定 DataCell 是 podCell 的一個實現
		if !ok {
			continue
		}
		deployments[i] = appsv1.Deployment(dc) // 進行類型斷言，將 DataCell 轉換為 podCell，然後轉為 corev1.Pod
	}
	return deployments
}
