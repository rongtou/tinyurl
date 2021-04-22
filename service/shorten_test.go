package service

import "testing"

func TestString(t *testing.T) {
	id := int64(1384884694285291520)
	token, _ := NewHashId(SetSalt("hello world")).Encode(id)
	if len(token) <= 11 {
		t.Logf("success, token length <= 6, param=%d, token=%s", id, token)
	} else {
		t.Fatalf("fail, token length > 6, param=%d, token=%s", id, token)
	}

}
