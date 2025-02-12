// This is generated code. Should not be manually modified.

package omega

import (
	"time"
	"github.com/zeusyf/btcd/chaincfg/chainhash"
	"github.com/zeusyf/btcd/wire"
	"github.com/zeusyf/omega/token"
)

var IntlDateLine = [][2]float64 {	// international date line
	{ 90.000000, 180.000000 },
	{ 75.000000, 180.000000 },
	{ 68.245600, -169.000000 },
	{ 65.518900, -169.000000 },
	{ 53.086300, 170.050000 },
	{ 47.835300, 180.000000 },
	{ -1.200000, 180.000000 },
	{ -1.200000, -159.650000 },
	{ 2.900000, -159.650000 },
	{ 2.900000, -162.850000 },
	{ 6.500000, -162.850000 },
	{ 6.500000, -155.950000 },
	{ -9.500000, -149.650000 },
	{ -11.700000, -149.650000 },
	{ -11.700000, -154.050000 },
	{ -10.700000, -154.050000 },
	{ -10.700000, -166.550000 },
	{ -15.600000, -172.700000 },
	{ -45.000000, -172.700000 },
	{ -51.181500, 180.000000 },
	{ -90.000000, 180.000000 },
}

var InitDefs = []token.Definition{
	&token.BorderDef {
		Father: chainhash.Hash{},
		Begin: * token.NewVertexDef(188743680, -377487360, 0),
		End: * token.NewVertexDef(-188743680, -377487360, 0),
	},
	&token.BorderDef {
		Father: chainhash.Hash{},
		Begin: * token.NewVertexDef(-188743680, -377487360, 0),
		End: * token.NewVertexDef(-188743680, 377487360, 0),
	},
	&token.BorderDef {
		Father: chainhash.Hash{},
		Begin: * token.NewVertexDef(-188743680, 377487360, 0),
		End: * token.NewVertexDef(188743680, 377487360, 0),
	},
	&token.BorderDef {
		Father: chainhash.Hash{},
		Begin: * token.NewVertexDef(188743680, 377487360, 0),
		End: * token.NewVertexDef(188743680, -377487360, 0),
	},
	&token.BorderDef {
		Father: chainhash.Hash{
		0x36, 0x10, 0xa4, 0xdf, 0xe7, 0xbc, 0xb2, 0x5c, 
		0x0e, 0xa5, 0xfc, 0x6f, 0xb2, 0x33, 0x9b, 0x23, 
		0x65, 0x7a, 0x55, 0x3d, 0x6b, 0x81, 0x0f, 0xac, 
		0x76, 0x31, 0xa2, 0x77, 0x6a, 0x23, 0x16, 0x8a, 
	},
		Begin: * token.NewVertexDef(188743680, -377487360, 0),
		End: * token.NewVertexDef(157286400, -377487360, 0),
	},
	&token.BorderDef {
		Father: chainhash.Hash{
		0x96, 0x3d, 0xdc, 0x10, 0xa8, 0x36, 0x68, 0xe5, 
		0x0b, 0xfb, 0xdd, 0x0b, 0xf5, 0x1a, 0xc8, 0x69, 
		0x7d, 0x92, 0xdb, 0x26, 0x92, 0x69, 0x56, 0x3f, 
		0x49, 0x0c, 0xfc, 0x0a, 0x22, 0xf1, 0x90, 0x92, 
	},
		Begin: * token.NewVertexDef(-188743680, 377487360, 0),
		End: * token.NewVertexDef(-107335385, 377487360, 0),
	},
	&token.BorderDef {
		Father: chainhash.Hash{
		0x36, 0x10, 0xa4, 0xdf, 0xe7, 0xbc, 0xb2, 0x5c, 
		0x0e, 0xa5, 0xfc, 0x6f, 0xb2, 0x33, 0x9b, 0x23, 
		0x65, 0x7a, 0x55, 0x3d, 0x6b, 0x81, 0x0f, 0xac, 
		0x76, 0x31, 0xa2, 0x77, 0x6a, 0x23, 0x16, 0x8a, 
	},
		Begin: * token.NewVertexDef(157286400, -377487360, 0),
		End: * token.NewVertexDef(143121396, -354418688, 0),
	},
	&token.BorderDef {
		Father: chainhash.Hash{
		0x96, 0x3d, 0xdc, 0x10, 0xa8, 0x36, 0x68, 0xe5, 
		0x0b, 0xfb, 0xdd, 0x0b, 0xf5, 0x1a, 0xc8, 0x69, 
		0x7d, 0x92, 0xdb, 0x26, 0x92, 0x69, 0x56, 0x3f, 
		0x49, 0x0c, 0xfc, 0x0a, 0x22, 0xf1, 0x90, 0x92, 
	},
		Begin: * token.NewVertexDef(-107335385, 377487360, 0),
		End: * token.NewVertexDef(-94371840, 392796569, 0),
	},
	&token.BorderDef {
		Father: chainhash.Hash{
		0x36, 0x10, 0xa4, 0xdf, 0xe7, 0xbc, 0xb2, 0x5c, 
		0x0e, 0xa5, 0xfc, 0x6f, 0xb2, 0x33, 0x9b, 0x23, 
		0x65, 0x7a, 0x55, 0x3d, 0x6b, 0x81, 0x0f, 0xac, 
		0x76, 0x31, 0xa2, 0x77, 0x6a, 0x23, 0x16, 0x8a, 
	},
		Begin: * token.NewVertexDef(143121396, -354418688, 0),
		End: * token.NewVertexDef(137403092, -354418688, 0),
	},
	&token.BorderDef {
		Father: chainhash.Hash{
		0x96, 0x3d, 0xdc, 0x10, 0xa8, 0x36, 0x68, 0xe5, 
		0x0b, 0xfb, 0xdd, 0x0b, 0xf5, 0x1a, 0xc8, 0x69, 
		0x7d, 0x92, 0xdb, 0x26, 0x92, 0x69, 0x56, 0x3f, 
		0x49, 0x0c, 0xfc, 0x0a, 0x22, 0xf1, 0x90, 0x92, 
	},
		Begin: * token.NewVertexDef(-94371840, 392796569, 0),
		End: * token.NewVertexDef(-32715571, 392796569, 0),
	},
	&token.BorderDef {
		Father: chainhash.Hash{
		0x36, 0x10, 0xa4, 0xdf, 0xe7, 0xbc, 0xb2, 0x5c, 
		0x0e, 0xa5, 0xfc, 0x6f, 0xb2, 0x33, 0x9b, 0x23, 
		0x65, 0x7a, 0x55, 0x3d, 0x6b, 0x81, 0x0f, 0xac, 
		0x76, 0x31, 0xa2, 0x77, 0x6a, 0x23, 0x16, 0x8a, 
	},
		Begin: * token.NewVertexDef(137403092, -354418688, 0),
		End: * token.NewVertexDef(111330040, -398354022, 0),
	},
	&token.BorderDef {
		Father: chainhash.Hash{
		0x96, 0x3d, 0xdc, 0x10, 0xa8, 0x36, 0x68, 0xe5, 
		0x0b, 0xfb, 0xdd, 0x0b, 0xf5, 0x1a, 0xc8, 0x69, 
		0x7d, 0x92, 0xdb, 0x26, 0x92, 0x69, 0x56, 0x3f, 
		0x49, 0x0c, 0xfc, 0x0a, 0x22, 0xf1, 0x90, 0x92, 
	},
		Begin: * token.NewVertexDef(-32715571, 392796569, 0),
		End: * token.NewVertexDef(-22439526, 405694054, 0),
	},
	&token.BorderDef {
		Father: chainhash.Hash{
		0x36, 0x10, 0xa4, 0xdf, 0xe7, 0xbc, 0xb2, 0x5c, 
		0x0e, 0xa5, 0xfc, 0x6f, 0xb2, 0x33, 0x9b, 0x23, 
		0x65, 0x7a, 0x55, 0x3d, 0x6b, 0x81, 0x0f, 0xac, 
		0x76, 0x31, 0xa2, 0x77, 0x6a, 0x23, 0x16, 0x8a, 
	},
		Begin: * token.NewVertexDef(111330040, -398354022, 0),
		End: * token.NewVertexDef(100317895, -377487360, 0),
	},
	&token.BorderDef {
		Father: chainhash.Hash{
		0x96, 0x3d, 0xdc, 0x10, 0xa8, 0x36, 0x68, 0xe5, 
		0x0b, 0xfb, 0xdd, 0x0b, 0xf5, 0x1a, 0xc8, 0x69, 
		0x7d, 0x92, 0xdb, 0x26, 0x92, 0x69, 0x56, 0x3f, 
		0x49, 0x0c, 0xfc, 0x0a, 0x22, 0xf1, 0x90, 0x92, 
	},
		Begin: * token.NewVertexDef(-22439526, 405694054, 0),
		End: * token.NewVertexDef(-22439526, 431908454, 0),
	},
	&token.BorderDef {
		Father: chainhash.Hash{
		0x36, 0x10, 0xa4, 0xdf, 0xe7, 0xbc, 0xb2, 0x5c, 
		0x0e, 0xa5, 0xfc, 0x6f, 0xb2, 0x33, 0x9b, 0x23, 
		0x65, 0x7a, 0x55, 0x3d, 0x6b, 0x81, 0x0f, 0xac, 
		0x76, 0x31, 0xa2, 0x77, 0x6a, 0x23, 0x16, 0x8a, 
	},
		Begin: * token.NewVertexDef(100317895, -377487360, 0),
		End: * token.NewVertexDef(-2516582, -377487360, 0),
	},
	&token.BorderDef {
		Father: chainhash.Hash{
		0x96, 0x3d, 0xdc, 0x10, 0xa8, 0x36, 0x68, 0xe5, 
		0x0b, 0xfb, 0xdd, 0x0b, 0xf5, 0x1a, 0xc8, 0x69, 
		0x7d, 0x92, 0xdb, 0x26, 0x92, 0x69, 0x56, 0x3f, 
		0x49, 0x0c, 0xfc, 0x0a, 0x22, 0xf1, 0x90, 0x92, 
	},
		Begin: * token.NewVertexDef(-22439526, 431908454, 0),
		End: * token.NewVertexDef(-24536678, 431908454, 0),
	},
	&token.BorderDef {
		Father: chainhash.Hash{
		0x36, 0x10, 0xa4, 0xdf, 0xe7, 0xbc, 0xb2, 0x5c, 
		0x0e, 0xa5, 0xfc, 0x6f, 0xb2, 0x33, 0x9b, 0x23, 
		0x65, 0x7a, 0x55, 0x3d, 0x6b, 0x81, 0x0f, 0xac, 
		0x76, 0x31, 0xa2, 0x77, 0x6a, 0x23, 0x16, 0x8a, 
	},
		Begin: * token.NewVertexDef(-2516582, -377487360, 0),
		End: * token.NewVertexDef(-2516582, -334810316, 0),
	},
	&token.BorderDef {
		Father: chainhash.Hash{
		0x96, 0x3d, 0xdc, 0x10, 0xa8, 0x36, 0x68, 0xe5, 
		0x0b, 0xfb, 0xdd, 0x0b, 0xf5, 0x1a, 0xc8, 0x69, 
		0x7d, 0x92, 0xdb, 0x26, 0x92, 0x69, 0x56, 0x3f, 
		0x49, 0x0c, 0xfc, 0x0a, 0x22, 0xf1, 0x90, 0x92, 
	},
		Begin: * token.NewVertexDef(-24536678, 431908454, 0),
		End: * token.NewVertexDef(-24536678, 441135923, 0),
	},
	&token.BorderDef {
		Father: chainhash.Hash{
		0x36, 0x10, 0xa4, 0xdf, 0xe7, 0xbc, 0xb2, 0x5c, 
		0x0e, 0xa5, 0xfc, 0x6f, 0xb2, 0x33, 0x9b, 0x23, 
		0x65, 0x7a, 0x55, 0x3d, 0x6b, 0x81, 0x0f, 0xac, 
		0x76, 0x31, 0xa2, 0x77, 0x6a, 0x23, 0x16, 0x8a, 
	},
		Begin: * token.NewVertexDef(-2516582, -334810316, 0),
		End: * token.NewVertexDef(6081740, -334810316, 0),
	},
	&token.BorderDef {
		Father: chainhash.Hash{
		0x96, 0x3d, 0xdc, 0x10, 0xa8, 0x36, 0x68, 0xe5, 
		0x0b, 0xfb, 0xdd, 0x0b, 0xf5, 0x1a, 0xc8, 0x69, 
		0x7d, 0x92, 0xdb, 0x26, 0x92, 0x69, 0x56, 0x3f, 
		0x49, 0x0c, 0xfc, 0x0a, 0x22, 0xf1, 0x90, 0x92, 
	},
		Begin: * token.NewVertexDef(-24536678, 441135923, 0),
		End: * token.NewVertexDef(-19922944, 441135923, 0),
	},
	&token.BorderDef {
		Father: chainhash.Hash{
		0x36, 0x10, 0xa4, 0xdf, 0xe7, 0xbc, 0xb2, 0x5c, 
		0x0e, 0xa5, 0xfc, 0x6f, 0xb2, 0x33, 0x9b, 0x23, 
		0x65, 0x7a, 0x55, 0x3d, 0x6b, 0x81, 0x0f, 0xac, 
		0x76, 0x31, 0xa2, 0x77, 0x6a, 0x23, 0x16, 0x8a, 
	},
		Begin: * token.NewVertexDef(6081740, -334810316, 0),
		End: * token.NewVertexDef(6081740, -341521203, 0),
	},
	&token.BorderDef {
		Father: chainhash.Hash{
		0x96, 0x3d, 0xdc, 0x10, 0xa8, 0x36, 0x68, 0xe5, 
		0x0b, 0xfb, 0xdd, 0x0b, 0xf5, 0x1a, 0xc8, 0x69, 
		0x7d, 0x92, 0xdb, 0x26, 0x92, 0x69, 0x56, 0x3f, 
		0x49, 0x0c, 0xfc, 0x0a, 0x22, 0xf1, 0x90, 0x92, 
	},
		Begin: * token.NewVertexDef(-19922944, 441135923, 0),
		End: * token.NewVertexDef(13631488, 427923865, 0),
	},
	&token.BorderDef {
		Father: chainhash.Hash{
		0x36, 0x10, 0xa4, 0xdf, 0xe7, 0xbc, 0xb2, 0x5c, 
		0x0e, 0xa5, 0xfc, 0x6f, 0xb2, 0x33, 0x9b, 0x23, 
		0x65, 0x7a, 0x55, 0x3d, 0x6b, 0x81, 0x0f, 0xac, 
		0x76, 0x31, 0xa2, 0x77, 0x6a, 0x23, 0x16, 0x8a, 
	},
		Begin: * token.NewVertexDef(6081740, -341521203, 0),
		End: * token.NewVertexDef(13631488, -341521203, 0),
	},
	&token.BorderDef {
		Father: chainhash.Hash{
		0x96, 0x3d, 0xdc, 0x10, 0xa8, 0x36, 0x68, 0xe5, 
		0x0b, 0xfb, 0xdd, 0x0b, 0xf5, 0x1a, 0xc8, 0x69, 
		0x7d, 0x92, 0xdb, 0x26, 0x92, 0x69, 0x56, 0x3f, 
		0x49, 0x0c, 0xfc, 0x0a, 0x22, 0xf1, 0x90, 0x92, 
	},
		Begin: * token.NewVertexDef(13631488, 427923865, 0),
		End: * token.NewVertexDef(13631488, 413453516, 0),
	},
	&token.BorderDef {
		Father: chainhash.Hash{
		0x36, 0x10, 0xa4, 0xdf, 0xe7, 0xbc, 0xb2, 0x5c, 
		0x0e, 0xa5, 0xfc, 0x6f, 0xb2, 0x33, 0x9b, 0x23, 
		0x65, 0x7a, 0x55, 0x3d, 0x6b, 0x81, 0x0f, 0xac, 
		0x76, 0x31, 0xa2, 0x77, 0x6a, 0x23, 0x16, 0x8a, 
	},
		Begin: * token.NewVertexDef(13631488, -341521203, 0),
		End: * token.NewVertexDef(13631488, -327050854, 0),
	},
	&token.BorderDef {
		Father: chainhash.Hash{
		0x96, 0x3d, 0xdc, 0x10, 0xa8, 0x36, 0x68, 0xe5, 
		0x0b, 0xfb, 0xdd, 0x0b, 0xf5, 0x1a, 0xc8, 0x69, 
		0x7d, 0x92, 0xdb, 0x26, 0x92, 0x69, 0x56, 0x3f, 
		0x49, 0x0c, 0xfc, 0x0a, 0x22, 0xf1, 0x90, 0x92, 
	},
		Begin: * token.NewVertexDef(13631488, 413453516, 0),
		End: * token.NewVertexDef(6081740, 413453516, 0),
	},
	&token.BorderDef {
		Father: chainhash.Hash{
		0x36, 0x10, 0xa4, 0xdf, 0xe7, 0xbc, 0xb2, 0x5c, 
		0x0e, 0xa5, 0xfc, 0x6f, 0xb2, 0x33, 0x9b, 0x23, 
		0x65, 0x7a, 0x55, 0x3d, 0x6b, 0x81, 0x0f, 0xac, 
		0x76, 0x31, 0xa2, 0x77, 0x6a, 0x23, 0x16, 0x8a, 
	},
		Begin: * token.NewVertexDef(13631488, -327050854, 0),
		End: * token.NewVertexDef(-19922944, -313838796, 0),
	},
	&token.BorderDef {
		Father: chainhash.Hash{
		0x96, 0x3d, 0xdc, 0x10, 0xa8, 0x36, 0x68, 0xe5, 
		0x0b, 0xfb, 0xdd, 0x0b, 0xf5, 0x1a, 0xc8, 0x69, 
		0x7d, 0x92, 0xdb, 0x26, 0x92, 0x69, 0x56, 0x3f, 
		0x49, 0x0c, 0xfc, 0x0a, 0x22, 0xf1, 0x90, 0x92, 
	},
		Begin: * token.NewVertexDef(6081740, 413453516, 0),
		End: * token.NewVertexDef(6081740, 420164403, 0),
	},
	&token.BorderDef {
		Father: chainhash.Hash{
		0x36, 0x10, 0xa4, 0xdf, 0xe7, 0xbc, 0xb2, 0x5c, 
		0x0e, 0xa5, 0xfc, 0x6f, 0xb2, 0x33, 0x9b, 0x23, 
		0x65, 0x7a, 0x55, 0x3d, 0x6b, 0x81, 0x0f, 0xac, 
		0x76, 0x31, 0xa2, 0x77, 0x6a, 0x23, 0x16, 0x8a, 
	},
		Begin: * token.NewVertexDef(-19922944, -313838796, 0),
		End: * token.NewVertexDef(-24536678, -313838796, 0),
	},
	&token.BorderDef {
		Father: chainhash.Hash{
		0x96, 0x3d, 0xdc, 0x10, 0xa8, 0x36, 0x68, 0xe5, 
		0x0b, 0xfb, 0xdd, 0x0b, 0xf5, 0x1a, 0xc8, 0x69, 
		0x7d, 0x92, 0xdb, 0x26, 0x92, 0x69, 0x56, 0x3f, 
		0x49, 0x0c, 0xfc, 0x0a, 0x22, 0xf1, 0x90, 0x92, 
	},
		Begin: * token.NewVertexDef(6081740, 420164403, 0),
		End: * token.NewVertexDef(-2516582, 420164403, 0),
	},
	&token.BorderDef {
		Father: chainhash.Hash{
		0x36, 0x10, 0xa4, 0xdf, 0xe7, 0xbc, 0xb2, 0x5c, 
		0x0e, 0xa5, 0xfc, 0x6f, 0xb2, 0x33, 0x9b, 0x23, 
		0x65, 0x7a, 0x55, 0x3d, 0x6b, 0x81, 0x0f, 0xac, 
		0x76, 0x31, 0xa2, 0x77, 0x6a, 0x23, 0x16, 0x8a, 
	},
		Begin: * token.NewVertexDef(-24536678, -313838796, 0),
		End: * token.NewVertexDef(-24536678, -323066265, 0),
	},
	&token.BorderDef {
		Father: chainhash.Hash{
		0x96, 0x3d, 0xdc, 0x10, 0xa8, 0x36, 0x68, 0xe5, 
		0x0b, 0xfb, 0xdd, 0x0b, 0xf5, 0x1a, 0xc8, 0x69, 
		0x7d, 0x92, 0xdb, 0x26, 0x92, 0x69, 0x56, 0x3f, 
		0x49, 0x0c, 0xfc, 0x0a, 0x22, 0xf1, 0x90, 0x92, 
	},
		Begin: * token.NewVertexDef(-2516582, 420164403, 0),
		End: * token.NewVertexDef(-2516582, 377487360, 0),
	},
	&token.BorderDef {
		Father: chainhash.Hash{
		0x36, 0x10, 0xa4, 0xdf, 0xe7, 0xbc, 0xb2, 0x5c, 
		0x0e, 0xa5, 0xfc, 0x6f, 0xb2, 0x33, 0x9b, 0x23, 
		0x65, 0x7a, 0x55, 0x3d, 0x6b, 0x81, 0x0f, 0xac, 
		0x76, 0x31, 0xa2, 0x77, 0x6a, 0x23, 0x16, 0x8a, 
	},
		Begin: * token.NewVertexDef(-24536678, -323066265, 0),
		End: * token.NewVertexDef(-22439526, -323066265, 0),
	},
	&token.BorderDef {
		Father: chainhash.Hash{
		0x96, 0x3d, 0xdc, 0x10, 0xa8, 0x36, 0x68, 0xe5, 
		0x0b, 0xfb, 0xdd, 0x0b, 0xf5, 0x1a, 0xc8, 0x69, 
		0x7d, 0x92, 0xdb, 0x26, 0x92, 0x69, 0x56, 0x3f, 
		0x49, 0x0c, 0xfc, 0x0a, 0x22, 0xf1, 0x90, 0x92, 
	},
		Begin: * token.NewVertexDef(-2516582, 377487360, 0),
		End: * token.NewVertexDef(100317895, 377487360, 0),
	},
	&token.BorderDef {
		Father: chainhash.Hash{
		0x36, 0x10, 0xa4, 0xdf, 0xe7, 0xbc, 0xb2, 0x5c, 
		0x0e, 0xa5, 0xfc, 0x6f, 0xb2, 0x33, 0x9b, 0x23, 
		0x65, 0x7a, 0x55, 0x3d, 0x6b, 0x81, 0x0f, 0xac, 
		0x76, 0x31, 0xa2, 0x77, 0x6a, 0x23, 0x16, 0x8a, 
	},
		Begin: * token.NewVertexDef(-22439526, -323066265, 0),
		End: * token.NewVertexDef(-22439526, -349280665, 0),
	},
	&token.BorderDef {
		Father: chainhash.Hash{
		0x96, 0x3d, 0xdc, 0x10, 0xa8, 0x36, 0x68, 0xe5, 
		0x0b, 0xfb, 0xdd, 0x0b, 0xf5, 0x1a, 0xc8, 0x69, 
		0x7d, 0x92, 0xdb, 0x26, 0x92, 0x69, 0x56, 0x3f, 
		0x49, 0x0c, 0xfc, 0x0a, 0x22, 0xf1, 0x90, 0x92, 
	},
		Begin: * token.NewVertexDef(100317895, 377487360, 0),
		End: * token.NewVertexDef(111330040, 356620697, 0),
	},
	&token.BorderDef {
		Father: chainhash.Hash{
		0x36, 0x10, 0xa4, 0xdf, 0xe7, 0xbc, 0xb2, 0x5c, 
		0x0e, 0xa5, 0xfc, 0x6f, 0xb2, 0x33, 0x9b, 0x23, 
		0x65, 0x7a, 0x55, 0x3d, 0x6b, 0x81, 0x0f, 0xac, 
		0x76, 0x31, 0xa2, 0x77, 0x6a, 0x23, 0x16, 0x8a, 
	},
		Begin: * token.NewVertexDef(-22439526, -349280665, 0),
		End: * token.NewVertexDef(-32715571, -362178150, 0),
	},
	&token.BorderDef {
		Father: chainhash.Hash{
		0x96, 0x3d, 0xdc, 0x10, 0xa8, 0x36, 0x68, 0xe5, 
		0x0b, 0xfb, 0xdd, 0x0b, 0xf5, 0x1a, 0xc8, 0x69, 
		0x7d, 0x92, 0xdb, 0x26, 0x92, 0x69, 0x56, 0x3f, 
		0x49, 0x0c, 0xfc, 0x0a, 0x22, 0xf1, 0x90, 0x92, 
	},
		Begin: * token.NewVertexDef(111330040, 356620697, 0),
		End: * token.NewVertexDef(137403092, 400556032, 0),
	},
	&token.BorderDef {
		Father: chainhash.Hash{
		0x36, 0x10, 0xa4, 0xdf, 0xe7, 0xbc, 0xb2, 0x5c, 
		0x0e, 0xa5, 0xfc, 0x6f, 0xb2, 0x33, 0x9b, 0x23, 
		0x65, 0x7a, 0x55, 0x3d, 0x6b, 0x81, 0x0f, 0xac, 
		0x76, 0x31, 0xa2, 0x77, 0x6a, 0x23, 0x16, 0x8a, 
	},
		Begin: * token.NewVertexDef(-32715571, -362178150, 0),
		End: * token.NewVertexDef(-94371840, -362178150, 0),
	},
	&token.BorderDef {
		Father: chainhash.Hash{
		0x96, 0x3d, 0xdc, 0x10, 0xa8, 0x36, 0x68, 0xe5, 
		0x0b, 0xfb, 0xdd, 0x0b, 0xf5, 0x1a, 0xc8, 0x69, 
		0x7d, 0x92, 0xdb, 0x26, 0x92, 0x69, 0x56, 0x3f, 
		0x49, 0x0c, 0xfc, 0x0a, 0x22, 0xf1, 0x90, 0x92, 
	},
		Begin: * token.NewVertexDef(137403092, 400556032, 0),
		End: * token.NewVertexDef(143121396, 400556032, 0),
	},
	&token.BorderDef {
		Father: chainhash.Hash{
		0x36, 0x10, 0xa4, 0xdf, 0xe7, 0xbc, 0xb2, 0x5c, 
		0x0e, 0xa5, 0xfc, 0x6f, 0xb2, 0x33, 0x9b, 0x23, 
		0x65, 0x7a, 0x55, 0x3d, 0x6b, 0x81, 0x0f, 0xac, 
		0x76, 0x31, 0xa2, 0x77, 0x6a, 0x23, 0x16, 0x8a, 
	},
		Begin: * token.NewVertexDef(-94371840, -362178150, 0),
		End: * token.NewVertexDef(-107335385, -377487360, 0),
	},
	&token.BorderDef {
		Father: chainhash.Hash{
		0x96, 0x3d, 0xdc, 0x10, 0xa8, 0x36, 0x68, 0xe5, 
		0x0b, 0xfb, 0xdd, 0x0b, 0xf5, 0x1a, 0xc8, 0x69, 
		0x7d, 0x92, 0xdb, 0x26, 0x92, 0x69, 0x56, 0x3f, 
		0x49, 0x0c, 0xfc, 0x0a, 0x22, 0xf1, 0x90, 0x92, 
	},
		Begin: * token.NewVertexDef(143121396, 400556032, 0),
		End: * token.NewVertexDef(157286400, 377487360, 0),
	},
	&token.BorderDef {
		Father: chainhash.Hash{
		0x36, 0x10, 0xa4, 0xdf, 0xe7, 0xbc, 0xb2, 0x5c, 
		0x0e, 0xa5, 0xfc, 0x6f, 0xb2, 0x33, 0x9b, 0x23, 
		0x65, 0x7a, 0x55, 0x3d, 0x6b, 0x81, 0x0f, 0xac, 
		0x76, 0x31, 0xa2, 0x77, 0x6a, 0x23, 0x16, 0x8a, 
	},
		Begin: * token.NewVertexDef(-107335385, -377487360, 0),
		End: * token.NewVertexDef(-188743680, -377487360, 0),
	},
	&token.BorderDef {
		Father: chainhash.Hash{
		0x96, 0x3d, 0xdc, 0x10, 0xa8, 0x36, 0x68, 0xe5, 
		0x0b, 0xfb, 0xdd, 0x0b, 0xf5, 0x1a, 0xc8, 0x69, 
		0x7d, 0x92, 0xdb, 0x26, 0x92, 0x69, 0x56, 0x3f, 
		0x49, 0x0c, 0xfc, 0x0a, 0x22, 0xf1, 0x90, 0x92, 
	},
		Begin: * token.NewVertexDef(157286400, 377487360, 0),
		End: * token.NewVertexDef(188743680, 377487360, 0),
	},
	&token.PolygonDef {	Loops: []token.LoopDef{{	// Loop 0:
		chainhash.Hash{
		0x36, 0x10, 0xa4, 0xdf, 0xe7, 0xbc, 0xb2, 0x5c, 
		0x0e, 0xa5, 0xfc, 0x6f, 0xb2, 0x33, 0x9b, 0x23, 
		0x65, 0x7a, 0x55, 0x3d, 0x6b, 0x81, 0x0f, 0xac, 
		0x76, 0x31, 0xa2, 0x77, 0x6a, 0x23, 0x16, 0x8a, 
	},
		chainhash.Hash{
		0x70, 0x7d, 0xb9, 0x39, 0xc1, 0x91, 0xd2, 0x35, 
		0x0f, 0x45, 0x62, 0x2c, 0x0b, 0x3c, 0x72, 0x82, 
		0x75, 0x7c, 0x11, 0x5e, 0x72, 0x81, 0x4d, 0x58, 
		0xf2, 0xb3, 0xea, 0x40, 0x49, 0xe5, 0x33, 0x42, 
	},
		chainhash.Hash{
		0x96, 0x3d, 0xdc, 0x10, 0xa8, 0x36, 0x68, 0xe5, 
		0x0b, 0xfb, 0xdd, 0x0b, 0xf5, 0x1a, 0xc8, 0x69, 
		0x7d, 0x92, 0xdb, 0x26, 0x92, 0x69, 0x56, 0x3f, 
		0x49, 0x0c, 0xfc, 0x0a, 0x22, 0xf1, 0x90, 0x92, 
	},
		chainhash.Hash{
		0x68, 0x10, 0xc4, 0xce, 0xed, 0x20, 0xee, 0xc8, 
		0x8b, 0xda, 0x7a, 0x11, 0xf6, 0xbb, 0x68, 0xaf, 
		0x76, 0xcb, 0x35, 0xdb, 0xca, 0x3e, 0xa9, 0xf1, 
		0x35, 0x37, 0x51, 0x97, 0x57, 0xf1, 0xf2, 0x1d, 
	},
			},
		},
	},
	&token.RightDef {Father: chainhash.Hash{},
		Desc: []byte("All Rights"),
		Attrib: 128,
	},
}

