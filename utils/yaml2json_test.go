package utils

import (
	"reflect"
	"testing"
)

func Test_yamlToJson(t *testing.T) {
	type args struct {
		filename string
	}
	tests := []struct {
		name string
		args args
		want []map[string]interface{}
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := yamlToJson(tt.args.filename); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("yamlToJson() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_checkError(t *testing.T) {
	type args struct {
		err error
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := checkError(tt.args.err); got != tt.want {
				t.Errorf("checkError() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetConfigurations(t *testing.T) {
	tests := []struct {
		name string
		want []map[string]interface{}
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetConfigurations(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetConfigurations() = %v, want %v", got, tt.want)
			}
		})
	}
}
