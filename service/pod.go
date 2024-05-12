package service

import corev1 "k8s.io/api/core/v1"

// Pod 是全局可變的 pod 實例。
var Pod pod

// pod 結構體是空的，可能是因為實際的實現細節在其他地方或尚未實現。
type pod struct {
}

// toCells 方法將 corev1.Pod 的切片轉換成 DataCell 的切片。
func (p *pod) toCells(pods []corev1.Pod) []DataCell {
	cells := make([]DataCell, len(pods)) // 創建與 pods 相同長度的 DataCell 切片
	for i := range pods {
		cells[i] = podCell(pods[i]) // 假定 podCell 是 DataCell 的一個實現
	}
	return cells
}

// fromCells 方法將 DataCell 的切片轉換回 corev1.Pod 的切片。
func (p *pod) fromCells(cells []DataCell) []corev1.Pod {
	pods := make([]corev1.Pod, len(cells)) // 創建與 cells 相同長度的 corev1.Pod 切片
	for i, cell := range cells {
		pc, ok := cell.(podCell) // 假定 DataCell 是 podCell 的一個實現
		if !ok {
			continue
		}
		pods[i] = corev1.Pod(pc) // 進行類型斷言，將 DataCell 轉換為 podCell，然後轉為 corev1.Pod
	}
	return pods
}
