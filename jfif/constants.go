package jfif

type marker uint16

const (
	soi  marker = 0xffd8
	sof0        = 0xffc0
	sof2        = 0xffc2
	dht         = 0xffc4
	dqt         = 0xffdb
	dri         = 0xffdd
	sos         = 0xffda
	rstn        = 0xffd0 // 0xffd0 to 0xffd7
	appn        = 0xffe0 // 0xffe0 to 0ffxef
	com         = 0xfffe
	eoi         = 0xffd9
)
