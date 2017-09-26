// Code generated by easyjson for marshaling/unmarshaling. DO NOT EDIT.

package esi

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

func easyjson6491e41cDecodeGithubComAntihaxGoesiEsi(in *jlexer.Lexer, out *GetUniverseSystemJumps200OkList) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		in.Skip()
		*out = nil
	} else {
		in.Delim('[')
		if *out == nil {
			if !in.IsDelim(']') {
				*out = make(GetUniverseSystemJumps200OkList, 0, 8)
			} else {
				*out = GetUniverseSystemJumps200OkList{}
			}
		} else {
			*out = (*out)[:0]
		}
		for !in.IsDelim(']') {
			var v1 GetUniverseSystemJumps200Ok
			(v1).UnmarshalEasyJSON(in)
			*out = append(*out, v1)
			in.WantComma()
		}
		in.Delim(']')
	}
	if isTopLevel {
		in.Consumed()
	}
}
func easyjson6491e41cEncodeGithubComAntihaxGoesiEsi(out *jwriter.Writer, in GetUniverseSystemJumps200OkList) {
	if in == nil && (out.Flags&jwriter.NilSliceAsEmpty) == 0 {
		out.RawString("null")
	} else {
		out.RawByte('[')
		for v2, v3 := range in {
			if v2 > 0 {
				out.RawByte(',')
			}
			(v3).MarshalEasyJSON(out)
		}
		out.RawByte(']')
	}
}

// MarshalJSON supports json.Marshaler interface
func (v GetUniverseSystemJumps200OkList) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson6491e41cEncodeGithubComAntihaxGoesiEsi(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v GetUniverseSystemJumps200OkList) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson6491e41cEncodeGithubComAntihaxGoesiEsi(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *GetUniverseSystemJumps200OkList) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson6491e41cDecodeGithubComAntihaxGoesiEsi(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *GetUniverseSystemJumps200OkList) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson6491e41cDecodeGithubComAntihaxGoesiEsi(l, v)
}
func easyjson6491e41cDecodeGithubComAntihaxGoesiEsi1(in *jlexer.Lexer, out *GetUniverseSystemJumps200Ok) {
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
		key := in.UnsafeString()
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "ship_jumps":
			out.ShipJumps = int32(in.Int32())
		case "system_id":
			out.SystemId = int32(in.Int32())
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
func easyjson6491e41cEncodeGithubComAntihaxGoesiEsi1(out *jwriter.Writer, in GetUniverseSystemJumps200Ok) {
	out.RawByte('{')
	first := true
	_ = first
	if in.ShipJumps != 0 {
		if !first {
			out.RawByte(',')
		}
		first = false
		out.RawString("\"ship_jumps\":")
		out.Int32(int32(in.ShipJumps))
	}
	if in.SystemId != 0 {
		if !first {
			out.RawByte(',')
		}
		first = false
		out.RawString("\"system_id\":")
		out.Int32(int32(in.SystemId))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v GetUniverseSystemJumps200Ok) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson6491e41cEncodeGithubComAntihaxGoesiEsi1(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v GetUniverseSystemJumps200Ok) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson6491e41cEncodeGithubComAntihaxGoesiEsi1(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *GetUniverseSystemJumps200Ok) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson6491e41cDecodeGithubComAntihaxGoesiEsi1(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *GetUniverseSystemJumps200Ok) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson6491e41cDecodeGithubComAntihaxGoesiEsi1(l, v)
}
