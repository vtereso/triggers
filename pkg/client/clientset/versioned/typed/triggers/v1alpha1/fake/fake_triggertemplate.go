/*
Copyright 2019 The Tekton Authors

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
	v1alpha1 "github.com/tektoncd/triggers/pkg/apis/triggers/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeTriggerTemplates implements TriggerTemplateInterface
type FakeTriggerTemplates struct {
	Fake *FakeTriggersV1alpha1
	ns   string
}

var triggertemplatesResource = schema.GroupVersionResource{Group: "triggers", Version: "v1alpha1", Resource: "triggertemplates"}

var triggertemplatesKind = schema.GroupVersionKind{Group: "triggers", Version: "v1alpha1", Kind: "TriggerTemplate"}

// Get takes name of the triggerTemplate, and returns the corresponding triggerTemplate object, and an error if there is any.
func (c *FakeTriggerTemplates) Get(name string, options v1.GetOptions) (result *v1alpha1.TriggerTemplate, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(triggertemplatesResource, c.ns, name), &v1alpha1.TriggerTemplate{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.TriggerTemplate), err
}

// List takes label and field selectors, and returns the list of TriggerTemplates that match those selectors.
func (c *FakeTriggerTemplates) List(opts v1.ListOptions) (result *v1alpha1.TriggerTemplateList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(triggertemplatesResource, triggertemplatesKind, c.ns, opts), &v1alpha1.TriggerTemplateList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.TriggerTemplateList{ListMeta: obj.(*v1alpha1.TriggerTemplateList).ListMeta}
	for _, item := range obj.(*v1alpha1.TriggerTemplateList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested triggerTemplates.
func (c *FakeTriggerTemplates) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(triggertemplatesResource, c.ns, opts))

}

// Create takes the representation of a triggerTemplate and creates it.  Returns the server's representation of the triggerTemplate, and an error, if there is any.
func (c *FakeTriggerTemplates) Create(triggerTemplate *v1alpha1.TriggerTemplate) (result *v1alpha1.TriggerTemplate, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(triggertemplatesResource, c.ns, triggerTemplate), &v1alpha1.TriggerTemplate{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.TriggerTemplate), err
}

// Update takes the representation of a triggerTemplate and updates it. Returns the server's representation of the triggerTemplate, and an error, if there is any.
func (c *FakeTriggerTemplates) Update(triggerTemplate *v1alpha1.TriggerTemplate) (result *v1alpha1.TriggerTemplate, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(triggertemplatesResource, c.ns, triggerTemplate), &v1alpha1.TriggerTemplate{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.TriggerTemplate), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeTriggerTemplates) UpdateStatus(triggerTemplate *v1alpha1.TriggerTemplate) (*v1alpha1.TriggerTemplate, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(triggertemplatesResource, "status", c.ns, triggerTemplate), &v1alpha1.TriggerTemplate{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.TriggerTemplate), err
}

// Delete takes name of the triggerTemplate and deletes it. Returns an error if one occurs.
func (c *FakeTriggerTemplates) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(triggertemplatesResource, c.ns, name), &v1alpha1.TriggerTemplate{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeTriggerTemplates) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(triggertemplatesResource, c.ns, listOptions)

	_, err := c.Fake.Invokes(action, &v1alpha1.TriggerTemplateList{})
	return err
}

// Patch applies the patch and returns the patched triggerTemplate.
func (c *FakeTriggerTemplates) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.TriggerTemplate, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(triggertemplatesResource, c.ns, name, data, subresources...), &v1alpha1.TriggerTemplate{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.TriggerTemplate), err
}
