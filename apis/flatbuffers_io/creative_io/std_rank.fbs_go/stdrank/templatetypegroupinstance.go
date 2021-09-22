// Code generated by the FlatBuffers compiler. DO NOT EDIT.

package stdrank

import (
	flatbuffers "github.com/google/flatbuffers/go"
)

type templatetypegroupinstance struct {
	_tab flatbuffers.Table
}

func GetRootAstemplatetypegroupinstance(buf []byte, offset flatbuffers.UOffsetT) *templatetypegroupinstance {
	n := flatbuffers.GetUOffsetT(buf[offset:])
	x := &templatetypegroupinstance{}
	x.Init(buf, n+offset)
	return x
}

func GetSizePrefixedRootAstemplatetypegroupinstance(buf []byte, offset flatbuffers.UOffsetT) *templatetypegroupinstance {
	n := flatbuffers.GetUOffsetT(buf[offset+flatbuffers.SizeUint32:])
	x := &templatetypegroupinstance{}
	x.Init(buf, n+offset+flatbuffers.SizeUint32)
	return x
}

func (rcv *templatetypegroupinstance) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *templatetypegroupinstance) Table() flatbuffers.Table {
	return rcv._tab
}

func (rcv *templatetypegroupinstance) TemplateGroup() int32 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		return rcv._tab.GetInt32(o + rcv._tab.Pos)
	}
	return 0
}

func (rcv *templatetypegroupinstance) MutateTemplateGroup(n int32) bool {
	return rcv._tab.MutateInt32Slot(4, n)
}

func (rcv *templatetypegroupinstance) TemplateIdTuple(j int) int32 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(6))
	if o != 0 {
		a := rcv._tab.Vector(o)
		return rcv._tab.GetInt32(a + flatbuffers.UOffsetT(j*4))
	}
	return 0
}

func (rcv *templatetypegroupinstance) TemplateIdTupleLength() int {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(6))
	if o != 0 {
		return rcv._tab.VectorLen(o)
	}
	return 0
}

func (rcv *templatetypegroupinstance) MutateTemplateIdTuple(j int, n int32) bool {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(6))
	if o != 0 {
		a := rcv._tab.Vector(o)
		return rcv._tab.MutateInt32(a+flatbuffers.UOffsetT(j*4), n)
	}
	return false
}

func (rcv *templatetypegroupinstance) TrynewGroup() bool {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(8))
	if o != 0 {
		return rcv._tab.GetBool(o + rcv._tab.Pos)
	}
	return false
}

func (rcv *templatetypegroupinstance) MutateTrynewGroup(n bool) bool {
	return rcv._tab.MutateBoolSlot(8, n)
}

func templatetypegroupinstanceStart(builder *flatbuffers.Builder) {
	builder.StartObject(3)
}
func templatetypegroupinstanceAddTemplateGroup(builder *flatbuffers.Builder, templateGroup int32) {
	builder.PrependInt32Slot(0, templateGroup, 0)
}
func templatetypegroupinstanceAddTemplateIdTuple(builder *flatbuffers.Builder, templateIdTuple flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(1, flatbuffers.UOffsetT(templateIdTuple), 0)
}
func templatetypegroupinstanceStartTemplateIdTupleVector(builder *flatbuffers.Builder, numElems int) flatbuffers.UOffsetT {
	return builder.StartVector(4, numElems, 4)
}
func templatetypegroupinstanceAddTrynewGroup(builder *flatbuffers.Builder, trynewGroup bool) {
	builder.PrependBoolSlot(2, trynewGroup, false)
}
func templatetypegroupinstanceEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	return builder.EndObject()
}
