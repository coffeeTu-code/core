// Code generated by the FlatBuffers compiler. DO NOT EDIT.

package std_rank

import (
	flatbuffers "github.com/google/flatbuffers/go"
)

type Traffic struct {
	_tab flatbuffers.Table
}

func GetRootAsTraffic(buf []byte, offset flatbuffers.UOffsetT) *Traffic {
	n := flatbuffers.GetUOffsetT(buf[offset:])
	x := &Traffic{}
	x.Init(buf, n+offset)
	return x
}

func GetSizePrefixedRootAsTraffic(buf []byte, offset flatbuffers.UOffsetT) *Traffic {
	n := flatbuffers.GetUOffsetT(buf[offset+flatbuffers.SizeUint32:])
	x := &Traffic{}
	x.Init(buf, n+offset+flatbuffers.SizeUint32)
	return x
}

func (rcv *Traffic) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *Traffic) Table() flatbuffers.Table {
	return rcv._tab
}

func (rcv *Traffic) Adx() []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		return rcv._tab.ByteVector(o + rcv._tab.Pos)
	}
	return nil
}

func (rcv *Traffic) Adtype() []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(6))
	if o != 0 {
		return rcv._tab.ByteVector(o + rcv._tab.Pos)
	}
	return nil
}

func (rcv *Traffic) Publisher() []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(8))
	if o != 0 {
		return rcv._tab.ByteVector(o + rcv._tab.Pos)
	}
	return nil
}

func (rcv *Traffic) Publishid() []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(10))
	if o != 0 {
		return rcv._tab.ByteVector(o + rcv._tab.Pos)
	}
	return nil
}

func (rcv *Traffic) Packagenmae() []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(12))
	if o != 0 {
		return rcv._tab.ByteVector(o + rcv._tab.Pos)
	}
	return nil
}

func (rcv *Traffic) Appid() []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(14))
	if o != 0 {
		return rcv._tab.ByteVector(o + rcv._tab.Pos)
	}
	return nil
}

func (rcv *Traffic) Placementid() []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(16))
	if o != 0 {
		return rcv._tab.ByteVector(o + rcv._tab.Pos)
	}
	return nil
}

func (rcv *Traffic) Unitid() []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(18))
	if o != 0 {
		return rcv._tab.ByteVector(o + rcv._tab.Pos)
	}
	return nil
}

func (rcv *Traffic) Size() []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(20))
	if o != 0 {
		return rcv._tab.ByteVector(o + rcv._tab.Pos)
	}
	return nil
}

func (rcv *Traffic) Companionsize() []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(22))
	if o != 0 {
		return rcv._tab.ByteVector(o + rcv._tab.Pos)
	}
	return nil
}

func (rcv *Traffic) Orientation() []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(24))
	if o != 0 {
		return rcv._tab.ByteVector(o + rcv._tab.Pos)
	}
	return nil
}

func (rcv *Traffic) Traffictype() []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(26))
	if o != 0 {
		return rcv._tab.ByteVector(o + rcv._tab.Pos)
	}
	return nil
}

func (rcv *Traffic) Appcontentcategory(j int) []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(28))
	if o != 0 {
		a := rcv._tab.Vector(o)
		return rcv._tab.ByteVector(a + flatbuffers.UOffsetT(j*4))
	}
	return nil
}

func (rcv *Traffic) AppcontentcategoryLength() int {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(28))
	if o != 0 {
		return rcv._tab.VectorLen(o)
	}
	return 0
}

func (rcv *Traffic) Appkeywords(j int) []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(30))
	if o != 0 {
		a := rcv._tab.Vector(o)
		return rcv._tab.ByteVector(a + flatbuffers.UOffsetT(j*4))
	}
	return nil
}

func (rcv *Traffic) AppkeywordsLength() int {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(30))
	if o != 0 {
		return rcv._tab.VectorLen(o)
	}
	return 0
}

func (rcv *Traffic) Country() []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(32))
	if o != 0 {
		return rcv._tab.ByteVector(o + rcv._tab.Pos)
	}
	return nil
}

func (rcv *Traffic) City() []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(34))
	if o != 0 {
		return rcv._tab.ByteVector(o + rcv._tab.Pos)
	}
	return nil
}

func (rcv *Traffic) Sdkversion() []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(36))
	if o != 0 {
		return rcv._tab.ByteVector(o + rcv._tab.Pos)
	}
	return nil
}

func (rcv *Traffic) Bidfloor() float32 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(38))
	if o != 0 {
		return rcv._tab.GetFloat32(o + rcv._tab.Pos)
	}
	return 0.0
}

