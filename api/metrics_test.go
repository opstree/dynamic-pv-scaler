package api

import (
	"net/http"
	"reflect"
	"testing"
)

func TestGetPersistentVolumeList(t *testing.T) {
	type args struct {
		nameSpace            string
		persistentVolumeName string
	}
	tests := []struct {
		name string
		args args
		want PersistentVolumeList
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetPersistentVolumeList(tt.args.nameSpace, tt.args.persistentVolumeName); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetPersistentVolumeList() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetPeristentVolumeUsage(t *testing.T) {
	type args struct {
		nameSpace            string
		persistentVolumeName string
	}
	tests := []struct {
		name string
		args args
		want PersistentVolumeUsage
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetPeristentVolumeUsage(tt.args.nameSpace, tt.args.persistentVolumeName); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetPeristentVolumeUsage() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetVolumeListQueryResponse(t *testing.T) {
	type args struct {
		nameSpace            string
		persistentVolumeName string
	}
	tests := []struct {
		name string
		args args
		want *http.Response
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetVolumeListQueryResponse(tt.args.nameSpace, tt.args.persistentVolumeName); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetVolumeListQueryResponse() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetVolumeUsageQueryResponse(t *testing.T) {
	type args struct {
		nameSpace            string
		persistentVolumeName string
	}
	tests := []struct {
		name string
		args args
		want *http.Response
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetVolumeUsageQueryResponse(tt.args.nameSpace, tt.args.persistentVolumeName); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetVolumeUsageQueryResponse() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGenerateVolumeListQuery(t *testing.T) {
	type args struct {
		nameSpace            string
		persistentVolumeName string
	}
	tests := []struct {
		name string
		args args
		want *http.Request
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GenerateVolumeListQuery(tt.args.nameSpace, tt.args.persistentVolumeName); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GenerateVolumeListQuery() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGenerateVolumeUsageQuery(t *testing.T) {
	type args struct {
		nameSpace            string
		persistentVolumeName string
	}
	tests := []struct {
		name string
		args args
		want *http.Request
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GenerateVolumeUsageQuery(tt.args.nameSpace, tt.args.persistentVolumeName); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GenerateVolumeUsageQuery() = %v, want %v", got, tt.want)
			}
		})
	}
}
