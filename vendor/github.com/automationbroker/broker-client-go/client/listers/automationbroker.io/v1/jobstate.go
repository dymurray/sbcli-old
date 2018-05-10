/*
Copyright (c) 2018 Red Hat, Inc.

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

// This file was automatically generated by lister-gen

package v1

import (
	v1 "github.com/automationbroker/broker-client-go/pkg/apis/automationbroker.io/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// JobStateLister helps list JobStates.
type JobStateLister interface {
	// List lists all JobStates in the indexer.
	List(selector labels.Selector) (ret []*v1.JobState, err error)
	// JobStates returns an object that can list and get JobStates.
	JobStates(namespace string) JobStateNamespaceLister
	JobStateListerExpansion
}

// jobStateLister implements the JobStateLister interface.
type jobStateLister struct {
	indexer cache.Indexer
}

// NewJobStateLister returns a new JobStateLister.
func NewJobStateLister(indexer cache.Indexer) JobStateLister {
	return &jobStateLister{indexer: indexer}
}

// List lists all JobStates in the indexer.
func (s *jobStateLister) List(selector labels.Selector) (ret []*v1.JobState, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1.JobState))
	})
	return ret, err
}

// JobStates returns an object that can list and get JobStates.
func (s *jobStateLister) JobStates(namespace string) JobStateNamespaceLister {
	return jobStateNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// JobStateNamespaceLister helps list and get JobStates.
type JobStateNamespaceLister interface {
	// List lists all JobStates in the indexer for a given namespace.
	List(selector labels.Selector) (ret []*v1.JobState, err error)
	// Get retrieves the JobState from the indexer for a given namespace and name.
	Get(name string) (*v1.JobState, error)
	JobStateNamespaceListerExpansion
}

// jobStateNamespaceLister implements the JobStateNamespaceLister
// interface.
type jobStateNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all JobStates in the indexer for a given namespace.
func (s jobStateNamespaceLister) List(selector labels.Selector) (ret []*v1.JobState, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1.JobState))
	})
	return ret, err
}

// Get retrieves the JobState from the indexer for a given namespace and name.
func (s jobStateNamespaceLister) Get(name string) (*v1.JobState, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1.Resource("jobstate"), name)
	}
	return obj.(*v1.JobState), nil
}
