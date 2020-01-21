/* Copyright 2020 The FedLearner Authors. All Rights Reserved.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *    http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	v1alpha1 "github.com/bytedance/fedlearner/deploy/kubernetes_operator/pkg/apis/fedlearner.k8s.io/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeFLApps implements FLAppInterface
type FakeFLApps struct {
	Fake *FakeFedlearnerV1alpha1
	ns   string
}

var flappsResource = schema.GroupVersionResource{Group: "fedlearner.k8s.io", Version: "v1alpha1", Resource: "flapps"}

var flappsKind = schema.GroupVersionKind{Group: "fedlearner.k8s.io", Version: "v1alpha1", Kind: "FLApp"}

// Get takes name of the fLApp, and returns the corresponding fLApp object, and an error if there is any.
func (c *FakeFLApps) Get(name string, options v1.GetOptions) (result *v1alpha1.FLApp, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(flappsResource, c.ns, name), &v1alpha1.FLApp{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.FLApp), err
}

// List takes label and field selectors, and returns the list of FLApps that match those selectors.
func (c *FakeFLApps) List(opts v1.ListOptions) (result *v1alpha1.FLAppList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(flappsResource, flappsKind, c.ns, opts), &v1alpha1.FLAppList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.FLAppList{ListMeta: obj.(*v1alpha1.FLAppList).ListMeta}
	for _, item := range obj.(*v1alpha1.FLAppList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested fLApps.
func (c *FakeFLApps) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(flappsResource, c.ns, opts))

}

// Create takes the representation of a fLApp and creates it.  Returns the server's representation of the fLApp, and an error, if there is any.
func (c *FakeFLApps) Create(fLApp *v1alpha1.FLApp) (result *v1alpha1.FLApp, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(flappsResource, c.ns, fLApp), &v1alpha1.FLApp{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.FLApp), err
}

// Update takes the representation of a fLApp and updates it. Returns the server's representation of the fLApp, and an error, if there is any.
func (c *FakeFLApps) Update(fLApp *v1alpha1.FLApp) (result *v1alpha1.FLApp, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(flappsResource, c.ns, fLApp), &v1alpha1.FLApp{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.FLApp), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeFLApps) UpdateStatus(fLApp *v1alpha1.FLApp) (*v1alpha1.FLApp, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(flappsResource, "status", c.ns, fLApp), &v1alpha1.FLApp{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.FLApp), err
}

// Delete takes name of the fLApp and deletes it. Returns an error if one occurs.
func (c *FakeFLApps) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(flappsResource, c.ns, name), &v1alpha1.FLApp{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeFLApps) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(flappsResource, c.ns, listOptions)

	_, err := c.Fake.Invokes(action, &v1alpha1.FLAppList{})
	return err
}

// Patch applies the patch and returns the patched fLApp.
func (c *FakeFLApps) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.FLApp, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(flappsResource, c.ns, name, pt, data, subresources...), &v1alpha1.FLApp{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.FLApp), err
}