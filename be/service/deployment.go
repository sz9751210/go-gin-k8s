package service

import (
	"context"
	"errors"

	"github.com/wonderivan/logger"
	appsv1 "k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/resource"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/util/intstr"
)

var Deployment deployment

type deployment struct{}

// 定義列表的返回內容，Item是 deployment 元素列表，Total 為 deployment 總數。
type DeploymentsResp struct {
	Total int                 `json:"total"`
	Items []appsv1.Deployment `json:"items"`
}

type DeployCreate struct {
	Name            string            `json:"name"`
	Namespace       string            `json:"namespace"`
	Replicas        int32             `json:"replicas"`
	Image           string            `json:"image"`
	Label           map[string]string `json:"label"`
	Cpu             string            `json:"cpu"`
	Memory          string            `json:"memory"`
	ContainerPort   int32             `json:"container_port"`
	HealthCheck     bool              `json:"health_check"`
	HealthCheckPath string            `json:"health_path"`
}

func (d *deployment) GetDeployments(filterName, namespace string, limit, page int) (deploymentsResp *DeploymentsResp, err error) {
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

	data := filtered.Sort().Paginate()
	deployments := d.fromCells(data.GenericDataList)
	deploymentsResp = &DeploymentsResp{
		Total: total,
		Items: deployments,
	}

	return deploymentsResp, nil
}

func (d *deployment) GetDeploymentDetail(deloymentName, namespace string) (deployment *appsv1.Deployment, err error) {
	deployment, err = K8s.ClientSet.AppsV1().Deployments(namespace).Get(context.TODO(), deloymentName, metav1.GetOptions{})
	if err != nil {
		logger.Error("get deployment detail error: ", err.Error())
		return nil, errors.New("get deployment detail error" + err.Error())
	}
	return deployment, nil
}

func (d *deployment) ScaleDeployment(deploymentName, namespace string, scaleNum int) (replica int32, err error) {
	scale, err := K8s.ClientSet.AppsV1().Deployments(namespace).GetScale(context.TODO(), deploymentName, metav1.GetOptions{})
	if err != nil {
		logger.Error("get deployment scale error: ", err.Error())
		return 0, errors.New("get deployment scale error" + err.Error())
	}
	// 修改 scale 的副本數
	scale.Spec.Replicas = int32(scaleNum)

	// 更新 scale 的副本數，返回更新後的副本數
	newScale, err := K8s.ClientSet.AppsV1().Deployments(namespace).UpdateScale(context.TODO(), deploymentName, scale, metav1.UpdateOptions{})
	if err != nil {
		logger.Error("update deployment scale error: ", err.Error())
		return 0, errors.New("update deployment scale error" + err.Error())
	}
	return newScale.Spec.Replicas, nil
}

func (d *deployment) CreateDeployment(data *DeployCreate) error {
	replicas := data.Replicas
	deployment := &appsv1.Deployment{
		ObjectMeta: metav1.ObjectMeta{
			Name:      data.Name,
			Namespace: data.Namespace,
			Labels:    data.Label,
		},
		Spec: appsv1.DeploymentSpec{
			Replicas: &replicas,
			Selector: &metav1.LabelSelector{
				MatchLabels: data.Label,
			},
			Template: corev1.PodTemplateSpec{
				ObjectMeta: metav1.ObjectMeta{
					Labels: data.Label,
				},
				Spec: corev1.PodSpec{
					Containers: []corev1.Container{
						{
							Name:  data.Name,
							Image: data.Image,
							Ports: []corev1.ContainerPort{
								{
									Name:          "http",
									ContainerPort: data.ContainerPort,
								},
							},
							Resources: corev1.ResourceRequirements{
								Limits: corev1.ResourceList{
									corev1.ResourceCPU:    resource.MustParse(data.Cpu),
									corev1.ResourceMemory: resource.MustParse(data.Memory),
								},
								Requests: corev1.ResourceList{
									corev1.ResourceCPU:    resource.MustParse(data.Cpu),
									corev1.ResourceMemory: resource.MustParse(data.Memory),
								},
							},
						},
					},
				},
			},
		},
	}

	if data.HealthCheck {
		probe := &corev1.Probe{
			ProbeHandler: corev1.ProbeHandler{
				HTTPGet: &corev1.HTTPGetAction{
					Path: data.HealthCheckPath,
					Port: intstr.IntOrString{
						Type:   intstr.Int,
						IntVal: data.ContainerPort,
					},
				},
			},
			InitialDelaySeconds: 5,
			PeriodSeconds:       5,
			TimeoutSeconds:      5,
			SuccessThreshold:    1,
			FailureThreshold:    3,
		}
		deployment.Spec.Template.Spec.Containers[0].ReadinessProbe = probe
		deployment.Spec.Template.Spec.Containers[0].LivenessProbe = probe
	}

	_, err := K8s.ClientSet.AppsV1().Deployments(data.Namespace).Create(context.TODO(), deployment, metav1.CreateOptions{})
	if err != nil {
		logger.Error("create deployment error: ", err.Error())
		return errors.New("create deployment error: " + err.Error())
	}

	return nil
}

func (d *deployment) DeleteDeployment(deploymentName, namespace string) error {
	err := K8s.ClientSet.AppsV1().Deployments(namespace).Delete(context.TODO(), deploymentName, metav1.DeleteOptions{})
	if err != nil {
		logger.Error("delete deployment error: ", err.Error())
		return errors.New("delete deployment error: " + err.Error())
	}
	return nil
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
