package repositories

import (
	"app/project/datasource"
	"fmt"
	"testing"
)

func TestSportstarRepository_GetAll(t *testing.T) {
	dao := NewSportstarRepository(datasource.InstanceMaster())
	if dao.GetAll() == nil {
		t.Error("数据为空")
	} else {
		fmt.Println(dao.GetAll())
	}
}
