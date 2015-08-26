package models

import (
	//"adexchange/lib"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	//"strings"
	"bytes"
	"time"
)

func ProcessPartition(schemaName, tableName string, delFlg bool) (err error) {

	o := orm.NewOrm()

	beego.Info("Start process partition for:" + tableName)

	var list orm.ParamsList

	sql := "select REPLACE(partition_name,'p','') from INFORMATION_SCHEMA.PARTITIONS where TABLE_SCHEMA=? and table_name=? order by partition_ordinal_position "

	paramList := []interface{}{schemaName, tableName}

	_, err = o.Raw(sql, paramList).ValuesFlat(&list)

	if err != nil {
		beego.Critical(err.Error())
		return
	}
	if list == nil {
		beego.Info("No Partition for table :" + tableName)
		return
	}

	strFirst := fmt.Sprintf("%v", list[0])
	strLast := fmt.Sprintf("%v", list[len(list)-1])

	layout := "20060102"
	first, err := time.Parse(layout, strFirst)
	last, err := time.Parse(layout, strLast)
	beforeDuration := time.Now().Sub(first)
	afterDuration := last.Sub(time.Now())

	if delFlg {
		if beforeDuration.Hours() > 360.0 {
			deleteSql := "ALTER TABLE " + tableName + " DROP PARTITION " + "p" + strFirst
			_, err = o.Raw(deleteSql).Exec()
			if err != nil {
				beego.Critical(err.Error())
			} else {
				beego.Info("Drop Table Partition:" + tableName + ":" + "p" + strFirst)
			}
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

func UpdateForeverAllocation(lastDay string, currentDay string) (err error) {
	o := orm.NewOrm()

	type IdResult struct {
		Id int
	}

	beego.Info("Process forever adspace last day:" + lastDay + " current day:" + currentDay)
	var ids []IdResult
	paramList := []interface{}{currentDay, lastDay}

	var buffer bytes.Buffer
	buffer.WriteString("UPDATE pmp_daily_allocation SET ad_date = ? WHERE ad_date =? ")

	num, err := o.Raw("SELECT id FROM pmp_adspace WHERE forever_flg = 1").QueryRows(&ids)
	if err == nil && num > 0 {
		//aryStrId := make([]string, num)
		buffer.WriteString("and pmp_adspace_id in (0")
		for _, record := range ids {
			//aryStrId[i] = lib.ConvertIntToString(record.Id)
			buffer.WriteString(",?")
			paramList = append(paramList, record.Id)
		}
		buffer.WriteString(")")

		//strIds := strings.Join(aryStrId, ",")
		//p, err := o.Raw("UPDATE pmp_daily_allocation SET ad_date = ? WHERE ad_date = ? and pmp_adspace_id in (?)").Prepare()
		//res, err := p.Exec(currentDay, lastDay, strIds)
		//res, err := o.Raw("UPDATE pmp_daily_allocation SET ad_date = ? WHERE ad_date =? and pmp_adspace_id in (?)", currentDay, lastDay, 7).Exec()
		res, err := o.Raw(buffer.String(), paramList).Exec()
		if err == nil {
			num, _ := res.RowsAffected()
			beego.Info("Update allocation records: ", num)

		} else {
			beego.Critical(err.Error())

		}
	}

	return err
}
