package models

import (
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"time"
)

func ProcessPartition(schemaName, tableName string) (err error) {

	o := orm.NewOrm()

	beego.Info("Start process partition for:" + tableName)

	var list orm.ParamsList

	sql := "select REPLACE(partition_name,'p','') from INFORMATION_SCHEMA.PARTITIONS where TABLE_SCHEMA=? and table_name=? order by partition_ordinal_position "

	paramList := []interface{}{schemaName, tableName}

	_, err = o.Raw(sql, paramList).ValuesFlat(&list)

	if err != nil {
		beego.Critical(err.Error())
		return err
	}
	strFirst := fmt.Sprintf("%v", list[0])
	strLast := fmt.Sprintf("%v", list[len(list)-1])

	layout := "20060102"
	first, err := time.Parse(layout, strFirst)
	last, err := time.Parse(layout, strLast)
	beforeDuration := time.Now().Sub(first)
	afterDuration := last.Sub(time.Now())

	if beforeDuration.Hours() > 360.0 {
		deleteSql := "ALTER TABLE " + tableName + " DROP PARTITION " + "p" + strFirst
		_, err = o.Raw(deleteSql).Exec()
		if err != nil {
			beego.Critical(err.Error())
		} else {
			beego.Info("Drop Table Partition:" + tableName + ":" + "p" + strFirst)
		}
	}

	if afterDuration.Hours() < 144.0 {
		newLast := last.Add(time.Hour * 24)

		addSql := "ALTER TABLE " + tableName + " ADD PARTITION (PARTITION " + "p" + newLast.Format("20060102") + " VALUES LESS THAN (TO_DAYS ('" + newLast.Format("2006-01-02") + "')))"
		_, err = o.Raw(addSql).Exec()
		if err != nil {
			beego.Critical(err.Error())
		} else {
			beego.Info("Add Table Partition:" + tableName + ":" + "p" + newLast.Format("20060102"))

		}
	}

	return err
}
