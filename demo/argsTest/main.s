"".main STEXT nosplit size=1 args=0x0 locals=0x0
	0x0000 00000 (main.go:3)	TEXT	"".main(SB), NOSPLIT|ABIInternal, $0-0
	0x0000 00000 (main.go:3)	FUNCDATA	$0, gclocals·33cdeccccebe80329f1fdbee7f5874cb(SB)
	0x0000 00000 (main.go:3)	FUNCDATA	$1, gclocals·33cdeccccebe80329f1fdbee7f5874cb(SB)
	0x0000 00000 (main.go:5)	RET
	0x0000 c3                                               .
"".test STEXT nosplit size=91 args=0x60 locals=0x0
	0x0000 00000 (main.go:8)	TEXT	"".test(SB), NOSPLIT|ABIInternal, $0-96
	0x0000 00000 (main.go:8)	FUNCDATA	$0, gclocals·33cdeccccebe80329f1fdbee7f5874cb(SB)
	0x0000 00000 (main.go:8)	FUNCDATA	$1, gclocals·33cdeccccebe80329f1fdbee7f5874cb(SB)
	0x0000 00000 (main.go:9)	MOVQ	"".a+8(SP), AX
	0x0005 00005 (main.go:9)	MOVQ	"".b+16(SP), CX
	0x000a 00010 (main.go:9)	ADDQ	CX, AX
	0x000d 00013 (main.go:9)	MOVQ	"".c+24(SP), CX
	0x0012 00018 (main.go:9)	ADDQ	CX, AX
	0x0015 00021 (main.go:9)	MOVQ	"".d+32(SP), CX
	0x001a 00026 (main.go:9)	ADDQ	CX, AX
	0x001d 00029 (main.go:9)	MOVQ	"".e+40(SP), CX
	0x0022 00034 (main.go:9)	ADDQ	CX, AX
	0x0025 00037 (main.go:9)	MOVQ	"".f+48(SP), CX
	0x002a 00042 (main.go:9)	ADDQ	CX, AX
	0x002d 00045 (main.go:9)	MOVQ	"".g+56(SP), CX
	0x0032 00050 (main.go:9)	ADDQ	CX, AX
	0x0035 00053 (main.go:9)	MOVQ	"".h+64(SP), CX
	0x003a 00058 (main.
	:9)	ADDQ	CX, AX
	0x003d 00061 (main.go:9)	MOVQ	"".i+72(SP), CX
	0x0042 00066 (main.go:9)	ADDQ	CX, AX
	0x0045 00069 (main.go:9)	MOVQ	"".j+80(SP), CX
	0x004a 00074 (main.go:9)	ADDQ	CX, AX
	0x004d 00077 (main.go:9)	MOVQ	"".k+88(SP), CX
	0x0052 00082 (main.go:9)	ADDQ	CX, AX
	0x0055 00085 (main.go:9)	MOVQ	AX, "".~r11+96(SP)
	0x005a 00090 (main.go:9)	RET
	0x0000 48 8b 44 24 08 48 8b 4c 24 10 48 01 c8 48 8b 4c  H.D$.H.L$.H..H.L
	0x0010 24 18 48 01 c8 48 8b 4c 24 20 48 01 c8 48 8b 4c  $.H..H.L$ H..H.L
	0x0020 24 28 48 01 c8 48 8b 4c 24 30 48 01 c8 48 8b 4c  $(H..H.L$0H..H.L
	0x0030 24 38 48 01 c8 48 8b 4c 24 40 48 01 c8 48 8b 4c  $8H..H.L$@H..H.L
	0x0040 24 48 48 01 c8 48 8b 4c 24 50 48 01 c8 48 8b 4c  $HH..H.L$PH..H.L
	0x0050 24 58 48 01 c8 48 89 44 24 60 c3                 $XH..H.D$`.
go.cuinfo.packagename. SDWARFINFO dupok size=0
	0x0000 6d 61 69 6e                                      main
go.info."".test$abstract SDWARFINFO dupok size=98
	0x0000 04 2e 74 65 73 74 00 01 01 11 61 00 00 00 00 00  ..test....a.....
	0x0010 00 11 62 00 00 00 00 00 00 11 63 00 00 00 00 00  ..b.......c.....
	0x0020 00 11 64 00 00 00 00 00 00 11 65 00 00 00 00 00  ..d.......e.....
	0x0030 00 11 66 00 00 00 00 00 00 11 67 00 00 00 00 00  ..f.......g.....
	0x0040 00 11 68 00 00 00 00 00 00 11 69 00 00 00 00 00  ..h.......i.....
	0x0050 00 11 6a 00 00 00 00 00 00 11 6b 00 00 00 00 00  ..j.......k.....
	0x0060 00 00                                            ..
	rel 0+0 t=24 type.int+0
	rel 13+4 t=29 go.info.int+0
	rel 21+4 t=29 go.info.int+0
	rel 29+4 t=29 go.info.int+0
	rel 37+4 t=29 go.info.int+0
	rel 45+4 t=29 go.info.int+0
	rel 53+4 t=29 go.info.int+0
	rel 61+4 t=29 go.info.int+0
	rel 69+4 t=29 go.info.int+0
	rel 77+4 t=29 go.info.int+0
	rel 85+4 t=29 go.info.int+0
	rel 93+4 t=29 go.info.int+0
""..inittask SNOPTRDATA size=24
	0x0000 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	0x0010 00 00 00 00 00 00 00 00                          ........
gclocals·33cdeccccebe80329f1fdbee7f5874cb SRODATA dupok size=8
	0x0000 01 00 00 00 00 00 00 00                          ........
