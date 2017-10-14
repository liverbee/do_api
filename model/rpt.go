package model

import (
	"github.com/jmoiron/sqlx"
	"log"
	"fmt"
)

type Rpt struct {
	DocYear 		int64 	`json:"doc_year" db:"DocYear"`
	DocMonth 		int64 	`json:"doc_month" db:"DocMonth"`
	DocWeek 		int64 	`json:"doc_week" db:"DocWeek"`
	DocDay 			int64 	`json:"doc_day" db:"DocDay"`
	DocNo 			string 	`json:"doc_no" db:"DocNo"`
	DocDate 		string 	`json:"doc_date" db:"DocDate"`
	DateSend 		string 	`json:"date_send" db:"DateSend"`
	DateReturn 		string 	`json:"date_return" db:"DateReturn"`
	TimeSend  		string 	`json:"time_send" db:"TimeSend"`
	TimeReturn  	string 	`json:"time_return" db:"TimeReturn"`
	MeasureStart  	string 	`json:"measure_start" db:"MeasureStart"`
	MeasureStop  	string 	`json:"measure_stop" db:"MeasureStop"`
	IsCancel  		int64 	`json:"is_cancel" db:"IsCancel"`
	CarLicence  	string 	`json:"car_licence" db:"CarLicence"`
	IsReturn  		int64 	`json:"is_return" db:"IsReturn"`
	SendResult  	int64 	`json:"send_result" db:"SendResult"`
	Branch  		string 	`json:"branch" db:"Branch"`
	EmpCode  		string 	`json:"emp_code" db:"EmpCode"`
	EmpName  		string 	`json:"emp_name" db:"EmpName"`
	Menus     		[]*RptSub `json:"menu"`
}

type RptSub struct {
	SoNo			string `json:"so_no" db:"SoNo"`
	ReceiveName 	string `json:"receive_name" db:"ReceiveName"`
	SaleCode 		string `json:"sale_code" db:"SaleCode"`
	SaleName 		string `json:"sale_name" db:"SaleName"`
	InvoiceNo 		string `json:"invoice_no" db:"InvoiceNo"`
	ArCode  		string `json:"ar_code" db:"ArCode"`
	ArName  		string `json:"ar_name" db:"ArName"`
}

func (r *Rpt) RptGetDoDetail(db *sqlx.DB, access_token string,date_start string,date_stop string,branch string) (err error) {
		sql := `SET datestyle = "ISO, DMY"; ` +
			` SELECT date_part('year'::text, a.docdate) AS docyear,date_part('month'::text, a.docdate) AS docmonth,` +
			` date_part('week'::text, a.docdate) AS docweek,date_part('day'::text, a.docdate) AS docday,` +
			` a.docno,a.docdate,a.datesend,a.timesend,a.datereturn,a.timereturn,a.measurestart,` +
			` a.measurestop,a.iscancel,a.carlicence,a.isreturn,a.sendresult,a.branch,d.empcode,e.name1 as empname` +
			` FROM sm_do.tb_do_delivery a` +
			` LEFT JOIN sm_do.tb_do_deliverysub b ON a.docno = b.docno` +
			` LEFT JOIN sm_do.tb_do_queue c ON b.sono = c.sono` +
			` LEFT JOIN sm_do.tb_do_delivery_empbplus as d on a.docno=d.docno and d.emp_position='1'` +
			` LEFT JOIN sm_do.tb_do_empbplus as e on d.empcode=e.code` +
			` WHERE a.iscancel='0' and a.docdate between '1/10/2017' and '2/10/2017'` +
			` group by date_part('year'::text, a.docdate),date_part('month'::text, a.docdate),` +
			` date_part('week'::text, a.docdate),date_part('day'::text, a.docdate),` +
			` a.docno,a.docdate,a.datesend,a.timesend,a.datereturn,a.timereturn,a.measurestart,` +
			` a.measurestop,a.iscancel,a.carlicence,a.isreturn,a.sendresult,a.branch,d.empcode,e.name1 order by a.docno`
		r.DocDate = date_start
		r.DocDate = date_stop
		r.Branch = branch
		err = db.Get(r, sql, r.DocDate, r.DocDate, r.Branch)
		log.Println(sql)

		if err != nil {
			log.Println("Error ", err.Error())
		}

		sqlsub := `select f.Id as MenuID,f.MenuCode,f.MenuName,ifnull(g.Id,0) as PermissionID,ifnull(g.IsCreate,0) as IsCreate` +
			` ,ifnull(g.IsUpdate,0) as IsUpdate,ifnull(g.IsRead,0) as IsRead,ifnull(g.IsDelete,0) as IsDelete` +
			` from User as a` +
			` left join UserRole as b on a.Id=b.UserId` +
			` left join Role as c on b.RoleId=c.Id` +
			` left join App as d on b.AppId=d.Id` +
			` left join Menu as f on d.Id=f.AppId` +
			` left join Permission as g on c.Id=g.RoleId and d.Id=g.AppId and f.Id=g.MenuId` +
			` where a.UserCode=? and a.Password=? and b.AppID=? and f.activestatus=1`

		fmt.Println(sqlsub)
		err = db.Select(&r.Menus, sqlsub, r.DocDate, r.DocDate, r.Branch)

		fmt.Println("Menus = ", r.DocDate, r.DocDate, r.Branch)
		if err != nil {
			log.Println("Error ", err.Error())
		}
		fmt.Println(r)
		return nil
}