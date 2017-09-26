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

func easyjson9cd7a0b1DecodeGithubComAntihaxGoesiEsi(in *jlexer.Lexer, out *GetCharactersCharacterIdPlanetsPlanetIdLinkList) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		in.Skip()
		*out = nil
	} else {
		in.Delim('[')
		if *out == nil {
			if !in.IsDelim(']') {
				*out = make(GetCharactersCharacterIdPlanetsPlanetIdLinkList, 0, 2)
			} else {
				*out = GetCharactersCharacterIdPlanetsPlanetIdLinkList{}
			}
		} else {
			*out = (*out)[:0]
		}
		for !in.IsDelim(']') {
			var v1 GetCharactersCharacterIdPlanetsPlanetIdLink
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
func easyjson9cd7a0b1EncodeGithubComAntihaxGoesiEsi(out *jwriter.Writer, in GetCharactersCharacterIdPlanetsPlanetIdLinkList) {
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
func (v GetCharactersCharacterIdPlanetsPlanetIdLinkList) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson9cd7a0b1EncodeGithubComAntihaxGoesiEsi(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v GetCharactersCharacterIdPlanetsPlanetIdLinkList) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson9cd7a0b1EncodeGithubComAntihaxGoesiEsi(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *GetCharactersCharacterIdPlanetsPlanetIdLinkList) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson9cd7a0b1DecodeGithubComAntihaxGoesiEsi(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *GetCharactersCharacterIdPlanetsPlanetIdLinkList) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson9cd7a0b1DecodeGithubComAntihaxGoesiEsi(l, v)
}
func easyjson9cd7a0b1DecodeGithubComAntihaxGoesiEsi1(in *jlexer.Lexer, out *GetCharactersCharacterIdPlanetsPlanetIdLink) {
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
		case "destination_pin_id":
			out.DestinationPinId = int64(in.Int64())
		case "link_level":
			out.LinkLevel = int32(in.Int32())
		case "source_pin_id":
			out.SourcePinId = int64(in.Int64())
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
func easyjson9cd7a0b1EncodeGithubComAntihaxGoesiEsi1(out *jwriter.Writer, in GetCharactersCharacterIdPlanetsPlanetIdLink) {
	out.RawByte('{')
	first := true
	_ = first
	if in.DestinationPinId != 0 {
		if !first {
			out.RawByte(',')
		}
		first = false
		out.RawString("\"destination_pin_id\":")
		out.Int64(int64(in.DestinationPinId))
	}
	if in.LinkLevel != 0 {
		if !first {
			out.RawByte(',')
		}
		first = false
		out.RawString("\"link_level\":")
		out.Int32(int32(in.LinkLevel))
	}
	if in.SourcePinId != 0 {
		if !first {
			out.RawByte(',')
		}
		first = false
		out.RawString("\"source_pin_id\":")
		out.Int64(int64(in.SourcePinId))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v GetCharactersCharacterIdPlanetsPlanetIdLink) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson9cd7a0b1EncodeGithubComAntihaxGoesiEsi1(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v GetCharactersCharacterIdPlanetsPlanetIdLink) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson9cd7a0b1EncodeGithubComAntihaxGoesiEsi1(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *GetCharactersCharacterIdPlanetsPlanetIdLink) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson9cd7a0b1DecodeGithubComAntihaxGoesiEsi1(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *GetCharactersCharacterIdPlanetsPlanetIdLink) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson9cd7a0b1DecodeGithubComAntihaxGoesiEsi1(l, v)
}
