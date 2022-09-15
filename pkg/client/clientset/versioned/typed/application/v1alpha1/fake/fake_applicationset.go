// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	"context"

	v1alpha1 "github.com/argoproj/argo-cd/v2/pkg/apis/application/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeApplicationSets implements ApplicationSetInterface
type FakeApplicationSets struct {
	Fake *FakeArgoprojV1alpha1
	ns   string
}

var applicationsetsResource = schema.GroupVersionResource{Group: "argoproj.io", Version: "v1alpha1", Resource: "applicationsets"}

var applicationsetsKind = schema.GroupVersionKind{Group: "argoproj.io", Version: "v1alpha1", Kind: "ApplicationSet"}

// Get takes name of the applicationSet, and returns the corresponding applicationSet object, and an error if there is any.
func (c *FakeApplicationSets) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.ApplicationSet, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(applicationsetsResource, c.ns, name), &v1alpha1.ApplicationSet{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ApplicationSet), err
}

// List takes label and field selectors, and returns the list of ApplicationSets that match those selectors.
func (c *FakeApplicationSets) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.ApplicationSetList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(applicationsetsResource, applicationsetsKind, c.ns, opts), &v1alpha1.ApplicationSetList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.ApplicationSetList{ListMeta: obj.(*v1alpha1.ApplicationSetList).ListMeta}
	for _, item := range obj.(*v1alpha1.ApplicationSetList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested applicationSets.
func (c *FakeApplicationSets) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(applicationsetsResource, c.ns, opts))

}

// Create takes the representation of a applicationSet and creates it.  Returns the server's representation of the applicationSet, and an error, if there is any.
func (c *FakeApplicationSets) Create(ctx context.Context, applicationSet *v1alpha1.ApplicationSet, opts v1.CreateOptions) (result *v1alpha1.ApplicationSet, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(applicationsetsResource, c.ns, applicationSet), &v1alpha1.ApplicationSet{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ApplicationSet), err
}

// Update takes the representation of a applicationSet and updates it. Returns the server's representation of the applicationSet, and an error, if there is any.
func (c *FakeApplicationSets) Update(ctx context.Context, applicationSet *v1alpha1.ApplicationSet, opts v1.UpdateOptions) (result *v1alpha1.ApplicationSet, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(applicationsetsResource, c.ns, applicationSet), &v1alpha1.ApplicationSet{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ApplicationSet), err
}

// Delete takes name of the applicationSet and deletes it. Returns an error if one occurs.
func (c *FakeApplicationSets) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteActionWithOptions(applicationsetsResource, c.ns, name, opts), &v1alpha1.ApplicationSet{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeApplicationSets) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(applicationsetsResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.ApplicationSetList{})
	return err
}

// Patch applies the patch and returns the patched applicationSet.
func (c *FakeApplicationSets) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.ApplicationSet, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(applicationsetsResource, c.ns, name, pt, data, subresources...), &v1alpha1.ApplicationSet{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ApplicationSet), err
}