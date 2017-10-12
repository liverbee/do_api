package model

import (
	"github.com/jmoiron/sqlx"
	"log"
	"fmt"
)

type Rpt struct {
	Id 	 int64 `json:"id" db:"Id"`
	UserCode string `json:"usercode" db:"UserCode"`
	UserName string `json:"username" db:"UserName"`
	Password string `json:"password" db:"Password"`
	UserActiveStatus int64 `json:"usercctivestatus" db:"UserActiveStatus"`
	RoleId 	 int64 `json:"roleid" db:"RoleId"`
	RoleCode string `json:"rolecode" db:"RoleCode"`
	RoleName string `json:"rolename" db:"RoleName"`
	AppID  int64 `json:"appid" db:"AppID"`
	AppCode  string `json:"appcode" db:"AppCode"`
	AppName  string `json:"appname" db:"AppName"`
	Menus     []*RptSub `json:"menu"`
}

type RptSub struct {
	MenuID int64 `json:"menuid" db:"MenuID"`
	MenuCode string `json:"menucode" db:"MenuCode"`
	MenuName string `json:"menuname" db:"MenuName"`
	PermissionID int64 `json:"permissionid" db:"PermissionID"`
	IsCreate int64 `json:"is_create" db:"IsCreate"`
	IsUpdate  int64 `json:"is_update" db:"IsUpdate"`
	IsRead  int64 `json:"is_read" db:"IsRead"`
	IsDelete  int64 `json:"is_delete" db:"IsDelete"`
}


func (r *Rpt) RptGetDoDetail(db *sqlx.DB, access_token string,user_code string,password string,appid int64) (err error) {
	sql := `select a.Id,a.UserCode,a.UserName,a.Password,a.ActiveStatus as UserActiveStatus,c.id as RoleId,c.RoleCode,c.RoleName,b.AppID,d.AppCode,d.AppName`+
		` from User as a`+
		` left join UserRole as b on a.Id=b.UserId`+
		` left join Role as c on b.RoleId=c.Id`+
		` left join App as d on b.AppId=d.Id`+
		` where a.UserCode=? and a.Password=? and b.AppID=? limit 1`
	r.UserCode = user_code
	r.Password = password
	r.AppID = appid
	err = db.Get(r,sql,r.UserCode,r.Password,r.AppID)
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
	err = db.Select(&r.Menus,sqlsub,r.UserCode,r.Password,r.AppID)

	fmt.Println("Menus = ", r.UserCode,r.Password,r.AppID)
	if err != nil {
		log.Println("Error ", err.Error())
	}
	fmt.Println(l)

	return nil
}