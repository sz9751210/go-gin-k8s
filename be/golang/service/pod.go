package service

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"k8s-go-gin/config"

	"github.com/wonderivan/logger"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// Pod 是全局可變的 pod 實例。
var Pod pod

// pod 結構體是空的，可能是因為實際的實現細節在其他地方或尚未實現。
type pod struct {
}

type PodsResp struct {
	Total int          `json:"total"`
	Items []corev1.Pod `json:"items"`
}

type PodsNp struct {
	Namespace string
	PodNum    int
}

func (p *pod) GetPods(filterName, namespace string, limit, page int) (podsResp *PodsResp, err error) {
	podList, err := K8s.ClientSet.CoreV1().Pods(namespace).List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		logger.Info("get podList error: ", err.Error())
		return nil, errors.New("get podList error" + err.Error())
	}
	selectableData := &dataSelector{
		GenericDataList: p.toCells(podList.Items),
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
	pods := p.fromCells(data.GenericDataList)

	fmt.Println("處理前的 pods 列表: ")
	for _, pod := range podList.Items {
		fmt.Println(pod.Name, pod.CreationTimestamp.Time)
	}

	fmt.Println("處理後的 pods 列表: ")
	for _, pod := range pods {
		fmt.Println(pod.Name, pod.CreationTimestamp.Time)
	}

	podsResp = &PodsResp{
		Total: total,
		Items: pods,
	}
	return podsResp, nil

}

func (p *pod) GetPodDetail(podName, namespace string) (podDetail *corev1.Pod, err error) {
	pod, err := K8s.ClientSet.CoreV1().Pods(namespace).Get(context.TODO(), podName, metav1.GetOptions{})
	if err != nil {
		logger.Error("get pod error: ", err.Error())
		return nil, errors.New("get pod error" + err.Error())
	}
	return pod, nil
}

func (p *pod) DeletePod(podName, namespace string) (err error) {
	err = K8s.ClientSet.CoreV1().Pods(namespace).Delete(context.TODO(), podName, metav1.DeleteOptions{})
	if err != nil {
		logger.Error("delete pod error: ", err.Error())
		return errors.New("delete error" + err.Error())
	}
	return nil
}

func (p *pod) UpdatePod(podName, namespace, content string) (err error) {
	pod := &corev1.Pod{}
	err = json.Unmarshal([]byte(content), pod)
	if err != nil {
		logger.Error("unmarshal pod error," + err.Error())
		return errors.New("unmarshal pod error" + err.Error())
	}

	_, err = K8s.ClientSet.CoreV1().Pods(namespace).Update(context.TODO(), pod, metav1.UpdateOptions{})
	if err != nil {
		logger.Error("update pod error, " + err.Error())
		return errors.New("update pod error" + err.Error())
	}
	return nil
}

func (p *pod) GetPodContainers(podName, namespace string) (containers []string, err error) {
	pod, err := p.GetPodDetail(podName, namespace)
	if err != nil {
		return nil, err
	}
	for _, container := range pod.Spec.Containers {
		containers = append(containers, container.Name)
	}
	return containers, nil
}

func (p *pod) GetPodLog(podName, containerName, namespace string) (log string, err error) {
	lineLimit := int64(config.PodLogTailLine)
	option := &corev1.PodLogOptions{
		Container: containerName,
		TailLines: &lineLimit,
	}

	req := K8s.ClientSet.CoreV1().Pods(namespace).GetLogs(podName, option)
	podLogs, err := req.Stream(context.TODO())
	if err != nil {
		logger.Error(errors.New("get pod log error" + err.Error()))
		return "", errors.New("get pod log error" + err.Error())
	}
	defer podLogs.Close()
	buf := new(bytes.Buffer)
	_, err = io.Copy(buf, podLogs)
	// _, err = buf.ReadFrom(podLogs)
	if err != nil {
		logger.Error(errors.New("copy pod log error" + err.Error()))
		return "", errors.New("copy pod log error" + err.Error())
	}
	return buf.String(), nil
}

// get pod number per namespace
func (p *pod) GetPodNumPerNp() (podsNps []*PodsNp, err error) {
	namespaceList, err := K8s.ClientSet.CoreV1().Namespaces().List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		logger.Error("get namespace list error: ", err.Error())
		return nil, errors.New("get namespace list error" + err.Error())
	}
	for _, namesapce := range namespaceList.Items {
		podList, err := K8s.ClientSet.CoreV1().Pods(namesapce.Name).List(context.TODO(), metav1.ListOptions{})
		if err != nil {
			logger.Error("get pod list error: ", err.Error())
			return nil, errors.New("get pod list error" + err.Error())
		}
		podsNp := &PodsNp{
			Namespace: namesapce.Name,
			PodNum:    len(podList.Items),
		}
		podsNps = append(podsNps, podsNp)
	}

	return podsNps, nil
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
