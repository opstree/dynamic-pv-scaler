package api

import (
	"reflect"
	"testing"

	"k8s.io/client-go/kubernetes"
)

func TestCreateClient(t *testing.T) {
	tests := []struct {
		name string
		want *kubernetes.Clientset
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CreateClient(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CreateClient() = %v, want %v", got, tt.want)
			}
		})
	}
}