func (rcv *Traffic) MutateBidfloor(n float32) bool {
	return rcv._tab.MutateFloat32Slot(38, n)
}

func (rcv *Traffic) Ispreaudit() bool {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(40))
	if o != 0 {
		return rcv._tab.GetBool(o + rcv._tab.Pos)
	}
	return false
}

func (rcv *Traffic) MutateIspreaudit(n bool) bool {
	return rcv._tab.MutateBoolSlot(40, n)
}

func (rcv *Traffic) Appversion() []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(42))
	if o != 0 {
		return rcv._tab.ByteVector(o + rcv._tab.Pos)
	}
	return nil
}

func (rcv *Traffic) Requesttype() []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(44))
	if o != 0 {
		return rcv._tab.ByteVector(o + rcv._tab.Pos)
	}
	return nil
}

func TrafficStart(builder *flatbuffers.Builder) {
	builder.StartObject(21)
}
func TrafficAddAdx(builder *flatbuffers.Builder, adx flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(0, flatbuffers.UOffsetT(adx), 0)
}
func TrafficAddAdtype(builder *flatbuffers.Builder, adtype flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(1, flatbuffers.UOffsetT(adtype), 0)
}
func TrafficAddPublisher(builder *flatbuffers.Builder, publisher flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(2, flatbuffers.UOffsetT(publisher), 0)
}
func TrafficAddPublishid(builder *flatbuffers.Builder, publishid flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(3, flatbuffers.UOffsetT(publishid), 0)
}
func TrafficAddPackagenmae(builder *flatbuffers.Builder, packagenmae flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(4, flatbuffers.UOffsetT(packagenmae), 0)
}
func TrafficAddAppid(builder *flatbuffers.Builder, appid flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(5, flatbuffers.UOffsetT(appid), 0)
}
func TrafficAddPlacementid(builder *flatbuffers.Builder, placementid flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(6, flatbuffers.UOffsetT(placementid), 0)
}
func TrafficAddUnitid(builder *flatbuffers.Builder, unitid flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(7, flatbuffers.UOffsetT(unitid), 0)
}
func TrafficAddSize(builder *flatbuffers.Builder, size flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(8, flatbuffers.UOffsetT(size), 0)
}
func TrafficAddCompanionsize(builder *flatbuffers.Builder, companionsize flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(9, flatbuffers.UOffsetT(companionsize), 0)
}
func TrafficAddOrientation(builder *flatbuffers.Builder, orientation flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(10, flatbuffers.UOffsetT(orientation), 0)
}
func TrafficAddTraffictype(builder *flatbuffers.Builder, traffictype flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(11, flatbuffers.UOffsetT(traffictype), 0)
}
func TrafficAddAppcontentcategory(builder *flatbuffers.Builder, appcontentcategory flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(12, flatbuffers.UOffsetT(appcontentcategory), 0)
}
func TrafficStartAppcontentcategoryVector(builder *flatbuffers.Builder, numElems int) flatbuffers.UOffsetT {
	return builder.StartVector(4, numElems, 4)
}
func TrafficAddAppkeywords(builder *flatbuffers.Builder, appkeywords flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(13, flatbuffers.UOffsetT(appkeywords), 0)
}
func TrafficStartAppkeywordsVector(builder *flatbuffers.Builder, numElems int) flatbuffers.UOffsetT {
	return builder.StartVector(4, numElems, 4)
}
func TrafficAddCountry(builder *flatbuffers.Builder, country flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(14, flatbuffers.UOffsetT(country), 0)
}
func TrafficAddCity(builder *flatbuffers.Builder, city flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(15, flatbuffers.UOffsetT(city), 0)
}
func TrafficAddSdkversion(builder *flatbuffers.Builder, sdkversion flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(16, flatbuffers.UOffsetT(sdkversion), 0)
}
func TrafficAddBidfloor(builder *flatbuffers.Builder, bidfloor float32) {
	builder.PrependFloat32Slot(17, bidfloor, 0.0)
}
func TrafficAddIspreaudit(builder *flatbuffers.Builder, ispreaudit bool) {
	builder.PrependBoolSlot(18, ispreaudit, false)
}
func TrafficAddAppversion(builder *flatbuffers.Builder, appversion flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(19, flatbuffers.UOffsetT(appversion), 0)
}
func TrafficAddRequesttype(builder *flatbuffers.Builder, requesttype flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(20, flatbuffers.UOffsetT(requesttype), 0)
}
func TrafficEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	return builder.EndObject()
}
