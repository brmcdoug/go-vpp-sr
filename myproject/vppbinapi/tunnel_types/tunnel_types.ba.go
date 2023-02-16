// Code generated by GoVPP's binapi-generator. DO NOT EDIT.
// versions:
//  binapi-generator: v0.5.0
//  VPP:              23.02-rc0~230-g1d9780a43~b3131
// source: /usr/share/vpp/api/core/tunnel_types.api.json

// Package tunnel_types contains generated bindings for API file tunnel_types.api.
//
// Contents:
//
//	3 enums
//	1 struct
package tunnel_types

import (
	"strconv"

	api "git.fd.io/govpp.git/api"
	interface_types "mygit.com/myproject/vppbinapi/interface_types"
	ip_types "mygit.com/myproject/vppbinapi/ip_types"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the GoVPP api package it is being compiled against.
// A compilation error at this line likely means your copy of the
// GoVPP api package needs to be updated.
const _ = api.GoVppAPIPackageIsVersion2

const (
	APIFile    = "tunnel_types"
	APIVersion = "1.0.1"
	VersionCrc = 0x882f6758
)

// TunnelEncapDecapFlags defines enum 'tunnel_encap_decap_flags'.
type TunnelEncapDecapFlags uint8

const (
	TUNNEL_API_ENCAP_DECAP_FLAG_NONE                  TunnelEncapDecapFlags = 0
	TUNNEL_API_ENCAP_DECAP_FLAG_ENCAP_COPY_DF         TunnelEncapDecapFlags = 1
	TUNNEL_API_ENCAP_DECAP_FLAG_ENCAP_SET_DF          TunnelEncapDecapFlags = 2
	TUNNEL_API_ENCAP_DECAP_FLAG_ENCAP_COPY_DSCP       TunnelEncapDecapFlags = 4
	TUNNEL_API_ENCAP_DECAP_FLAG_ENCAP_COPY_ECN        TunnelEncapDecapFlags = 8
	TUNNEL_API_ENCAP_DECAP_FLAG_DECAP_COPY_ECN        TunnelEncapDecapFlags = 16
	TUNNEL_API_ENCAP_DECAP_FLAG_ENCAP_INNER_HASH      TunnelEncapDecapFlags = 32
	TUNNEL_API_ENCAP_DECAP_FLAG_ENCAP_COPY_HOP_LIMIT  TunnelEncapDecapFlags = 64
	TUNNEL_API_ENCAP_DECAP_FLAG_ENCAP_COPY_FLOW_LABEL TunnelEncapDecapFlags = 128
)

var (
	TunnelEncapDecapFlags_name = map[uint8]string{
		0:   "TUNNEL_API_ENCAP_DECAP_FLAG_NONE",
		1:   "TUNNEL_API_ENCAP_DECAP_FLAG_ENCAP_COPY_DF",
		2:   "TUNNEL_API_ENCAP_DECAP_FLAG_ENCAP_SET_DF",
		4:   "TUNNEL_API_ENCAP_DECAP_FLAG_ENCAP_COPY_DSCP",
		8:   "TUNNEL_API_ENCAP_DECAP_FLAG_ENCAP_COPY_ECN",
		16:  "TUNNEL_API_ENCAP_DECAP_FLAG_DECAP_COPY_ECN",
		32:  "TUNNEL_API_ENCAP_DECAP_FLAG_ENCAP_INNER_HASH",
		64:  "TUNNEL_API_ENCAP_DECAP_FLAG_ENCAP_COPY_HOP_LIMIT",
		128: "TUNNEL_API_ENCAP_DECAP_FLAG_ENCAP_COPY_FLOW_LABEL",
	}
	TunnelEncapDecapFlags_value = map[string]uint8{
		"TUNNEL_API_ENCAP_DECAP_FLAG_NONE":                  0,
		"TUNNEL_API_ENCAP_DECAP_FLAG_ENCAP_COPY_DF":         1,
		"TUNNEL_API_ENCAP_DECAP_FLAG_ENCAP_SET_DF":          2,
		"TUNNEL_API_ENCAP_DECAP_FLAG_ENCAP_COPY_DSCP":       4,
		"TUNNEL_API_ENCAP_DECAP_FLAG_ENCAP_COPY_ECN":        8,
		"TUNNEL_API_ENCAP_DECAP_FLAG_DECAP_COPY_ECN":        16,
		"TUNNEL_API_ENCAP_DECAP_FLAG_ENCAP_INNER_HASH":      32,
		"TUNNEL_API_ENCAP_DECAP_FLAG_ENCAP_COPY_HOP_LIMIT":  64,
		"TUNNEL_API_ENCAP_DECAP_FLAG_ENCAP_COPY_FLOW_LABEL": 128,
	}
)

func (x TunnelEncapDecapFlags) String() string {
	s, ok := TunnelEncapDecapFlags_name[uint8(x)]
	if ok {
		return s
	}
	str := func(n uint8) string {
		s, ok := TunnelEncapDecapFlags_name[uint8(n)]
		if ok {
			return s
		}
		return "TunnelEncapDecapFlags(" + strconv.Itoa(int(n)) + ")"
	}
	for i := uint8(0); i <= 8; i++ {
		val := uint8(x)
		if val&(1<<i) != 0 {
			if s != "" {
				s += "|"
			}
			s += str(1 << i)
		}
	}
	if s == "" {
		return str(uint8(x))
	}
	return s
}

// TunnelMode defines enum 'tunnel_mode'.
type TunnelMode uint8

const (
	TUNNEL_API_MODE_P2P TunnelMode = 0
	TUNNEL_API_MODE_MP  TunnelMode = 1
)

var (
	TunnelMode_name = map[uint8]string{
		0: "TUNNEL_API_MODE_P2P",
		1: "TUNNEL_API_MODE_MP",
	}
	TunnelMode_value = map[string]uint8{
		"TUNNEL_API_MODE_P2P": 0,
		"TUNNEL_API_MODE_MP":  1,
	}
)

func (x TunnelMode) String() string {
	s, ok := TunnelMode_name[uint8(x)]
	if ok {
		return s
	}
	return "TunnelMode(" + strconv.Itoa(int(x)) + ")"
}

// TunnelFlags defines enum 'tunnel_flags'.
type TunnelFlags uint8

const (
	TUNNEL_API_FLAG_TRACK_MTU TunnelFlags = 1
)

var (
	TunnelFlags_name = map[uint8]string{
		1: "TUNNEL_API_FLAG_TRACK_MTU",
	}
	TunnelFlags_value = map[string]uint8{
		"TUNNEL_API_FLAG_TRACK_MTU": 1,
	}
)

func (x TunnelFlags) String() string {
	s, ok := TunnelFlags_name[uint8(x)]
	if ok {
		return s
	}
	str := func(n uint8) string {
		s, ok := TunnelFlags_name[uint8(n)]
		if ok {
			return s
		}
		return "TunnelFlags(" + strconv.Itoa(int(n)) + ")"
	}
	for i := uint8(0); i <= 8; i++ {
		val := uint8(x)
		if val&(1<<i) != 0 {
			if s != "" {
				s += "|"
			}
			s += str(1 << i)
		}
	}
	if s == "" {
		return str(uint8(x))
	}
	return s
}

// Tunnel defines type 'tunnel'.
type Tunnel struct {
	Instance        uint32                         `binapi:"u32,name=instance" json:"instance,omitempty"`
	Src             ip_types.Address               `binapi:"address,name=src" json:"src,omitempty"`
	Dst             ip_types.Address               `binapi:"address,name=dst" json:"dst,omitempty"`
	SwIfIndex       interface_types.InterfaceIndex `binapi:"interface_index,name=sw_if_index" json:"sw_if_index,omitempty"`
	TableID         uint32                         `binapi:"u32,name=table_id" json:"table_id,omitempty"`
	EncapDecapFlags TunnelEncapDecapFlags          `binapi:"tunnel_encap_decap_flags,name=encap_decap_flags" json:"encap_decap_flags,omitempty"`
	Mode            TunnelMode                     `binapi:"tunnel_mode,name=mode" json:"mode,omitempty"`
	Flags           TunnelFlags                    `binapi:"tunnel_flags,name=flags" json:"flags,omitempty"`
	Dscp            ip_types.IPDscp                `binapi:"ip_dscp,name=dscp" json:"dscp,omitempty"`
	HopLimit        uint8                          `binapi:"u8,name=hop_limit" json:"hop_limit,omitempty"`
}
