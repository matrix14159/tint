//go:build !js && !wasm

package tint

// WasmConsole represent as browser Window.WasmConsole
type WasmConsole struct {
}

func NewWasmConsole() *WasmConsole {
	return &WasmConsole{}
}

// https://developer.mozilla.org/en-US/docs/Web/API/console/assert
func (p *WasmConsole) Assert(condition bool, a ...any) {
}

// https://developer.mozilla.org/en-US/docs/Web/API/console/clear
func (p *WasmConsole) Clear() {
}

// https://developer.mozilla.org/en-US/docs/Web/API/console/count
func (p *WasmConsole) Count(label string) {
}

// https://developer.mozilla.org/en-US/docs/Web/API/console/countReset
func (p *WasmConsole) CountReset(label string) {
}

// https://developer.mozilla.org/en-US/docs/Web/API/console/dir
func (p *WasmConsole) Dir(obj any) {
}

// https://developer.mozilla.org/en-US/docs/Web/API/console/dirxml
func (p *WasmConsole) Dirxml(obj any) {
}

// https://developer.mozilla.org/en-US/docs/Web/API/console/group
func (p *WasmConsole) Group(label string) {
}

// https://developer.mozilla.org/en-US/docs/Web/API/console/groupCollapsed
func (p *WasmConsole) GroupCollapsed(label string) {
}

// https://developer.mozilla.org/en-US/docs/Web/API/console/groupEnd
func (p *WasmConsole) groupEnd() {
}

// https://developer.mozilla.org/en-US/docs/Web/API/console/table
func (p *WasmConsole) Table(data any) {
}

// https://developer.mozilla.org/en-US/docs/Web/API/console/time
func (p *WasmConsole) Time(label string) {
}

// https://developer.mozilla.org/en-US/docs/Web/API/console/timeEnd
func (p *WasmConsole) TimeEnd(label string) {
}

// https://developer.mozilla.org/en-US/docs/Web/API/console/timeLog
func (p *WasmConsole) TimeLog(label string) {
}

// https://developer.mozilla.org/en-US/docs/Web/API/console/trace
func (p *WasmConsole) Trace(msg string, a ...any) {
}

// https://developer.mozilla.org/en-US/docs/Web/API/console/debug
func (p *WasmConsole) Debug(msg []byte) {
}

// https://developer.mozilla.org/en-US/docs/Web/API/console/info
func (p *WasmConsole) Info(msg []byte) {
}

// https://developer.mozilla.org/en-US/docs/Web/API/console/warn
func (p *WasmConsole) Warn(msg []byte) {
}

// https://developer.mozilla.org/en-US/docs/Web/API/console/error
func (p *WasmConsole) Error(msg []byte) {
}

// https://developer.mozilla.org/en-US/docs/Web/API/console/log
func (p *WasmConsole) Log(msg []byte) {
}

func (p *WasmConsole) Write(buf []byte) (n int, err error) {
	return
}
