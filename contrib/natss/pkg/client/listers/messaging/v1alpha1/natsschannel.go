/*
Copyright 2019 The Knative Authors

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

// Code generated by lister-gen. DO NOT EDIT.

package v1alpha1

import (
	v1alpha1 "github.com/knative/eventing/contrib/natss/pkg/apis/messaging/v1alpha1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// NatssChannelLister helps list NatssChannels.
type NatssChannelLister interface {
	// List lists all NatssChannels in the indexer.
	List(selector labels.Selector) (ret []*v1alpha1.NatssChannel, err error)
	// NatssChannels returns an object that can list and get NatssChannels.
	NatssChannels(namespace string) NatssChannelNamespaceLister
	NatssChannelListerExpansion
}

// natssChannelLister implements the NatssChannelLister interface.
type natssChannelLister struct {
	indexer cache.Indexer
}

// NewNatssChannelLister returns a new NatssChannelLister.
func NewNatssChannelLister(indexer cache.Indexer) NatssChannelLister {
	return &natssChannelLister{indexer: indexer}
}

// List lists all NatssChannels in the indexer.
func (s *natssChannelLister) List(selector labels.Selector) (ret []*v1alpha1.NatssChannel, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.NatssChannel))
	})
	return ret, err
}

// NatssChannels returns an object that can list and get NatssChannels.
func (s *natssChannelLister) NatssChannels(namespace string) NatssChannelNamespaceLister {
	return natssChannelNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// NatssChannelNamespaceLister helps list and get NatssChannels.
type NatssChannelNamespaceLister interface {
	// List lists all NatssChannels in the indexer for a given namespace.
	List(selector labels.Selector) (ret []*v1alpha1.NatssChannel, err error)
	// Get retrieves the NatssChannel from the indexer for a given namespace and name.
	Get(name string) (*v1alpha1.NatssChannel, error)
	NatssChannelNamespaceListerExpansion
}

// natssChannelNamespaceLister implements the NatssChannelNamespaceLister
// interface.
type natssChannelNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all NatssChannels in the indexer for a given namespace.
func (s natssChannelNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.NatssChannel, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.NatssChannel))
	})
	return ret, err
}

// Get retrieves the NatssChannel from the indexer for a given namespace and name.
func (s natssChannelNamespaceLister) Get(name string) (*v1alpha1.NatssChannel, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("natsschannel"), name)
	}
	return obj.(*v1alpha1.NatssChannel), nil
}
