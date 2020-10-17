package pkg

import (
	"dynamic-pv-scaling/api"
	"dynamic-pv-scaling/logger"
	log "github.com/sirupsen/logrus"
	"k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/resource"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"strconv"
)

// PodList defines struct for json key values as mentioned below.
type PodList struct {
	PodName              string `json:"pod_name"`
	PersistentVolumeName string `json:"persistent_volume_name"`
}

// ListPods function takes namespace as input and returns PodList with list of all the pods running in the given namespace */
func ListPods(namespace string) []PodList {
	var podLists []PodList
	var podInformation PodList
	logger.LogStdout()

	clientset := api.CreateClient()
	pods, err := clientset.CoreV1().Pods(namespace).List(metav1.ListOptions{})
	if err != nil {
		log.WithFields(log.Fields{
			"namespace": namespace,
		}).Error(err.Error())
	}
	for _, pod := range pods.Items {
		for _, volume := range pod.Spec.Volumes {
			if volume.PersistentVolumeClaim != nil {
				podInformation = PodList{
					PodName:              pod.GetName(),
					PersistentVolumeName: volume.PersistentVolumeClaim.ClaimName,
				}
				podLists = append(podLists, podInformation)
			}
		}
	}

	log.WithFields(log.Fields{
		"namespace": namespace,
	}).Info("Successfully listed pods with persistent volumes")

	return podLists
}

// DeletePod function takes podName and nameSpace as input and deleted the said pod in the given namespace */
func DeletePod(podName string, nameSpace string) {
	logger.LogStdout()

	clientset := api.CreateClient()
	deletePolicy := metav1.DeletePropagationForeground

	podClient := clientset.CoreV1().Pods(nameSpace)

	err := podClient.Delete(podName, &metav1.DeleteOptions{PropagationPolicy: &deletePolicy})

	if err != nil {
		log.WithFields(log.Fields{
			"pod": podName,
		}).Error(err.Error())
	}

	log.WithFields(log.Fields{
		"pod": podName,
	}).Info("Successfully deleted the pod to resize persistent volume")
}

// ResizePersistentVolume takes pvcName to increase size of  in nameSpace given  and with the value requred  */
func ResizePersistentVolume(pvcName string, nameSpace string, value int) {
	logger.LogStdout()

	clientset := api.CreateClient()
	rawPersistentVolumeClaim, err := clientset.CoreV1().PersistentVolumeClaims(nameSpace).Get(pvcName, metav1.GetOptions{})
	if err != nil {
		log.WithFields(log.Fields{
			"pvc": pvcName,
		}).Error(err.Error())
	}

	updateValue := strconv.Itoa(value) + "Gi"
	rawPersistentVolumeClaim.Spec.Resources.Requests[v1.ResourceStorage], err = resource.ParseQuantity(updateValue)
	if err != nil {
		log.WithFields(log.Fields{
			"pvc": pvcName,
		}).Error(err.Error())
	}
	_, err = clientset.CoreV1().PersistentVolumeClaims(nameSpace).Update(rawPersistentVolumeClaim)
	if err != nil {
		log.WithFields(log.Fields{
			"pvc": pvcName,
		}).Error(err.Error())
	}

	log.WithFields(log.Fields{
		"pvc": pvcName,
	}).Info("Successfully resized the pvc, the new size is " + updateValue)
}
