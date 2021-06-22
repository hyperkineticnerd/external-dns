// Code generated by protoc-gen-gogo. DO NOT EDIT.
// policy/v1beta1/http_response.proto is a deprecated file.

package v1beta1

import (
	bytes "bytes"
	fmt "fmt"
	github_com_gogo_protobuf_jsonpb "github.com/gogo/protobuf/jsonpb"
	proto "github.com/gogo/protobuf/proto"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// MarshalJSON is a custom marshaler for DirectHttpResponse
func (this *DirectHttpResponse) MarshalJSON() ([]byte, error) {
	str, err := HttpResponseMarshaler.MarshalToString(this)
	return []byte(str), err
}

// UnmarshalJSON is a custom unmarshaler for DirectHttpResponse
func (this *DirectHttpResponse) UnmarshalJSON(b []byte) error {
	return HttpResponseUnmarshaler.Unmarshal(bytes.NewReader(b), this)
}

var (
	HttpResponseMarshaler   = &github_com_gogo_protobuf_jsonpb.Marshaler{}
	HttpResponseUnmarshaler = &github_com_gogo_protobuf_jsonpb.Unmarshaler{AllowUnknownFields: true}
)
