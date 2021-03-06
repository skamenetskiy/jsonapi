// Code generated by easyjson for marshaling/unmarshaling. DO NOT EDIT.

package main

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

func easyjson89aae3efDecodeGitlabComNgrsLibGoJsonapiExamplesController(in *jlexer.Lexer, out *names) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		in.Skip()
		*out = nil
	} else {
		in.Delim('[')
		if *out == nil {
			if !in.IsDelim(']') {
				*out = make(names, 0, 8)
			} else {
				*out = names{}
			}
		} else {
			*out = (*out)[:0]
		}
		for !in.IsDelim(']') {
			var v1 *name
			if in.IsNull() {
				in.Skip()
				v1 = nil
			} else {
				if v1 == nil {
					v1 = new(name)
				}
				if data := in.Raw(); in.Ok() {
					in.AddError((*v1).UnmarshalJSON(data))
				}
			}
			*out = append(*out, v1)
			in.WantComma()
		}
		in.Delim(']')
	}
	if isTopLevel {
		in.Consumed()
	}
}
func easyjson89aae3efEncodeGitlabComNgrsLibGoJsonapiExamplesController(out *jwriter.Writer, in names) {
	if in == nil && (out.Flags&jwriter.NilSliceAsEmpty) == 0 {
		out.RawString("null")
	} else {
		out.RawByte('[')
		for v2, v3 := range in {
			if v2 > 0 {
				out.RawByte(',')
			}
			if v3 == nil {
				out.RawString("null")
			} else {
				out.Raw((*v3).MarshalJSON())
			}
		}
		out.RawByte(']')
	}
}

// MarshalJSON supports json.Marshaler interface
func (v names) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson89aae3efEncodeGitlabComNgrsLibGoJsonapiExamplesController(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v names) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson89aae3efEncodeGitlabComNgrsLibGoJsonapiExamplesController(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *names) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson89aae3efDecodeGitlabComNgrsLibGoJsonapiExamplesController(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *names) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson89aae3efDecodeGitlabComNgrsLibGoJsonapiExamplesController(l, v)
}
func easyjson89aae3efDecodeGitlabComNgrsLibGoJsonapiExamplesController1(in *jlexer.Lexer, out *name) {
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
		case "id":
			out.ID = int(in.Int())
		case "name":
			out.Name = string(in.String())
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
func easyjson89aae3efEncodeGitlabComNgrsLibGoJsonapiExamplesController1(out *jwriter.Writer, in name) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"id\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Int(int(in.ID))
	}
	{
		const prefix string = ",\"name\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.Name))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v name) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson89aae3efEncodeGitlabComNgrsLibGoJsonapiExamplesController1(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v name) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson89aae3efEncodeGitlabComNgrsLibGoJsonapiExamplesController1(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *name) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson89aae3efDecodeGitlabComNgrsLibGoJsonapiExamplesController1(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *name) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson89aae3efDecodeGitlabComNgrsLibGoJsonapiExamplesController1(l, v)
}
