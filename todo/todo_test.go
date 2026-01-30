package todo

import (
	"io/ioutil"
	"os"
	"testing"
)

func TestAdd(t *testing.T) {
	l := List{}

	taskName := "New Task"

	l.Add(taskName)

	if l[0].Task != taskName {
		t.Errorf("Add: task name should be %s, got %s", taskName, l[0].Task)
	}

}

func TestComplete(t *testing.T) {
	l := List{}
	taskName := "New Task"
	l.Add(taskName)

	if l[0].Task != taskName {
		t.Errorf(" task name should be %s, got %s", taskName, l[0].Task)
	}

	if l[0].Done {
		t.Errorf("New task should not be done")
	}

	err := l.Completed(1)

	if !l[0].Done {
		t.Errorf("New task should be done: %s", err)
	}
}

func TestDelete(t *testing.T) {
	l := List{}

	tasks := []string{
		"New Task1",
		"New Task2",
		"New Task3",
	}

	for _, task := range tasks {
		l.Add(task)
	}

	if l[0].Task != tasks[0] {
		t.Errorf("error1: task name should be %s, got %s", tasks[0], l[0].Task)
	}

	l.Delete(2)

	if len(l) != 2 {
		t.Errorf("Delete: length should be %d, got %d", 2, len(l))
	}

	if l[1].Task != tasks[2] {
		t.Errorf("error2: task name should be %s, got %s", tasks[2], l[1].Task)
	}
}

func TestSaveGet(t *testing.T) {
	l1 := List{}
	l2 := List{}

	taskName := "New Task"
	l1.Add(taskName)

	if l1[0].Task != taskName {
		t.Errorf(" task name should be %s, got %s", taskName, l1[0].Task)
	}

	tf, err := ioutil.TempFile("", "")
	if err != nil {
		t.Fatalf("Error creating temp file: %s", err)
	}
	defer os.Remove(tf.Name())
	if err := l1.Save(tf.Name()); err != nil {
		t.Fatalf("Error saving list to temp file: %s", err)
	}

	if err := l2.Get(tf.Name()); err != nil {
		t.Fatalf("Error getting list from temp file: %s", err)
	}

	if l1[0].Task != l2[0].Task {
		t.Errorf(" task %q should match %q task", l1[0].Task, l2[0].Task)
	}
}
