/*
-------------------------------------
# @Time    : 2022/5/14 17:47:37
# @Author  : Giyn
# @Email   : giyn.jy@gmail.com
# @File    : tinyweb_test.go
# @Software: GoLand
-------------------------------------
*/

package tinyweb

import "testing"

func TestNestedGroup(t *testing.T) {
	r := New()
	v1 := r.Group("/v1")
	v2 := v1.Group("/v2")
	v3 := v2.Group("/v3")
	if v2.prefix != "/v1/v2" {
		t.Fatal("v2 prefix should be /v1/v2")
	}
	if v3.prefix != "/v1/v2/v3" {
		t.Fatal("v2 prefix should be /v1/v2")
	}
}
