// Code generated by the FlatBuffers compiler. DO NOT EDIT.

package std_rank

import (
	flatbuffers "github.com/google/flatbuffers/go"
)

type CreativeIndex struct {
	_tab flatbuffers.Table
}

func GetRootAsCreativeIndex(buf []byte, offset flatbuffers.UOffsetT) *CreativeIndex {
	n := flatbuffers.GetUOffsetT(buf[offset:])
	x := &CreativeIndex{}
	x.Init(buf, n+offset)
	return x
}

func GetSizePrefixedRootAsCreativeIndex(buf []byte, offset flatbuffers.UOffsetT) *CreativeIndex {
	n := flatbuffers.GetUOffsetT(buf[offset+flatbuffers.SizeUint32:])
	x := &CreativeIndex{}
	x.Init(buf, n+offset+flatbuffers.SizeUint32)
	return x
}

func (rcv *CreativeIndex) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *CreativeIndex) Table() flatbuffers.Table {
	return rcv._tab
}

func (rcv *CreativeIndex) Creativeid() int64 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		return rcv._tab.GetInt64(o + rcv._tab.Pos)
	}
	return 0
}

func (rcv *CreativeIndex) MutateCreativeid(n int64) bool {
	return rcv._tab.MutateInt64Slot(4, n)
}

func (rcv *CreativeIndex) Mgodocid() []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(6))
	if o != 0 {
		return rcv._tab.ByteVector(o + rcv._tab.Pos)
	}
	return nil
}

func (rcv *CreativeIndex) Trynewcreative() bool {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(8))
	if o != 0 {
		return rcv._tab.GetBool(o + rcv._tab.Pos)
	}
	return false
}

func (rcv *CreativeIndex) MutateTrynewcreative(n bool) bool {
	return rcv._tab.MutateBoolSlot(8, n)
}

func CreativeIndexStart(builder *flatbuffers.Builder) {
	builder.StartObject(3)
}
func CreativeIndexAddCreativeid(builder *flatbuffers.Builder, creativeid int64) {
	builder.PrependInt64Slot(0, creativeid, 0)
}
func CreativeIndexAddMgodocid(builder *flatbuffers.Builder, mgodocid flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(1, flatbuffers.UOffsetT(mgodocid), 0)
}
func CreativeIndexAddTrynewcreative(builder *flatbuffers.Builder, trynewcreative bool) {
	builder.PrependBoolSlot(2, trynewcreative, false)
}
func CreativeIndexEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	return builder.EndObject()
}
