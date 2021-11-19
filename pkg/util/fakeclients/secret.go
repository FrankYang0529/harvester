package fakeclients

import (
	"context"

	ctlcorev1 "github.com/rancher/wrangler/pkg/generated/controllers/core/v1"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/apimachinery/pkg/types"
	"k8s.io/apimachinery/pkg/watch"

	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	corev1type "k8s.io/client-go/kubernetes/typed/core/v1"
)

type SecretClient func(namespace string) corev1type.SecretInterface

func (c SecretClient) Create(secret *v1.Secret) (*v1.Secret, error) {
	return c(secret.Namespace).Create(context.TODO(), secret, metav1.CreateOptions{})
}

func (c SecretClient) Update(secret *v1.Secret) (*v1.Secret, error) {
	return c(secret.Namespace).Update(context.TODO(), secret, metav1.UpdateOptions{})
}

func (c SecretClient) Delete(namespace, name string, options *metav1.DeleteOptions) error {
	panic("implement me")
}

func (c SecretClient) Get(namespace, name string, opts metav1.GetOptions) (*v1.Secret, error) {
	return c(namespace).Get(context.TODO(), name, opts)
}

func (c SecretClient) List(namespace string, opts metav1.ListOptions) (*v1.SecretList, error) {
	return c(namespace).List(context.TODO(), opts)
}

func (c SecretClient) Watch(namespace string, opts metav1.ListOptions) (watch.Interface, error) {
	panic("implement me")
}

func (c SecretClient) Patch(namespace, name string, pt types.PatchType, data []byte, subresources ...string) (result *v1.Secret, err error) {
	panic("implement me")
}

type SecretCache func(namespace string) corev1type.SecretInterface

func (c SecretCache) Get(namespace, name string) (*v1.Secret, error) {
	return c(namespace).Get(context.TODO(), name, metav1.GetOptions{})
}

func (c SecretCache) List(namespace string, selector labels.Selector) ([]*v1.Secret, error) {
	panic("implement me")
}

func (c SecretCache) AddIndexer(indexName string, indexer ctlcorev1.SecretIndexer) {
	panic("implement me")
}

func (c SecretCache) GetByIndex(indexName, key string) ([]*v1.Secret, error) {
	panic("implement me")
}
