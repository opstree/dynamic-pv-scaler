package pkg

import (
	"reflect"
	"testing"
)

func TestListPods(t *testing.T) {
	type args struct {
		namespace string
	}
	tests := []struct {
		name string
		args args
		want []PodList
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ListPods(tt.args.namespace); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ListPods() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDeletePod(t *testing.T) {
	type args struct {
		podName   string
		nameSpace string
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			DeletePod(tt.args.podName, tt.args.nameSpace)
		})
	}
}

func TestResizePersistentVolume(t *testing.T) {
	type args struct {
		pvcName   string
		nameSpace string
		value     int
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ResizePersistentVolume(tt.args.pvcName, tt.args.nameSpace, tt.args.value)
		})
	}
}