var coinToken = token.Token{
	TokenType: 0,
	Value: &token.NumToken{Val: 600000000},
	Rights: &chainhash.Hash{},
}

var polygonToken = token.Token{
	TokenType: 3,
	Value: &token.HashToken{Hash: chainhash.Hash{
		0xd5, 0x2f, 0xf4, 0x98, 0x24, 0x3d, 0xb2, 0xe7, 
		0x41, 0xa8, 0x34, 0xdc, 0x05, 0x01, 0x63, 0x7e, 
		0x0a, 0x4c, 0x2d, 0xb9, 0x31, 0x34, 0x6e, 0x60, 
		0xf6, 0xb0, 0x17, 0x23, 0xa8, 0xb5, 0x91, 0xec, 
	}},
	Rights: &chainhash.Hash{
		0xaf, 0xd4, 0xc8, 0x04, 0x55, 0x4a, 0x17, 0x47, 
		0x7c, 0xfb, 0xb0, 0xd8, 0x3c, 0xf1, 0x03, 0x39, 
		0xac, 0x5f, 0xe4, 0x97, 0x0d, 0x1a, 0x1a, 0x25, 
		0xaa, 0xb3, 0x83, 0xf3, 0xf6, 0x2e, 0x36, 0x39, 
	},
}

var mainnetcreator = [20]byte{
				0x5d, 0x3c, 0x36, 0xc2, 0x57, 0x87, 0x6c, 0x03, 
				0xe0, 0xee, 0xff, 0x8e, 0xbf, 0x4a, 0x9c, 0x82, 
				0x72, 0x1e, 0x34, 0x66, }


