// SPDX-License-Identifier: Apache-2.0
// Copyright Authors of Tetragon

// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	"context"

	v1alpha1 "github.com/cilium/tetragon/pkg/k8s/apis/cilium.io/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeTracingPoliciesNamespaced implements TracingPolicyNamespacedInterface
type FakeTracingPoliciesNamespaced struct {
	Fake *FakeCiliumV1alpha1
	ns   string
}

var tracingpoliciesnamespacedResource = schema.GroupVersionResource{Group: "cilium.io", Version: "v1alpha1", Resource: "tracingpoliciesnamespaced"}

var tracingpoliciesnamespacedKind = schema.GroupVersionKind{Group: "cilium.io", Version: "v1alpha1", Kind: "TracingPolicyNamespaced"}

// Get takes name of the tracingPolicyNamespaced, and returns the corresponding tracingPolicyNamespaced object, and an error if there is any.
func (c *FakeTracingPoliciesNamespaced) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.TracingPolicyNamespaced, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(tracingpoliciesnamespacedResource, c.ns, name), &v1alpha1.TracingPolicyNamespaced{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.TracingPolicyNamespaced), err
}

// List takes label and field selectors, and returns the list of TracingPoliciesNamespaced that match those selectors.
func (c *FakeTracingPoliciesNamespaced) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.TracingPolicyNamespacedList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(tracingpoliciesnamespacedResource, tracingpoliciesnamespacedKind, c.ns, opts), &v1alpha1.TracingPolicyNamespacedList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.TracingPolicyNamespacedList{ListMeta: obj.(*v1alpha1.TracingPolicyNamespacedList).ListMeta}
	for _, item := range obj.(*v1alpha1.TracingPolicyNamespacedList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested tracingPoliciesNamespaced.
func (c *FakeTracingPoliciesNamespaced) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(tracingpoliciesnamespacedResource, c.ns, opts))

}

// Create takes the representation of a tracingPolicyNamespaced and creates it.  Returns the server's representation of the tracingPolicyNamespaced, and an error, if there is any.
func (c *FakeTracingPoliciesNamespaced) Create(ctx context.Context, tracingPolicyNamespaced *v1alpha1.TracingPolicyNamespaced, opts v1.CreateOptions) (result *v1alpha1.TracingPolicyNamespaced, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(tracingpoliciesnamespacedResource, c.ns, tracingPolicyNamespaced), &v1alpha1.TracingPolicyNamespaced{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.TracingPolicyNamespaced), err
}

// Update takes the representation of a tracingPolicyNamespaced and updates it. Returns the server's representation of the tracingPolicyNamespaced, and an error, if there is any.
func (c *FakeTracingPoliciesNamespaced) Update(ctx context.Context, tracingPolicyNamespaced *v1alpha1.TracingPolicyNamespaced, opts v1.UpdateOptions) (result *v1alpha1.TracingPolicyNamespaced, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(tracingpoliciesnamespacedResource, c.ns, tracingPolicyNamespaced), &v1alpha1.TracingPolicyNamespaced{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.TracingPolicyNamespaced), err
}

// Delete takes name of the tracingPolicyNamespaced and deletes it. Returns an error if one occurs.
func (c *FakeTracingPoliciesNamespaced) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteActionWithOptions(tracingpoliciesnamespacedResource, c.ns, name, opts), &v1alpha1.TracingPolicyNamespaced{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeTracingPoliciesNamespaced) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(tracingpoliciesnamespacedResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.TracingPolicyNamespacedList{})
	return err
}

// Patch applies the patch and returns the patched tracingPolicyNamespaced.
func (c *FakeTracingPoliciesNamespaced) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.TracingPolicyNamespaced, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(tracingpoliciesnamespacedResource, c.ns, name, pt, data, subresources...), &v1alpha1.TracingPolicyNamespaced{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.TracingPolicyNamespaced), err
}
