// SPDX-License-Identifier: AGPL-3.0-only

// Code generated by client-gen. DO NOT EDIT.

package v0alpha1

import (
	"context"
	json "encoding/json"
	"fmt"
	"time"

	v0alpha1 "github.com/grafana/grafana/pkg/apis/alerting_notifications/v0alpha1"
	alertingnotificationsv0alpha1 "github.com/grafana/grafana/pkg/generated/applyconfiguration/alerting_notifications/v0alpha1"
	scheme "github.com/grafana/grafana/pkg/generated/clientset/versioned/scheme"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// TimeIntervalsGetter has a method to return a TimeIntervalInterface.
// A group's client should implement this interface.
type TimeIntervalsGetter interface {
	TimeIntervals(namespace string) TimeIntervalInterface
}

// TimeIntervalInterface has methods to work with TimeInterval resources.
type TimeIntervalInterface interface {
	Create(ctx context.Context, timeInterval *v0alpha1.TimeInterval, opts v1.CreateOptions) (*v0alpha1.TimeInterval, error)
	Update(ctx context.Context, timeInterval *v0alpha1.TimeInterval, opts v1.UpdateOptions) (*v0alpha1.TimeInterval, error)
	Delete(ctx context.Context, name string, opts v1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error
	Get(ctx context.Context, name string, opts v1.GetOptions) (*v0alpha1.TimeInterval, error)
	List(ctx context.Context, opts v1.ListOptions) (*v0alpha1.TimeIntervalList, error)
	Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v0alpha1.TimeInterval, err error)
	Apply(ctx context.Context, timeInterval *alertingnotificationsv0alpha1.TimeIntervalApplyConfiguration, opts v1.ApplyOptions) (result *v0alpha1.TimeInterval, err error)
	TimeIntervalExpansion
}

// timeIntervals implements TimeIntervalInterface
type timeIntervals struct {
	client rest.Interface
	ns     string
}

// newTimeIntervals returns a TimeIntervals
func newTimeIntervals(c *NotificationsV0alpha1Client, namespace string) *timeIntervals {
	return &timeIntervals{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the timeInterval, and returns the corresponding timeInterval object, and an error if there is any.
func (c *timeIntervals) Get(ctx context.Context, name string, options v1.GetOptions) (result *v0alpha1.TimeInterval, err error) {
	result = &v0alpha1.TimeInterval{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("timeintervals").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do(ctx).
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of TimeIntervals that match those selectors.
func (c *timeIntervals) List(ctx context.Context, opts v1.ListOptions) (result *v0alpha1.TimeIntervalList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v0alpha1.TimeIntervalList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("timeintervals").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do(ctx).
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested timeIntervals.
func (c *timeIntervals) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("timeintervals").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch(ctx)
}

// Create takes the representation of a timeInterval and creates it.  Returns the server's representation of the timeInterval, and an error, if there is any.
func (c *timeIntervals) Create(ctx context.Context, timeInterval *v0alpha1.TimeInterval, opts v1.CreateOptions) (result *v0alpha1.TimeInterval, err error) {
	result = &v0alpha1.TimeInterval{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("timeintervals").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(timeInterval).
		Do(ctx).
		Into(result)
	return
}

// Update takes the representation of a timeInterval and updates it. Returns the server's representation of the timeInterval, and an error, if there is any.
func (c *timeIntervals) Update(ctx context.Context, timeInterval *v0alpha1.TimeInterval, opts v1.UpdateOptions) (result *v0alpha1.TimeInterval, err error) {
	result = &v0alpha1.TimeInterval{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("timeintervals").
		Name(timeInterval.Name).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(timeInterval).
		Do(ctx).
		Into(result)
	return
}

// Delete takes name of the timeInterval and deletes it. Returns an error if one occurs.
func (c *timeIntervals) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("timeintervals").
		Name(name).
		Body(&opts).
		Do(ctx).
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *timeIntervals) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	var timeout time.Duration
	if listOpts.TimeoutSeconds != nil {
		timeout = time.Duration(*listOpts.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("timeintervals").
		VersionedParams(&listOpts, scheme.ParameterCodec).
		Timeout(timeout).
		Body(&opts).
		Do(ctx).
		Error()
}

// Patch applies the patch and returns the patched timeInterval.
func (c *timeIntervals) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v0alpha1.TimeInterval, err error) {
	result = &v0alpha1.TimeInterval{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("timeintervals").
		Name(name).
		SubResource(subresources...).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}

// Apply takes the given apply declarative configuration, applies it and returns the applied timeInterval.
func (c *timeIntervals) Apply(ctx context.Context, timeInterval *alertingnotificationsv0alpha1.TimeIntervalApplyConfiguration, opts v1.ApplyOptions) (result *v0alpha1.TimeInterval, err error) {
	if timeInterval == nil {
		return nil, fmt.Errorf("timeInterval provided to Apply must not be nil")
	}
	patchOpts := opts.ToPatchOptions()
	data, err := json.Marshal(timeInterval)
	if err != nil {
		return nil, err
	}
	name := timeInterval.Name
	if name == nil {
		return nil, fmt.Errorf("timeInterval.Name must be provided to Apply")
	}
	result = &v0alpha1.TimeInterval{}
	err = c.client.Patch(types.ApplyPatchType).
		Namespace(c.ns).
		Resource("timeintervals").
		Name(*name).
		VersionedParams(&patchOpts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}