var mainnetcoinbaseTx = wire.MsgTx{
	Version: 1,
	TxDef: []token.Definition{},
	TxIn: []*wire.TxIn{
		{
			PreviousOutPoint: wire.OutPoint{
			Hash:  chainhash.Hash{},
			Index: 0,
		},
		SignatureIndex: 0xffffffff,
		Sequence: 0xffffffff,
	},
	},
	TxOut: []*wire.TxOut{
		{
			Token:coinToken,
			PkScript: []byte{
				0x00, 0x5d, 0x3c, 0x36, 0xc2, 0x57, 0x87, 0x6c, 
				0x03, 0xe0, 0xee, 0xff, 0x8e, 0xbf, 0x4a, 0x9c, 
				0x82, 0x72, 0x1e, 0x34, 0x66, 0x41, 0x00, 0x00, 
				0x00, 
			},
		},
	},
	SignatureScripts: [][]byte { []byte{
		0x8c, 0x7e, 0x94, 0x70, 0x7d, 0x7c, 0x7b, 0x1c, 
		0x42, 0x5a, 0xe3, 0xcb, 0x7b, 0xdd, 0x3e, 0xe8, 
		0xd8, 0xa9, 0x46, 0x8b, 0x50, 0x51, 0xfd, 0x5b, 
		0x66, 0x5e, 0x62, 0x45, 0x4e, 0xec, 0xd8, 0x63, 
	} },
	LockTime: 0,
}

