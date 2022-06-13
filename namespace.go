package watch

import (
	corev1 "k8s.io/api/core/v1"
	"k8s.io/client-go/tools/cache"
	"k8s.io/klog"
)

func NamespaceAdd(obj interface{}) {
	ns := obj.(*corev1.Namespace)
	if ns.ResourceVersion <= NamespaceLastSyncVersion {
		return
	}
	key, err := cache.MetaNamespaceKeyFunc(obj)
	if err != nil {
		klog.Info(err.Error())
		return
	}
	klog.Infof("Namespace %s created, resourceVersion: %s", key, ns.ResourceVersion)
}

func NamespaceUpdate(old, new interface{})  {
	ns := new.(*corev1.Namespace)
	if ns.ResourceVersion <= NamespaceLastSyncVersion {
		return
	}
	key, err := cache.MetaNamespaceKeyFunc(new)
	if err != nil {
		klog.Infof(err.Error())
		return
	}
	klog.Infof("Namespace %s updated, resourceVersion: %s", key, ns.ResourceVersion)
}

func NamespaceDelete(obj interface{})  {
	ns := obj.(*corev1.Namespace)
	if ns.ResourceVersion <= NamespaceLastSyncVersion {
		return
	}
	key, err := cache.DeletionHandlingMetaNamespaceKeyFunc(obj)
	if err != nil {
		klog.Info(err.Error())
		return
	}
	klog.Info("namespace ", key, " deleted!")
}
