// Code generated by helmet-generate. DO NOT EDIT.

package v1

import (
	"github.com/onosproject/helmet/pkg/kubernetes/resource"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"time"
)

var ConfigMapKind = resource.Kind{
	Group:   "core",
	Version: "v1",
	Kind:    "ConfigMap",
	Scoped:  true,
}

var ConfigMapResource = resource.Type{
	Kind: ConfigMapKind,
	Name: "configmaps",
}

func NewConfigMap(configMap *corev1.ConfigMap, client resource.Client) *ConfigMap {
	return &ConfigMap{
		Resource: resource.NewResource(configMap.ObjectMeta, ConfigMapKind, client),
		Object:   configMap,
	}
}

type ConfigMap struct {
	*resource.Resource
	Object *corev1.ConfigMap
}

func (r *ConfigMap) Delete() error {
	return r.Clientset().
		CoreV1().
		RESTClient().
		Delete().
		NamespaceIfScoped(r.Namespace, ConfigMapKind.Scoped).
		Resource(ConfigMapResource.Name).
		Name(r.Name).
		VersionedParams(&metav1.DeleteOptions{}, metav1.ParameterCodec).
		Timeout(time.Minute).
		Do().
		Error()
}