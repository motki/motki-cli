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

func easyjsonEe3cb066DecodeGithubComAntihaxGoesiEsi(in *jlexer.Lexer, out *GetFwLeaderboardsCharactersLastWeekList) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		in.Skip()
		*out = nil
	} else {
		in.Delim('[')
		if *out == nil {
			if !in.IsDelim(']') {
				*out = make(GetFwLeaderboardsCharactersLastWeekList, 0, 8)
			} else {
				*out = GetFwLeaderboardsCharactersLastWeekList{}
			}
		} else {
			*out = (*out)[:0]
		}
		for !in.IsDelim(']') {
			var v1 GetFwLeaderboardsCharactersLastWeek
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
func easyjsonEe3cb066EncodeGithubComAntihaxGoesiEsi(out *jwriter.Writer, in GetFwLeaderboardsCharactersLastWeekList) {
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
func (v GetFwLeaderboardsCharactersLastWeekList) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonEe3cb066EncodeGithubComAntihaxGoesiEsi(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v GetFwLeaderboardsCharactersLastWeekList) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonEe3cb066EncodeGithubComAntihaxGoesiEsi(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *GetFwLeaderboardsCharactersLastWeekList) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonEe3cb066DecodeGithubComAntihaxGoesiEsi(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *GetFwLeaderboardsCharactersLastWeekList) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonEe3cb066DecodeGithubComAntihaxGoesiEsi(l, v)
}
func easyjsonEe3cb066DecodeGithubComAntihaxGoesiEsi1(in *jlexer.Lexer, out *GetFwLeaderboardsCharactersLastWeek) {
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
		case "amount":
			out.Amount = int32(in.Int32())
		case "character_id":
			out.CharacterId = int32(in.Int32())
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
func easyjsonEe3cb066EncodeGithubComAntihaxGoesiEsi1(out *jwriter.Writer, in GetFwLeaderboardsCharactersLastWeek) {
	out.RawByte('{')
	first := true
	_ = first
	if in.Amount != 0 {
		if !first {
			out.RawByte(',')
		}
		first = false
		out.RawString("\"amount\":")
		out.Int32(int32(in.Amount))
	}
	if in.CharacterId != 0 {
		if !first {
			out.RawByte(',')
		}
		first = false
		out.RawString("\"character_id\":")
		out.Int32(int32(in.CharacterId))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v GetFwLeaderboardsCharactersLastWeek) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonEe3cb066EncodeGithubComAntihaxGoesiEsi1(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v GetFwLeaderboardsCharactersLastWeek) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonEe3cb066EncodeGithubComAntihaxGoesiEsi1(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *GetFwLeaderboardsCharactersLastWeek) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonEe3cb066DecodeGithubComAntihaxGoesiEsi1(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *GetFwLeaderboardsCharactersLastWeek) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonEe3cb066DecodeGithubComAntihaxGoesiEsi1(l, v)
}
