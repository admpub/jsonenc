package json

import (
	"reflect"
	"strings"
)

func (e *encodeState) Reset() {
	e.Buffer.Reset()
	e.resetPath()
}

func (e *encodeState) resetPath() {
	e.path = ``
}

type Filter interface {
	Filter(string, reflect.Value) bool
}

type Selector interface {
	Select(string, reflect.Value) bool
}

type excludeFilters map[string]struct{}

func (f excludeFilters) Add(names ...string) excludeFilters {
	for _, name := range names {
		f[name] = struct{}{}
	}
	return f
}

func (f excludeFilters) Filter(name string, v reflect.Value) bool {
	_, ok := f[name]
	return ok
}

type includeFilters map[string]struct{}

func (f includeFilters) Add(names ...string) includeFilters {
	for _, name := range names {
		f[name] = struct{}{}
	}
	return f
}

func (f includeFilters) Select(name string, v reflect.Value) bool {
	_, ok := f[name]
	if ok {
		return ok
	}
	parts := strings.Split(name, `.`)
	if len(parts) == 1 {
		_, ok = f[name+`.*`]
		return ok
	}
	for index := range parts[1:] {
		prefix := strings.Join(parts[:index+1], `.`) + `.*`
		_, ok = f[prefix]
		if ok {
			return ok
		}
	}
	return ok
}

func Include(names ...string) Selector {
	f := make(includeFilters)
	return f.Add(names...)
}

func Exclude(names ...string) Filter {
	f := make(excludeFilters)
	return f.Add(names...)
}

func OptionFilter(f Filter) Option {
	return func(o *encOpts) {
		o.filter = f
	}
}

func OptionSelector(f Selector) Option {
	return func(o *encOpts) {
		o.selector = f
	}
}

func OptionEscapeHTML(escapeHTML bool) Option {
	return func(o *encOpts) {
		o.escapeHTML = escapeHTML
	}
}

func MarshalFilter(v any, f Filter) ([]byte, error) {
	return MarshalWithOption(v, OptionFilter(f))
}

func MarshalSelector(v any, f Selector) ([]byte, error) {
	return MarshalWithOption(v, OptionSelector(f))
}

type Option func(*encOpts)

func MarshalWithOption(v any, opts ...Option) ([]byte, error) {
	e := newEncodeState()
	defer encodeStatePool.Put(e)

	option := encOpts{escapeHTML: true}
	for _, opt := range opts {
		opt(&option)
	}
	err := e.marshal(v, option)
	if err != nil {
		return nil, err
	}
	buf := append([]byte(nil), e.Bytes()...)

	return buf, nil
}
