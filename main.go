package main

import (
	"dynamic-pv-scaling/api"
	"dynamic-pv-scaling/logger"
	"dynamic-pv-scaling/pkg"
	"dynamic-pv-scaling/utils"
	"fmt"
	"github.com/jasonlvhit/gocron"
	log "github.com/sirupsen/logrus"
	"strconv"
	"time"
)

var (
	nameSpace           string
	scalePercentage     int
	thresholdPercentage int
	pvcName             string
)

func ResizeFunction() {
	infos := utils.GetConfigurations()
	logger.LogStdout()
	for _, info := range infos {
		nameSpace = fmt.Sprintf("%v", info["namespace"])
		scalePercentage, _ = strconv.Atoi(fmt.Sprintf("%v", info["scale_percentage"]))
		thresholdPercentage, _ = strconv.Atoi(fmt.Sprintf("%v", info["threshold_percentage"]))
		pvcName = fmt.Sprintf("%v", info["pvc_name"])
		if (100 - api.GetPersistentVolumeList(nameSpace, pvcName).Value) >= thresholdPercentage {
			updatedSize := utils.CalculateUpdatedSize(api.GetPeristentVolumeUsage(nameSpace, pvcName).Value, scalePercentage)
			pkg.ResizePersistentVolume(pvcName, nameSpace, updatedSize)
			for _, pod := range pkg.ListPods(nameSpace) {
				if pod.PersistentVolumeName == pvcName {
					time.Sleep(100 * time.Second)
					pkg.DeletePod(pod.PodName, nameSpace)
				}
			}
		} else {
			log.WithFields(log.Fields{
				"namespace":              nameSpace,
				"persistent_volume_name": pvcName,
			}).Info("Threshold is under control, no need of resizing")
		}
	}
}

func main() {
	scheduler := gocron.NewScheduler()
	scheduler.Every(100).Seconds().Do(ResizeFunction)
	<-scheduler.Start()
}
