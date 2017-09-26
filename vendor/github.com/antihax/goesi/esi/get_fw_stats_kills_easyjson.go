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

func easyjson5bbe7ca8DecodeGithubComAntihaxGoesiEsi(in *jlexer.Lexer, out *GetFwStatsKillsList) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		in.Skip()
		*out = nil
	} else {
		in.Delim('[')
		if *out == nil {
			if !in.IsDelim(']') {
				*out = make(GetFwStatsKillsList, 0, 5)
			} else {
				*out = GetFwStatsKillsList{}
			}
		} else {
			*out = (*out)[:0]
		}
		for !in.IsDelim(']') {
			var v1 GetFwStatsKills
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
func easyjson5bbe7ca8EncodeGithubComAntihaxGoesiEsi(out *jwriter.Writer, in GetFwStatsKillsList) {
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
func (v GetFwStatsKillsList) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson5bbe7ca8EncodeGithubComAntihaxGoesiEsi(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v GetFwStatsKillsList) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson5bbe7ca8EncodeGithubComAntihaxGoesiEsi(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *GetFwStatsKillsList) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson5bbe7ca8DecodeGithubComAntihaxGoesiEsi(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *GetFwStatsKillsList) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson5bbe7ca8DecodeGithubComAntihaxGoesiEsi(l, v)
}
func easyjson5bbe7ca8DecodeGithubComAntihaxGoesiEsi1(in *jlexer.Lexer, out *GetFwStatsKills) {
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
		case "last_week":
			out.LastWeek = int32(in.Int32())
		case "total":
			out.Total = int32(in.Int32())
		case "yesterday":
			out.Yesterday = int32(in.Int32())
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
func easyjson5bbe7ca8EncodeGithubComAntihaxGoesiEsi1(out *jwriter.Writer, in GetFwStatsKills) {
	out.RawByte('{')
	first := true
	_ = first
	if in.LastWeek != 0 {
		if !first {
			out.RawByte(',')
		}
		first = false
		out.RawString("\"last_week\":")
		out.Int32(int32(in.LastWeek))
	}
	if in.Total != 0 {
		if !first {
			out.RawByte(',')
		}
		first = false
		out.RawString("\"total\":")
		out.Int32(int32(in.Total))
	}
	if in.Yesterday != 0 {
		if !first {
			out.RawByte(',')
		}
		first = false
		out.RawString("\"yesterday\":")
		out.Int32(int32(in.Yesterday))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v GetFwStatsKills) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson5bbe7ca8EncodeGithubComAntihaxGoesiEsi1(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v GetFwStatsKills) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson5bbe7ca8EncodeGithubComAntihaxGoesiEsi1(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *GetFwStatsKills) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson5bbe7ca8DecodeGithubComAntihaxGoesiEsi1(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *GetFwStatsKills) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson5bbe7ca8DecodeGithubComAntihaxGoesiEsi1(l, v)
}
