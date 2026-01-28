package gosmapi

import (
	"net/url"
	"strings"
)

// ["servers"], ["include"]:"subServers" -> /servers?include=subServers
// ["servers", "ID"] -> /servers/ID
func buildRequestPath(elems []string, params map[string]string) string {
	var s strings.Builder

	s.Grow(100)

	for _, v := range elems {
		s.WriteByte('/')
		s.WriteString(url.PathEscape(v))
	}

	if len(params) > 0 {
		s.WriteByte('?')
		paramCount := 0
		for k, v := range params {
			if paramCount > 0 {
				s.WriteByte('&')
			}
			s.WriteString(url.QueryEscape(k))
			s.WriteByte('=')
			s.WriteString(url.QueryEscape(v))
			paramCount++
		}
	}

	return s.String()
}