var mainnetPolygonTx = wire.MsgTx{
	Version: 1,
	TxDef: InitDefs,
	TxIn: []*wire.TxIn{},
	TxOut: []*wire.TxOut{
		{
			Token: polygonToken,
			PkScript: []byte{
				0x00, 0x46, 0x76, 0x8d, 0x06, 0x5d, 0x31, 0x27, 
				0xc0, 0x42, 0x14, 0xdf, 0x75, 0x8e, 0x51, 0x2b, 
				0x54, 0x48, 0x80, 0x95, 0xc1, 0x41, 0x00, 0x00, 
				0x00, 
			},
		},
	},
	LockTime: 0,
}

var MainNetGenesisMerkleRoot = chainhash.Hash{
		0x8c, 0x7e, 0x94, 0x70, 0x7d, 0x7c, 0x7b, 0x1c, 
		0x42, 0x5a, 0xe3, 0xcb, 0x7b, 0xdd, 0x3e, 0xe8, 
		0xd8, 0xa9, 0x46, 0x8b, 0x50, 0x51, 0xfd, 0x5b, 
		0x66, 0x5e, 0x62, 0x45, 0x4e, 0xec, 0xd8, 0x63, 
	}

var MainNetGenesisBlock = wire.MsgBlock{
	Header: wire.BlockHeader{
		Version:    0x10000,
		PrevBlock:  chainhash.Hash{},
		MerkleRoot: MainNetGenesisMerkleRoot,
		Timestamp:  time.Unix(0x5effa4cf, 0), 
		Nonce:      560458838,
	},
	Transactions: []*wire.MsgTx{&mainnetcoinbaseTx, &mainnetPolygonTx},
}

