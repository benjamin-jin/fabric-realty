package model

type Storage struct {
	ObjectType      string  `json:"docType"`
	UserID          string  `json:"user_id"`
	GoodCode        string  `json:"good_code"`
	StorageStatus   int     `json:"storage_status"`
	Count           int     `json:"count"`
	GoodName        string  `json:"good_name"`
	GoodGrossWeight float64 `json:"good_gross_weight"`
	UserName        string  `json:"user_name"`
	UserTel         string  `json:"user_tel"`
	UserAddress     string  `json:"user_address"`
	LocationCode    string  `json:"location_code"`
	LocationName    string  `json:"location_name"`
	OrgID           string  `json:"org_id"`
	OrgName         string  `json:"org_name"`
}

type Demand struct {
	ObjectType      string  `json:"docType"`
	UserID          string  `json:"user_id"`
	GoodCode        string  `json:"good_code"`
	DemandStatus    int     `json:"storage_status"`
	Count           int     `json:"count"`
	GoodName        string  `json:"good_name"`
	GoodGrossWeight float64 `json:"good_gross_weight"`
	UserName        string  `json:"user_name"`
	UserTel         string  `json:"user_tel"`
	UserAddress     string  `json:"user_address"`
	LocationCode    string  `json:"location_code"`
	LocationName    string  `json:"location_name"`
	OrgID           string  `json:"org_id"`
	OrgName         string  `json:"org_name"`
}

type Transport struct {
	ObjectType        string  `json:"docType"`
	OrgID             string  `json:"org_id"`
	OrgName           string  `json:"org_name"`
	StartLocationCode string  `json:"start_location_code"`
	StartLocationName string  `json:"start_location_name"`
	EndLocationCode   string  `json:"end_location_code"`
	EndLocationName   string  `json:"end_location_name"`
	TransportStatus   int     `json:"transport_status"`
	TransportType     string  `json:"transport_type"`
	TransportCount    int     `json:"transport_count"`
	TransportTonnage  float64 `json:"transport_tonnage"`
	TransportCapacity int     `json:"transport_capacity"`
	TransportDuration float64 `json:"transport_duration"`
	TransportCost     int     `json:"transport_cost"`
}

type Proposal struct {
	ObjectType        string        `json:"docType"`
	AdminUserID       string        `json:"admin_user_id"`
	AdminUserName     string        `json:"admin_user_name"`
	SubProposalList   []SubProposal `json:"sub_proposal_list"`
	AlgSubList        []SubProposal `json:"alg_sub_list"`
	AlgDuration       int           `json:"alg_duration"`
	AlgCost           int           `json:"alg_cost"`
	IsProposalConfirm bool          `json:"is_proposal_confirm"`
	ProposalDuration  float64       `json:"proposal_duration"`
	ProposalCost      int           `json:"proposal_cost"`
	CreateTime        string        `json:"create_time"`
}

type SubProposal struct {
	SubProposalID        int          `json:"sub_proposal_id"`
	StorageUserId        string       `json:"storage_user_id"`
	StorageUserName      string       `json:"storage_user_name"`
	StorageUserTel       string       `json:"storage_user_tel"`
	StorageUserAddress   string       `json:"storage_user_address"`
	StorageOrgID         string       `json:"storage_org_id"`
	StorageOrgName       string       `json:"storage_org_name"`
	IsStorageConfirm     bool         `json:"is_storage_confirm"`
	TransportOrgID       string       `json:"transport_org_id"`
	TransportOrgName     string       `json:"transport_org_name"`
	IsTransportConfirm   bool         `json:"is_transport_confirm"`
	StartLocationCode    string       `json:"start_location_code"`
	StartLocationName    string       `json:"start_location_name"`
	EndLocationCode      string       `json:"end_location_code"`
	EndLocationName      string       `json:"end_location_name"`
	TransportType        string       `json:"transport_type"`
	TransportCount       int          `json:"transport_count"`
	TransportTonnage     float64      `json:"transport_tonnage"`
	TransportCapacity    int          `json:"transport_capacity"`
	SubProposalDuration  float64      `json:"sub_proposal_duration"`
	SubProposalCost      int          `json:"sub_proposal_cost"`
	DemandList           []DemandItem `json:"demand_list"`
	AllDemandConfirm     bool         `json:"all_demand_confirm"`
	IsSubProposalConfirm bool         `json:"is_sub_proposal_confirm"`
}

type DemandItem struct {
	DemandUserID      string  `json:"demand_user_id"`
	DemandUserName    string  `json:"demand_user_name"`
	DemandUserTel     string  `json:"demand_user_tel"`
	DemandUserAddress string  `json:"demand_user_address"`
	DemandOrgID       string  `json:"demand_org_id"`
	DemandOrgName     string  `json:"demand_org_name"`
	IsDemandConfirm   bool    `json:"is_demand_confirm"`
	GoodCode          string  `json:"good_code"`
	GoodName          string  `json:"good_name"`
	GoodGrossWeight   float64 `json:"good_gross_weight"`
	DemandCount       int     `json:"demand_count"`
}

// Selling 销售要约
// 需要确定ObjectOfSale是否属于Seller
// 买家初始为空
// Seller和ObjectOfSale一起作为复合键,保证可以通过seller查询到名下所有发起的销售
type Selling struct {
	ObjectOfSale  string  `json:"objectOfSale"`  //销售对象(正在出售的房地产RealEstateID)
	Seller        string  `json:"seller"`        //发起销售人、卖家(卖家AccountId)
	Buyer         string  `json:"buyer"`         //参与销售人、买家(买家AccountId)
	Price         float64 `json:"price"`         //价格
	CreateTime    string  `json:"createTime"`    //创建时间
	SalePeriod    int     `json:"salePeriod"`    //智能合约的有效期(单位为天)
	SellingStatus string  `json:"sellingStatus"` //销售状态
}

// SellingStatusConstant 销售状态
var SellingStatusConstant = func() map[string]string {
	return map[string]string{
		"saleStart": "销售中", //正在销售状态,等待买家光顾
		"cancelled": "已取消", //被卖家取消销售或买家退款操作导致取消
		"expired":   "已过期", //销售期限到期
		"delivery":  "交付中", //买家买下并付款,处于等待卖家确认收款状态,如若卖家未能确认收款，买家可以取消并退款
		"done":      "完成",  //卖家确认接收资金，交易完成
	}
}

// Donating 捐赠要约
// 需要确定ObjectOfDonating是否属于Donor
// 需要指定受赠人Grantee，并等待受赠人同意接收
type Donating struct {
	ObjectOfDonating string `json:"objectOfDonating"` //捐赠对象(正在捐赠的房地产RealEstateID)
	Donor            string `json:"donor"`            //捐赠人(捐赠人AccountId)
	Grantee          string `json:"grantee"`          //受赠人(受赠人AccountId)
	CreateTime       string `json:"createTime"`       //创建时间
	DonatingStatus   string `json:"donatingStatus"`   //捐赠状态
}

// DonatingStatusConstant 捐赠状态
var DonatingStatusConstant = func() map[string]string {
	return map[string]string{
		"donatingStart": "捐赠中", //捐赠人发起捐赠合约，等待受赠人确认受赠
		"cancelled":     "已取消", //捐赠人在受赠人确认受赠之前取消捐赠或受赠人取消接收受赠
		"done":          "完成",  //受赠人确认接收，交易完成
	}
}
