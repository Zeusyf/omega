; contract construction
; Bitcoin Address (Compressed)        1LHxyYZY6NvaigH67DvbLDoyr1XRFsnwMX
; Bitcoin Testnet Address (Compressed)        mzovGbeWuQMqVnkhpntyA92Ji1888Vej1A
; Public Key Bytes (Compressed)       022c7c48dc315b2ad8e96614d13446b37809231a53155bbb48e2ed2919ae58bcd3
; Public Key Base64 (Compressed)      Aix8SNwxWyrY6WYU0TRGs3gJIxpTFVu7SOLtKRmuWLzT
; 
; Bitcoin Address (Uncompressed)      17CzR9btarqquMhY97toRQQZZmGPXssAA5
; Bitcoin Testnet Address (Uncompressed)      mmiwiCgsPtH6gUB9rgsBFKctRks6WeMjCW
; Public Key Bytes (Uncompressed)     042C7C48DC315B2AD8E96614D13446B37809231A53155BBB48E2ED2919AE58BCD
;                                     35FC41414DF2E43C0A43804001A828564E7DCB7EA05F49C78FBC9B3BB574D3B24
; Public Key Base64 (Uncompressed)    BCx8SNwxWyrY6WYU0TRGs3gJIxpTFVu7SOLtKRmuWLzTX8QUFN8uQ8CkOAQAGoKFZOfct+oF9Jx4+8mzu1dNOyQ=
; 
; Private Key WIFC (Compressed)       L4dnAe1xmVbbvH6kHqXhyS77ev8q7yd9pgeDPz3cGFJHohmZo7w5
; Private Key WIF (Uncompressed)      5KVixCZPc18qhAYD5GXX8MBuFwBPEMz8SJbG3VLAoW4vCXHjEjg
; Private Key Bytes                   DD3E5D138FB3F0F1DBED9B464DBF433E95EF9D17C19DA3EEE3CF1B08FCADB088
; Private Key Base64                  3T5dE4+z8PHb7ZtGTb9DPpXvnRfBnaPu488bCPytsIg=
; Bitcoin TestNet Address Private Key WIF (Uncompressed)        93GMXwNwCECyfE3VhcRRzwjrubY6PXXKnFTD87gg9EoxyaEG1EV
; Bitcoin TestNet Address Private Key WIF (Compressed)        cUzmdZ1pCZHs5ia1gFLqLkcBH9SEnRiqtingWQW7mMxJ4SomhfBw
;

define RECEIVING x17d5f1c868d2c64a5af9e0bba1efad0ee49c5957	; public contract for receiving money
define PUBKEY kx022c7c48dc315b2ad8e96614d13446b37809231a53155bbb48e2ed2919ae58bcd3

STORE abi("sequence"),D0,
STORE abi("utxotid"),Hx0,
STORE abi("utxotseq"),D0,
STORE abi("lastid"),Hx0,
STORE abi("lastseq"),D0,
STORE abi("balance"),Q0,
MALLOC gi0,128,		; alloc mem len of store code

; TBD bug fix: don't use i0 for data; always use ii0"xxx

define NETID 0		; 0 for main net, x6f for test net

define tokentype i20
define tokenval i28

GETCOIN tokentype,

EVAL64 i64,tokentype,0,=
IF i64,2,
REVERT

EVAL64 i64,tokenval,0,=
IF i64,.doneinit,

define OUTPOINT i64
define OUTPOINTSEQ i96

RECEIVED OUTPOINT,
STORE abi("utxotid"),HOUTPOINT,
STORE abi("utxotseq"),DOUTPOINT"32,
STORE abi("balance"),Qtokenval,

define doneinit .
EVAL32 gi0,4,		; length of result
EVAL32 gi4,BODY,	; result = the first instruction of contract body code.
STOP
;
; contract body
define BODY .
define pubkey gii0"260
define pubkeyd gii0"259
define contract gii0"8

define BALANCE gii0"100
define BALANCEDATA gii0"104

MALLOC gii0"8,1032,	; storage to use. enough for all cases.

EVAL32 gii0"16,abi("getseq()uint32"),gi8,=
IF gii0"16,.getownerseq,
EVAL32 gii0"16,abi("issuspended()bool"),gi8,=
IF gii0"16,.suspended,
EVAL32 gii0"16,abi("balance()int64bool"),gi8,=
IF gii0"16,.balance,

; above code are contract calls (except for suicide)
; codes below are transactions. do common work for txs
LIBLOAD 0,RECEIVING,

