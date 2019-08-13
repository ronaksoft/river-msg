package admin

import (
	"encoding/json"
	"testing"
)

func TestAdmin(t *testing.T) {
	usr := User{
		UserID:   3435,
		Username: "username",
		Phone:    "+9823232",
	}
	b, err := usr.Marshal()
	if err != nil {
		t.Fatal(err)
	}

	tmpusr1 := &User{}
	if err := tmpusr1.Unmarshal(b); err != nil {
		t.Fatal(err)
	}
	t.Logf("%#v", tmpusr1)

	b, err = json.Marshal(usr)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(string(b))

	tmpusr2 := &User{}
	if err = json.Unmarshal(b, tmpusr2); err != nil {
		t.Fatal(err)
	}
	t.Logf("%#v", tmpusr2)
}