var MainNetGenesisHash = []chainhash.Hash{
chainhash.Hash{
		0x63, 0xe9, 0x5b, 0x2f, 0x3a, 0x45, 0x5c, 0x82, 
		0x7e, 0x59, 0x52, 0x53, 0xde, 0xd8, 0x91, 0x43, 
		0xc8, 0x1b, 0x21, 0x6c, 0x07, 0x0c, 0x5d, 0xb8, 
		0xce, 0xd9, 0x59, 0xca, 0x0f, 0x00, 0x00, 0x00, 
	},
chainhash.Hash{
		0x26, 0xa8, 0x23, 0x26, 0xb3, 0x11, 0xac, 0x1f, 
		0x6b, 0x2b, 0xc2, 0x8f, 0x31, 0x9b, 0xde, 0xfc, 
		0x3d, 0x80, 0x69, 0x85, 0x9f, 0xed, 0xdf, 0x66, 
		0xf0, 0x34, 0xc4, 0x37, 0x3c, 0x00, 0x00, 0x00, 
	},
}

var MainNetGenesisMinerBlock = wire.MingingRightBlock{
	Version: 0x10000,
	PrevBlock:  chainhash.Hash{},
	BestBlock: MainNetGenesisHash[0],
		Timestamp:  time.Unix(0x5effa4cf, 0), 
	Bits:      0x1e00fff0,
	Nonce:      31458741,
	Connection:      []byte{
				0x6f, 0x6d, 0x65, 0x67, 0x61, 0x73, 0x75, 0x69, 
				0x74, 0x65, 0x2e, 0x6f, 0x72, 0x67, },
	ViolationReport: []*wire.Violations{},
	Utxos:           nil,
	Miner:           mainnetcreator,
}

