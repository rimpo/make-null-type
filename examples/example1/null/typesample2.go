package null

import (
	"encoding/json"
	log "github.com/Sirupsen/logrus"
	"runtime/debug"

	"github.com/rimpo/go-null/examples/example1/enum"
)

//Auto-generated code; DONT EDIT THIS CODE

type TypeSample2 struct {
	val   enum.TypeSample2
	valid bool
}

func NewTypeSample2(val enum.TypeSample2) TypeSample2 {
	return TypeSample2{val: val, valid: true}
}

func (t *TypeSample2) Set(val enum.TypeSample2) {
	t.val = val
	t.valid = true
}

//Logs error message
func (t *TypeSample2) Get() enum.TypeSample2 {
	if t.IsNull() {
		log.WithFields(log.Fields{"type": "TypeSample2", "stack": string(debug.Stack()[:])}).Warn("null value used !!!.")
	}
	return t.val
}

func (t *TypeSample2) GetUnsafe() enum.TypeSample2 {
	return t.val
}

func (t *TypeSample2) GetPtr() *enum.TypeSample2 {
	return &t.val
}

func (t *TypeSample2) IsNull() bool {
	return !t.valid
}

func (t *TypeSample2) Reset() {
	t.valid = false
}

//Must for loading from external data (i.e. database, elastic, redis, etc.). logs error message
func (t *TypeSample2) SetSafe(val enum.TypeSample2) {
	if !IsValueTypeSample2(val) {
		log.WithFields(log.Fields{"type": "TypeSample2", "value": val, "stack": string(debug.Stack()[:])}).Warn("unknown value assigned !!!.")
	}
	t.val = val
	t.valid = true
}

var (
	mapTypeSample2IDToText = map[enum.TypeSample2]string{
		enum.Sample2_1:  "A",
		enum.Sample2_2:  "B",
		enum.Sample2_3:  "C",
		enum.Sample2_4:  "D",
		enum.Sample2_5:  "E",
		enum.Sample2_6:  "F",
		enum.Sample2_7:  "G",
		enum.Sample2_8:  "H",
		enum.Sample2_9:  "I",
		enum.Sample2_10: "J",
		enum.Sample2_11: "K",
	}
)

func _LookupTypeSample2IDToText(val enum.TypeSample2) (string, bool) {
	res, ok := mapTypeSample2IDToText[val]
	return res, ok

}

func IsValueTypeSample2(val enum.TypeSample2) bool {
	_, ok := _LookupTypeSample2IDToText(val)
	return ok
}

func (t *TypeSample2) GetDisplay() string {
	if !t.valid {
		return ""
	}
	val, ok := _LookupTypeSample2IDToText(t.val)
	if ok {
		return val
	}
	return ""
}

func (t *TypeSample2) MarshalJSON() ([]byte, error) {
	val, ok := _LookupTypeSample2IDToText(t.val)
	if ok {
		return json.Marshal(val)
	}
	return json.Marshal(t.val)
}
