package service

import (
	"k8s-go-gin/config"

	"github.com/wonderivan/logger"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
)

var K8s k8s

type k8s struct {
	ClientSet *kubernetes.Clientset
}

func (k *k8s) Init() {
	conf, err := clientcmd.BuildConfigFromFlags("", config.Kubeconfig)
	if err != nil {
		panic(err)
	}

	clientSet, err := kubernetes.NewForConfig(conf)
	if err != nil {
		panic("create k8s clientSet error" + err.Error())
	} else {
		logger.Info("create k8s clientSet success")
	}
	k.ClientSet = clientSet
}