var regtestcreator = [20]byte{
				0x18, 0x2c, 0xbd, 0xcd, 0xf5, 0x31, 0xc1, 0xaf, 
				0xeb, 0x8f, 0x68, 0x57, 0xd7, 0x87, 0xa8, 0x8e, 
				0xc6, 0x2f, 0x48, 0x59, }


var regtestcoinbaseTx = wire.MsgTx{
	Version: 1,
	TxDef: []token.Definition{},
	TxIn: []*wire.TxIn{
		{
			PreviousOutPoint: wire.OutPoint{
			Hash:  chainhash.Hash{},
			Index: 0,
		},
		SignatureIndex: 0xffffffff,
		Sequence: 0xffffffff,
	},
	},
	TxOut: []*wire.TxOut{
		{
			Token:coinToken,
			PkScript: []byte{
				0x6f, 0x18, 0x2c, 0xbd, 0xcd, 0xf5, 0x31, 0xc1, 
				0xaf, 0xeb, 0x8f, 0x68, 0x57, 0xd7, 0x87, 0xa8, 
				0x8e, 0xc6, 0x2f, 0x48, 0x59, 0x41, 0x00, 0x00, 
				0x00, 
			},
		},
	},
	SignatureScripts: [][]byte { []byte{
		0x6f, 0x41, 0xca, 0x23, 0x8f, 0x90, 0x2b, 0x2f, 
		0x09, 0x14, 0xe6, 0x97, 0x80, 0xd3, 0x5a, 0x5b, 
		0x21, 0x63, 0x64, 0xef, 0x4b, 0x9e, 0xfe, 0x04, 
		0x0d, 0x79, 0x6b, 0xb0, 0x07, 0x6e, 0xf5, 0xf9, 
	} },
	LockTime: 0,
}

