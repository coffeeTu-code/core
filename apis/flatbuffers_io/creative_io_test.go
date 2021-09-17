package flatbuffers_io

import (
	"testing"

	flatbuffers "github.com/google/flatbuffers/go"

	"github.com/crawler-007/core/apis/flatbuffers_io/creative_io/creative.fbs_go/Creative"
)

func TestFBS_Creative(t *testing.T) {
	builder := flatbuffers.NewBuilder(0)

	{
		Creative.CreativeAttrStart(builder)
		Creative.CreativeAttrAddId(builder, 999999999)
		Creative.CreativeAttrEnd(builder)
	}

}
