package pgp_test

import (
	"strings"
	"testing"

	"salsa.debian.org/autodeb-team/autodeb/internal/pgp"

	"github.com/stretchr/testify/assert"
)

func TestVerifyMesageSignature(t *testing.T) {
	msg, entity, err := pgp.VerifySignatureClearsigned(
		strings.NewReader(signedMessage),
		strings.NewReader(testKeyPublic),
	)

	assert.NoError(t, err)
	assert.Equal(t, "this is a test\n", msg)
	assert.Equal(t, "0x8B5F287532CF42DF1CF7F7EAD31EBBBFC2E40CF3", pgp.EntityFingerprint(entity))
}

const signedMessage = `
-----BEGIN PGP SIGNED MESSAGE-----
Hash: SHA512

this is a test
-----BEGIN PGP SIGNATURE-----

iQEzBAEBCgAdFiEEi18odTLPQt8c9/fq0x67v8LkDPMFAlsDbBMACgkQ0x67v8Lk
DPMY5gf/XNzYK5CjFrUFMgfEujMffHBpU76LH9uc9iGi4W6wDE25wwM9o+ncqF9N
D50P5A1o5zhlWC0DgGwX/i2MKjdna05bjWTrgG6GPIRqsylPOsznFSjtuOOQymAa
+kqCTyqOByrrwYFChqdWbAXxzftsZMUA1H5M3P9hQFWnYMy8WUKTx/n+0DbebzYn
2iJsk2ZmkzwRRbx/y7oWv7Zl7DUjH8czdN6TZ7u2/kjJMAtMeLnO2BmdgmzGho76
O1Uk2WsCTL9skUyVgvgaBxwqcFkwTx+POX0hsx/14jebZrZPnJxWV0f3OlmtVtcP
XKFCklrFE+4eVrLTn0BoyomfNiTGQA==
=SmYp
-----END PGP SIGNATURE-----
`

const testKeyPublic = `
-----BEGIN PGP PUBLIC KEY BLOCK-----

mQENBFsDYR4BCADdQiKJCDpbxOKY03duqLUewMBQ2nyyzTEzEpo/sd1xSF7EL8ln
DQHLNp45sVBCbVYHav5NIq+hms4hBywK4AJ9031aXDPMTp1pkn8Pqt/M7Jf6+TSc
oyx+sXQaoHzFqEufV4IFJhz5XywSHpm4BC+VHDi7ooopyDGtOqAkB/7j9AXSMw17
k/RPVc4C9/Ifgz6Hu4sXK+KyMOPHav2r7BeSL4QVDr890KpMeX2LfVnh+xTn7K9/
3NHEwsO7iZnA2KvWU4VrdfClgtkHe6w8MByzKFWyDfuvruH47+GID0hFECp9FLNA
ENWu/Sszw0GXGo/5yzWWni67370HY59zb2X5ABEBAAG0EEF1dG9kZWIgVGVzdCBL
ZXmJAVQEEwEKAD4WIQSLXyh1Ms9C3xz39+rTHru/wuQM8wUCWwNhHgIbAwUJA8Jn
AAULCQgHAwUVCgkICwUWAgMBAAIeAQIXgAAKCRDTHru/wuQM850PCADFfM8ljLJg
TAbc+ecSzkG/Qt3Ll+BxQa2Yjuu8G1hy9fMsYULnLl9BGXit5sammRHFmH8GexuG
WXg9vaHUqBCyGqk7m8/VTmNE41VqH5kb3kwp4ev+7QbhGy0YSU7KW1+6BZ4FkU8P
x32w4JqK7wjAFZgo4QgGG+iblrOcfPvIbMkHv27FCQAMKUZ4MK2UGt8zDs22xL6y
RmqPVspizQ10Z+w65fUEP2IMb/F05JiBMIE6Z5t2xrcOL6wXp9zMRGjRFo2UKuAN
OWaU1AUU8tz9rY1WCYNveeLp3Yn+6LmLz6sIzHg5QXQeLB7weoMycBYrDH4Wn7uQ
WHjoIlsIlttmuQENBFsDYR4BCADFdiMZstgdDf0hbVMo6m3149BXkbE8IbJ4dmVD
KrkPc25QjI5cqU7DG8T9VlQTAqd37rnfGBS8hY7MXaFTEOjtFePm3Dvq8f6eb9SQ
VSFN0lLQOiERtv/UA1emheS2pTvrt8VqPlT9qbLUY76iAqrFQ3HPAg0Zs7P9tHEK
OnUgAV2ThQjVqulpekQbud0GvUkI2A+xyuPOdzh8su5lma6LT4EKGb18qXbbD6RY
AAZwg8/PXCdcQzW+753I/k4/y/zNQ0prdCNKsJqmWoHtOHvtB7ZAmC/sMRcmMT8r
NfJqby2C56gHFCgu4k6jfQ0HijvPAN1MZCKpUw8ecX0cnX1LABEBAAGJATwEGAEK
ACYWIQSLXyh1Ms9C3xz39+rTHru/wuQM8wUCWwNhHgIbDAUJA8JnAAAKCRDTHru/
wuQM8270B/0Yz0sWBp0GrGAaKE2tNGIOXliCavFx9mWR759gsKlgtrX4YW407Cgv
kTj986sau2VJxsPtmUogv/iR2xxAXAYnvpyBubIgXiSxGTxim2/cuBiziX1pAAU5
wZHTR7s+6MaLt9vtG0gnBOV0G6adQ29sVqKljNjbaV9yZAGFn75zS9Xk5GYNepvW
iKmdTOv/NgdnA3c7+3KNzhN9DoPBipD1OdfF8nhHfzGUE5kCeuF8qvTF83i61Pux
nlUNlIYGC7XdWqtxOR/4/lsecj6d3zPcnfTeqRQe6beKiVHrr/ciY4NUi2HeVuil
YfCBhFxCGhE3eni9G4RKZdT6CWTqanDy
=JZyf
-----END PGP PUBLIC KEY BLOCK-----
`

