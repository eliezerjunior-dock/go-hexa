package pkg

import (
	"github.com/golang/protobuf/jsonpb"
	"github.com/golang/protobuf/proto"

)

func ProtoToJSON(msg proto.Message) (string, error) {
	marshaler := jsonpb.Marshaler{
		EnumsAsInts:  false,
		EmitDefaults: true,
		Indent:       "  ",
		OrigName:     true,
	}

	return marshaler.MarshalToString(msg)
}

func JSONToProto(data string, msg proto.Message) error {
	return jsonpb.UnmarshalString(data, msg)
}