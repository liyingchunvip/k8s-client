package watch

import (
	v1 "k8s.io/api/apps/v1"
	"k8s.io/client-go/tools/cache"
	"k8s.io/klog"
)

func DeploymentAdd(obj interface{}) {
	// Cast the obj as node
	dp := obj.(*v1.Deployment)
	if dp.ResourceVersion <= DeploymentLastSyncVersion {
		return
	}
	key, err := cache.MetaNamespaceKeyFunc(obj)
	if err != nil {
		klog.Info(err.Error())
		return
	}
	klog.Infof("Deployment %s created, resourceVersion: %s", key, dp.ResourceVersion)
}

func DeploymentUpdate(old, new interface{}) {
	dp := new.(*v1.Deployment)
	if dp.ResourceVersion <= DeploymentLastSyncVersion {
		return
	}
	key, err := cache.MetaNamespaceKeyFunc(new)
	if err != nil {
		klog.Info(err.Error())
		return
	}
	klog.Infof("Deployment %s updated, resourceVersion: %s", key, dp.ResourceVersion)
	oldDeployment := old.(*v1.Deployment)
	newDeployment := new.(*v1.Deployment)
	LogChange(oldDeployment, newDeployment)
}

func DeploymentDelete(obj interface{}) {
	dp := obj.(*v1.Deployment)
	if dp.ResourceVersion <= DeploymentLastSyncVersion {
		return
	}
	key, err := cache.DeletionHandlingMetaNamespaceKeyFunc(obj)
	if err != nil {
		klog.Info(err.Error())
		return
	}
	klog.Info("Deployment ", key, " deleted!")
}

func LogChange(old, new *v1.Deployment) {
	if nil == old || nil == new {
		return
	}
	oldLimit := old.DeepCopy().Spec.Template.Spec.Containers[0].Resources.Limits
	oldRequest := old.DeepCopy().Spec.Template.Spec.Containers[0].Resources.Requests
	newLimit := new.DeepCopy().Spec.Template.Spec.Containers[0].Resources.Limits
	newRequest := new.DeepCopy().Spec.Template.Spec.Containers[0].Resources.Requests

	if oldLimit != nil && newLimit != nil {
		if oldLimit.Cpu().MilliValue() != newLimit.Cpu().MilliValue() {
			klog.Infof("cpu limit:%dm change to %dm\n", oldLimit.Cpu().MilliValue(), newLimit.Cpu().MilliValue())
		}
		if oldLimit.Memory().MilliValue() != newLimit.Memory().MilliValue() {
			klog.Infof("memory limit:%d change to %d\n", oldLimit.Memory().MilliValue(), newLimit.Memory().MilliValue())
		}

	}
	if oldRequest != nil && newRequest != nil {
		if oldRequest.Cpu().MilliValue() != newRequest.Cpu().MilliValue() {
			klog.Infof("cpu request:%dm change to %dm\n", oldRequest.Cpu().MilliValue(), newRequest.Cpu().MilliValue())
		}
		if oldRequest.Memory().MilliValue() != newRequest.Memory().MilliValue() {
			klog.Infof("memory request:%d change to %d\n", oldRequest.Memory().MilliValue(), newRequest.Memory().MilliValue())
		}

	}
}
