package badwolf

import (
	"time"
	"context"
	"testing"
)

const (
	TestPath = "/var/tmp/mytest.tv"
	TestToolName = "mytoool"
	TestToolName2 = "mytoool"
	TestCategoryName = "testcategory"
)

func TestCreate(t *testing.T) {
	if err := CreateTimeVortex(TestPath); err != nil {
		t.Fatal("Create Failed. :", err)
	}
}

func TestExistCreate(t *testing.T) {
	if err := CreateTimeVortex(TestPath); err == nil {
		t.Fatal("Create Successed.")
	}
	return
}

func TestOpen(t *testing.T) {
	tv, err := OpenTimeVortex(nil, TestPath)
	if err != nil {
		t.Fatal("Open Failed. :", err)
	}
	tv.Close()
}

func TestAddData(t *testing.T) {
	tv, err := OpenTimeVortex(nil, TestPath)
	if err != nil {
		t.Fatal("Open Failed. :", err)
	}
	defer tv.Close()

	ns := testDummyNews()
	if _, err := tv.AddNewEvent(TestToolName, ns); err != nil {
		t.Fatal("Failed add news. :", err)
	}
}

func TestLookupAndDeleteData(t *testing.T) {
	tv, err := OpenTimeVortex(nil, TestPath)
	if err != nil {
		t.Fatal("Open Failed. :", err)
	}
	defer tv.Close()

	task_ctx := context.Background()
	st := time.Now().AddDate(0, 0, -1)
	et := time.Now().AddDate(0, 0, 1)

	evts, err := tv.Find(task_ctx, st, et, nil)
	if err != nil {
		t.Fatal("Failed event not found. :", err)
	}
	if len(evts) < 1 {
		t.Fatal("Failed event not found. : less than 1")
	}

	for _, evt := range evts {
		if err := tv.DeleteEvent(evt.Id()); err != nil {
			t.Fatal("Failed remove event. : ", err)
		}
	}
}

func TestUpdateCategory(t *testing.T) {
	tv, err := OpenTimeVortex(nil, TestPath)
	if err != nil {
		t.Fatal("Open Failed. :", err)
	}
	defer tv.Close()

	ns := testDummyNews()
	evt, err := tv.AddNewEvent(TestToolName, ns)
	if err != nil {
		t.Fatal("Failed add news. :", err)
	}
	eid := evt.Id()

	if err := tv.UpdateCategory(TestToolName2, TestCategoryName, [][]byte{eid}); err != nil {
		t.Fatal("Failed update category. : ", err)
	}
}

func TestDelete(t *testing.T) {
	if err := DeleteTimeVortex(TestPath); err != nil {
		t.Fatal("Delete Failed. :", err)
	}
}

func TestNotExistOpen(t *testing.T) {
	tv, err := OpenTimeVortex(nil, TestPath)
	if err == nil {
		tv.Close()
		t.Fatal("Open Successed.")
	}
	return
}
