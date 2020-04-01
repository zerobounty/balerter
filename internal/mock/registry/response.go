package registry

import (
	"github.com/balerter/balerter/internal/lua_formatter"
	lua "github.com/yuin/gopher-lua"
)

func (r *Registry) Response(AnyValue, method string, args []lua.LValue) ([]lua.LValue, error) {
	e, ok := r.responseEntries[method]
	if !ok {
		return nil, ErrEntryIsNotRegistered
	}

	for _, a := range args {
		s, err := lua_formatter.ValueToString(a)
		if err != nil {
			return nil, err
		}
		e1, ok := e.entries[a.Type().String()+"@"+s]
		if !ok {
			e1, ok = e.entries[AnyValue]
			if !ok {
				return nil, ErrEntryIsNotRegistered
			}
		}
		e = e1
	}

	if len(e.responses) == 0 {
		return nil, ErrEntryIsNotRegistered
	}

	resp := e.responses[len(e.responses)-1]
	e.responses = e.responses[:len(e.responses)-1]

	return resp, nil
}
