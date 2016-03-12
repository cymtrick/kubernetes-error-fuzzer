/*
Copyright 2014 The Kubernetes Authors All rights reserved.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package testing

// You can use cfssl tool to generate certificates, please refer
// https://github.com/coreos/etcd/tree/master/hack/tls-setup for more details.
//
// ca-config.json:
//   expiry was changed from 1 year to 100 years (876000h)
// ca-csr.json:
//   ca expiry was set to 100 years (876000h) ("ca":{"expiry":"876000h"})
//   key was changed from ecdsa,384 to rsa,2048
// req-csr.json:
//   key was changed from ecdsa,384 to rsa,2048
//   hosts were changed to "localhost","127.0.0.1"
const CAFileContent = `
-----BEGIN CERTIFICATE-----
MIIEUDCCAzigAwIBAgIUKfV5+qwlw3JneAPdJS7JCO8xIlYwDQYJKoZIhvcNAQEL
BQAwgawxCzAJBgNVBAYTAlVTMSowKAYDVQQKEyFIb25lc3QgQWNobWVkJ3MgVXNl
ZCBDZXJ0aWZpY2F0ZXMxKTAnBgNVBAsTIEhhc3RpbHktR2VuZXJhdGVkIFZhbHVl
cyBEaXZpc29uMRYwFAYDVQQHEw1TYW4gRnJhbmNpc2NvMRMwEQYDVQQIEwpDYWxp
Zm9ybmlhMRkwFwYDVQQDExBBdXRvZ2VuZXJhdGVkIENBMCAXDTE2MDMxMjIzMTQw
MFoYDzIxMTYwMjE3MjMxNDAwWjCBrDELMAkGA1UEBhMCVVMxKjAoBgNVBAoTIUhv
bmVzdCBBY2htZWQncyBVc2VkIENlcnRpZmljYXRlczEpMCcGA1UECxMgSGFzdGls
eS1HZW5lcmF0ZWQgVmFsdWVzIERpdmlzb24xFjAUBgNVBAcTDVNhbiBGcmFuY2lz
Y28xEzARBgNVBAgTCkNhbGlmb3JuaWExGTAXBgNVBAMTEEF1dG9nZW5lcmF0ZWQg
Q0EwggEiMA0GCSqGSIb3DQEBAQUAA4IBDwAwggEKAoIBAQDP+acpr1USrObZFu+6
v+Bk6rYw+sWynP373cNUUiHfnZ3D7f9yJsDscV0Mo4R8DddqkxawrA5fK2Fm2Z9G
vvY5par4/JbwRIEkXmeM4e52Mqv0Yuoz62O+0jQvRawnCCJMcKuo+ijHMjmm0AF1
JdhTpTgvUwEP9WtY9JVTkfMCnDqZiqOU5D+d4YWUtkKqgQNvbZRs6wGubhMCZe8X
m+3bK8YAsWWtoFgr7plxXk4D8MLh+PqJ3oJjfxfW5A9dHbnSEmdZ3vrYwrKgyfNf
bvHE5qQmiSZUbUaCw3mKfaEMCNesPT46nBHxhAWc5aiL1tOXzvV5Uze7A7huPoI9
a3etAgMBAAGjZjBkMA4GA1UdDwEB/wQEAwIBBjASBgNVHRMBAf8ECDAGAQH/AgEC
MB0GA1UdDgQWBBQYc0xXQ6VNjFvIOqWfXorxx9rKRzAfBgNVHSMEGDAWgBQYc0xX
Q6VNjFvIOqWfXorxx9rKRzANBgkqhkiG9w0BAQsFAAOCAQEAaKyHDWYVjEyEKTXJ
qS9r46ehL5FZlWD2ZytBP8aHE307l9AfQ+DFWldCNaqMXLZozsresVaSzSOI6UUD
lCIQLDpPyxbpR320u8mC08+lhhwR/YRkrEqKHk56Wl4OaqoyWmguqYU9p0DiQeTU
sZsxOwG7cyEEvvs+XmZ/vBLBOr59xyjwn4seQqzwZj3VYeiKLw40iQt1yT442rcP
CfdlE9wTEONvWT+kBGMt0JlalXH3jFvlfcGQdDfRmDeTJtA+uIbvJhwJuGCNHHAc
xqC+4mAGBPN/dMPXpjayHD5dOXIKLfrNpqse6jImYlY9zduvwIHRDK/zvqTyPlNZ
uR84Nw==
-----END CERTIFICATE-----
`
const CertFileContent = `
-----BEGIN CERTIFICATE-----
MIIELzCCAxegAwIBAgIUcjkJA3cmHeoBQggaKZmfKebFL9cwDQYJKoZIhvcNAQEL
BQAwgawxCzAJBgNVBAYTAlVTMSowKAYDVQQKEyFIb25lc3QgQWNobWVkJ3MgVXNl
ZCBDZXJ0aWZpY2F0ZXMxKTAnBgNVBAsTIEhhc3RpbHktR2VuZXJhdGVkIFZhbHVl
cyBEaXZpc29uMRYwFAYDVQQHEw1TYW4gRnJhbmNpc2NvMRMwEQYDVQQIEwpDYWxp
Zm9ybmlhMRkwFwYDVQQDExBBdXRvZ2VuZXJhdGVkIENBMCAXDTE2MDMxMjIzMTQw
MFoYDzIxMTYwMjE3MjMxNDAwWjBVMRYwFAYDVQQKEw1hdXRvZ2VuZXJhdGVkMRUw
EwYDVQQLEwxldGNkIGNsdXN0ZXIxFTATBgNVBAcTDHRoZSBpbnRlcm5ldDENMAsG
A1UEAxMEZXRjZDCCASIwDQYJKoZIhvcNAQEBBQADggEPADCCAQoCggEBAOiW5A65
hWGbnwceoZHM0+OexU4cPF/FpP+7BOK5i7ymSWAqfKfNuio2TB1lAErC1oX7bgTX
ieP10uz3FYWQNrlDn0I4KSA888rFPtx8GwoxH/52fGlE80BUV9PNeOVP+mYza0ih
oFj2+PhXVL/JZbx9P/2RLSNbEnq+OPk8AN82SkNtpFzanwtpb3f+kt73878KNoQu
xYZaCF1sK45Kn7mjKSDu/b3xUbTrNwnyVAGOdLzI7CCWOu+ECoZYAH4ZNHHakbyY
eWQ7U9leocEOPlqxsQAKodaCYjuAaOFIcz8/W81q+3qNw/6GbZ4znjRKQ3OtIPZ4
JH1iNofCudWDp+0CAwEAAaOBnDCBmTAOBgNVHQ8BAf8EBAMCBaAwHQYDVR0lBBYw
FAYIKwYBBQUHAwEGCCsGAQUFBwMCMAwGA1UdEwEB/wQCMAAwHQYDVR0OBBYEFMJE
43qLCWhyZAE/wxNneSJw7aUVMB8GA1UdIwQYMBaAFBhzTFdDpU2MW8g6pZ9eivHH
2spHMBoGA1UdEQQTMBGCCWxvY2FsaG9zdIcEfwAAATANBgkqhkiG9w0BAQsFAAOC
AQEAuELC8tbmpyKlA4HLSDHOUquypNyiE6ftBIifJtp8bvBd+jiv4Pr8oVGxHoqq
48X7lamvDirLV5gmK0CxO+EXkIUHhULzPyYPynqsR7KZlk1PWghqsF65nwqcjS3b
tykLttD1AUDIozYvujVYBKXGxb6jcGM1rBF1XtslciFZ5qQnj6dTUujo9/xBA2ql
kOKiVXBNU8KFzq4c20RzHFLfWkbc30Q4XG4dTDVBeGupnFQRkZ0y2dSSU82QcLA/
HgAyQSO7+csN13r84zbmDuRpUgo6eTXzJ+77G19KDkEL7XEtlw2jB2L6/o+3RGtw
JLOpEsgi7hsvOYCuTA3Krw52Mw==
-----END CERTIFICATE-----
`
const KeyFileContent = `
-----BEGIN RSA PRIVATE KEY-----
MIIEpAIBAAKCAQEA6JbkDrmFYZufBx6hkczT457FThw8X8Wk/7sE4rmLvKZJYCp8
p826KjZMHWUASsLWhftuBNeJ4/XS7PcVhZA2uUOfQjgpIDzzysU+3HwbCjEf/nZ8
aUTzQFRX08145U/6ZjNrSKGgWPb4+FdUv8llvH0//ZEtI1sSer44+TwA3zZKQ22k
XNqfC2lvd/6S3vfzvwo2hC7FhloIXWwrjkqfuaMpIO79vfFRtOs3CfJUAY50vMjs
IJY674QKhlgAfhk0cdqRvJh5ZDtT2V6hwQ4+WrGxAAqh1oJiO4Bo4UhzPz9bzWr7
eo3D/oZtnjOeNEpDc60g9ngkfWI2h8K51YOn7QIDAQABAoIBAQCj88Fc08++x0kp
ZqEzunPebsvcTLEOPa8aiUVfYLWszHbKsAhg7Pb+zHmI+upiyMcZeOvLw/eyVlVR
rrZgCRFaNN2texMaY3zigXnXSDBzVb+cyv7V4cGqpgmnBp7i3ia/Jh3I/A2gyK8l
t8HI03nAjXWvE0gDNS5okXBt16sxq6ZWyzHHVbN3UYtCDxnyh2Ibck4b+K8I8Bn1
mwMsSqPXJS1UQ3U5UqcaMs7WOEGx+xmaPJTWm5Lb//BkakGuBTQj+7wotyXQYG5U
uZdPPcFRk6cqgjzUeKVUtGkdmfgHSTdIwZowkKibB4rdrudsRnSwfeB+83Jp9JwG
JPrGvsbNAoGBAPULIO+vVBZXVpUEAhvNSXtmOi/hAbQhOuix8iwHbJ5EbrWaDn4B
Reb2cw/fZGgGG4jtAOXdiY8R1XGGP8+RPZ5El10ZWnNrKQfpZ27gK/5yeq5dfGBG
4JLUpcrT180FJo00rgiQYJnHCk1fWrnzXNV6K08ZZHGr6yv4S/jbq/7vAoGBAPL9
NTN/UWXWFlSHVcb2dFHcvIiPwRj9KwhuMu90b/CilBbSJ1av13xtf2ar5zkrEtWH
CB3q3wBaklQP9MfOqEWGZeOUcd9AbYWtxHjHmP5fJA9RjErjlTtqGkusNtZJbchU
UWfT/Tl9pREpCvJ/8iawc1hx7sHHKzYwnDnMaQbjAoGAfJdd9cBltr5NjZLuJ4in
dhCyQSncnePPegUQJwbXWVleGQPtnm+zRQ3Fzyo8eQ+x7Frk+/s6N/5PUlt6EmW8
uL4TYAjGDq1LvXQVXTCp7cPzULjDxogDI2Tvr0MrFFksEtvYKQ6Pr2CeglybWrS8
XOazIpK8mXdaKY8jwbKfrw0CgYAFnfrb3OaZzxAnFhXSiqH3vn2RPpl9JWUYRcvh
ozRvQKLhwCvuohP+KV3XlsO6m5dM3lk+r85F6NIXJWNINyvGp6u1ThovygJ+I502
GY8c2kAwJndyx74MaJCBDVMbMwlZpzFWkBz7dj8ZnXRGVNTZNh0Ef2XAjwUdtJP3
9hS7dwKBgQDCzq0RIxFyy3F5baGHWLVICxmhNExQ2+Vebh+DvsPKtnz6OrWdRbGX
wgGVLrn53s6eCblnXLtKr/Li+t7fS8IkQkvu5guOvI9VeVUmZhFET3GVmUxu+JTb
iQY4uBgaf8Fgay4dkOfjvlOpFDR4E7UbJpg8/cFKTrpwgOiUVyFVdQ==
-----END RSA PRIVATE KEY-----
`
