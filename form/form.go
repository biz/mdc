// Package form implements the Form interface used in the mdc package
package form

import "net/url"

func New() *Form {
	return &Form{
		Values: url.Values{},
		Errors: map[string]string{},
	}
}

type Form struct {
	Values url.Values
	Errors map[string]string
}

func (f *Form) SetFormValue(key, value string) {
	f.Values[key] = []string{value}
}

func (f *Form) SetFormError(key, value string) {
	f.Errors[key] = value
}

func (f *Form) FormValue(key string) string {
	v, ok := f.Values[key]
	if ok && len(v) > 0 {
		return v[0]
	}
	return ""
}

func (f *Form) FormValues(key string) []string {
	v, ok := f.Values[key]
	if !ok {
		return []string{}
	}
	return v
}

func (f *Form) FormError(key string) string {
	v, ok := f.Errors[key]
	if ok {
		return v
	}
	return ""
}
