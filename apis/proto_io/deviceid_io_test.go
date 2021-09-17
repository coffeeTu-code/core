package proto_io

import (
	"fmt"
	"testing"

	"github.com/golang/protobuf/proto"

	"github.com/crawler-007/core/apis/proto_io/deviceid_io"
)

func TestPROTO_DeviceId(t *testing.T) {
	deviceid := &deviceid_io.DeviceId{
		Gaid: "gaid",
	}

	pbys, perr := proto.Marshal(deviceid)
	if perr != nil {
		t.Fatal(perr)
	}

	deviceidC := &deviceid_io.DeviceIdCollector{
		DeviceIds: string(pbys),
	}

	pbys, perr = proto.Marshal(deviceidC)
	if perr != nil {
		t.Fatal(perr)
	}

	deviceidC_output := &deviceid_io.DeviceIdCollector{}
	perr = proto.Unmarshal(pbys, deviceidC_output)
	if perr != nil {
		t.Fatal(perr)
	}
	fmt.Println(deviceidC_output)

	deviceid_output := &deviceid_io.DeviceId{}
	perr = proto.Unmarshal([]byte(deviceidC_output.DeviceIds), deviceid_output)
	if perr != nil {
		t.Fatal(perr)
	}
	fmt.Println(deviceid_output)
}
