package flatbuffers_io

import (
	"fmt"
	"testing"

	flatbuffers "github.com/google/flatbuffers/go"

	"github.com/crawler-007/core/apis/flatbuffers_io/creative_io/std_rank.fbs_go/stdrank"
)

func TestFBS_Creative(t *testing.T) {
	builder := flatbuffers.NewBuilder(0)

	audienceIds := []int32{1, 1, 2, 2, 3}
	keywords := []string{"k", "e", "y"}
	keyword_offset := []flatbuffers.UOffsetT{}
	for i := range keywords {
		uoffset := builder.CreateString(keywords[i])
		keyword_offset = append(keyword_offset, uoffset)
	}

	stdrank.DeviceStartAudienceidsVector(builder, len(audienceIds))
	for i := len(audienceIds) - 1; i >= 0; i-- {
		builder.PrependInt32(audienceIds[i])
	}
	audienceIds_offset := builder.EndVector(len(audienceIds))

	stdrank.DeviceStartKeywordsVector(builder, len(keyword_offset))
	for i := len(keyword_offset) - 1; i >= 0; i-- {
		builder.PrependUOffsetT(keyword_offset[i])
	}
	keywords_offset := builder.EndVector(len(keywords))

	brand_uoffset := builder.CreateString("brand")

	//
	stdrank.DeviceStart(builder)
	stdrank.DeviceAddDevicetype(builder, 2)
	stdrank.DeviceAddBrand(builder, brand_uoffset)
	stdrank.DeviceAddModel(builder, 2)
	stdrank.DeviceAddOs(builder, 2)
	stdrank.DeviceAddOsv(builder, 2)
	stdrank.DeviceAddMajorosv(builder, 2)
	stdrank.DeviceAddLanguage(builder, 2)
	stdrank.DeviceAddNetworktype(builder, 2)
	stdrank.DeviceAddCarrier(builder, 2)
	stdrank.DeviceAddScreensize(builder, 2)
	stdrank.DeviceAddAudienceids(builder, audienceIds_offset)
	stdrank.DeviceAddDeviceip(builder, 2)
	stdrank.DeviceAddKeywords(builder, keywords_offset)
	stdrank.DeviceAddAge(builder, 2)
	stdrank.DeviceAddGender(builder, 2)
	stdrank.DeviceAddDominantdevtype(builder, 2)
	stdrank.DeviceAddUa(builder, 2)
	dev := stdrank.DeviceEnd(builder)
	builder.Finish(dev)
	buf := builder.FinishedBytes()

	_ = buf // send bytes

	dev_read := stdrank.GetRootAsDevice(buf, 0)
	fmt.Println(string(dev_read.Brand()))

	fmt.Println("dev_read.KeywordsLength()=", dev_read.KeywordsLength())
	for i := 0; i < dev_read.KeywordsLength(); i++ {
		fmt.Println(string(dev_read.Keywords(i)))
	}

	fmt.Println("dev_read.AudienceidsLength()=", dev_read.AudienceidsLength())
	for i := 0; i < dev_read.AudienceidsLength(); i++ {
		fmt.Println(dev_read.Audienceids(i))
	}

}
