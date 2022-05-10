package util

import (
	"context"
	"fmt"
	"sync"
	"time"

	"github.com/go-kratos/kratos/v2/log"
	etcd "go.etcd.io/etcd/client/v3"
)

type EtcdWatcher struct {
	client         *etcd.Client
	watchKey       string
	watchKeyModRev sync.Map
	log            log.Helper
}

type waitChangeInfo struct {
	key              string
	currentKeyModRev int64
}

type TaskEtcdPrefix string

func NewEtcdWatcher(client *etcd.Client, watchKey TaskEtcdPrefix, logger log.Logger) *EtcdWatcher {
	return &EtcdWatcher{
		client:         client,
		watchKey:       string(watchKey),
		log:            *log.NewHelper(logger),
		watchKeyModRev: sync.Map{},
	}
}

func (e *EtcdWatcher) Watch(opts ...etcd.OpOption) <-chan etcd.WatchResponse {
	watchChan := e.client.Watch(context.Background(), e.watchKey, opts...)
	keyChangedChForExternal := make(chan etcd.WatchResponse)
	go func() {
		for watchResponse := range watchChan {
			response := watchResponse
			for _, event := range response.Events {
				e.watchKeyModRev.Store(string(event.Kv.Key), event.Kv.ModRevision)
			}
			go func() {
				keyChangedChForExternal <- response
			}()
		}
	}()
	return keyChangedChForExternal
}

func (e *EtcdWatcher) WaitUntilWatchKeyChanged(ctx context.Context, key string, baseRev int64) error {
	for {
		select {
		case <-ctx.Done():
			return fmt.Errorf("wait watch key %s change timeout", e.watchKey)
		default:
			value, ok := e.watchKeyModRev.Load(key)
			if !ok {
				time.Sleep(1 * time.Second)
				continue
			}
			newRev, ok := value.(int64)
			if !ok {
				panic(fmt.Errorf("convert revision %v to int64 error", value))
			}
			e.log.Infof("EtcdWatcher.WaitUntilWatchKeyChanged %s baseRev=%d newRev=%d", key, baseRev, newRev)
			if newRev < baseRev {
				time.Sleep(1 * time.Second)
				continue
			}
			return nil
		}
	}
}
