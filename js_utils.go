// +build js,wasm

package main

import (
	"sync"
	"syscall/js"
)

type (
	_JS struct{}
)

var (
	JS _JS
)

func (_ _JS) waitForUnload() (closeCh chan struct{}) {
	closeCh = make(chan struct{})
	var closeOnce sync.Once

	beforeUnloadCb := js.NewEventCallback(0, func(event js.Value) {
		closeOnce.Do(func() {
			close(closeCh)
		})
	})
	//defer beforeUnloadCb.Release() // Это и так вызывается при выгрузке страницы
	addEventListener := js.Global().Get(`addEventListener`)
	addEventListener.Invoke(`beforeunload`, beforeUnloadCb)

	return
}
