package queue

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_queue_Push_Correct_Length(t *testing.T) {

	q := New()
	q.Push(1)
	q.Push(2)
	q.Push(3)

	want := 3
	got := q.Len()

	if want != got {
		t.Errorf("incorrect length of the Queue: we want %v but got %v", want, got)
	}
}

func Test_queue_Push_Correct_Length_With_Nil(t *testing.T) {

	q := New()
	q.Push(1)
	q.Push(2)
	q.Push(nil)
	q.Push(3)

	want := 3
	got := q.Len()

	if want != got {
		t.Errorf("incorrect length of the Queue: we want %v but got %v", want, got)
	}
}

func Test_queue_Push_Correct_First_Element(t *testing.T) {

	q := New()
	q.Push(4)
	q.Push(3)
	q.Push(2)
	q.Push(1)

	want := 4
	got := q.List.Front().Value

	if want != got {
		t.Errorf("incorrect first element in the Queue: we want %v but got %v", want, got)
	}
}

func Test_queue_Pop_Nil(t *testing.T) {

	q := New()

	got := q.Pop()
	assert.EqualValues(t, nil, got, "incorrect element on pop: we want %v but got %v", nil, got)

	q.Pop()
	want := 0
	gotLen := q.Len()

	if want != gotLen {
		t.Errorf("incorrect length of the Queue: we want %v but got %v", want, got)
	}

}

func Test_queue_Pop_Length(t *testing.T) {

	q := New()
	q.Push(1)
	q.Push(2)
	q.Push(3)
	q.Push(4)
	q.Pop()

	want := 3
	got := q.Len()

	if want != got {
		t.Errorf("incorrect length of the Queue: we want %v but got %v", want, got)
	}
}

func Test_queue_Pop_Removes_Elements(t *testing.T) {

	q := New()
	q.Push(1)
	q.Push(2)
	q.Push(3)
	q.Push(4)

	want := 1
	got := q.Pop()

	if want != got {
		t.Errorf("incorrect element on pop: we want %v but got %v", want, got)
	}
	want = 2
	got = q.Pop()

	if want != got {
		t.Errorf("incorrect element on pop: we want %v but got %v", want, got)
	}

	want = 2
	got = q.Len()

	if want != got {
		t.Errorf("incorrect length of the Queue: we want %v but got %v", want, got)
	}

}

func Test_Duplicate_Is_Pushed_Back(t *testing.T) {

	q := New()
	q.Push(2)
	q.Push(5)
	q.Push(1)
	q.Push(3)
	q.Push(5)

	want := 2
	got := q.Pop()

	if want != got {
		t.Errorf("incorrect element on pop: we want %v but got %v", want, got)
	}

	want = 1
	got = q.Pop()

	if want != got {
		t.Errorf("incorrect element on pop: we want %v but got %v", want, got)
	}

	want = 3
	got = q.Pop()

	if want != got {
		t.Errorf("incorrect element on pop: we want %v but got %v", want, got)
	}

	want = 5
	got = q.Pop()

	if want != got {
		t.Errorf("incorrect element on pop: we want %v but got %v", want, got)
	}

	want = 0
	got = q.Len()

	if want != got {
		t.Errorf("incorrect length of the Queue after popping: we want %v but got %v", want, got)
	}

}

func Test_Unique_For_Repeating_Two_Objects(t *testing.T) {

	q := New()
	q.Push(1)
	q.Push(2)
	q.Push(1)
	q.Push(2)
	q.Push(1)

	want := 2
	gotLen := q.Len()

	if want != gotLen {
		t.Errorf("incorrect element on pop: we want %v but got %v", want, gotLen)
	}

	want = 2
	got := q.Pop()

	if want != got {
		t.Errorf("incorrect element on pop: we want %v but got %v", want, got)
	}

	want = 1
	got = q.Pop()

	if want != got {
		t.Errorf("incorrect element on pop: we want %v but got %v", want, got)
	}

	got = q.Pop()
	assert.EqualValues(t, nil, got, "incorrect element on pop: we want %v but got %v", nil, got)

}

func Test_Unique_For_Repeating_One_Object(t *testing.T) {

	q := New()
	q.Push(1)
	q.Push(1)
	q.Push(1)
	q.Push(1)
	q.Push(1)

	want := 1
	gotLen := q.Len()

	if want != gotLen {
		t.Errorf("incorrect length of the Queue: we want %v but got %v", want, gotLen)
	}

	q.Pop()
	got := q.Pop()
	assert.EqualValues(t, nil, got, "incorrect element on pop: we want %v but got %v", nil, got)

}
