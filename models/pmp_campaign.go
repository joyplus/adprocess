package models

import (
	"github.com/astaxie/beego/orm"
)

type PmpCampaign struct {
	Id              int     `orm:"column(id);auto"`
	GroupId         int     `orm:"column(group_id)"`
	Name            string  `orm:"column(name);size(45)"`
	StartDate       string  `orm:"column(start_date);type(date);null"`
	EndDate         string  `orm:"column(end_date);type(date);null"`
	CampaignStatus  int     `orm:"column(campaign_status)"`
	DemandAdspaceId int     `orm:"column(demand_adspace_id)"`
	ImpTrackingUrl  string  `orm:"column(imp_tracking_url);size(1000);null"`
	ClkTrackingUrl  string  `orm:"column(clk_tracking_url);size(1000);null"`
	LandingUrl      string  `orm:"column(landing_url);size(1000);null"`
	AdType          int     `orm:"column(ad_type);null"`
	CampaignType    int     `orm:"column(campaign_type);null"`
	AccurateType    int     `orm:"column(accurate_type);null"`
	PricingType     int     `orm:"column(pricing_type);null"`
	StrategyType    int     `orm:"column(strategy_type);null"`
	BudgetType      int     `orm:"column(budget_type);null"`
	Budget          int     `orm:"column(budget);null"`
	BidPrice        float32 `orm:"column(bid_price);null"`
	Imp             int     `orm:"column(imp)"`
	Clk             int     `orm:"column(clk)"`
}

func init() {
	orm.RegisterModel(new(PmpCampaign))
}
