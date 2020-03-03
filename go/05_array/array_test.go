package _5_array

import "testing"

func TestInsert(t *testing.T) {
	capacity := 10
	arr := NewArray(uint(capacity))
	for i := 0; i < capacity-2; i++ {
		err := arr.Insert(uint(i), i+1)
		if err != nil {
			t.Fatal(err.Error())
		}
	}
	arr.Print()
	t.Log("*****")
	err := arr.Insert(uint(10), 999)
	if err != nil {
		t.Log("insert err", err.Error())
	}
	arr.Print()
	err = arr.InsertToTail(666)
	if err != nil {
		t.Log("insert err", err.Error())
	}
	arr.Print()
}

func TestDelete(t *testing.T) {
	capacity := 10
	arr := NewArray(uint(capacity))
	for i :=0; i < capacity; i++ {
		err := arr.Insert(uint(i), i + 1)
		if err != nil {
			t.Fatal(err.Error())
		}
	}
	arr.Print()
	for i := 9; i >= 0; i-- {
		_, err := arr.Delete(uint(i))
		if err != nil {
			t.Fatal(err)
		}
		arr.Print()
	}
}

func TestFind(t *testing.T) {
	capacity := 10
	arr := NewArray(uint(capacity))
	for i := 0; i < capacity; i++ {
		err := arr.Insert(uint(i), i+1)
		if err != nil {
			t.Fatal(err.Error())
		}
	}
	arr.Print()
	t.Log(arr.Find(0))
	t.Log(arr.Find(9))
	t.Log(arr.Find(11))

}

