// Code generated by the FlatBuffers compiler. DO NOT EDIT.

package stdrank

import (
	flatbuffers "github.com/google/flatbuffers/go"
)

type templatedetail struct {
	_tab flatbuffers.Table
}

func GetRootAstemplatedetail(buf []byte, offset flatbuffers.UOffsetT) *templatedetail {
	n := flatbuffers.GetUOffsetT(buf[offset:])
	x := &templatedetail{}
	x.Init(buf, n+offset)
	return x
}

func GetSizePrefixedRootAstemplatedetail(buf []byte, offset flatbuffers.UOffsetT) *templatedetail {
	n := flatbuffers.GetUOffsetT(buf[offset+flatbuffers.SizeUint32:])
	x := &templatedetail{}
	x.Init(buf, n+offset+flatbuffers.SizeUint32)
	return x
}

func (rcv *templatedetail) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *templatedetail) Table() flatbuffers.Table {
	return rcv._tab
}

func (rcv *templatedetail) Ele() []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		return rcv._tab.ByteVector(o + rcv._tab.Pos)
	}
	return nil
}

func (rcv *templatedetail) Crid() int64 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(6))
	if o != 0 {
		return rcv._tab.GetInt64(o + rcv._tab.Pos)
	}
	return 0
}

func (rcv *templatedetail) MutateCrid(n int64) bool {
	return rcv._tab.MutateInt64Slot(6, n)
}

func templatedetailStart(builder *flatbuffers.Builder) {
	builder.StartObject(2)
}
func templatedetailAddEle(builder *flatbuffers.Builder, ele flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(0, flatbuffers.UOffsetT(ele), 0)
}
func templatedetailAddCrid(builder *flatbuffers.Builder, crid int64) {
	builder.PrependInt64Slot(1, crid, 0)
}
func templatedetailEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	return builder.EndObject()
}
