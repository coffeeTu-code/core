// Code generated by the FlatBuffers compiler. DO NOT EDIT.

package stdrank

import (
	flatbuffers "github.com/google/flatbuffers/go"
)

type respcampaignext struct {
	_tab flatbuffers.Table
}

func GetRootAsrespcampaignext(buf []byte, offset flatbuffers.UOffsetT) *respcampaignext {
	n := flatbuffers.GetUOffsetT(buf[offset:])
	x := &respcampaignext{}
	x.Init(buf, n+offset)
	return x
}

func GetSizePrefixedRootAsrespcampaignext(buf []byte, offset flatbuffers.UOffsetT) *respcampaignext {
	n := flatbuffers.GetUOffsetT(buf[offset+flatbuffers.SizeUint32:])
	x := &respcampaignext{}
	x.Init(buf, n+offset+flatbuffers.SizeUint32)
	return x
}

func (rcv *respcampaignext) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *respcampaignext) Table() flatbuffers.Table {
	return rcv._tab
}

func respcampaignextStart(builder *flatbuffers.Builder) {
	builder.StartObject(0)
}
func respcampaignextEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	return builder.EndObject()
}
