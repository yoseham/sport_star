package repositories

import (
	"app/project/datamodels"
	"fmt"
	"github.com/go-xorm/xorm"
	"log"
)

type SportstarRepository struct {
	engine *xorm.Engine
}

func NewSportstarRepository(engine *xorm.Engine) *SportstarRepository {
	return &SportstarRepository{engine: engine}
}
func (d *SportstarRepository) Get(id int) *datamodels.StarInfo {
	data := &datamodels.StarInfo{Id:id}
	fmt.Println(data)
	ok, err := d.engine.Get(data)
	if ok && err == nil {
		return data
	} else {
		return nil
	}
}
func (d *SportstarRepository) GetAll() []datamodels.StarInfo {
	dataList := make([]datamodels.StarInfo, 0)
	//dataList := []datamodels.StarInfo{}
	//id排序
	err := d.engine.Desc("id").Find(&dataList)
	if err == nil {
		return dataList
	}else {
		log.Fatalf("GetAll error: %s", err)
		return nil
	}
}
func (d *SportstarRepository) Search(country string) []datamodels.StarInfo {
	dataList := []datamodels.StarInfo{}
	err := d.engine.Where("country=?", country).Desc("id").Find(&dataList)
	if err == nil {
		return dataList
	}else {
		log.Fatalf("Search error: %s", err)
		return nil
	}
}
func (d *SportstarRepository) Delete(id int) error {
	data := &datamodels.StarInfo{Id: id, SysStatus: 1}
	_, err := d.engine.Id(data.Id).Update(data)
	return err
}
func (d *SportstarRepository) Update(info *datamodels.StarInfo, columns []string) error {
	_, err := d.engine.Id(info.Id).MustCols(columns...).Update(info)
	return err
}
func (d *SportstarRepository) Create(info *datamodels.StarInfo) error {
	_, err := d.engine.Insert(info)
	return err
}


