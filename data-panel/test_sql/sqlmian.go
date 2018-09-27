package main

import (
	"fmt"
	"time"
)

// var mysql, _ = utils.NewMysql(false, false)

type DeliverOrderDataa struct {
	Classification  string `json:"classification"`
	DeliverOrderNum int    `json:"deliverOrderNum"`
}

func main() {
	// msa.BuildService((*DataPanel)(nil), "DataPanelInstance", "dataPane", "../")
	// a := GetMemberData()
	// log.Printf("%v", a)
	beginData, endData := getDataRange()
	StatusDeliverSql = fmt.Sprintf(StatusDeliverSql, beginData, endData)
	// rowsTwo, err := mysql.Query(StatusDeliverSql)
	fmt.Println(StatusDeliverSql)
	// if err != nil {
	// 	log.Printf("%v", err)
	// }
	// defer rowsTwo.Close()
	// for rowsTwo.Next() {
	// 	var data DeliverOrderDataa
	// 	rowsTwo.Scan(&data.DeliverOrderNum, &data.DeliverOrderNum)
	// 	log.Printf("Classification : %s, DeliverOrderNum : %v \n", data.Classification, data.DeliverOrderNum)
	// }

}

var StatusDeliverSql = `select case when order_payment_status='0' then '未支付' 
when order_payment_status='1' then '已支付'
when order_payment_status='2' then '已取消' 
else null end as 'Classification',
count(id) as 'DeliverOrderNum'
from m_deliver
where deleted<>'1'
and create_time between '%s' and '%s'
group by case when order_payment_status='0' then '未支付' 
when order_payment_status='1' then '已支付'
when order_payment_status='2' then '已取消' 
else null end
`

func getDataRange() (string, string) {
	yesterday := time.Now().AddDate(0, 0, -1).Format("2006-01-02")
	beginDate := yesterday + " 00:00:00"
	endDate := yesterday + " 23:59:59"
	return beginDate, endDate
}
