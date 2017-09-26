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

func easyjson8d4e6c8bDecodeGithubComAntihaxGoesiEsi(in *jlexer.Lexer, out *GetCharactersCharacterIdBookmarksItemList) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		in.Skip()
		*out = nil
	} else {
		in.Delim('[')
		if *out == nil {
			if !in.IsDelim(']') {
				*out = make(GetCharactersCharacterIdBookmarksItemList, 0, 4)
			} else {
				*out = GetCharactersCharacterIdBookmarksItemList{}
			}
		} else {
			*out = (*out)[:0]
		}
		for !in.IsDelim(']') {
			var v1 GetCharactersCharacterIdBookmarksItem
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
func easyjson8d4e6c8bEncodeGithubComAntihaxGoesiEsi(out *jwriter.Writer, in GetCharactersCharacterIdBookmarksItemList) {
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
func (v GetCharactersCharacterIdBookmarksItemList) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson8d4e6c8bEncodeGithubComAntihaxGoesiEsi(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v GetCharactersCharacterIdBookmarksItemList) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson8d4e6c8bEncodeGithubComAntihaxGoesiEsi(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *GetCharactersCharacterIdBookmarksItemList) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson8d4e6c8bDecodeGithubComAntihaxGoesiEsi(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *GetCharactersCharacterIdBookmarksItemList) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson8d4e6c8bDecodeGithubComAntihaxGoesiEsi(l, v)
}
func easyjson8d4e6c8bDecodeGithubComAntihaxGoesiEsi1(in *jlexer.Lexer, out *GetCharactersCharacterIdBookmarksItem) {
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
		case "item_id":
			out.ItemId = int64(in.Int64())
		case "type_id":
			out.TypeId = int32(in.Int32())
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
func easyjson8d4e6c8bEncodeGithubComAntihaxGoesiEsi1(out *jwriter.Writer, in GetCharactersCharacterIdBookmarksItem) {
	out.RawByte('{')
	first := true
	_ = first
	if in.ItemId != 0 {
		if !first {
			out.RawByte(',')
		}
		first = false
		out.RawString("\"item_id\":")
		out.Int64(int64(in.ItemId))
	}
	if in.TypeId != 0 {
		if !first {
			out.RawByte(',')
		}
		first = false
		out.RawString("\"type_id\":")
		out.Int32(int32(in.TypeId))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v GetCharactersCharacterIdBookmarksItem) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson8d4e6c8bEncodeGithubComAntihaxGoesiEsi1(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v GetCharactersCharacterIdBookmarksItem) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson8d4e6c8bEncodeGithubComAntihaxGoesiEsi1(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *GetCharactersCharacterIdBookmarksItem) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson8d4e6c8bDecodeGithubComAntihaxGoesiEsi1(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *GetCharactersCharacterIdBookmarksItem) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson8d4e6c8bDecodeGithubComAntihaxGoesiEsi1(l, v)
}
