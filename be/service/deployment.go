package service

import (
	appsv1 "k8s.io/api/apps/v1"
)

var Deployment deployment

type deployment struct{}

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
