package boltutil

import (
	"bytes"
	"testing"
)

func TestToBytesFromString(t *testing.T) {
	b, err := ToBytes("aa")
	if err != nil {
		t.Error(err)
	}

	//t.Log(b)
	if bytes.Compare([]byte("aa"), b) == -1 {
		t.Errorf("expected %v but result: %v", []byte("aa"), b)
	}
}

func TestToBytesFromInt(t *testing.T) {
	b, err := ToBytes(1)
	if err != nil {
		t.Error(err)
	}

	//t.Log(b)
	if bytes.Compare([]byte{0, 0, 0, 0, 0, 0, 0, 1}, b) == -1 {
		t.Errorf("expected %v but result: %v", []byte{0, 0, 0, 0, 0, 0, 0, 1}, b)
	}
}

type example struct {
	Str string
	Num int
	B   bool
	M   map[string]string
}

func TestSerializeStruct(t *testing.T) {
	e := example{
		Str: "aaa",
		Num: 100,
		B:   true,
		M: map[string]string{
			"hoge": "hoge111",
		},
	}

	b, err := Serialize(e)
	if err != nil {
		t.Error(err)
	}

	//t.Log(b)

	var ret example
	err = Deserialize(b, &ret)
	if err != nil {
		t.Error(err)
	}

	//t.Log(ret)
	if e.Str != ret.Str {
		t.Errorf("expected %v but result: %v", e.Str, ret.Str)
	}
}
