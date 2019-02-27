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

// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	v1alpha1 "github.com/kudobuilder/kudo/pkg/apis/kudo/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakePlanExecutions implements PlanExecutionInterface
type FakePlanExecutions struct {
	Fake *FakeKudoV1alpha1
	ns   string
}

var planexecutionsResource = schema.GroupVersionResource{Group: "kudo.k8s.io", Version: "v1alpha1", Resource: "planexecutions"}

var planexecutionsKind = schema.GroupVersionKind{Group: "kudo.k8s.io", Version: "v1alpha1", Kind: "PlanExecution"}

// Get takes name of the planExecution, and returns the corresponding planExecution object, and an error if there is any.
func (c *FakePlanExecutions) Get(name string, options v1.GetOptions) (result *v1alpha1.PlanExecution, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(planexecutionsResource, c.ns, name), &v1alpha1.PlanExecution{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.PlanExecution), err
}

// List takes label and field selectors, and returns the list of PlanExecutions that match those selectors.
func (c *FakePlanExecutions) List(opts v1.ListOptions) (result *v1alpha1.PlanExecutionList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(planexecutionsResource, planexecutionsKind, c.ns, opts), &v1alpha1.PlanExecutionList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.PlanExecutionList{ListMeta: obj.(*v1alpha1.PlanExecutionList).ListMeta}
	for _, item := range obj.(*v1alpha1.PlanExecutionList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested planExecutions.
func (c *FakePlanExecutions) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(planexecutionsResource, c.ns, opts))

}

// Create takes the representation of a planExecution and creates it.  Returns the server's representation of the planExecution, and an error, if there is any.
func (c *FakePlanExecutions) Create(planExecution *v1alpha1.PlanExecution) (result *v1alpha1.PlanExecution, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(planexecutionsResource, c.ns, planExecution), &v1alpha1.PlanExecution{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.PlanExecution), err
}

// Update takes the representation of a planExecution and updates it. Returns the server's representation of the planExecution, and an error, if there is any.
func (c *FakePlanExecutions) Update(planExecution *v1alpha1.PlanExecution) (result *v1alpha1.PlanExecution, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(planexecutionsResource, c.ns, planExecution), &v1alpha1.PlanExecution{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.PlanExecution), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakePlanExecutions) UpdateStatus(planExecution *v1alpha1.PlanExecution) (*v1alpha1.PlanExecution, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(planexecutionsResource, "status", c.ns, planExecution), &v1alpha1.PlanExecution{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.PlanExecution), err
}

// Delete takes name of the planExecution and deletes it. Returns an error if one occurs.
func (c *FakePlanExecutions) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(planexecutionsResource, c.ns, name), &v1alpha1.PlanExecution{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakePlanExecutions) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(planexecutionsResource, c.ns, listOptions)

	_, err := c.Fake.Invokes(action, &v1alpha1.PlanExecutionList{})
	return err
}

// Patch applies the patch and returns the patched planExecution.
func (c *FakePlanExecutions) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.PlanExecution, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(planexecutionsResource, c.ns, name, pt, data, subresources...), &v1alpha1.PlanExecution{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.PlanExecution), err
}
