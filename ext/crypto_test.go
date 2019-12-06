package msg

import (
	"crypto/rand"
	"git.ronaksoftware.com/river/pkg/db/scylla"
	"git.ronaksoftware.com/river/pkg/metrics"
	"git.ronaksoftware.com/river/pkg/tools"
	"testing"
	"time"
)

/*
   Creation Time: 2019 - Jan - 16
   Created by:  (ehsan)
   Maintainers:
      1.  Ehsan N. Moosa (E2)
   Auditor: Ehsan N. Moosa (E2)
   Copyright Ronak Software Group 2018
*/

func BenchmarkGenerateMessageKey(b *testing.B) {
	total := 10000
	plain := []byte("This is just a text to be encrypted. We are not changing the text")
	authKeys := make([][]byte, 0, total)
	for i := 0; i < total; i++ {
		x := make([]byte, 256)
		rand.Read(x)
		authKeys = append(authKeys, x)
	}
	b.StartTimer()
	b.RunParallel(func(pb *testing.PB) {
		var msgKey []byte
		for pb.Next() {
			_ = GenerateMessageKey(authKeys[tools.RandomInt(total)], plain, msgKey)
		}
	})

}

func BenchmarkEncrypt(b *testing.B) {
	total := 10000
	plain := []byte("This is just a text to be encrypted. We are not changing the text")
	authKeys := make([][]byte, 0, total)
	for i := 0; i < total; i++ {
		x := make([]byte, 256)
		rand.Read(x)
		authKeys = append(authKeys, x)
	}
	b.ReportAllocs()
	b.StartTimer()
	b.RunParallel(func(pb *testing.PB) {
		var msgKey []byte
		for pb.Next() {
			_ = GenerateMessageKey(authKeys[0], plain, msgKey)
			_, _ = Encrypt(authKeys[0], msgKey, plain)
		}
	})
}

func BenchmarkPool(b *testing.B) {
	metrics.Run("Test", "01", 2374)
	dbConf := scylla.DefaultConfig
	dbConf.Host = "localhost"
	dbConf.Keyspace = "river"
	scylla.Init(dbConf)

	bs := map[string]func(*testing.B){
		"1.Bytes":  benchBytes,
		"2.String": benchString,
	}
	for n, t := range bs {
		b.Run(n, t)
	}
}
func benchBytes(b *testing.B) {
	x := TestRequest{
		Payload: tools.StrToByte(tools.RandomID(1024)),
		Hash:    tools.StrToByte(tools.RandomID(100)),
	}
	xb, _ := x.Marshal()

	b.ReportAllocs()
	b.SetParallelism(10)
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			y := PoolTestRequest.Get()
			err := y.Unmarshal(xb)
			if err != nil {
				b.Fatal(err)
			}
			q := scylla.UsersUpdatesInsert.GetQuery()
			// CnUserID, CnUpdateID, CnConstructor, CnObject, CnCreatedOn
			q.Bind(tools.RandomInt64(0), tools.RandomInt64(0), 0, y.Payload, time.Now().Unix())
			err = q.Exec()
			if err != nil {
				b.Fatal(err)
			}
			scylla.UsersUpdatesInsert.Put(q)
			PoolTestRequest.Put(y)
		}
	})
}
func benchString(b *testing.B) {
	x := TestRequestWithString{
		Payload: tools.RandomID(1024),
		Hash:    tools.RandomID(100),
	}
	xb, _ := x.Marshal()

	b.ReportAllocs()
	b.SetParallelism(10)
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			y := PoolTestRequestWithString.Get()
			err := y.Unmarshal(xb)
			if err != nil {
				b.Fatal(err)
			}
			q := scylla.UsersUpdatesInsert.GetQuery()
			// CnUserID, CnUpdateID, CnConstructor, CnObject, CnCreatedOn
			q.Bind(tools.RandomInt64(0), tools.RandomInt64(0), 0, y.Payload, time.Now().Unix())
			err = scylla.Exec(q)
			if err != nil {
				b.Fatal(err)
			}
			scylla.UsersUpdatesInsert.Put(q)
			PoolTestRequestWithString.Put(y)
		}
	})
}
