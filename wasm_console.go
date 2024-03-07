package tint

import (
	"fmt"
	"syscall/js"
)

// WasmConsole represent as browser Window.WasmConsole
type WasmConsole struct {
	js.Value
}

func NewWasmConsole() *WasmConsole {
	return &WasmConsole{js.Global().Get("console")}
}

// https://developer.mozilla.org/en-US/docs/Web/API/console/assert
func (p *WasmConsole) Assert(condition bool, a ...any) {
	p.Call("assert", condition, js.ValueOf(a))
}

// https://developer.mozilla.org/en-US/docs/Web/API/console/clear
func (p *WasmConsole) Clear() {
	p.Call("clear")
}

// https://developer.mozilla.org/en-US/docs/Web/API/console/count
func (p *WasmConsole) Count(label string) {
	p.Call("count", label)
}

// https://developer.mozilla.org/en-US/docs/Web/API/console/countReset
func (p *WasmConsole) CountReset(label string) {
	p.Call("countReset", label)
}

// https://developer.mozilla.org/en-US/docs/Web/API/console/dir
func (p *WasmConsole) Dir(obj any) {
	p.Call("dir", js.ValueOf(obj))
}

// https://developer.mozilla.org/en-US/docs/Web/API/console/dirxml
func (p *WasmConsole) Dirxml(obj any) {
	p.Call("dirxml", obj)
}

// https://developer.mozilla.org/en-US/docs/Web/API/console/group
func (p *WasmConsole) Group(label string) {
	p.Call("group", label)
}

// https://developer.mozilla.org/en-US/docs/Web/API/console/groupCollapsed
func (p *WasmConsole) GroupCollapsed(label string) {
	p.Call("groupCollapsed", label)
}

// https://developer.mozilla.org/en-US/docs/Web/API/console/groupEnd
func (p *WasmConsole) groupEnd() {
	p.Call("groupEnd")
}

// https://developer.mozilla.org/en-US/docs/Web/API/console/table
func (p *WasmConsole) Table(data any) {
	p.Call("table", js.ValueOf(data))
}

// https://developer.mozilla.org/en-US/docs/Web/API/console/time
func (p *WasmConsole) Time(label string) {
	p.Call("time", label)
}

// https://developer.mozilla.org/en-US/docs/Web/API/console/timeEnd
func (p *WasmConsole) TimeEnd(label string) {
	p.Call("timeEnd", label)
}

// https://developer.mozilla.org/en-US/docs/Web/API/console/timeLog
func (p *WasmConsole) TimeLog(label string) {
	p.Call("timeLog", label)
}

// https://developer.mozilla.org/en-US/docs/Web/API/console/trace
func (p *WasmConsole) Trace(msg string, a ...any) {
	p.Call("trace", fmt.Sprintf(msg, a...))
}

// https://developer.mozilla.org/en-US/docs/Web/API/console/debug
func (p *WasmConsole) Debug(msg []byte) {
	p.Call("debug", string(msg))
}

// https://developer.mozilla.org/en-US/docs/Web/API/console/info
func (p *WasmConsole) Info(msg []byte) {
	p.Call("info", string(msg))
}

// https://developer.mozilla.org/en-US/docs/Web/API/console/warn
func (p *WasmConsole) Warn(msg []byte) {
	p.Call("warn", string(msg))
}

// https://developer.mozilla.org/en-US/docs/Web/API/console/error
func (p *WasmConsole) Error(msg []byte) {
	p.Call("error", string(msg))
}

// https://developer.mozilla.org/en-US/docs/Web/API/console/log
func (p *WasmConsole) Log(msg []byte) {
	p.Call("log", string(msg))
}

func (p *WasmConsole) Write(buf []byte) (n int, err error) {
	p.Log(buf)
	return len(buf), nil
}
