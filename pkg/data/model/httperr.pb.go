// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.12
// source: httperr.proto

package model

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type HttpErr int32

const (
	HttpErr_HTTPERR_NONE                            HttpErr = 0
	HttpErr_HTTPERR_INTERNAL                        HttpErr = 1
	HttpErr_HTTPERR_OK                              HttpErr = 200
	HttpErr_HTTPERR_BAD_REQUEST                     HttpErr = 400
	HttpErr_HTTPERR_UNAUTHORIZED                    HttpErr = 401
	HttpErr_HTTPERR_PAYMENT_REQUIRED                HttpErr = 402
	HttpErr_HTTPERR_FORBIDDEN                       HttpErr = 403
	HttpErr_HTTPERR_NOT_FOUND                       HttpErr = 404
	HttpErr_HTTPERR_METHOD_NOT_ALLOWED              HttpErr = 405
	HttpErr_HTTPERR_NOT_ACCEPTABLE                  HttpErr = 406
	HttpErr_HTTPERR_PROXY_AUTH_REQUIRED             HttpErr = 407
	HttpErr_HTTPERR_REQUEST_TIMEOUT                 HttpErr = 408
	HttpErr_HTTPERR_CONFLICT                        HttpErr = 409
	HttpErr_HTTPERR_GONE                            HttpErr = 410
	HttpErr_HTTPERR_LENGTH_REQUIRED                 HttpErr = 411
	HttpErr_HTTPERR_PRECONDITION_FAILED             HttpErr = 412
	HttpErr_HTTPERR_REQUEST_ENTITY_TOO_LARGE        HttpErr = 413
	HttpErr_HTTPERR_REQUEST_URI_TOO_LONG            HttpErr = 414
	HttpErr_HTTPERR_UNSUPPORTEDMEDIATYPE            HttpErr = 415
	HttpErr_HTTPERR_REQUESTED_RANGE_NOT_SATISFIABLE HttpErr = 416
	HttpErr_HTTPERR_EXPECTATION_FAILED              HttpErr = 417
	HttpErr_HTTPERR_TEAPOT                          HttpErr = 418
	HttpErr_HTTPERR_MISDIRECTED_REQUEST             HttpErr = 421
	HttpErr_HTTPERR_UNPROCESSABLE_ENTITY            HttpErr = 422
	HttpErr_HTTPERR_LOCKED                          HttpErr = 423
	HttpErr_HTTPERR_FAILED_DEPENDENCY               HttpErr = 424
	HttpErr_HTTPERR_TOO_EARLY                       HttpErr = 425
	HttpErr_HTTPERR_UPGRADE_REQUIRED                HttpErr = 426
	HttpErr_HTTPERR_PRECONDITION_REQUIRED           HttpErr = 428
	HttpErr_HTTPERR_TOO_MANY_REQUESTS               HttpErr = 429
	HttpErr_HTTPERR_REQUEST_HEADER_FIELDS_TOO_LARGE HttpErr = 431
	HttpErr_HTTPERR_UNAVAILABLE_FOR_LEGAL_REASONS   HttpErr = 451
	HttpErr_HTTPERR_INTERNAL_SERVER_ERROR           HttpErr = 500
	HttpErr_HTTPERR_NOT_IMPLEMENTED                 HttpErr = 501
	HttpErr_HTTPERR_BAD_GATEWAY                     HttpErr = 502
	HttpErr_HTTPERR_SERVICE_UNAVAILABLE             HttpErr = 503
	HttpErr_HTTPERR_GATEWAY_TIMEOUT                 HttpErr = 504
	HttpErr_HTTPERR_HTTP_VERSION_NOT_SUPPORTED      HttpErr = 505
	HttpErr_HTTPERR_VARIANT_ALSO_NEGOTIATES         HttpErr = 506
	HttpErr_HTTPERR_INSUFFICIENT_STORAGE            HttpErr = 507
	HttpErr_HTTPERR_LOOP_DETECTED                   HttpErr = 508
	HttpErr_HTTPERR_NOT_EXTENDED                    HttpErr = 510
	HttpErr_HTTPERR_NETWORK_AUTHENTICATION_REQUIRED HttpErr = 511
)

