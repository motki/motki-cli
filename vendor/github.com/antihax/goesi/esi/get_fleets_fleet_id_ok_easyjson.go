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

func easyjson19c0c1b4DecodeGithubComAntihaxGoesiEsi(in *jlexer.Lexer, out *GetFleetsFleetIdOkList) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		in.Skip()
		*out = nil
	} else {
		in.Delim('[')
		if *out == nil {
			if !in.IsDelim(']') {
				*out = make(GetFleetsFleetIdOkList, 0, 2)
			} else {
				*out = GetFleetsFleetIdOkList{}
			}
		} else {
			*out = (*out)[:0]
		}
		for !in.IsDelim(']') {
			var v1 GetFleetsFleetIdOk
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
func easyjson19c0c1b4EncodeGithubComAntihaxGoesiEsi(out *jwriter.Writer, in GetFleetsFleetIdOkList) {
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
func (v GetFleetsFleetIdOkList) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson19c0c1b4EncodeGithubComAntihaxGoesiEsi(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v GetFleetsFleetIdOkList) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson19c0c1b4EncodeGithubComAntihaxGoesiEsi(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *GetFleetsFleetIdOkList) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson19c0c1b4DecodeGithubComAntihaxGoesiEsi(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *GetFleetsFleetIdOkList) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson19c0c1b4DecodeGithubComAntihaxGoesiEsi(l, v)
}
func easyjson19c0c1b4DecodeGithubComAntihaxGoesiEsi1(in *jlexer.Lexer, out *GetFleetsFleetIdOk) {
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
		case "is_free_move":
			out.IsFreeMove = bool(in.Bool())
		case "is_registered":
			out.IsRegistered = bool(in.Bool())
		case "is_voice_enabled":
			out.IsVoiceEnabled = bool(in.Bool())
		case "motd":
			out.Motd = string(in.String())
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
func easyjson19c0c1b4EncodeGithubComAntihaxGoesiEsi1(out *jwriter.Writer, in GetFleetsFleetIdOk) {
	out.RawByte('{')
	first := true
	_ = first
	if in.IsFreeMove {
		if !first {
			out.RawByte(',')
		}
		first = false
		out.RawString("\"is_free_move\":")
		out.Bool(bool(in.IsFreeMove))
	}
	if in.IsRegistered {
		if !first {
			out.RawByte(',')
		}
		first = false
		out.RawString("\"is_registered\":")
		out.Bool(bool(in.IsRegistered))
	}
	if in.IsVoiceEnabled {
		if !first {
			out.RawByte(',')
		}
		first = false
		out.RawString("\"is_voice_enabled\":")
		out.Bool(bool(in.IsVoiceEnabled))
	}
	if in.Motd != "" {
		if !first {
			out.RawByte(',')
		}
		first = false
		out.RawString("\"motd\":")
		out.String(string(in.Motd))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v GetFleetsFleetIdOk) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson19c0c1b4EncodeGithubComAntihaxGoesiEsi1(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v GetFleetsFleetIdOk) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson19c0c1b4EncodeGithubComAntihaxGoesiEsi1(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *GetFleetsFleetIdOk) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson19c0c1b4DecodeGithubComAntihaxGoesiEsi1(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *GetFleetsFleetIdOk) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson19c0c1b4DecodeGithubComAntihaxGoesiEsi1(l, v)
}