var regtestPolygonTx = wire.MsgTx{
	Version: 1,
	TxDef: InitDefs,
	TxIn: []*wire.TxIn{},
	TxOut: []*wire.TxOut{
		{
			Token: polygonToken,
			PkScript: []byte{
				0x6f, 0xca, 0x7e, 0x74, 0xdb, 0xff, 0x00, 0xfd, 
				0xd7, 0x1e, 0xc1, 0xec, 0xb6, 0x6d, 0x02, 0x2c, 
				0x0b, 0xe4, 0x79, 0x2d, 0x19, 0x41, 0x00, 0x00, 
				0x00, 
			},
		},
	},
	LockTime: 0,
}

var RegNetGenesisMerkleRoot = chainhash.Hash{
		0x6f, 0x41, 0xca, 0x23, 0x8f, 0x90, 0x2b, 0x2f, 
		0x09, 0x14, 0xe6, 0x97, 0x80, 0xd3, 0x5a, 0x5b, 
		0x21, 0x63, 0x64, 0xef, 0x4b, 0x9e, 0xfe, 0x04, 
		0x0d, 0x79, 0x6b, 0xb0, 0x07, 0x6e, 0xf5, 0xf9, 
	}

var RegNetGenesisBlock = wire.MsgBlock{
	Header: wire.BlockHeader{
		Version:    0x10000,
		PrevBlock:  chainhash.Hash{},
		MerkleRoot: RegNetGenesisMerkleRoot,
		Timestamp:  time.Unix(0x5effb32a, 0), 
		Nonce:      44,
	},
	Transactions: []*wire.MsgTx{&regtestcoinbaseTx, &regtestPolygonTx},
}

var RegNetGenesisHash = []chainhash.Hash{
chainhash.Hash{
		0x56, 0xaa, 0x23, 0x03, 0xdf, 0xed, 0xe1, 0x81, 
		0xe1, 0x55, 0x9e, 0xb9, 0x9d, 0x75, 0x45, 0xa0, 
		0x31, 0xc0, 0x22, 0xe4, 0x51, 0x61, 0xa3, 0xe7, 
		0xd3, 0x69, 0x95, 0xf6, 0x9e, 0xe8, 0x3e, 0x02, 
	},
chainhash.Hash{
		0xd4, 0x82, 0xe5, 0x9e, 0xd4, 0xe3, 0x0a, 0x44, 
		0xbe, 0x46, 0xf3, 0x07, 0x3a, 0x16, 0xbe, 0xf8, 
		0x5b, 0x01, 0x02, 0xd3, 0x6e, 0x16, 0xfe, 0x5c, 
		0xd9, 0xa9, 0xb1, 0xc7, 0x49, 0xe7, 0xf2, 0x43, 
	},
}

var RegNetGenesisMinerBlock = wire.MingingRightBlock{
	Version: 0x10000,
	PrevBlock:  chainhash.Hash{},
	BestBlock: RegNetGenesisHash[0],
		Timestamp:  time.Unix(0x5effb32a, 0), 
	Bits:      0x207fffff,
	Nonce:      4,
	Connection:      []byte{
				0x6f, 0x6d, 0x65, 0x67, 0x61, 0x73, 0x75, 0x69, 
				0x74, 0x65, 0x2e, 0x6f, 0x72, 0x67, },
	ViolationReport: []*wire.Violations{},
	Utxos:           nil,
	Miner:           regtestcreator,
}

var testnetcreator = [20]byte{
				0x18, 0x2c, 0xbd, 0xcd, 0xf5, 0x31, 0xc1, 0xaf, 
				0xeb, 0x8f, 0x68, 0x57, 0xd7, 0x87, 0xa8, 0x8e, 
				0xc6, 0x2f, 0x48, 0x59, }


var testnetcoinbaseTx = wire.MsgTx{
	Version: 1,
	TxDef: []token.Definition{},
	TxIn: []*wire.TxIn{
		{
			PreviousOutPoint: wire.OutPoint{
			Hash:  chainhash.Hash{},
			Index: 0,
		},
		SignatureIndex: 0xffffffff,
		Sequence: 0xffffffff,
	},
	},
	TxOut: []*wire.TxOut{
		{
			Token:coinToken,
			PkScript: []byte{
				0x6f, 0x18, 0x2c, 0xbd, 0xcd, 0xf5, 0x31, 0xc1, 
				0xaf, 0xeb, 0x8f, 0x68, 0x57, 0xd7, 0x87, 0xa8, 
				0x8e, 0xc6, 0x2f, 0x48, 0x59, 0x41, 0x00, 0x00, 
				0x00, 
			},
		},
	},
	SignatureScripts: [][]byte { []byte{
		0x6f, 0x41, 0xca, 0x23, 0x8f, 0x90, 0x2b, 0x2f, 
		0x09, 0x14, 0xe6, 0x97, 0x80, 0xd3, 0x5a, 0x5b, 
		0x21, 0x63, 0x64, 0xef, 0x4b, 0x9e, 0xfe, 0x04, 
		0x0d, 0x79, 0x6b, 0xb0, 0x07, 0x6e, 0xf5, 0xf9, 
	} },
	LockTime: 0,
}

var testnetPolygonTx = wire.MsgTx{
	Version: 1,
	TxDef: InitDefs,
	TxIn: []*wire.TxIn{},
	TxOut: []*wire.TxOut{
		{
			Token: polygonToken,
			PkScript: []byte{
				0x6f, 0xca, 0x7e, 0x74, 0xdb, 0xff, 0x00, 0xfd, 
				0xd7, 0x1e, 0xc1, 0xec, 0xb6, 0x6d, 0x02, 0x2c, 
				0x0b, 0xe4, 0x79, 0x2d, 0x19, 0x41, 0x00, 0x00, 
				0x00, 
			},
		},
	},
	LockTime: 0,
}

var TestNetGenesisMerkleRoot = chainhash.Hash{
		0x6f, 0x41, 0xca, 0x23, 0x8f, 0x90, 0x2b, 0x2f, 
		0x09, 0x14, 0xe6, 0x97, 0x80, 0xd3, 0x5a, 0x5b, 
		0x21, 0x63, 0x64, 0xef, 0x4b, 0x9e, 0xfe, 0x04, 
		0x0d, 0x79, 0x6b, 0xb0, 0x07, 0x6e, 0xf5, 0xf9, 
	}

var TestNetGenesisBlock = wire.MsgBlock{
	Header: wire.BlockHeader{
		Version:    0x10000,
		PrevBlock:  chainhash.Hash{},
		MerkleRoot: TestNetGenesisMerkleRoot,
		Timestamp:  time.Unix(0x5effb32a, 0), 
		Nonce:      22855,
	},
	Transactions: []*wire.MsgTx{&testnetcoinbaseTx, &testnetPolygonTx},
}