// Enum value maps for HttpErr.
var (
	HttpErr_name = map[int32]string{
		0:   "HTTPERR_NONE",
		1:   "HTTPERR_INTERNAL",
		200: "HTTPERR_OK",
		400: "HTTPERR_BAD_REQUEST",
		401: "HTTPERR_UNAUTHORIZED",
		402: "HTTPERR_PAYMENT_REQUIRED",
		403: "HTTPERR_FORBIDDEN",
		404: "HTTPERR_NOT_FOUND",
		405: "HTTPERR_METHOD_NOT_ALLOWED",
		406: "HTTPERR_NOT_ACCEPTABLE",
		407: "HTTPERR_PROXY_AUTH_REQUIRED",
		408: "HTTPERR_REQUEST_TIMEOUT",
		409: "HTTPERR_CONFLICT",
		410: "HTTPERR_GONE",
		411: "HTTPERR_LENGTH_REQUIRED",
		412: "HTTPERR_PRECONDITION_FAILED",
		413: "HTTPERR_REQUEST_ENTITY_TOO_LARGE",
		414: "HTTPERR_REQUEST_URI_TOO_LONG",
		415: "HTTPERR_UNSUPPORTEDMEDIATYPE",
		416: "HTTPERR_REQUESTED_RANGE_NOT_SATISFIABLE",
		417: "HTTPERR_EXPECTATION_FAILED",
		418: "HTTPERR_TEAPOT",
		421: "HTTPERR_MISDIRECTED_REQUEST",
		422: "HTTPERR_UNPROCESSABLE_ENTITY",
		423: "HTTPERR_LOCKED",
		424: "HTTPERR_FAILED_DEPENDENCY",
		425: "HTTPERR_TOO_EARLY",
		426: "HTTPERR_UPGRADE_REQUIRED",
		428: "HTTPERR_PRECONDITION_REQUIRED",
		429: "HTTPERR_TOO_MANY_REQUESTS",
		431: "HTTPERR_REQUEST_HEADER_FIELDS_TOO_LARGE",
		451: "HTTPERR_UNAVAILABLE_FOR_LEGAL_REASONS",
		500: "HTTPERR_INTERNAL_SERVER_ERROR",
		501: "HTTPERR_NOT_IMPLEMENTED",
		502: "HTTPERR_BAD_GATEWAY",
		503: "HTTPERR_SERVICE_UNAVAILABLE",
		504: "HTTPERR_GATEWAY_TIMEOUT",
		505: "HTTPERR_HTTP_VERSION_NOT_SUPPORTED",
		506: "HTTPERR_VARIANT_ALSO_NEGOTIATES",
		507: "HTTPERR_INSUFFICIENT_STORAGE",
		508: "HTTPERR_LOOP_DETECTED",
		510: "HTTPERR_NOT_EXTENDED",
		511: "HTTPERR_NETWORK_AUTHENTICATION_REQUIRED",
	}
	HttpErr_value = map[string]int32{
		"HTTPERR_NONE":                            0,
		"HTTPERR_INTERNAL":                        1,
		"HTTPERR_OK":                              200,
		"HTTPERR_BAD_REQUEST":                     400,
		"HTTPERR_UNAUTHORIZED":                    401,
		"HTTPERR_PAYMENT_REQUIRED":                402,
		"HTTPERR_FORBIDDEN":                       403,
		"HTTPERR_NOT_FOUND":                       404,
		"HTTPERR_METHOD_NOT_ALLOWED":              405,
		"HTTPERR_NOT_ACCEPTABLE":                  406,
		"HTTPERR_PROXY_AUTH_REQUIRED":             407,
		"HTTPERR_REQUEST_TIMEOUT":                 408,
		"HTTPERR_CONFLICT":                        409,
		"HTTPERR_GONE":                            410,
		"HTTPERR_LENGTH_REQUIRED":                 411,
		"HTTPERR_PRECONDITION_FAILED":             412,
		"HTTPERR_REQUEST_ENTITY_TOO_LARGE":        413,
		"HTTPERR_REQUEST_URI_TOO_LONG":            414,
		"HTTPERR_UNSUPPORTEDMEDIATYPE":            415,
		"HTTPERR_REQUESTED_RANGE_NOT_SATISFIABLE": 416,
		"HTTPERR_EXPECTATION_FAILED":              417,
		"HTTPERR_TEAPOT":                          418,
		"HTTPERR_MISDIRECTED_REQUEST":             421,
		"HTTPERR_UNPROCESSABLE_ENTITY":            422,
		"HTTPERR_LOCKED":                          423,
		"HTTPERR_FAILED_DEPENDENCY":               424,
		"HTTPERR_TOO_EARLY":                       425,
		"HTTPERR_UPGRADE_REQUIRED":                426,
		"HTTPERR_PRECONDITION_REQUIRED":           428,
		"HTTPERR_TOO_MANY_REQUESTS":               429,
		"HTTPERR_REQUEST_HEADER_FIELDS_TOO_LARGE": 431,
		"HTTPERR_UNAVAILABLE_FOR_LEGAL_REASONS":   451,
		"HTTPERR_INTERNAL_SERVER_ERROR":           500,
		"HTTPERR_NOT_IMPLEMENTED":                 501,
		"HTTPERR_BAD_GATEWAY":                     502,
		"HTTPERR_SERVICE_UNAVAILABLE":             503,
		"HTTPERR_GATEWAY_TIMEOUT":                 504,
		"HTTPERR_HTTP_VERSION_NOT_SUPPORTED":      505,
		"HTTPERR_VARIANT_ALSO_NEGOTIATES":         506,
		"HTTPERR_INSUFFICIENT_STORAGE":            507,
		"HTTPERR_LOOP_DETECTED":                   508,
		"HTTPERR_NOT_EXTENDED":                    510,
		"HTTPERR_NETWORK_AUTHENTICATION_REQUIRED": 511,
	}
)