define seq gii0"416
define seqtext gii0"420		; don't change this address w/o considering signature
define adr gii0"396
define func gii0"424
define ftext gii0"400
define issueto gii0"428
define amount gii0"449
define adrtext gii0"400

define hashed gii0"300

; define hashed160 ii0"340
COPYIMM pubkey,PUBKEY,			; issue([21]byte,uint64,[]byte)
EVAL8 pubkeyd,33,

; increase in all cases, if exec fails, it won't be saved anyway
LOAD seq,abi("sequence"),		; build text for hashing
EVAL32 seqtext,seqtext,1,+
STORE abi("sequence"),Dseqtext,		; contract address + seq + abi
META adr,7,"address",
EVAL32 func,gi8,

EVAL32 gii0"16,abi("resume([]byte)"),gi8,=
IF gii0"16,.resume,
EVAL32 gii0"16,abi("suspend([]byte)"),gi8,=
IF gii0"16,.suspend,
EVAL32 gii0"16,abi("issue([21]byte,uint64,[]byte)"),gi8,=	; func(address, amount, signature)
IF gii0"16,.issue,
EVAL32 gii0"16,abi("suicide([]byte)"),gi8,=
IF gii0"16,.suicide,
CALL RECEIVING,abi("generic(bool,int64)byte"),0,0,0,		; receiveing money, sender pay fees
STOP

define suspend .
CALL 0,.sigcheck2,@gi12,
CALL RECEIVING,abi("generic(bool,int64)byte"),0,1,0,		; contract pay fees. if fees is low, there is no need to suspend
STORE abi("suspended"),B1,		; suspend()bool
EVAL32 gi0,0,
STOP

define resume .
CALL 0,.sigcheck2,@gi12,
CALL RECEIVING,abi("generic(bool,int64)byte"),0,0,0,		; caller pay fees because contract might have no money
STORE abi("suspended"),B0,		; resume()bool
EVAL32 gi0,0,
STOP

define sigcheck2 .
HASH hashed,ftext,28,
SIGCHECK tmp,hashed,pubkeyd,ii8,
IF tmp,2,
REVERT
RETURN

define suspendcheck .
EVAL8 tmp"4,0,
LOAD tmp,abi("suspended"),		; suspendcheck()
IF tmp"4,2,
RETURN
REVERT

define issue .
CALL 0,.suspendcheck,
LOAD BALANCE,abi("balance"),
EVAL64 tmp,BALANCEDATA,gi33,)
IF tmp,2,
REVERT
CALL 0,.sigcheck,
CALL 0,.issuetoken,
EVAL32 gi0,0,
STOP

define sigcheck .
COPY issueto,gi12,29,			; addr + uint64
HASH hashed,ftext,57,
SIGCHECK i0,hashed,pubkeyd,gi41,
IF i0,2,
REVERT
RETURN

define issuetoken .
EVAL64 tokentype,0,
EVAL64 tokenval,amount,
EVAL32 scriptlen,25,
COPY scriptver,gi12,21,
EVAL32 scriptfunc,x41,
ADDTXOUT tmp,tokentype,
CALL RECEIVING,abi("generic(bool,int64)byte"),0,1,amount,
EVAL32 gi0,0,
RETURN

define tokentype gii0"28
define tokenval gii0"36
define scriptlen gii0"44
define scriptver gii0"48
define scriptstr gii0"49
define scriptfunc gii0"69
define tmp gii0'80

define suicide .
GETCOIN tokentype,		; for suicide, must take no value in
EVAL64 tmp,tokentype,0,=
IF tmp,2,
REVERT
EVAL64 tmp,tokenval,0,=
IF tmp,2,
REVERT

CALL 0,.sigcheck2,@gi12,

COPYIMM pubkey,PUBKEY,
HASH160 hashed,pubkey,33,
EVAL8 scriptver,NETID,
COPY scriptstr,hashed,20,

CALL RECEIVING,abi("disburse(pointer,pointer)"),@tmp,@scriptstr,

LOAD LASTSEQ,abi("lastseq"),
LOAD LASTID,abi("lastid"),
LOAD UTXOTSEQ,abi("utxotseq"),
LOAD UTXOTID,abi("utxotid"),
LOAD BALANCE,abi("balance"),

EVAL64 tmp,BALANCEDATA,0,=		; no balance, go die directly
IF tmp,.kill,

define UTXOTID gii0"108
define UTXOTSEQ gii0"140
define LASTID gii0"144
define LASTSEQ gii0"176

define UTXOTIDDATA gii0"112
define UTXOTSEQDATA gii0"144
define LASTIDDATA gii0"148
define LASTSEQDATA gii0"180

