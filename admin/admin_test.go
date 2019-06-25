package admin

import (
	"encoding/json"
	"testing"
)

func TestAdmin(t *testing.T) {
	usr := User{
		ID:        "someid",
		UserID:    3435,
		Username:  "username",
		FirstName: "fname",
		Phone:     "+9823232",
		PhotoBig: &FileLocation{
			FileID:     22,
			ClusterID:  9393,
			AccessHash: 8322,
		},
		AuthIDs: []int64{1, 2, 3},
		Bio:     "bio my",
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

	usr.AuthIDs = nil
	usr.PhotoBig = nil
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
