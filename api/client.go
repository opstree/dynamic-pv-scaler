package api

import (
	"dynamic-pv-scaling/logger"
	log "github.com/sirupsen/logrus"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
)

const (
	kubeConfig = "InClusterConfig"
)
	// CreateClient function returns are kubernetes client */
func CreateClient() *kubernetes.Clientset {
	config, err := rest.InClusterConfig()
	logger.LogStdout()

	if err != nil {
		log.WithFields(log.Fields{
			"kubeconfig": kubeConfig,
		}).Error(err.Error())
	}

	log.WithFields(log.Fields{
		"kubeconfig": kubeConfig,
	}).Info("Successfully authenticated with K8s cluster")

	clientset, err := kubernetes.NewForConfig(config)

	if err != nil {
		log.WithFields(log.Fields{
			"kubeconfig": kubeConfig,
		}).Error(err.Error())
	}

	log.WithFields(log.Fields{
		"kubeconfig": kubeConfig,
	}).Info("Successfully created the K8s client")

	return clientset
}