EVAL64 tmp,UTXOTIDDATA,0,=		; spend all
IF tmp,2,
SPEND UTXOTIDDATA,UTXOTSEQDATA,
EVAL64 tmp,LASTIDDATA,0,=
IF tmp,2,
SPEND LASTIDDATA,LASTSEQDATA,

define txfees ii0"190
define outval ii0"198

TXFEE txfees,2,				; pay fees
EVAL64 outval,BALANCEDATA,txfees,-
EVAL64 tmp,outval,0,(
IF tmp,.kill,

COPYIMM pubkey,PUBKEY,
HASH160 hashed,pubkey,33,
EVAL64 tokenval,outval,
EVAL32 scriptlen,25,
EVAL8 scriptver,NETID,
COPY scriptstr,hashed,20,
EVAL32 scriptfunc,x41,			; pkhpubkey
ADDTXOUT i232,tokentype,

define kill .				; die. delete storage
DEL abi("balance"),
DEL abi("utxotid"),
DEL abi("utxotseq"),
DEL abi("lastid"),
DEL abi("lastseq"),
DEL abi("sequence"),
SELFDESTRUCT

define getownerseq .
LOAD gi0,abi("sequence"),		; getseq()uint32
STOP

define suspended .
EVAL8 gi4,0,
LOAD gi0,abi("suspended"),		; suspended()bool
STOP

define balance .
LOAD gi0,abi("balance"),		; suspended()bool
STOP





Ox07030000,D0,
Ox06030000,Hx0,
Ox00060000,D0,
Ox02000205,Hx0,
Ox00070000,D0,
Ox01050705,Q0,
Rgi0,128,
cgii0"28,
Di64,gii0"28,0,=
Ki64,2,
X
Di64,gii0"36,0,=
Ki64,5,
ai64,
Ox06030000,Hi64,
Ox00060000,Di64"32,
Ox01050705,Qgii0"36,
Cgi0,4,
Cgi4,20,
z
Rgii0"8,1032,
Cgii0"16,x00060000,gi8,=
Kgii0"16,107,
Cgii0"16,x05000004,gi8,=
Kgii0"16,107,
Cgii0"16,x00070907,gi8,=
Kgii0"16,108,
Q0,xceddcbdf1bc15a0feb4473df6578ec61445a48cb,
Ugii0"260,kx022c7c48dc315b2ad8e96614d13446b37809231a53155bbb48e2ed2919ae58bcd3,
Agii0"259,33,
Ngii0"416,x07030000,
Cgii0"420,gii0"420,1,+
Ox07030000,Dgii0"420,
kgii0"396,7,x61646472657373,
Cgii0"424,gi8,
Cgii0"16,x09030604,gi8,=
Kgii0"16,9,
Cgii0"16,x06060007,gi8,=
Kgii0"16,12,
Cgii0"16,x02000400,gi8,=
Kgii0"16,25,
Cgii0"16,x00060200,gi8,=
Kgii0"16,47,
Lxceddcbdf1bc15a0feb4473df6578ec61445a48cb,x00010000,0,0,0,
z
L0,10,@gi12,
Lxceddcbdf1bc15a0feb4473df6578ec61445a48cb,x00010000,0,0,0,
Ox00020000,B0,
Cgi0,0,
z
L0,5,@gi12,
Lxceddcbdf1bc15a0feb4473df6578ec61445a48cb,x00010000,0,1,0,
Ox00020000,B1,
Cgi0,0,
z
Ggii0"300,gii0"400,28,
Igii0'80,gii0"300,gii0"259,ii8,
Kgii0'80,2,
X
Y
Agii0'80"4,0,
Ngii0'80,x00020000,
Kgii0'80"4,2,
Y
X
L0,n5,
Ngii0"100,x01050705,
Dgii0'80,gii0"104,gi33,)
Kgii0'80,2,
X
L0,4,
L0,9,
Cgi0,0,
z
Tgii0"428,gi12,29,
Ggii0"300,gii0"400,57,
Ii0,gii0"300,gii0"259,gi41,
Ki0,2,
X
Y
Dgii0"28,0,
Dgii0"36,gii0"449,
Cgii0"44,25,
Tgii0"48,gi12,21,
Cgii0"69,x41,
ggii0'80,gii0"28,
Lxceddcbdf1bc15a0feb4473df6578ec61445a48cb,x00010000,0,1,gii0"449,
Cgi0,0,
Y
cgii0"28,
Dgii0'80,gii0"28,0,=
Kgii0'80,2,
X
Dgii0'80,gii0"36,0,=
Kgii0'80,2,
X
L0,n41,@gi12,
Ngii0"176,x00070000,
Ngii0"144,x02000205,
Ngii0"140,x00060000,
Ngii0"108,x06030000,
Ngii0"100,x01050705,
Dgii0'80,gii0"104,0,=
Kgii0'80,19,
Dgii0'80,gii0"112,0,=
Kgii0'80,2,
egii0"112,gii0"144,
Dgii0'80,gii0"148,0,=
Kgii0'80,2,
egii0"148,gii0"180,
bii0"190,2,
Dii0"198,gii0"104,ii0"190,-
Dgii0'80,ii0"198,0,(
Kgii0'80,9,
Ugii0"260,kx022c7c48dc315b2ad8e96614d13446b37809231a53155bbb48e2ed2919ae58bcd3,
Hgii0"300,gii0"260,33,
Dgii0"36,ii0"198,
Cgii0"44,25,
Agii0"48,x6f,
Tgii0"49,gii0"300,20,
Cgii0"69,x41,
gi232,gii0"28,
Px01050705,
Px06030000,
Px00060000,
Px02000205,
Px00070000,
Px07030000,
W
Ngi0,x07030000,
z
Agi4,0,
Ngi0,x00020000,
z
Ngi0,x01050705,
z

eee899eaec03ca9b35a1e3130cd90339a5e1eafc
88eee899eaec03ca9b35a1e3130cd90339a5e1eafc000000004f7830373033303030302c44302c0a4f7830363033303030302c4878302c0a4f7830303036303030302c44302c0a4f7830323030303230352c4878302c0a4f7830303037303030302c44302c0a4f7830313035303730352c51302c0a526769302c3132382c0a63676969302232382c0a446936342c676969302232382c302c3d0a4b6936342c322c0a580a446936342c676969302233362c302c3d0a4b6936342c352c0a616936342c0a4f7830363033303030302c486936342c0a4f7830303036303030302c446936342233322c0a4f7830313035303730352c51676969302233362c0a436769302c342c0a436769342c32302c0a7a0a526769693022382c313033322c0a43676969302231362c7830303036303030302c6769382c3d0a4b676969302231362c3130372c0a43676969302231362c7830353030303030342c6769382c3d0a4b676969302231362c3130372c0a43676969302231362c7830303037303930372c6769382c3d0a4b676969302231362c3130382c0a51302c78636564646362646631626331356130666562343437336466363537386563363134343561343863622c0a5567696930223236302c6b783032326337633438646333313562326164386539363631346431333434366233373830393233316135333135356262623438653265643239313961653538626364332c0a4167696930223235392c33332c0a4e67696930223431362c7830373033303030302c0a4367696930223432302c67696930223432302c312c2b0a4f7830373033303030302c4467696930223432302c0a6b67696930223339362c372c7836313634363437323635373337332c0a4367696930223432342c6769382c0a43676969302231362c7830393033303630342c6769382c3d0a4b676969302231362c392c0a43676969302231362c7830363036303030372c6769382c3d0a4b676969302231362c31322c0a43676969302231362c7830323030303430302c6769382c3d0a4b676969302231362c32352c0a43676969302231362c7830303036303230302c6769382c3d0a4b676969302231362c34372c0a4c78636564646362646631626331356130666562343437336466363537386563363134343561343863622c7830303031303030302c302c302c302c0a7a0a4c302c31302c40676931322c0a4c78636564646362646631626331356130666562343437336466363537386563363134343561343863622c7830303031303030302c302c302c302c0a4f7830303032303030302c42302c0a436769302c302c0a7a0a4c302c352c40676931322c0a4c78636564646362646631626331356130666562343437336466363537386563363134343561343863622c7830303031303030302c302c312c302c0a4f7830303032303030302c42312c0a436769302c302c0a7a0a4767696930223330302c67696930223430302c32382c0a49676969302738302c67696930223330302c67696930223235392c6969382c0a4b676969302738302c322c0a580a590a416769693027383022342c302c0a4e676969302738302c7830303032303030302c0a4b6769693027383022342c322c0a590a580a4c302c6e352c0a4e67696930223130302c7830313035303730352c0a44676969302738302c67696930223130342c676933332c290a4b676969302738302c322c0a580a4c302c342c0a4c302c392c0a436769302c302c0a7a0a5467696930223432382c676931322c32392c0a4767696930223330302c67696930223430302c35372c0a4969302c67696930223330302c67696930223235392c676934312c0a4b69302c322c0a580a590a44676969302232382c302c0a44676969302233362c67696930223434392c0a43676969302234342c32352c0a54676969302234382c676931322c32312c0a43676969302236392c7834312c0a67676969302738302c676969302232382c0a4c78636564646362646631626331356130666562343437336466363537386563363134343561343863622c7830303031303030302c302c312c67696930223434392c0a436769302c302c0a590a63676969302232382c0a44676969302738302c676969302232382c302c3d0a4b676969302738302c322c0a580a44676969302738302c676969302233362c302c3d0a4b676969302738302c322c0a580a4c302c6e34312c40676931322c0a4e67696930223137362c7830303037303030302c0a4e67696930223134342c7830323030303230352c0a4e67696930223134302c7830303036303030302c0a4e67696930223130382c7830363033303030302c0a4e67696930223130302c7830313035303730352c0a44676969302738302c67696930223130342c302c3d0a4b676969302738302c31392c0a44676969302738302c67696930223131322c302c3d0a4b676969302738302c322c0a6567696930223131322c67696930223134342c0a44676969302738302c67696930223134382c302c3d0a4b676969302738302c322c0a6567696930223134382c67696930223138302c0a62696930223139302c322c0a44696930223139382c67696930223130342c696930223139302c2d0a44676969302738302c696930223139382c302c280a4b676969302738302c392c0a5567696930223236302c6b783032326337633438646333313562326164386539363631346431333434366233373830393233316135333135356262623438653265643239313961653538626364332c0a4867696930223330302c67696930223236302c33332c0a44676969302233362c696930223139382c0a43676969302234342c32352c0a41676969302234382c7836662c0a54676969302234392c67696930223330302c32302c0a43676969302236392c7834312c0a67693233322c676969302232382c0a507830313035303730352c0a507830363033303030302c0a507830303036303030302c0a507830323030303230352c0a507830303037303030302c0a507830373033303030302c0a570a4e6769302c7830373033303030302c0a7a0a416769342c302c0a4e6769302c7830303032303030302c0a7a0a4e6769302c7830313035303730352c0a7a0a0a

















Depolyment:
payment: 8eef8283a6b37ae15fcbb3616bd56958d7e6d0e3ef96152ca6ec0535e736d809 : 0 =>
	 (mz8JyMrJNGnGd7Jyz3s7BQabpKu89VHHoQ) c5247299f739933901413888e30bc0ce51f1ca8e9d026fafd4f08eaf5419bc84 : 0

contract: eee899eaec03ca9b35a1e3130cd90339a5e1eafc
deposit: 88eee899eaec03ca9b35a1e3130cd90339a5e1eafc52360000
balance: 07090700
sequence: 00000600
suspended: 04000005

suspend: 07000606
resume: 04060309
issue: 00040002
suicide: 00020600

make of signature for issue token: text=eee899eaec03ca9b35a1e3130cd90339a5e1eafc + (seq+1) + func abi + [21]byte address + uint64 amount
hash of text

eee899eaec03ca9b35a1e3130cd90339a5e1eafc06000000000400026fcc2215697a1215cbe6941c4da703e190c7edfcd50000010000000000


eee899eaec03ca9b35a1e3130cd90339a5e1eafc
06000000
00040002
6fcc2215697a1215cbe6941c4da703e190c7edfcd5
0000010000000000

0e70c13ec82f7c5804227f5f8f7c402c2070012a7afb6c30529310c0a533ab79


sign on hash
priv k: cUzmdZ1pCZHs5ia1gFLqLkcBH9SEnRiqtingWQW7mMxJ4SomhfBw	(test)
L4dnAe1xmVbbvH6kHqXhyS77ev8q7yd9pgeDPz3cGFJHohmZo7w5	(main)

3045022100ef29c89b309364aad5068e3c1e2bb9f6dbcdbbca10f1a8d9c2c42705277c5ed502205c2319f57ad0de34d54dc793a66c0cc377e2fe00454b21ae73087830fdb061f6


issue: 88eee899eaec03ca9b35a1e3130cd90339a5e1eafc00040002[21]byte address,uint64 amount,[]byte signature
[21]byte address: 6fcc2215697a1215cbe6941c4da703e190c7edfcd5

88eee899eaec03ca9b35a1e3130cd90339a5e1eafc000400026fcc2215697a1215cbe6941c4da703e190c7edfcd500000100000000003045022100881f1ea7bfeae388504c87808ef99d1763876efba1964aecd2673cb0f4af55e102205863d25687f2c68289ad7c47851aeebe47a9053c83e1fbe16671b573c4b03094













make of signature for suspend/resume: text=eee899eaec03ca9b35a1e3130cd90339a5e1eafc + (seq+1) + func abi
hash of text

0a18ebc5a1ef99f0b0cf991ffdab230b2c9a6f5a0c00000000020600




0a18ebc5a1ef99f0b0cf991ffdab230b2c9a6f5a
0c000000
00020600

9641e52fb635bb1d11d7760cbdd287662113c846a6e78dcc52e4d6b533f06442


sign on hash
priv k: cUzmdZ1pCZHs5ia1gFLqLkcBH9SEnRiqtingWQW7mMxJ4SomhfBw

3045022100ef29c89b309364aad5068e3c1e2bb9f6dbcdbbca10f1a8d9c2c42705277c5ed502205c2319f57ad0de34d54dc793a66c0cc377e2fe00454b21ae73087830fdb061f6


suspend: 88eee899eaec03ca9b35a1e3130cd90339a5e1eafc07000606,[]byte signature
88eee899eaec03ca9b35a1e3130cd90339a5e1eafc070006063045022100dafe87a19f7e83acebc589d7b57aa4e949546461b5aa6348f66f9c59f3d90fe402201e0d6efc52875a10630591d76080d470f31d947d5dd14d1e09ef965d38930766



resume:88eee899eaec03ca9b35a1e3130cd90339a5e1eafc04060309,[]byte signature
88eee899eaec03ca9b35a1e3130cd90339a5e1eafc040603093045022100c0d6c1e0f4bf912608d401adf6edf425c32ed5d84ef776f178f997a2669d9d3002202f00631b7e4e9f377d0794aba797f7f8105aa4b1eade7e67e831cb208001bbc5


suicide: 88eee899eaec03ca9b35a1e3130cd90339a5e1eafc00020600,[]byte signature
880a18ebc5a1ef99f0b0cf991ffdab230b2c9a6f5a000206003045022100ef29c89b309364aad5068e3c1e2bb9f6dbcdbbca10f1a8d9c2c42705277c5ed502205c2319f57ad0de34d54dc793a66c0cc377e2fe00454b21ae73087830fdb061f6





mz8JyMrJNGnGd7Jyz3s7BQabpKu89VHHoQ
88eee899eaec03ca9b35a1e3130cd90339a5e1eafc52360000
deposit: c5247299f739933901413888e30bc0ce51f1ca8e9d026fafd4f08eaf5419bc84 : 0 =>
	 (mz8JyMrJNGnGd7Jyz3s7BQabpKu89VHHoQ) f20701b265eb3a6f5c7923e4c7da626b6524e98b5ce272aae1b490d0c542357b : 0 => all used
	 
	 95d971b501a0bd7ab940516e9f1efeaf9d3195e6f70ed5770be304c28fafc64d

suspend: 6de8cde5f9e0791179cb6ec5ba9c0faba82dd2f6f248bb1cefb80164eccb4581
suspension verified, issue reverted
resume: c0aaacbe8ea7458b25500ed3bb8fc1c2e1096999d7c1767a66a24893298d79a6
deposit: 95d971b501a0bd7ab940516e9f1efeaf9d3195e6f70ed5770be304c28fafc64d
issue: ea5f14bf681da36caa96739ab283448abd6e9644994e30667ad7ef238c1b0360, success

1c9128d2b675b65430d61209aae1953909846cef65fd2507afb812b09c19b220
56f44d2fa827231f423b3e1210dbdb51d30f36c1a7310c23d77392fd6a441d5c

suicide: 118290eade85fa4ee20dfcde478df9c214a975e68e25852e28300c7463c1339f, success
contract call gets Contract does not exist

OK for mainnet depoly



Ox07030000,D0,
Ox06030000,Hx0,
Ox00060000,D0,
Ox02000205,Hx0,
Ox00070000,D0,
Ox01050705,Q0,
Rgi0,128,
cgii0"28,
Di64,gii0"28,0,=
Ki64,2,
X
Di64,gii0"36,0,=
Ki64,5,
ai64,
Ox06030000,Hi64,
Ox00060000,Di64"32,
Ox01050705,Qgii0"36,
Cgi0,4,
Cgi4,20,
z
Rgii0"8,1032,
Cgii0"16,x00060000,gi8,=
Kgii0"16,107,
Cgii0"16,x05000004,gi8,=
Kgii0"16,107,
Cgii0"16,x00070907,gi8,=
Kgii0"16,108,
Q0,xceddcbdf1bc15a0feb4473df6578ec61445a48cb,
Ugii0"260,kx022c7c48dc315b2ad8e96614d13446b37809231a53155bbb48e2ed2919ae58bcd3,
Agii0"259,33,
Ngii0"416,x07030000,
Cgii0"420,gii0"420,1,+
Ox07030000,Dgii0"420,
kgii0"396,7,x61646472657373,
Cgii0"424,gi8,
Cgii0"16,x09030604,gi8,=
Kgii0"16,14,
Cgii0"16,x06060007,gi8,=
Kgii0"16,7,
Cgii0"16,x02000400,gi8,=
Kgii0"16,25,
Cgii0"16,x00060200,gi8,=
Kgii0"16,47,
Lxceddcbdf1bc15a0feb4473df6578ec61445a48cb,x00010000,0,0,0,
z
L0,10,@gi12,
Lxceddcbdf1bc15a0feb4473df6578ec61445a48cb,x00010000,0,1,0,
Ox00020000,B1,
Cgi0,0,
z
L0,5,@gi12,
Lxceddcbdf1bc15a0feb4473df6578ec61445a48cb,x00010000,0,0,0,
Ox00020000,B0,
Cgi0,0,
z
Ggii0"300,gii0"400,28,
Igii0'80,gii0"300,gii0"259,ii8,
Kgii0'80,2,
X
Y
Agii0'80"4,0,
Ngii0'80,x00020000,
Kgii0'80"4,2,
Y
X
L0,n5,
Ngii0"100,x01050705,
Dgii0'80,gii0"104,gi33,)
Kgii0'80,2,
X
L0,4,
L0,9,
Cgi0,0,
z
Tgii0"428,gi12,29,
Ggii0"300,gii0"400,57,
Ii0,gii0"300,gii0"259,gi41,
Ki0,2,
X
Y
Dgii0"28,0,
Dgii0"36,gii0"449,
Cgii0"44,25,
Tgii0"48,gi12,21,
Cgii0"69,x41,
ggii0'80,gii0"28,
Lxceddcbdf1bc15a0feb4473df6578ec61445a48cb,x00010000,0,1,gii0"449,
Cgi0,0,
Y
cgii0"28,
Dgii0'80,gii0"28,0,=
Kgii0'80,2,
X
Dgii0'80,gii0"36,0,=
Kgii0'80,2,
X
L0,n41,@gi12,
Ngii0"176,x00070000,
Ngii0"144,x02000205,
Ngii0"140,x00060000,
Ngii0"108,x06030000,
Ngii0"100,x01050705,
Dgii0'80,gii0"104,0,=
Kgii0'80,19,
Dgii0'80,gii0"112,0,=
Kgii0'80,2,
egii0"112,gii0"144,
Dgii0'80,gii0"148,0,=
Kgii0'80,2,
egii0"148,gii0"180,
bii0"190,2,
Dii0"198,gii0"104,ii0"190,-
Dgii0'80,ii0"198,0,(
Kgii0'80,9,
Ugii0"260,kx022c7c48dc315b2ad8e96614d13446b37809231a53155bbb48e2ed2919ae58bcd3,
Hgii0"300,gii0"260,33,
Dgii0"36,ii0"198,
Cgii0"44,25,
Agii0"48,0,
Tgii0"49,gii0"300,20,
Cgii0"69,x41,
gi232,gii0"28,
Px01050705,
Px06030000,
Px00060000,
Px02000205,
Px00070000,
Px07030000,
W
Ngi0,x07030000,
z
Agi4,0,
Ngi0,x00020000,
z
Ngi0,x01050705,
z

0a18ebc5a1ef99f0b0cf991ffdab230b2c9a6f5a
880a18ebc5a1ef99f0b0cf991ffdab230b2c9a6f5a000000004f7830373033303030302c44302c0a4f7830363033303030302c4878302c0a4f7830303036303030302c44302c0a4f7830323030303230352c4878302c0a4f7830303037303030302c44302c0a4f7830313035303730352c51302c0a526769302c3132382c0a63676969302232382c0a446936342c676969302232382c302c3d0a4b6936342c322c0a580a446936342c676969302233362c302c3d0a4b6936342c352c0a616936342c0a4f7830363033303030302c486936342c0a4f7830303036303030302c446936342233322c0a4f7830313035303730352c51676969302233362c0a436769302c342c0a436769342c32302c0a7a0a526769693022382c313033322c0a43676969302231362c7830303036303030302c6769382c3d0a4b676969302231362c3130372c0a43676969302231362c7830353030303030342c6769382c3d0a4b676969302231362c3130372c0a43676969302231362c7830303037303930372c6769382c3d0a4b676969302231362c3130382c0a51302c78636564646362646631626331356130666562343437336466363537386563363134343561343863622c0a5567696930223236302c6b783032326337633438646333313562326164386539363631346431333434366233373830393233316135333135356262623438653265643239313961653538626364332c0a4167696930223235392c33332c0a4e67696930223431362c7830373033303030302c0a4367696930223432302c67696930223432302c312c2b0a4f7830373033303030302c4467696930223432302c0a6b67696930223339362c372c7836313634363437323635373337332c0a4367696930223432342c6769382c0a43676969302231362c7830393033303630342c6769382c3d0a4b676969302231362c31342c0a43676969302231362c7830363036303030372c6769382c3d0a4b676969302231362c372c0a43676969302231362c7830323030303430302c6769382c3d0a4b676969302231362c32352c0a43676969302231362c7830303036303230302c6769382c3d0a4b676969302231362c34372c0a4c78636564646362646631626331356130666562343437336466363537386563363134343561343863622c7830303031303030302c302c302c302c0a7a0a4c302c31302c40676931322c0a4c78636564646362646631626331356130666562343437336466363537386563363134343561343863622c7830303031303030302c302c312c302c0a4f7830303032303030302c42312c0a436769302c302c0a7a0a4c302c352c40676931322c0a4c78636564646362646631626331356130666562343437336466363537386563363134343561343863622c7830303031303030302c302c302c302c0a4f7830303032303030302c42302c0a436769302c302c0a7a0a4767696930223330302c67696930223430302c32382c0a49676969302738302c67696930223330302c67696930223235392c6969382c0a4b676969302738302c322c0a580a590a416769693027383022342c302c0a4e676969302738302c7830303032303030302c0a4b6769693027383022342c322c0a590a580a4c302c6e352c0a4e67696930223130302c7830313035303730352c0a44676969302738302c67696930223130342c676933332c290a4b676969302738302c322c0a580a4c302c342c0a4c302c392c0a436769302c302c0a7a0a5467696930223432382c676931322c32392c0a4767696930223330302c67696930223430302c35372c0a4969302c67696930223330302c67696930223235392c676934312c0a4b69302c322c0a580a590a44676969302232382c302c0a44676969302233362c67696930223434392c0a43676969302234342c32352c0a54676969302234382c676931322c32312c0a43676969302236392c7834312c0a67676969302738302c676969302232382c0a4c78636564646362646631626331356130666562343437336466363537386563363134343561343863622c7830303031303030302c302c312c67696930223434392c0a436769302c302c0a590a63676969302232382c0a44676969302738302c676969302232382c302c3d0a4b676969302738302c322c0a580a44676969302738302c676969302233362c302c3d0a4b676969302738302c322c0a580a4c302c6e34312c40676931322c0a4e67696930223137362c7830303037303030302c0a4e67696930223134342c7830323030303230352c0a4e67696930223134302c7830303036303030302c0a4e67696930223130382c7830363033303030302c0a4e67696930223130302c7830313035303730352c0a44676969302738302c67696930223130342c302c3d0a4b676969302738302c31392c0a44676969302738302c67696930223131322c302c3d0a4b676969302738302c322c0a6567696930223131322c67696930223134342c0a44676969302738302c67696930223134382c302c3d0a4b676969302738302c322c0a6567696930223134382c67696930223138302c0a62696930223139302c322c0a44696930223139382c67696930223130342c696930223139302c2d0a44676969302738302c696930223139382c302c280a4b676969302738302c392c0a5567696930223236302c6b783032326337633438646333313562326164386539363631346431333434366233373830393233316135333135356262623438653265643239313961653538626364332c0a4867696930223330302c67696930223236302c33332c0a44676969302233362c696930223139382c0a43676969302234342c32352c0a41676969302234382c302c0a54676969302234392c67696930223330302c32302c0a43676969302236392c7834312c0a67693233322c676969302232382c0a507830313035303730352c0a507830363033303030302c0a507830303036303030302c0a507830323030303230352c0a507830303037303030302c0a507830373033303030302c0a570a4e6769302c7830373033303030302c0a7a0a416769342c302c0a4e6769302c7830303032303030302c0a7a0a4e6769302c7830313035303730352c0a7a0a0a


deposit script: 880a18ebc5a1ef99f0b0cf991ffdab230b2c9a6f5a52360000

destruction: 9abd0921b005db15e8cc8a51ec5c92f7f67559dccfce9f21ccae26d178650b58
