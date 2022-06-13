package watch

import (
	corev1 "k8s.io/api/core/v1"
	"k8s.io/client-go/tools/cache"
	"k8s.io/klog"
)

func PodAdd(obj interface{}) {
	// Cast the obj as node
	pod := obj.(*corev1.Pod)
	if pod.ResourceVersion <= PodLastSyncVerison {
		return
	}
	key, err := cache.MetaNamespaceKeyFunc(obj)
	if err != nil {
		klog.Info(err.Error())
		return
	}
	klog.Infof("pod %s created, resourceVersion:%s", key, pod.ResourceVersion)
}

func PodUpdate(old, new interface{})  {
	pod := new.(*corev1.Pod)
	if pod.ResourceVersion <= PodLastSyncVerison {
		return
	}
	key, err := cache.MetaNamespaceKeyFunc(new)
	if err != nil {
		klog.Info(err.Error())
		return
	}
	klog.Infof("pod %s updated, resourceVersion:%s", key, pod.ResourceVersion)
}

func PodDelete(obj interface{})  {
	pod := obj.(*corev1.Pod)
	if pod.ResourceVersion <= PodLastSyncVerison {
		return
	}
	key, err := cache.DeletionHandlingMetaNamespaceKeyFunc(obj)
	if err != nil {
		klog.Info(err.Error())
		return
	}
	klog.Info("pod ", key, " deleted!")
}
