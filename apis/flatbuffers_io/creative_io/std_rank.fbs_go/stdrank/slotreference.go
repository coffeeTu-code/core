// Code generated by the FlatBuffers compiler. DO NOT EDIT.

package stdrank

import (
	flatbuffers "github.com/google/flatbuffers/go"
)

type slotreference struct {
	_tab flatbuffers.Table
}

func GetRootAsslotreference(buf []byte, offset flatbuffers.UOffsetT) *slotreference {
	n := flatbuffers.GetUOffsetT(buf[offset:])
	x := &slotreference{}
	x.Init(buf, n+offset)
	return x
}

func GetSizePrefixedRootAsslotreference(buf []byte, offset flatbuffers.UOffsetT) *slotreference {
	n := flatbuffers.GetUOffsetT(buf[offset+flatbuffers.SizeUint32:])
	x := &slotreference{}
	x.Init(buf, n+offset+flatbuffers.SizeUint32)
	return x
}

func (rcv *slotreference) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *slotreference) Table() flatbuffers.Table {
	return rcv._tab
}

func (rcv *slotreference) SlotIndex() int32 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		return rcv._tab.GetInt32(o + rcv._tab.Pos)
	}
	return 0
}

func (rcv *slotreference) MutateSlotIndex(n int32) bool {
	return rcv._tab.MutateInt32Slot(4, n)
}

func (rcv *slotreference) SlotId() int32 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(6))
	if o != 0 {
		return rcv._tab.GetInt32(o + rcv._tab.Pos)
	}
	return 0
}

func (rcv *slotreference) MutateSlotId(n int32) bool {
	return rcv._tab.MutateInt32Slot(6, n)
}

func (rcv *slotreference) Required() bool {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(8))
	if o != 0 {
		return rcv._tab.GetBool(o + rcv._tab.Pos)
	}
	return false
}

func (rcv *slotreference) MutateRequired(n bool) bool {
	return rcv._tab.MutateBoolSlot(8, n)
}

func slotreferenceStart(builder *flatbuffers.Builder) {
	builder.StartObject(3)
}
func slotreferenceAddSlotIndex(builder *flatbuffers.Builder, slotIndex int32) {
	builder.PrependInt32Slot(0, slotIndex, 0)
}
func slotreferenceAddSlotId(builder *flatbuffers.Builder, slotId int32) {
	builder.PrependInt32Slot(1, slotId, 0)
}
func slotreferenceAddRequired(builder *flatbuffers.Builder, required bool) {
	builder.PrependBoolSlot(2, required, false)
}
func slotreferenceEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	return builder.EndObject()
}
