package September

import (
	"encoding/json"
	"testing"
	"time"
)

type People interface {
	Say()
}
type Man struct{}

func (stu *Man) Say() {}

func Live() People {
	var man Man
	return &man
}

func TestInterface(t *testing.T) {
	obj := Live()
	if obj == nil {
		t.Log("nil")
	} else {
		//ob = (September.People|*September.Man)
		t.Log("not nil")
	}
}

func Test(t *testing.T) {

	t1 := struct {
		time.Time
		N int
	}{
		time.Date(2023, 9, 19, 11, 9, 0, 0, time.UTC),
		6,
	}
	b, _ := json.Marshal(t1)
	t.Logf("b=%v\n", string(b))
	t.Logf("b=%s", b)
}
