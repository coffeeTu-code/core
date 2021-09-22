// Code generated by the FlatBuffers compiler. DO NOT EDIT.

package stdrank

import (
	flatbuffers "github.com/google/flatbuffers/go"
)

type dateinfo struct {
	_tab flatbuffers.Table
}

func GetRootAsdateinfo(buf []byte, offset flatbuffers.UOffsetT) *dateinfo {
	n := flatbuffers.GetUOffsetT(buf[offset:])
	x := &dateinfo{}
	x.Init(buf, n+offset)
	return x
}

func GetSizePrefixedRootAsdateinfo(buf []byte, offset flatbuffers.UOffsetT) *dateinfo {
	n := flatbuffers.GetUOffsetT(buf[offset+flatbuffers.SizeUint32:])
	x := &dateinfo{}
	x.Init(buf, n+offset+flatbuffers.SizeUint32)
	return x
}

func (rcv *dateinfo) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *dateinfo) Table() flatbuffers.Table {
	return rcv._tab
}

func (rcv *dateinfo) Hour() []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		return rcv._tab.ByteVector(o + rcv._tab.Pos)
	}
	return nil
}

func (rcv *dateinfo) Whatday() []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(6))
	if o != 0 {
		return rcv._tab.ByteVector(o + rcv._tab.Pos)
	}
	return nil
}

func (rcv *dateinfo) Remainholidays() []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(8))
	if o != 0 {
		return rcv._tab.ByteVector(o + rcv._tab.Pos)
	}
	return nil
}

func dateinfoStart(builder *flatbuffers.Builder) {
	builder.StartObject(3)
}
func dateinfoAddHour(builder *flatbuffers.Builder, hour flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(0, flatbuffers.UOffsetT(hour), 0)
}
func dateinfoAddWhatday(builder *flatbuffers.Builder, whatday flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(1, flatbuffers.UOffsetT(whatday), 0)
}
func dateinfoAddRemainholidays(builder *flatbuffers.Builder, remainholidays flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(2, flatbuffers.UOffsetT(remainholidays), 0)
}
func dateinfoEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	return builder.EndObject()
}