package main

import (
	"context"
	"fmt"
	"time"

	"k8s.io/apimachinery/pkg/fields"
	"k8s.io/client-go/tools/cache"

	log "github.com/sirupsen/logrus"
	"github.com/uswitch/vault-webhook/pkg/apis/vaultwebhook.uswitch.com/v1alpha1"
	webhookclient "github.com/uswitch/vault-webhook/pkg/client/clientset/versioned"
)

type bindingAggregator struct {
	store      cache.Store
	controller cache.Controller
}

func NewListWatch(client *webhookclient.Clientset) *bindingAggregator {
	binder := &bindingAggregator{}
	watcher := cache.NewListWatchFromClient(client.VaultwebhookV1alpha1().RESTClient(), "databasecredentialbindings", "", fields.Everything())
	binder.store, binder.controller = cache.NewIndexerInformer(watcher, &v1alpha1.DatabaseCredentialBinding{}, time.Minute, binder, cache.Indexers{})
	return binder
}

func (b *bindingAggregator) OnAdd(obj interface{}) {
	log.Debugf("adding %+v", obj)
}

func (b *bindingAggregator) OnDelete(obj interface{}) {
	log.Debugf("deleting %+v", obj)
}

func (b *bindingAggregator) OnUpdate(old, new interface{}) {
	log.Debugf("updating %+v", new)
}

func (b *bindingAggregator) Run(ctx context.Context) error {
	go b.controller.Run(ctx.Done())
	cache.WaitForCacheSync(ctx.Done(), b.controller.HasSynced)
	log.Debugf("cache controller synced")

	return nil
}

func (b *bindingAggregator) List() ([]v1alpha1.DatabaseCredentialBinding, error) {
	bindingList := make([]v1alpha1.DatabaseCredentialBinding, 0)
	bindings := b.store.List()
	for _, obj := range bindings {
		binding, ok := obj.(*v1alpha1.DatabaseCredentialBinding)
		if !ok {
			return nil, fmt.Errorf("unexpected object in store: %+v", obj)
		}
		bindingList = append(bindingList, *binding)
	}
	return bindingList, nil

}
