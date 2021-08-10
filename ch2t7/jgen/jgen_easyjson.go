// Code generated by easyjson for marshaling/unmarshaling. DO NOT EDIT.

package jgen

import (
	json "encoding/json"
	easyjson "github.com/mailru/easyjson"
	jlexer "github.com/mailru/easyjson/jlexer"
	jwriter "github.com/mailru/easyjson/jwriter"
)

// suppress unused package warning
var (
	_ *json.RawMessage
	_ *jlexer.Lexer
	_ *jwriter.Writer
	_ easyjson.Marshaler
)

func easyjson725a3914DecodeGithubComQdrowTchGolangTreeCh2t7Jgen(in *jlexer.Lexer, out *MyStruct) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		if isTopLevel {
			in.Consumed()
		}
		in.Skip()
		return
	}
	in.Delim('{')
	for !in.IsDelim('}') {
		key := in.UnsafeFieldName(false)
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "Fname":
			out.Fname = string(in.String())
		case "Sname":
			out.Sname = string(in.String())
		case "Age":
			out.Age = int(in.Int())
		case "Performance":
			out.Performance = float32(in.Float32())
		default:
			in.SkipRecursive()
		}
		in.WantComma()
	}
	in.Delim('}')
	if isTopLevel {
		in.Consumed()
	}
}
func easyjson725a3914EncodeGithubComQdrowTchGolangTreeCh2t7Jgen(out *jwriter.Writer, in MyStruct) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"Fname\":"
		out.RawString(prefix[1:])
		out.String(string(in.Fname))
	}
	{
		const prefix string = ",\"Sname\":"
		out.RawString(prefix)
		out.String(string(in.Sname))
	}
	{
		const prefix string = ",\"Age\":"
		out.RawString(prefix)
		out.Int(int(in.Age))
	}
	{
		const prefix string = ",\"Performance\":"
		out.RawString(prefix)
		out.Float32(float32(in.Performance))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v MyStruct) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson725a3914EncodeGithubComQdrowTchGolangTreeCh2t7Jgen(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v MyStruct) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson725a3914EncodeGithubComQdrowTchGolangTreeCh2t7Jgen(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *MyStruct) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson725a3914DecodeGithubComQdrowTchGolangTreeCh2t7Jgen(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *MyStruct) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson725a3914DecodeGithubComQdrowTchGolangTreeCh2t7Jgen(l, v)
}
