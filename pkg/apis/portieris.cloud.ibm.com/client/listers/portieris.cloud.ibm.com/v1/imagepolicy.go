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

// Code generated by lister-gen. DO NOT EDIT.

package v1

import (
	v1 "github.com/IBM/portieris/pkg/apis/portieris.cloud.ibm.com/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// ImagePolicyLister helps list ImagePolicies.
// All objects returned here must be treated as read-only.
type ImagePolicyLister interface {
	// List lists all ImagePolicies in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1.ImagePolicy, err error)
	// ImagePolicies returns an object that can list and get ImagePolicies.
	ImagePolicies(namespace string) ImagePolicyNamespaceLister
	ImagePolicyListerExpansion
}

// imagePolicyLister implements the ImagePolicyLister interface.
type imagePolicyLister struct {
	indexer cache.Indexer
}

// NewImagePolicyLister returns a new ImagePolicyLister.
func NewImagePolicyLister(indexer cache.Indexer) ImagePolicyLister {
	return &imagePolicyLister{indexer: indexer}
}

// List lists all ImagePolicies in the indexer.
func (s *imagePolicyLister) List(selector labels.Selector) (ret []*v1.ImagePolicy, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1.ImagePolicy))
	})
	return ret, err
}

// ImagePolicies returns an object that can list and get ImagePolicies.
func (s *imagePolicyLister) ImagePolicies(namespace string) ImagePolicyNamespaceLister {
	return imagePolicyNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// ImagePolicyNamespaceLister helps list and get ImagePolicies.
// All objects returned here must be treated as read-only.
type ImagePolicyNamespaceLister interface {
	// List lists all ImagePolicies in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1.ImagePolicy, err error)
	// Get retrieves the ImagePolicy from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1.ImagePolicy, error)
	ImagePolicyNamespaceListerExpansion
}

// imagePolicyNamespaceLister implements the ImagePolicyNamespaceLister
// interface.
type imagePolicyNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all ImagePolicies in the indexer for a given namespace.
func (s imagePolicyNamespaceLister) List(selector labels.Selector) (ret []*v1.ImagePolicy, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1.ImagePolicy))
	})
	return ret, err
}

// Get retrieves the ImagePolicy from the indexer for a given namespace and name.
func (s imagePolicyNamespaceLister) Get(name string) (*v1.ImagePolicy, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1.Resource("imagepolicy"), name)
	}
	return obj.(*v1.ImagePolicy), nil
}
