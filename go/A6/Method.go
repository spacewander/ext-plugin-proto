// Code generated by the FlatBuffers compiler. DO NOT EDIT.

package A6

import "strconv"

type Method byte

const (
	MethodGET       Method = 0
	MethodHEAD      Method = 1
	MethodPOST      Method = 2
	MethodPUT       Method = 3
	MethodDELETE    Method = 4
	MethodMKCOL     Method = 5
	MethodCOPY      Method = 6
	MethodMOVE      Method = 7
	MethodOPTIONS   Method = 8
	MethodPROPFIND  Method = 9
	MethodPROPPATCH Method = 10
	MethodLOCK      Method = 11
	MethodUNLOCK    Method = 12
	MethodPATCH     Method = 13
	MethodTRACE     Method = 14
)

var EnumNamesMethod = map[Method]string{
	MethodGET:       "GET",
	MethodHEAD:      "HEAD",
	MethodPOST:      "POST",
	MethodPUT:       "PUT",
	MethodDELETE:    "DELETE",
	MethodMKCOL:     "MKCOL",
	MethodCOPY:      "COPY",
	MethodMOVE:      "MOVE",
	MethodOPTIONS:   "OPTIONS",
	MethodPROPFIND:  "PROPFIND",
	MethodPROPPATCH: "PROPPATCH",
	MethodLOCK:      "LOCK",
	MethodUNLOCK:    "UNLOCK",
	MethodPATCH:     "PATCH",
	MethodTRACE:     "TRACE",
}

var EnumValuesMethod = map[string]Method{
	"GET":       MethodGET,
	"HEAD":      MethodHEAD,
	"POST":      MethodPOST,
	"PUT":       MethodPUT,
	"DELETE":    MethodDELETE,
	"MKCOL":     MethodMKCOL,
	"COPY":      MethodCOPY,
	"MOVE":      MethodMOVE,
	"OPTIONS":   MethodOPTIONS,
	"PROPFIND":  MethodPROPFIND,
	"PROPPATCH": MethodPROPPATCH,
	"LOCK":      MethodLOCK,
	"UNLOCK":    MethodUNLOCK,
	"PATCH":     MethodPATCH,
	"TRACE":     MethodTRACE,
}

func (v Method) String() string {
	if s, ok := EnumNamesMethod[v]; ok {
		return s
	}
	return "Method(" + strconv.FormatInt(int64(v), 10) + ")"
}
