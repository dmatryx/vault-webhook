// Code generated by client-gen. DO NOT EDIT.

package v1alpha1

import (
	v1alpha1 "github.com/uswitch/vault-webhook/pkg/apis/vaultwebhook.uswitch.com/v1alpha1"
	scheme "github.com/uswitch/vault-webhook/pkg/client/clientset/versioned/scheme"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// DatabaseCredentialBindingsGetter has a method to return a DatabaseCredentialBindingInterface.
// A group's client should implement this interface.
type DatabaseCredentialBindingsGetter interface {
	DatabaseCredentialBindings(namespace string) DatabaseCredentialBindingInterface
}

// DatabaseCredentialBindingInterface has methods to work with DatabaseCredentialBinding resources.
type DatabaseCredentialBindingInterface interface {
	Create(*v1alpha1.DatabaseCredentialBinding) (*v1alpha1.DatabaseCredentialBinding, error)
	Update(*v1alpha1.DatabaseCredentialBinding) (*v1alpha1.DatabaseCredentialBinding, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1alpha1.DatabaseCredentialBinding, error)
	List(opts v1.ListOptions) (*v1alpha1.DatabaseCredentialBindingList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.DatabaseCredentialBinding, err error)
	DatabaseCredentialBindingExpansion
}

// databaseCredentialBindings implements DatabaseCredentialBindingInterface
type databaseCredentialBindings struct {
	client rest.Interface
	ns     string
}

// newDatabaseCredentialBindings returns a DatabaseCredentialBindings
func newDatabaseCredentialBindings(c *VaultwebhookV1alpha1Client, namespace string) *databaseCredentialBindings {
	return &databaseCredentialBindings{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the databaseCredentialBinding, and returns the corresponding databaseCredentialBinding object, and an error if there is any.
func (c *databaseCredentialBindings) Get(name string, options v1.GetOptions) (result *v1alpha1.DatabaseCredentialBinding, err error) {
	result = &v1alpha1.DatabaseCredentialBinding{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("databasecredentialbindings").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of DatabaseCredentialBindings that match those selectors.
func (c *databaseCredentialBindings) List(opts v1.ListOptions) (result *v1alpha1.DatabaseCredentialBindingList, err error) {
	result = &v1alpha1.DatabaseCredentialBindingList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("databasecredentialbindings").
		VersionedParams(&opts, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested databaseCredentialBindings.
func (c *databaseCredentialBindings) Watch(opts v1.ListOptions) (watch.Interface, error) {
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("databasecredentialbindings").
		VersionedParams(&opts, scheme.ParameterCodec).
		Watch()
}

// Create takes the representation of a databaseCredentialBinding and creates it.  Returns the server's representation of the databaseCredentialBinding, and an error, if there is any.
func (c *databaseCredentialBindings) Create(databaseCredentialBinding *v1alpha1.DatabaseCredentialBinding) (result *v1alpha1.DatabaseCredentialBinding, err error) {
	result = &v1alpha1.DatabaseCredentialBinding{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("databasecredentialbindings").
		Body(databaseCredentialBinding).
		Do().
		Into(result)
	return
}

// Update takes the representation of a databaseCredentialBinding and updates it. Returns the server's representation of the databaseCredentialBinding, and an error, if there is any.
func (c *databaseCredentialBindings) Update(databaseCredentialBinding *v1alpha1.DatabaseCredentialBinding) (result *v1alpha1.DatabaseCredentialBinding, err error) {
	result = &v1alpha1.DatabaseCredentialBinding{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("databasecredentialbindings").
		Name(databaseCredentialBinding.Name).
		Body(databaseCredentialBinding).
		Do().
		Into(result)
	return
}

// Delete takes name of the databaseCredentialBinding and deletes it. Returns an error if one occurs.
func (c *databaseCredentialBindings) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("databasecredentialbindings").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *databaseCredentialBindings) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("databasecredentialbindings").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched databaseCredentialBinding.
func (c *databaseCredentialBindings) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.DatabaseCredentialBinding, err error) {
	result = &v1alpha1.DatabaseCredentialBinding{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("databasecredentialbindings").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
