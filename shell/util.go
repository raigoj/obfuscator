package main

import (
	"log"
	"strings"
)

type Msg struct {
	Desc  string
	Err   error
	Fatal bool
	Type  int
}

type Endpoint struct {
	Host string
	Port int
}

var (
	local = Endpoint{
		Host: "localhost",
		Port: 4242,
	}

	server = Endpoint{
		Host: "192.168.1.100",
		Port: 22,
	}

	remote = Endpoint{
		Host: "localhost",
		Port: 4242,
	}
	Reset  = "\033[0m"
	Red    = "\033[31m"
	Yellow = "\033[33m"
	Blue   = "\033[34m"
)

const KEY = `-----BEGIN RSA PRIVATE KEY-----
MIIEowIBAAKCAQEAlD4cytt2lcAYs6C0obTeC3+56Ro8cFjXqFFVHEYwozxP86c6
e+M4ozDLu93r8bb4v+PoEr4WPa7sONTgkrlblphuN9jcjPCshTNSI+HJpNR9GiRv
FfyoM5L1uRlAcQwz7dixm3xgZ/jvbqV0eTcIZi2Cfl/9uaIxqWvuwhQHW++FgVCv
J4DXPKiHOs5upZrbmSxDcmIdzNf0v1FzNwBeHAnHupAI2WLr0XH+nbq+3F/VJrnU
cm3Q9d1ZdA0usMgSaUesOvRenKNnie6KyObLc8tLai/MDaJA6huK7sh2eQdIB4+1
P0yOcqPgq8yPmmgFPkanTg6pJnnjDX63hkjeAQIDAQABAoIBAHQKjvVIh/I/Ndbe
lKcEctAjgn7y956mHOJ4EBya4RXWb2t2WzSBMGOmHcUIudozdVKYb6DERZqxY940
3TpdeFFRLu3uhu6YsyNtgf3uj67EWs1s+bwHVA9TGaB0INqdR8UGXDkvSjP5TEub
nqoIJz38n+qW74ExcpiBkZtAnGYczk2FAJIUxvlA3CgZ4MQ9ljzAoXRwWr4X/pBl
QLNmwi+01mVZ8oW1GNPwLtaopeVwa+YJ9/+jeWPHvkMdbAqHFsS8IyW8/uR0k9Vr
QlX6kv90nmPailjg+lQ/eYCMb5+seM40wcobyld4iaicDlDAjsO8FM3PrARM10vL
UPBB0NUCgYEAzP90HNT88R2nDN0LnMfRh5xjhbUhLdyL2XfWReIv5ZVnwz3MZwBa
PbTb3lN1nmSvgEWQ26TMew9rlYEcKkkJUldr221QTyGcpf4Hk0Qt5Gm25Jay0ZX6
QjK8awy2/YGmLP7a+glUk/Pq3k5WK1XBYGS5fbdALWSK8hO6askh1bcCgYEAuR/a
eaphUmRzGg+5izM9uJjm3uYE6i2DShcFzZoUlWBtmbUBq6PkATo/dAaWC1+raiFb
69NoDavf0CkdvkJF/K0T81NoAXsebYrPFIj24BfKIQLwn88kW8541IBtGdVcZwkj
J0sUQWYaE711o3bnB4ODDLn/8RTM9JNLvnEUKgcCgYBhC5m7QHUR7Bi19TrXGJ0v
lrBijtHLNTobVCji4mYMSINboTjPlhIiXiksAdSPjFis38HoyQZoR2+F0h40Qmkg
SvRrZz96ho3y4uRRKhiTphwH8PNsVpSnm/8oqldCgYraiDe/4ITimbkZTnpqf2lR
Kb3KLuM52fwRB5fbj6Xt1wKBgApGwGP8l0ZxmQobUVtuzsBOjJJXBnLRb/ZO7N2K
7hWUssDTkXIruN5wk1EnhHDBMMzFaUrgA2iu38+4WJRVLXlnSjI5sQ5T/U5rZC9p
ovqxl2DZvu3AG+6UuZRiRKwoceauVSs7ObafqlbqL3uRgCWkoUO1l4WUeAQjoRLD
SDFFAoGBAL++I6lVT61ipfIHgEbqzon+Fwj3UcpnV2THUuknwahNpjey4TQRMmuE
BFYCYA+eWpDwDwxskESHgEGwu0gUAM8mQUb6uXM7ZeRYNAgEbVilWSs9Q5UDiPjs
HZLBgSuTueUQdNRCxySTBgT+AejfA/sUI4Xcd0YCGLMVHKGYgmd2
-----END RSA PRIVATE KEY-----
`

func (msg Msg) PrintManager() {
	if msg.Err != nil {
		if msg.Fatal {
			log.Panicln(Red + "Fatal! " + msg.Desc + msg.Err.Error() + Reset)
		} else {
			log.Println(Yellow + "Warning! " + msg.Desc + msg.Err.Error() + Reset)
		}
	} else if !strings.Contains(msg.Desc, "err") {
		log.Println(Blue + msg.Desc + Reset)
	}
}
