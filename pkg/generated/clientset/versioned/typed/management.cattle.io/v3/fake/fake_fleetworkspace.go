/*
Copyright 2025 Rancher Labs, Inc.

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

// Code generated by main. DO NOT EDIT.

package fake

import (
	"context"

	v3 "github.com/rancher/rancher/pkg/apis/management.cattle.io/v3"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeFleetWorkspaces implements FleetWorkspaceInterface
type FakeFleetWorkspaces struct {
	Fake *FakeManagementV3
}

var fleetworkspacesResource = v3.SchemeGroupVersion.WithResource("fleetworkspaces")

var fleetworkspacesKind = v3.SchemeGroupVersion.WithKind("FleetWorkspace")

// Get takes name of the fleetWorkspace, and returns the corresponding fleetWorkspace object, and an error if there is any.
func (c *FakeFleetWorkspaces) Get(ctx context.Context, name string, options v1.GetOptions) (result *v3.FleetWorkspace, err error) {
	emptyResult := &v3.FleetWorkspace{}
	obj, err := c.Fake.
		Invokes(testing.NewRootGetActionWithOptions(fleetworkspacesResource, name, options), emptyResult)
	if obj == nil {
		return emptyResult, err
	}
	return obj.(*v3.FleetWorkspace), err
}

// List takes label and field selectors, and returns the list of FleetWorkspaces that match those selectors.
func (c *FakeFleetWorkspaces) List(ctx context.Context, opts v1.ListOptions) (result *v3.FleetWorkspaceList, err error) {
	emptyResult := &v3.FleetWorkspaceList{}
	obj, err := c.Fake.
		Invokes(testing.NewRootListActionWithOptions(fleetworkspacesResource, fleetworkspacesKind, opts), emptyResult)
	if obj == nil {
		return emptyResult, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v3.FleetWorkspaceList{ListMeta: obj.(*v3.FleetWorkspaceList).ListMeta}
	for _, item := range obj.(*v3.FleetWorkspaceList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested fleetWorkspaces.
func (c *FakeFleetWorkspaces) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewRootWatchActionWithOptions(fleetworkspacesResource, opts))
}

// Create takes the representation of a fleetWorkspace and creates it.  Returns the server's representation of the fleetWorkspace, and an error, if there is any.
func (c *FakeFleetWorkspaces) Create(ctx context.Context, fleetWorkspace *v3.FleetWorkspace, opts v1.CreateOptions) (result *v3.FleetWorkspace, err error) {
	emptyResult := &v3.FleetWorkspace{}
	obj, err := c.Fake.
		Invokes(testing.NewRootCreateActionWithOptions(fleetworkspacesResource, fleetWorkspace, opts), emptyResult)
	if obj == nil {
		return emptyResult, err
	}
	return obj.(*v3.FleetWorkspace), err
}

// Update takes the representation of a fleetWorkspace and updates it. Returns the server's representation of the fleetWorkspace, and an error, if there is any.
func (c *FakeFleetWorkspaces) Update(ctx context.Context, fleetWorkspace *v3.FleetWorkspace, opts v1.UpdateOptions) (result *v3.FleetWorkspace, err error) {
	emptyResult := &v3.FleetWorkspace{}
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateActionWithOptions(fleetworkspacesResource, fleetWorkspace, opts), emptyResult)
	if obj == nil {
		return emptyResult, err
	}
	return obj.(*v3.FleetWorkspace), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeFleetWorkspaces) UpdateStatus(ctx context.Context, fleetWorkspace *v3.FleetWorkspace, opts v1.UpdateOptions) (result *v3.FleetWorkspace, err error) {
	emptyResult := &v3.FleetWorkspace{}
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateSubresourceActionWithOptions(fleetworkspacesResource, "status", fleetWorkspace, opts), emptyResult)
	if obj == nil {
		return emptyResult, err
	}
	return obj.(*v3.FleetWorkspace), err
}

// Delete takes name of the fleetWorkspace and deletes it. Returns an error if one occurs.
func (c *FakeFleetWorkspaces) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewRootDeleteActionWithOptions(fleetworkspacesResource, name, opts), &v3.FleetWorkspace{})
	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeFleetWorkspaces) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewRootDeleteCollectionActionWithOptions(fleetworkspacesResource, opts, listOpts)

	_, err := c.Fake.Invokes(action, &v3.FleetWorkspaceList{})
	return err
}

// Patch applies the patch and returns the patched fleetWorkspace.
func (c *FakeFleetWorkspaces) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v3.FleetWorkspace, err error) {
	emptyResult := &v3.FleetWorkspace{}
	obj, err := c.Fake.
		Invokes(testing.NewRootPatchSubresourceActionWithOptions(fleetworkspacesResource, name, pt, data, opts, subresources...), emptyResult)
	if obj == nil {
		return emptyResult, err
	}
	return obj.(*v3.FleetWorkspace), err
}
