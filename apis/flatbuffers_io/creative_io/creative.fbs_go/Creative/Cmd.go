// Code generated by the FlatBuffers compiler. DO NOT EDIT.

package Creative

import (
	flatbuffers "github.com/google/flatbuffers/go"
)

type Cmd struct {
	_tab flatbuffers.Table
}

func GetRootAsCmd(buf []byte, offset flatbuffers.UOffsetT) *Cmd {
	n := flatbuffers.GetUOffsetT(buf[offset:])
	x := &Cmd{}
	x.Init(buf, n+offset)
	return x
}

func GetSizePrefixedRootAsCmd(buf []byte, offset flatbuffers.UOffsetT) *Cmd {
	n := flatbuffers.GetUOffsetT(buf[offset+flatbuffers.SizeUint32:])
	x := &Cmd{}
	x.Init(buf, n+offset+flatbuffers.SizeUint32)
	return x
}

func (rcv *Cmd) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *Cmd) Table() flatbuffers.Table {
	return rcv._tab
}

func (rcv *Cmd) Cmd() int64 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		return rcv._tab.GetInt64(o + rcv._tab.Pos)
	}
	return 0
}

func (rcv *Cmd) MutateCmd(n int64) bool {
	return rcv._tab.MutateInt64Slot(4, n)
}

func CmdStart(builder *flatbuffers.Builder) {
	builder.StartObject(1)
}
func CmdAddCmd(builder *flatbuffers.Builder, cmd int64) {
	builder.PrependInt64Slot(0, cmd, 0)
}
func CmdEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	return builder.EndObject()
}