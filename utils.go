package watch

import (
	"encoding/json"
	"k8s.io/klog/v2"
)
var (
	PodLastSyncVerison = "0"
	DeploymentLastSyncVersion = "0"
	NamespaceLastSyncVersion = "0"
)

func StrOutJSON(obj interface{}) string{
	bytes, err := json.Marshal(obj)
	if err != nil {
		klog.Info(err.Error())
	}
	return string(bytes)
}
