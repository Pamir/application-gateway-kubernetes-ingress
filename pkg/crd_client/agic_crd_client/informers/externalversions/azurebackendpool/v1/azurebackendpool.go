/*
Copyright The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Code generated by informer-gen. DO NOT EDIT.

package v1

import (
	"context"
	time "time"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"

	azurebackendpoolv1 "github.com/Azure/application-gateway-kubernetes-ingress/pkg/apis/azurebackendpool/v1"
	versioned "github.com/Azure/application-gateway-kubernetes-ingress/pkg/crd_client/agic_crd_client/clientset/versioned"
	internalinterfaces "github.com/Azure/application-gateway-kubernetes-ingress/pkg/crd_client/agic_crd_client/informers/externalversions/internalinterfaces"
	v1 "github.com/Azure/application-gateway-kubernetes-ingress/pkg/crd_client/agic_crd_client/listers/azurebackendpool/v1"
)

// AzureBackendPoolInformer provides access to a shared informer and lister for
// AzureBackendPools.
type AzureBackendPoolInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() v1.AzureBackendPoolLister
}

type azureBackendPoolInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
}

// NewAzureBackendPoolInformer constructs a new informer for AzureBackendPool type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewAzureBackendPoolInformer(client versioned.Interface, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredAzureBackendPoolInformer(client, resyncPeriod, indexers, nil)
}

// NewFilteredAzureBackendPoolInformer constructs a new informer for AzureBackendPool type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredAzureBackendPoolInformer(client versioned.Interface, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options metav1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.AzurebackendpoolsV1().AzureBackendPools().List(context.TODO(), options)
			},
			WatchFunc: func(options metav1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.AzurebackendpoolsV1().AzureBackendPools().Watch(context.TODO(), options)
			},
		},
		&azurebackendpoolv1.AzureBackendPool{},
		resyncPeriod,
		indexers,
	)
}

func (f *azureBackendPoolInformer) defaultInformer(client versioned.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredAzureBackendPoolInformer(client, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *azureBackendPoolInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&azurebackendpoolv1.AzureBackendPool{}, f.defaultInformer)
}

func (f *azureBackendPoolInformer) Lister() v1.AzureBackendPoolLister {
	return v1.NewAzureBackendPoolLister(f.Informer().GetIndexer())
}
