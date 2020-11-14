// Package jsonpb provides a json codec
package jsonpb

import (
	"bytes"
	"io"
	"io/ioutil"

	// nolint: staticcheck
	oldjsonpb "github.com/golang/protobuf/jsonpb"
	// nolint: staticcheck
	oldproto "github.com/golang/protobuf/proto"
	"github.com/itzmanish/go-micro/v2/codec"
	bytesCodec "github.com/itzmanish/go-micro/v2/codec/bytes"
	jsonpb "google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

var (
	JsonpbMarshaler = jsonpb.MarshalOptions{
		UseEnumNumbers:  false,
		EmitUnpopulated: false,
		UseProtoNames:   true,
		AllowPartial:    false,
	}

	JsonpbUnmarshaler = jsonpb.UnmarshalOptions{
		DiscardUnknown: false,
		AllowPartial:   false,
	}

	OldJsonpbMarshaler = oldjsonpb.Marshaler{
		OrigName:     true,
		EmitDefaults: false,
	}

	OldJsonpbUnmarshaler = oldjsonpb.Unmarshaler{
		AllowUnknownFields: false,
	}
)

type jsonpbCodec struct {
	Conn io.ReadWriteCloser
}

func (c *jsonpbCodec) Marshal(v interface{}) ([]byte, error) {
	switch m := v.(type) {
	case nil:
		return nil, nil
	case *bytesCodec.Frame:
		return m.Data, nil
	case proto.Message:
		return JsonpbMarshaler.Marshal(m)
	case oldproto.Message:
		buf, err := OldJsonpbMarshaler.MarshalToString(m)
		return []byte(buf), err
	}
	return nil, codec.ErrInvalidMessage
}

func (c *jsonpbCodec) Unmarshal(d []byte, v interface{}) error {
	if d == nil {
		return nil
	}
	switch m := v.(type) {
	case nil:
		return nil
	case *bytesCodec.Frame:
		m.Data = d
		return nil
	case proto.Message:
		return JsonpbUnmarshaler.Unmarshal(d, m)
	case oldproto.Message:
		return OldJsonpbUnmarshaler.Unmarshal(bytes.NewReader(d), m)
	}
	return codec.ErrInvalidMessage
}
func (c *jsonpbCodec) ReadHeader(m *codec.Message, t codec.MessageType) error {
	return nil
}

func (c *jsonpbCodec) ReadBody(b interface{}) error {
	switch m := b.(type) {
	case nil:
		return nil
	case *bytesCodec.Frame:
		buf, err := ioutil.ReadAll(c.Conn)
		if err != nil {
			return err
		}
		m.Data = buf
		return nil
	case oldproto.Message:
		return OldJsonpbUnmarshaler.Unmarshal(c.Conn, m)
	case proto.Message:
		buf, err := ioutil.ReadAll(c.Conn)
		if err != nil {
			return err
		}
		return JsonpbUnmarshaler.Unmarshal(buf, m)
	}
	return codec.ErrInvalidMessage
}

func (c *jsonpbCodec) Write(m *codec.Message, b interface{}) error {
	switch m := b.(type) {
	case nil:
		return nil
	case *bytesCodec.Frame:
		_, err := c.Conn.Write(m.Data)
		return err
	case oldproto.Message:
		return OldJsonpbMarshaler.Marshal(c.Conn, m)
	case proto.Message:
		buf, err := JsonpbMarshaler.Marshal(m)
		if err != nil {
			return err
		}
		_, err = c.Conn.Write(buf)
		return err
	}
	return codec.ErrInvalidMessage
}

func (c *jsonpbCodec) Close() error {
	return c.Conn.Close()
}

func (c *jsonpbCodec) String() string {
	return "jsonpb"
}

func NewCodec(c io.ReadWriteCloser) codec.Codec {
	return &jsonpbCodec{
		Conn: c,
	}
}