var TestNetGenesisHash = []chainhash.Hash{
chainhash.Hash{
		0xbb, 0x29, 0xa2, 0x2d, 0xf2, 0xd0, 0x25, 0x39, 
		0xab, 0x3e, 0x00, 0x19, 0xa5, 0xdb, 0xdc, 0xa9, 
		0x24, 0x2e, 0xe7, 0x84, 0x97, 0xe8, 0xb0, 0x42, 
		0xca, 0xf0, 0xe7, 0xf0, 0x00, 0x68, 0x00, 0x00, 
	},
chainhash.Hash{
		0x7d, 0x68, 0x88, 0xb8, 0x44, 0xe0, 0x35, 0x27, 
		0xc1, 0x45, 0xe8, 0xfa, 0xca, 0x0a, 0x6b, 0x17, 
		0x02, 0x25, 0x7f, 0x3d, 0x51, 0x8c, 0x4c, 0x27, 
		0x4c, 0xfb, 0x22, 0x61, 0xd2, 0x25, 0x00, 0x00, 
	},
}

var TestNetGenesisMinerBlock = wire.MingingRightBlock{
	Version: 0x10000,
	PrevBlock:  chainhash.Hash{},
	BestBlock: TestNetGenesisHash[0],
		Timestamp:  time.Unix(0x5effb32a, 0), 
	Bits:      0x1f0fffff,
	Nonce:      2698,
	Connection:      []byte{
				0x6f, 0x6d, 0x65, 0x67, 0x61, 0x73, 0x75, 0x69, 
				0x74, 0x65, 0x2e, 0x6f, 0x72, 0x67, },
	ViolationReport: []*wire.Violations{},
	Utxos:           nil,
	Miner:           testnetcreator,
}

var simnetcreator = [20]byte{
				0x18, 0x2c, 0xbd, 0xcd, 0xf5, 0x31, 0xc1, 0xaf, 
				0xeb, 0x8f, 0x68, 0x57, 0xd7, 0x87, 0xa8, 0x8e, 
				0xc6, 0x2f, 0x48, 0x59, }


var simnetcoinbaseTx = wire.MsgTx{
	Version: 1,
	TxDef: []token.Definition{},
	TxIn: []*wire.TxIn{
		{
			PreviousOutPoint: wire.OutPoint{
			Hash:  chainhash.Hash{},
			Index: 0,
		},
		SignatureIndex: 0xffffffff,
		Sequence: 0xffffffff,
	},
	},
	TxOut: []*wire.TxOut{
		{
			Token:coinToken,
			PkScript: []byte{
				0x3f, 0x18, 0x2c, 0xbd, 0xcd, 0xf5, 0x31, 0xc1, 
				0xaf, 0xeb, 0x8f, 0x68, 0x57, 0xd7, 0x87, 0xa8, 
				0x8e, 0xc6, 0x2f, 0x48, 0x59, 0x41, 0x00, 0x00, 
				0x00, 
			},
		},
	},
	SignatureScripts: [][]byte { []byte{
		0x20, 0x60, 0x97, 0x8c, 0x42, 0xbc, 0x42, 0x85, 
		0x11, 0x3f, 0x5d, 0x03, 0x9d, 0x42, 0x2b, 0x34, 
		0x1d, 0x01, 0x40, 0xd3, 0xea, 0x60, 0xb6, 0x06, 
		0x27, 0xe6, 0xb0, 0xf6, 0xad, 0x93, 0x6c, 0x28, 
	} },
	LockTime: 0,
}

var simnetPolygonTx = wire.MsgTx{
	Version: 1,
	TxDef: InitDefs,
	TxIn: []*wire.TxIn{},
	TxOut: []*wire.TxOut{
		{
			Token: polygonToken,
			PkScript: []byte{
				0x3f, 0xca, 0x7e, 0x74, 0xdb, 0xff, 0x00, 0xfd, 
				0xd7, 0x1e, 0xc1, 0xec, 0xb6, 0x6d, 0x02, 0x2c, 
				0x0b, 0xe4, 0x79, 0x2d, 0x19, 0x41, 0x00, 0x00, 
				0x00, 
			},
		},
	},
	LockTime: 0,
}

var SimNetGenesisMerkleRoot = chainhash.Hash{
		0x20, 0x60, 0x97, 0x8c, 0x42, 0xbc, 0x42, 0x85, 
		0x11, 0x3f, 0x5d, 0x03, 0x9d, 0x42, 0x2b, 0x34, 
		0x1d, 0x01, 0x40, 0xd3, 0xea, 0x60, 0xb6, 0x06, 
		0x27, 0xe6, 0xb0, 0xf6, 0xad, 0x93, 0x6c, 0x28, 
	}

var SimNetGenesisBlock = wire.MsgBlock{
	Header: wire.BlockHeader{
		Version:    0x10000,
		PrevBlock:  chainhash.Hash{},
		MerkleRoot: SimNetGenesisMerkleRoot,
		Timestamp:  time.Unix(0x5effb32b, 0), 
		Nonce:      4,
	},
	Transactions: []*wire.MsgTx{&simnetcoinbaseTx, &simnetPolygonTx},
}

var SimNetGenesisHash = []chainhash.Hash{
chainhash.Hash{
		0x14, 0x0b, 0xde, 0x6d, 0xf1, 0x2f, 0xf4, 0xf6, 
		0x3c, 0xaa, 0x3b, 0x54, 0xa3, 0xae, 0x90, 0x1e, 
		0xbb, 0x88, 0x3d, 0x7b, 0x31, 0xda, 0xcd, 0xc8, 
		0x0b, 0x62, 0xf0, 0xd2, 0x59, 0x12, 0x0b, 0x00, 
	},
chainhash.Hash{
		0x0f, 0x7c, 0xe5, 0x17, 0xd9, 0x5e, 0x4e, 0x39, 
		0x76, 0xec, 0x4d, 0x81, 0x5f, 0x6c, 0x40, 0x87, 
		0x0b, 0xd9, 0x68, 0x8b, 0xbd, 0x7c, 0x8e, 0xef, 
		0x5d, 0x0e, 0x4b, 0x5a, 0x5a, 0x95, 0x10, 0x12, 
	},
}

var SimNetGenesisMinerBlock = wire.MingingRightBlock{
	Version: 0x10000,
	PrevBlock:  chainhash.Hash{},
	BestBlock: SimNetGenesisHash[0],
		Timestamp:  time.Unix(0x5effb32b, 0), 
	Bits:      0x207fffff,
	Nonce:      2,
	Connection:      []byte{
				0x6f, 0x6d, 0x65, 0x67, 0x61, 0x73, 0x75, 0x69, 
				0x74, 0x65, 0x2e, 0x6f, 0x72, 0x67, },
	ViolationReport: []*wire.Violations{},
	Utxos:           nil,
	Miner:           simnetcreator,
}