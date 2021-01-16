package main

import (
	"fmt"
	"syscall/js"
)

func main() {
  c := make(chan struct{}, 0)
	alert := js.Global().Get("alert")
	alert.Invoke("Hello wasm!")

  ws := js.Global().Get("WebSocket").New("ws://localhost:5001")

  ws.Call("addEventListener", "open", js.FuncOf(func(this js.Value, args []js.Value) interface{} {
    ws.Call("send", "message testing") 
    fmt.Printf("sent")
    return nil
  }))
  <-c
}