func (x HttpErr) Enum() *HttpErr {
	p := new(HttpErr)
	*p = x
	return p
}

func (x HttpErr) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (HttpErr) Descriptor() protoreflect.EnumDescriptor {
	return file_httperr_proto_enumTypes[0].Descriptor()
}

func (HttpErr) Type() protoreflect.EnumType {
	return &file_httperr_proto_enumTypes[0]
}

func (x HttpErr) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use HttpErr.Descriptor instead.
func (HttpErr) EnumDescriptor() ([]byte, []int) {
	return file_httperr_proto_rawDescGZIP(), []int{0}
}

var File_httperr_proto protoreflect.FileDescriptor

var file_httperr_proto_rawDesc = []byte{
	0x0a, 0x0d, 0x68, 0x74, 0x74, 0x70, 0x65, 0x72, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x0b, 0x64, 0x65, 0x66, 0x69, 0x6e, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2a, 0xc1, 0x0a, 0x0a,
	0x07, 0x48, 0x74, 0x74, 0x70, 0x45, 0x72, 0x72, 0x12, 0x10, 0x0a, 0x0c, 0x48, 0x54, 0x54, 0x50,
	0x45, 0x52, 0x52, 0x5f, 0x4e, 0x4f, 0x4e, 0x45, 0x10, 0x00, 0x12, 0x14, 0x0a, 0x10, 0x48, 0x54,
	0x54, 0x50, 0x45, 0x52, 0x52, 0x5f, 0x49, 0x4e, 0x54, 0x45, 0x52, 0x4e, 0x41, 0x4c, 0x10, 0x01,
	0x12, 0x0f, 0x0a, 0x0a, 0x48, 0x54, 0x54, 0x50, 0x45, 0x52, 0x52, 0x5f, 0x4f, 0x4b, 0x10, 0xc8,
	0x01, 0x12, 0x18, 0x0a, 0x13, 0x48, 0x54, 0x54, 0x50, 0x45, 0x52, 0x52, 0x5f, 0x42, 0x41, 0x44,
	0x5f, 0x52, 0x45, 0x51, 0x55, 0x45, 0x53, 0x54, 0x10, 0x90, 0x03, 0x12, 0x19, 0x0a, 0x14, 0x48,
	0x54, 0x54, 0x50, 0x45, 0x52, 0x52, 0x5f, 0x55, 0x4e, 0x41, 0x55, 0x54, 0x48, 0x4f, 0x52, 0x49,
	0x5a, 0x45, 0x44, 0x10, 0x91, 0x03, 0x12, 0x1d, 0x0a, 0x18, 0x48, 0x54, 0x54, 0x50, 0x45, 0x52,
	0x52, 0x5f, 0x50, 0x41, 0x59, 0x4d, 0x45, 0x4e, 0x54, 0x5f, 0x52, 0x45, 0x51, 0x55, 0x49, 0x52,
	0x45, 0x44, 0x10, 0x92, 0x03, 0x12, 0x16, 0x0a, 0x11, 0x48, 0x54, 0x54, 0x50, 0x45, 0x52, 0x52,
	0x5f, 0x46, 0x4f, 0x52, 0x42, 0x49, 0x44, 0x44, 0x45, 0x4e, 0x10, 0x93, 0x03, 0x12, 0x16, 0x0a,
	0x11, 0x48, 0x54, 0x54, 0x50, 0x45, 0x52, 0x52, 0x5f, 0x4e, 0x4f, 0x54, 0x5f, 0x46, 0x4f, 0x55,
	0x4e, 0x44, 0x10, 0x94, 0x03, 0x12, 0x1f, 0x0a, 0x1a, 0x48, 0x54, 0x54, 0x50, 0x45, 0x52, 0x52,
	0x5f, 0x4d, 0x45, 0x54, 0x48, 0x4f, 0x44, 0x5f, 0x4e, 0x4f, 0x54, 0x5f, 0x41, 0x4c, 0x4c, 0x4f,
	0x57, 0x45, 0x44, 0x10, 0x95, 0x03, 0x12, 0x1b, 0x0a, 0x16, 0x48, 0x54, 0x54, 0x50, 0x45, 0x52,
	0x52, 0x5f, 0x4e, 0x4f, 0x54, 0x5f, 0x41, 0x43, 0x43, 0x45, 0x50, 0x54, 0x41, 0x42, 0x4c, 0x45,
	0x10, 0x96, 0x03, 0x12, 0x20, 0x0a, 0x1b, 0x48, 0x54, 0x54, 0x50, 0x45, 0x52, 0x52, 0x5f, 0x50,
	0x52, 0x4f, 0x58, 0x59, 0x5f, 0x41, 0x55, 0x54, 0x48, 0x5f, 0x52, 0x45, 0x51, 0x55, 0x49, 0x52,
	0x45, 0x44, 0x10, 0x97, 0x03, 0x12, 0x1c, 0x0a, 0x17, 0x48, 0x54, 0x54, 0x50, 0x45, 0x52, 0x52,
	0x5f, 0x52, 0x45, 0x51, 0x55, 0x45, 0x53, 0x54, 0x5f, 0x54, 0x49, 0x4d, 0x45, 0x4f, 0x55, 0x54,
	0x10, 0x98, 0x03, 0x12, 0x15, 0x0a, 0x10, 0x48, 0x54, 0x54, 0x50, 0x45, 0x52, 0x52, 0x5f, 0x43,
	0x4f, 0x4e, 0x46, 0x4c, 0x49, 0x43, 0x54, 0x10, 0x99, 0x03, 0x12, 0x11, 0x0a, 0x0c, 0x48, 0x54,
	0x54, 0x50, 0x45, 0x52, 0x52, 0x5f, 0x47, 0x4f, 0x4e, 0x45, 0x10, 0x9a, 0x03, 0x12, 0x1c, 0x0a,
	0x17, 0x48, 0x54, 0x54, 0x50, 0x45, 0x52, 0x52, 0x5f, 0x4c, 0x45, 0x4e, 0x47, 0x54, 0x48, 0x5f,
	0x52, 0x45, 0x51, 0x55, 0x49, 0x52, 0x45, 0x44, 0x10, 0x9b, 0x03, 0x12, 0x20, 0x0a, 0x1b, 0x48,
	0x54, 0x54, 0x50, 0x45, 0x52, 0x52, 0x5f, 0x50, 0x52, 0x45, 0x43, 0x4f, 0x4e, 0x44, 0x49, 0x54,
	0x49, 0x4f, 0x4e, 0x5f, 0x46, 0x41, 0x49, 0x4c, 0x45, 0x44, 0x10, 0x9c, 0x03, 0x12, 0x25, 0x0a,
	0x20, 0x48, 0x54, 0x54, 0x50, 0x45, 0x52, 0x52, 0x5f, 0x52, 0x45, 0x51, 0x55, 0x45, 0x53, 0x54,
	0x5f, 0x45, 0x4e, 0x54, 0x49, 0x54, 0x59, 0x5f, 0x54, 0x4f, 0x4f, 0x5f, 0x4c, 0x41, 0x52, 0x47,
	0x45, 0x10, 0x9d, 0x03, 0x12, 0x21, 0x0a, 0x1c, 0x48, 0x54, 0x54, 0x50, 0x45, 0x52, 0x52, 0x5f,
	0x52, 0x45, 0x51, 0x55, 0x45, 0x53, 0x54, 0x5f, 0x55, 0x52, 0x49, 0x5f, 0x54, 0x4f, 0x4f, 0x5f,
	0x4c, 0x4f, 0x4e, 0x47, 0x10, 0x9e, 0x03, 0x12, 0x21, 0x0a, 0x1c, 0x48, 0x54, 0x54, 0x50, 0x45,
	0x52, 0x52, 0x5f, 0x55, 0x4e, 0x53, 0x55, 0x50, 0x50, 0x4f, 0x52, 0x54, 0x45, 0x44, 0x4d, 0x45,
	0x44, 0x49, 0x41, 0x54, 0x59, 0x50, 0x45, 0x10, 0x9f, 0x03, 0x12, 0x2c, 0x0a, 0x27, 0x48, 0x54,
	0x54, 0x50, 0x45, 0x52, 0x52, 0x5f, 0x52, 0x45, 0x51, 0x55, 0x45, 0x53, 0x54, 0x45, 0x44, 0x5f,
	0x52, 0x41, 0x4e, 0x47, 0x45, 0x5f, 0x4e, 0x4f, 0x54, 0x5f, 0x53, 0x41, 0x54, 0x49, 0x53, 0x46,
	0x49, 0x41, 0x42, 0x4c, 0x45, 0x10, 0xa0, 0x03, 0x12, 0x1f, 0x0a, 0x1a, 0x48, 0x54, 0x54, 0x50,
	0x45, 0x52, 0x52, 0x5f, 0x45, 0x58, 0x50, 0x45, 0x43, 0x54, 0x41, 0x54, 0x49, 0x4f, 0x4e, 0x5f,
	0x46, 0x41, 0x49, 0x4c, 0x45, 0x44, 0x10, 0xa1, 0x03, 0x12, 0x13, 0x0a, 0x0e, 0x48, 0x54, 0x54,
	0x50, 0x45, 0x52, 0x52, 0x5f, 0x54, 0x45, 0x41, 0x50, 0x4f, 0x54, 0x10, 0xa2, 0x03, 0x12, 0x20,
	0x0a, 0x1b, 0x48, 0x54, 0x54, 0x50, 0x45, 0x52, 0x52, 0x5f, 0x4d, 0x49, 0x53, 0x44, 0x49, 0x52,
	0x45, 0x43, 0x54, 0x45, 0x44, 0x5f, 0x52, 0x45, 0x51, 0x55, 0x45, 0x53, 0x54, 0x10, 0xa5, 0x03,
	0x12, 0x21, 0x0a, 0x1c, 0x48, 0x54, 0x54, 0x50, 0x45, 0x52, 0x52, 0x5f, 0x55, 0x4e, 0x50, 0x52,
	0x4f, 0x43, 0x45, 0x53, 0x53, 0x41, 0x42, 0x4c, 0x45, 0x5f, 0x45, 0x4e, 0x54, 0x49, 0x54, 0x59,
	0x10, 0xa6, 0x03, 0x12, 0x13, 0x0a, 0x0e, 0x48, 0x54, 0x54, 0x50, 0x45, 0x52, 0x52, 0x5f, 0x4c,
	0x4f, 0x43, 0x4b, 0x45, 0x44, 0x10, 0xa7, 0x03, 0x12, 0x1e, 0x0a, 0x19, 0x48, 0x54, 0x54, 0x50,
	0x45, 0x52, 0x52, 0x5f, 0x46, 0x41, 0x49, 0x4c, 0x45, 0x44, 0x5f, 0x44, 0x45, 0x50, 0x45, 0x4e,
	0x44, 0x45, 0x4e, 0x43, 0x59, 0x10, 0xa8, 0x03, 0x12, 0x16, 0x0a, 0x11, 0x48, 0x54, 0x54, 0x50,
	0x45, 0x52, 0x52, 0x5f, 0x54, 0x4f, 0x4f, 0x5f, 0x45, 0x41, 0x52, 0x4c, 0x59, 0x10, 0xa9, 0x03,
	0x12, 0x1d, 0x0a, 0x18, 0x48, 0x54, 0x54, 0x50, 0x45, 0x52, 0x52, 0x5f, 0x55, 0x50, 0x47, 0x52,
	0x41, 0x44, 0x45, 0x5f, 0x52, 0x45, 0x51, 0x55, 0x49, 0x52, 0x45, 0x44, 0x10, 0xaa, 0x03, 0x12,
	0x22, 0x0a, 0x1d, 0x48, 0x54, 0x54, 0x50, 0x45, 0x52, 0x52, 0x5f, 0x50, 0x52, 0x45, 0x43, 0x4f,
	0x4e, 0x44, 0x49, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x52, 0x45, 0x51, 0x55, 0x49, 0x52, 0x45, 0x44,
	0x10, 0xac, 0x03, 0x12, 0x1e, 0x0a, 0x19, 0x48, 0x54, 0x54, 0x50, 0x45, 0x52, 0x52, 0x5f, 0x54,
	0x4f, 0x4f, 0x5f, 0x4d, 0x41, 0x4e, 0x59, 0x5f, 0x52, 0x45, 0x51, 0x55, 0x45, 0x53, 0x54, 0x53,
	0x10, 0xad, 0x03, 0x12, 0x2c, 0x0a, 0x27, 0x48, 0x54, 0x54, 0x50, 0x45, 0x52, 0x52, 0x5f, 0x52,
	0x45, 0x51, 0x55, 0x45, 0x53, 0x54, 0x5f, 0x48, 0x45, 0x41, 0x44, 0x45, 0x52, 0x5f, 0x46, 0x49,
	0x45, 0x4c, 0x44, 0x53, 0x5f, 0x54, 0x4f, 0x4f, 0x5f, 0x4c, 0x41, 0x52, 0x47, 0x45, 0x10, 0xaf,
	0x03, 0x12, 0x2a, 0x0a, 0x25, 0x48, 0x54, 0x54, 0x50, 0x45, 0x52, 0x52, 0x5f, 0x55, 0x4e, 0x41,
	0x56, 0x41, 0x49, 0x4c, 0x41, 0x42, 0x4c, 0x45, 0x5f, 0x46, 0x4f, 0x52, 0x5f, 0x4c, 0x45, 0x47,
	0x41, 0x4c, 0x5f, 0x52, 0x45, 0x41, 0x53, 0x4f, 0x4e, 0x53, 0x10, 0xc3, 0x03, 0x12, 0x22, 0x0a,
	0x1d, 0x48, 0x54, 0x54, 0x50, 0x45, 0x52, 0x52, 0x5f, 0x49, 0x4e, 0x54, 0x45, 0x52, 0x4e, 0x41,
	0x4c, 0x5f, 0x53, 0x45, 0x52, 0x56, 0x45, 0x52, 0x5f, 0x45, 0x52, 0x52, 0x4f, 0x52, 0x10, 0xf4,
	0x03, 0x12, 0x1c, 0x0a, 0x17, 0x48, 0x54, 0x54, 0x50, 0x45, 0x52, 0x52, 0x5f, 0x4e, 0x4f, 0x54,
	0x5f, 0x49, 0x4d, 0x50, 0x4c, 0x45, 0x4d, 0x45, 0x4e, 0x54, 0x45, 0x44, 0x10, 0xf5, 0x03, 0x12,
	0x18, 0x0a, 0x13, 0x48, 0x54, 0x54, 0x50, 0x45, 0x52, 0x52, 0x5f, 0x42, 0x41, 0x44, 0x5f, 0x47,
	0x41, 0x54, 0x45, 0x57, 0x41, 0x59, 0x10, 0xf6, 0x03, 0x12, 0x20, 0x0a, 0x1b, 0x48, 0x54, 0x54,
	0x50, 0x45, 0x52, 0x52, 0x5f, 0x53, 0x45, 0x52, 0x56, 0x49, 0x43, 0x45, 0x5f, 0x55, 0x4e, 0x41,
	0x56, 0x41, 0x49, 0x4c, 0x41, 0x42, 0x4c, 0x45, 0x10, 0xf7, 0x03, 0x12, 0x1c, 0x0a, 0x17, 0x48,
	0x54, 0x54, 0x50, 0x45, 0x52, 0x52, 0x5f, 0x47, 0x41, 0x54, 0x45, 0x57, 0x41, 0x59, 0x5f, 0x54,
	0x49, 0x4d, 0x45, 0x4f, 0x55, 0x54, 0x10, 0xf8, 0x03, 0x12, 0x27, 0x0a, 0x22, 0x48, 0x54, 0x54,
	0x50, 0x45, 0x52, 0x52, 0x5f, 0x48, 0x54, 0x54, 0x50, 0x5f, 0x56, 0x45, 0x52, 0x53, 0x49, 0x4f,
	0x4e, 0x5f, 0x4e, 0x4f, 0x54, 0x5f, 0x53, 0x55, 0x50, 0x50, 0x4f, 0x52, 0x54, 0x45, 0x44, 0x10,
	0xf9, 0x03, 0x12, 0x24, 0x0a, 0x1f, 0x48, 0x54, 0x54, 0x50, 0x45, 0x52, 0x52, 0x5f, 0x56, 0x41,
	0x52, 0x49, 0x41, 0x4e, 0x54, 0x5f, 0x41, 0x4c, 0x53, 0x4f, 0x5f, 0x4e, 0x45, 0x47, 0x4f, 0x54,
	0x49, 0x41, 0x54, 0x45, 0x53, 0x10, 0xfa, 0x03, 0x12, 0x21, 0x0a, 0x1c, 0x48, 0x54, 0x54, 0x50,
	0x45, 0x52, 0x52, 0x5f, 0x49, 0x4e, 0x53, 0x55, 0x46, 0x46, 0x49, 0x43, 0x49, 0x45, 0x4e, 0x54,
	0x5f, 0x53, 0x54, 0x4f, 0x52, 0x41, 0x47, 0x45, 0x10, 0xfb, 0x03, 0x12, 0x1a, 0x0a, 0x15, 0x48,
	0x54, 0x54, 0x50, 0x45, 0x52, 0x52, 0x5f, 0x4c, 0x4f, 0x4f, 0x50, 0x5f, 0x44, 0x45, 0x54, 0x45,
	0x43, 0x54, 0x45, 0x44, 0x10, 0xfc, 0x03, 0x12, 0x19, 0x0a, 0x14, 0x48, 0x54, 0x54, 0x50, 0x45,
	0x52, 0x52, 0x5f, 0x4e, 0x4f, 0x54, 0x5f, 0x45, 0x58, 0x54, 0x45, 0x4e, 0x44, 0x45, 0x44, 0x10,
	0xfe, 0x03, 0x12, 0x2c, 0x0a, 0x27, 0x48, 0x54, 0x54, 0x50, 0x45, 0x52, 0x52, 0x5f, 0x4e, 0x45,
	0x54, 0x57, 0x4f, 0x52, 0x4b, 0x5f, 0x41, 0x55, 0x54, 0x48, 0x45, 0x4e, 0x54, 0x49, 0x43, 0x41,
	0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x52, 0x45, 0x51, 0x55, 0x49, 0x52, 0x45, 0x44, 0x10, 0xff, 0x03,
	0x42, 0x10, 0x5a, 0x0e, 0x70, 0x6b, 0x67, 0x2f, 0x64, 0x61, 0x74, 0x61, 0x2f, 0x6d, 0x6f, 0x64,
	0x65, 0x6c, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_httperr_proto_rawDescOnce sync.Once
	file_httperr_proto_rawDescData = file_httperr_proto_rawDesc
)

func file_httperr_proto_rawDescGZIP() []byte {
	file_httperr_proto_rawDescOnce.Do(func() {
		file_httperr_proto_rawDescData = protoimpl.X.CompressGZIP(file_httperr_proto_rawDescData)
	})
	return file_httperr_proto_rawDescData
}

var file_httperr_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_httperr_proto_goTypes = []interface{}{
	(HttpErr)(0), // 0: definitions.HttpErr
}
var file_httperr_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_httperr_proto_init() }
func file_httperr_proto_init() {
	if File_httperr_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_httperr_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_httperr_proto_goTypes,
		DependencyIndexes: file_httperr_proto_depIdxs,
		EnumInfos:         file_httperr_proto_enumTypes,
	}.Build()
	File_httperr_proto = out.File
	file_httperr_proto_rawDesc = nil
	file_httperr_proto_goTypes = nil
	file_httperr_proto_depIdxs = nil
}