// Kept here so that we can modify the tests without generating a new key
const testKeyPrivate = `
-----BEGIN PGP PRIVATE KEY BLOCK-----

lQOYBFsDYR4BCADdQiKJCDpbxOKY03duqLUewMBQ2nyyzTEzEpo/sd1xSF7EL8ln
DQHLNp45sVBCbVYHav5NIq+hms4hBywK4AJ9031aXDPMTp1pkn8Pqt/M7Jf6+TSc
oyx+sXQaoHzFqEufV4IFJhz5XywSHpm4BC+VHDi7ooopyDGtOqAkB/7j9AXSMw17
k/RPVc4C9/Ifgz6Hu4sXK+KyMOPHav2r7BeSL4QVDr890KpMeX2LfVnh+xTn7K9/
3NHEwsO7iZnA2KvWU4VrdfClgtkHe6w8MByzKFWyDfuvruH47+GID0hFECp9FLNA
ENWu/Sszw0GXGo/5yzWWni67370HY59zb2X5ABEBAAEAB/wOIhSXKLujk4S+6tkY
7DVfkMPoFkC7YiCHqONJ+MhSnWMSSJQmnT3VV3iNal5cU+QYJrU2Q6Yw3jTXBsCB
NFRsaac676NOZGIpNXagz1Mq4Z0Gpsb9z21+7UJjlgdKT+j+tWYEgq38nt/vi11Z
1AifD3WvRJ9rMlK8BD/o9UrKkDQnh3W1xi5VKhq0Zy6h3fZeUbp0m/i/Yw8zbNdZ
AwEHlFawxWcRBu3CWJG9FAYVQOEtZQwMQOJQxqDhs+END2Gt+CmUBtMGj0WN+rYi
W9HpKHdAt2cm+v1zlfvCbU6RplxLiDo9C1bVzZ5mqzYKG8dKpaXzbw1rJgb6rwh6
GzkBBADlzaRtDIyIOdwTdDmmpcNzpsr5RivoZYPzpRCNmK67jCnW1z2nGm008CUN
Y6w6kYpxOdWpkN22r1e5/4v1T4XJuh81X/6yGEKyiRT7hCwbzzTsqOmTiOMZkooX
T/sTusNg6sqX0InMjoV9BJ5FrtqfSsrAcGylO7REC7fs7FKvgQQA9nsgAJjdMJeF
M+4CvNl2ZR5sN/ZuA+i7wxyOdkK0AhvcmKRBxyFyfAZYMZ41IKSDBRwitPL9qhwa
OwSzV3xhl1B78zQfeM6DRgxlfara+Ao5fyGX/kXiJ03dNw+TmgKjI2FtXEGNILFD
rEVWBAF0rNL2k1KuisHdxTo6obezcnkEAO8e0abzZi+r9DTCCJ1PiFLioXCc9DLg
0V5hisSRfz+T+XKuKCcNcLNmaa3nVqrisCSdOHDwa+g5EH0RASGM/ba4s/Vpk7ME
MLlgS0d53SVB2AqG+arH04INPlS1mbOsfB0Va4Lv5pcHTnHYWAi6eV3Vn/rulFDN
OgxkBJ93aKByOYa0EEF1dG9kZWIgVGVzdCBLZXmJAVQEEwEKAD4WIQSLXyh1Ms9C
3xz39+rTHru/wuQM8wUCWwNhHgIbAwUJA8JnAAULCQgHAwUVCgkICwUWAgMBAAIe
AQIXgAAKCRDTHru/wuQM850PCADFfM8ljLJgTAbc+ecSzkG/Qt3Ll+BxQa2Yjuu8
G1hy9fMsYULnLl9BGXit5sammRHFmH8GexuGWXg9vaHUqBCyGqk7m8/VTmNE41Vq
H5kb3kwp4ev+7QbhGy0YSU7KW1+6BZ4FkU8Px32w4JqK7wjAFZgo4QgGG+iblrOc
fPvIbMkHv27FCQAMKUZ4MK2UGt8zDs22xL6yRmqPVspizQ10Z+w65fUEP2IMb/F0
5JiBMIE6Z5t2xrcOL6wXp9zMRGjRFo2UKuANOWaU1AUU8tz9rY1WCYNveeLp3Yn+
6LmLz6sIzHg5QXQeLB7weoMycBYrDH4Wn7uQWHjoIlsIlttmnQOYBFsDYR4BCADF
diMZstgdDf0hbVMo6m3149BXkbE8IbJ4dmVDKrkPc25QjI5cqU7DG8T9VlQTAqd3
7rnfGBS8hY7MXaFTEOjtFePm3Dvq8f6eb9SQVSFN0lLQOiERtv/UA1emheS2pTvr
t8VqPlT9qbLUY76iAqrFQ3HPAg0Zs7P9tHEKOnUgAV2ThQjVqulpekQbud0GvUkI
2A+xyuPOdzh8su5lma6LT4EKGb18qXbbD6RYAAZwg8/PXCdcQzW+753I/k4/y/zN
Q0prdCNKsJqmWoHtOHvtB7ZAmC/sMRcmMT8rNfJqby2C56gHFCgu4k6jfQ0HijvP
AN1MZCKpUw8ecX0cnX1LABEBAAEAB/0eGwzu9h2NEHzvg8OSEWwCeWFieIwVJu4W
/7Ygr3rXqDnBfiyWZBnmFW3LUkYvP9BYsMUWBo3i0FodPolCKOnae+PrZtib0ZmI
fnRiLRiOzOpjZPl1wfjvUjMi6HcegcLJBZPLxwUeR6lESJDgEpgGy2mmriFhMczX
eZNig9cnsckJw247Xesm41Ju7P2m9AXbkOjRdzD9PvFNdFsyKNnYSnBrS4gmbm22
zG1TL6IdTpFed/dhiHboKcx9w1BfWHGZxQ5d+Oy1hmbBN1jn64BpFdAK42j55B0e
nMMiEsVrFu8EbC5uPUiBgkSy08EJ5wDM6qHz+xptym7mKNmXo3YBBADdp+Md3yRD
b2X3Xozah/GB82npgTt+cMfOqy/EBaua6hqbicGBVhMXgtnkhVT6VFd2R5LW/6Wx
iZseaC7YiGIe5k0PrLdUg/mHpWWgeDA9iyfGXdyH4TPlDYqcnVn3+1NHh2QZAokU
/RIgXmYC7uHm0yZkn6A8iitS07XVW5BwqwQA5A6RHGPe0iA936P64ZmadoscdhCV
5FR2rO8WUCBJl81efGmyiSq5HC/7bxO4O/Hu5PtaEtWMSccKl17EB5Up5NBzQQV3
3wITMshvZRATtpQi7m7XBcfV3jI5I7YV/v9/QH2Vhbt7HrQZZhgw+IeKv3by5ByL
x48lXFqFjRilZeED/i6UcgfnnErltpEUMxYLXmHOCkj+p0L0ms984HywSyN1QjfP
yr5xpks41B9tzJ0NOrJVFRT0/rm4Fjm2xsm1lAooqSpyj8MuF8ZqNV1TuvldfRLt
cSunV6h5QJjVzQEDLTkcfz5Xp49E5s03iSgSsSD87yBAGNJU/TdRZd4KTvqXP5CJ
ATwEGAEKACYWIQSLXyh1Ms9C3xz39+rTHru/wuQM8wUCWwNhHgIbDAUJA8JnAAAK
CRDTHru/wuQM8270B/0Yz0sWBp0GrGAaKE2tNGIOXliCavFx9mWR759gsKlgtrX4
YW407CgvkTj986sau2VJxsPtmUogv/iR2xxAXAYnvpyBubIgXiSxGTxim2/cuBiz
iX1pAAU5wZHTR7s+6MaLt9vtG0gnBOV0G6adQ29sVqKljNjbaV9yZAGFn75zS9Xk
5GYNepvWiKmdTOv/NgdnA3c7+3KNzhN9DoPBipD1OdfF8nhHfzGUE5kCeuF8qvTF
83i61PuxnlUNlIYGC7XdWqtxOR/4/lsecj6d3zPcnfTeqRQe6beKiVHrr/ciY4NU
i2HeVuilYfCBhFxCGhE3eni9G4RKZdT6CWTqanDy
=LSir
-----END PGP PRIVATE KEY BLOCK-----
`
