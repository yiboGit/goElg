package main

type DataPanel interface {
	GetAllDataPanel() DataPanelResp

	GetDataPanel() DataPanelResp

	// GetCustomerData 获取 客户数据（客户数）
	GetCustomerData() CusDataResp

	// GetStoreData 获取 门店数据 (门店数)
	GetStoreData() StoreDataResp

	GetEmployeeData() EmployeeDataResp

	// GetMemberData 获取 会员数据 (会员数)
	GetMemberData() MemDataResp

	// GetOptometryRecord 验光记录
	GetOptometryRecord() OptometryResp

	// GetExpensesRecord 获取 消费记录数与销售额
	GetExpensesRecord() ExpensesResp

	// GetTopUpData 获取 充值总额和充值记录
	GetDepositData() DepositDataResp

	// GetCardVouchersData 获取 卡券相关
	GetCardVouchersData() CardVouchersResp

	// 易吸粉数据
	GetYiFansData() YiFansResp

	// 海报数据
	GetPosterData() PosterResp

	// GetProductsData 商品数量
	GetProductsData() ProductsResp

	// GetOrderData 获取订单数量
	GetOrderData() OrderResp

	// TODO
	//WorkOrder  获取加工单数量
	GetWorkOrderData() WorkOrderResp

	// TODO
	// GetDeliverOrderData 获取配送单
	GetDeliverOrderData() DeliverOrderResp

	// TODO
	// GetSalesData 获取商品销售量
	GetSalesData() SalesResp
}

type DataPanelResp struct {
	CustomerDataResp CusDataResp      `json:"customerDataResp"`
	StoreDataResp    StoreDataResp    `json:"storeDataResp"`
	EmployeeDataResp EmployeeDataResp `json:"employeeDataResp"`
	MemDataResp      MemDataResp      `json:"memDataResp"`
	OptometryResp    OptometryResp    `json:"optometryResp"`
	ExpensesResp     ExpensesResp     `json:"expensesResp"`
	DepositDataResp  DepositDataResp  `json:"depositDataResp"`
	CardVouchersResp CardVouchersResp `json:"cardVouchersResp"`
	YiFansResp       YiFansResp       `json:"yiFansResp"`
	PosterResp       PosterResp       `json:"posterResp"`
	ProductsResp     ProductsResp     `json:"productsResp"`
	OrderResp        OrderResp        `json:"orderResp"`
	WorkOrderResp    WorkOrderResp    `json:"workOrderResp"`
	DeliverOrderResp DeliverOrderResp `json:"deliverOrderResp"`
	SalesResp        SalesResp        `json:"salesResp"`
	Error            string           `json:"error"`
}

// 按地区划分的客户数量
type CustomerData struct {
	Area           string `db:"Area" json:"area"`                     // 地区/大区
	CusCreateNum   int    `db:"CusCreateNum" json:"cusCreateNum"`     // 客户创建数
	StoreCreateNum int    `db:"StoreCreateNum" json:"storeCreateNum"` // 门店创建数
}

// 按类型划分的客户数量
type CustomerTypeData struct {
	Type         string `db:"Area" json:"type"`                 // 类型
	CusCreateNum int    `db:"CusCreateNum" json:"cusCreateNum"` // 客户创建数
}

type CusDataResp struct {
	CustomerDataList []CustomerData     `json:"customerDataList"` // 按地区划分的客户数量
	CusTypeDataList  []CustomerTypeData `json:"cusTypeDataList"`  // 按类型划分的客户数量
	TotalNumber      int                `json:"totalNumber"`
	Error            string             `json:"error"`
}

// 门店数据
type StoreData struct {
	Area     string `json:"area"`     // 地区
	StoreNum int    `json:"storeNum"` // 门店数
}

// 门店按照加工中心
type StoreCenterData struct {
	CenterName string `json:"centerName"` // 加工中心名称
	StoreNum   int    `json:"storeNum"`   // 门店数
}

// 门店返回体
type StoreDataResp struct {
	StoreDataList       []StoreData       `json:"storeDataList"`       // 门店数量
	StoreCenterDataList []StoreCenterData `json:"storeCenterDataList"` // 按照加工中心分类门店数量
	TotalNumber         int               `json:"totalNumber"`         // 门店总数
	Error               string            `json:"error"`
}

// 员工数返回体
type EmployeeDataResp struct {
	EmployeeNum int    `json:"employeeNum"` // 员工总数
	Error       string `json:"error"`
}

// 会员数据
type MemberData struct {
	AllFans        int `json:"allFans"`        // 所有粉丝会员(包括人在关注与取消关注的)
	BeenFans       int `json:"beenFans"`       // 还在关注的会员数
	UnFans         int `json:"unFans"`         // 已经取关的会员数
	MembershipCard int `json:"membershipCard"` // 会员卡会员
	Mall           int `json:"mall"`           // 商城会员
	Records        int `json:"records"`        // 手工记录
}

// 会员数据
type MemDataResp struct {
	MemberData     MemberData `json:"memberData"`     // 会员数据
	TotalMemberNum int        `json:"totalMemberNum"` // 会员总数量
	Error          string     `json:"error"`
}

// 验光记录
type OptometryData struct {
	OneLevel   int `json:"oneLevel"`   // 0<数值=<-3.5
	TwoLevel   int `json:"towLevel"`   // -3.5<数值=<-5.5
	ThreeLevel int `json:"threeLevel"` // 数值<-5.5
	PlusLevel  int `json:"PlusLevel"`  //数值大于0
}

// 验光记录
type OptometryResp struct {
	OptometryData     OptometryData `json:"optometryDataList"` // 验光记录数据
	TotalOptometryNum int           `json:"totalOptometryNum"` // 验光记录总数
	Error             string        `json:"error"`
}

// 充值记录数与销售额
type DepositData struct {
	DepositMethod    string  `json:"depositMethod"`
	DepositRecordNum int     `json:"depositRecordNum"`
	DepositMoney     float64 `json:"depositMoney"`
}

// 充值记录数与销售额
type DepositDataResp struct {
	DepositDataList       []DepositData `json:"depositDataList"`
	TotalDepositRecordNum int           `json:"totalDepositRecordNum"`
	TotalDepositMoney     float64       `json:"totalDepositMoney"`
	Error                 string        `json:"error"`
}

type ExpensesData struct {
	PayMethod string  `json:"payMethod"` // PayMethod 付款方式
	RecordNum int     `json:"recordNum"` // ZcConsumptionNum 消费记录数
	RealPay   float64 `json:"realPay"`   // RealPay 实付金额_元
	Price     float64 `json:"price"`     // price 商品原价总金额_元
}

// 销售额和销售记录
type ExpensesResp struct {
	ExpensesDataList []ExpensesData `json:"expensesDataList"` // 销售额数据
	Error            string         `json:"error"`
	TotalRecordNum   int            `json:"totalRecordNum"` // 记录总数
	TotalRealPay     float64        `json:"totalrealPay"`   // 总实际额
	TotalPrice       float64        `json:"totalPrice"`     // 总金额
}

type CardVouchersData struct {
	CardEvent string `json:"cardEvent"` // 卡券事件
	CardNum   int    `json:"cardNum"`   // 兑换券数
}

// 卡券相关
type CardVouchersResp struct {
	CardVouchersList []CardVouchersData `json:"cardVouchersList"`
	TotalCardNum     int                `json:"totalCardNum"` // 卡券数量
	Error            string             `json:"error"`
}

type YiFansData struct {
	ActivityState string `json:"activityState"`
	ActivityNum   int    `json:"activityNum"`
	EChannelNum   int    `json:"eChannel"`
	ActFansNum    int    `json:"actFansNum"`
}

// TODO
// 吸引粉丝
type YiFansResp struct {
	YiFansDataList []YiFansData `json:"yiFansDataList"`
	TotalChannel   int          `json:"totalChannel"`
	Error          string       `json:"error"`
}

type PosterData struct {
	ProsterCresteNum int `json:"prosterCresteNum"`
	ChanCreateNum    int `json:"chanCreateNum"`
	NewFansNum       int `json:"newFansNum"`
}

// 海报数据
type PosterResp struct {
	PosterData PosterData `json:"posterData"`
	Error      string     `json:"error"`
}

type ProductsData struct {
	GoodsLibNum  int `json:"goodsLib"`  //平台商品数字
	ScoreGoodNum int `json:"scoreGood"` //客户商品数
	ShopCartNum  int `json:"shopCart"`  //购物车商品数
}

// 商品数量
type ProductsResp struct {
	ProductsData  ProductsData `json:"productsDataList"`
	Error         string       `json:"error"`
	TotalProducts int          `json:"totalProducts"`
}

type OrderData struct {
	OrderStatus string  `json:"orderStatus"`
	OrderNum    int     `json:"orderNum"`
	OrderMoney  float64 `json:"orderMoney"`
}

// 订单量和订单额
type OrderResp struct {
	OrderDataList   []OrderData `json:"orderDataList"`
	TotalOrderNum   int         `json:"totalOrderNum"`
	TotalOrderMoney float64     `json:"totalOrderMoney"`
	Error           string      `json:"error"`
}

// TODO
type WorkOrderData struct {
	Classification string `json:"classification"`
	WorkOrderNum   int    `json:"workOrderNum"`
}

// TODO
// 加工单数
type WorkOrderResp struct {
	SourceOrderDataList []WorkOrderData `json:"sourceOrderDataList"`
	TotalSourceOrderNum int             `json:"totalSourceOrderNum"`
	CenterOrderDataList []WorkOrderData `json:"centerOrderDataList"`
	TotalCenterOrderNum int             `json:"totalCenterOrderNum"`
	StatusOrderDataList []WorkOrderData `json:"statusOrderDataList"`
	TotalStatusOrderNum int             `json:"totalStatusOrderNum"`
	Error               string          `json:"error"`
}

// TODO
type DeliverOrderData struct {
	Classification  string `json:"classification"`
	DeliverOrderNum int    `json:"deliverOrderNum"`
}

// 配送单
type DeliverOrderResp struct {
	CenterDeliverDataList []DeliverOrderData `json:"centerDeliverDataList"`
	StatusDeliverDataList []DeliverOrderData `json:"statusDeliverDataList"`
	TotalStatusDeliverNum int                `json:"totalStatusDeliverNum"`
	TotalCenterDeliverNum int                `json:"totalCenterDeliverNum"`
	Error                 string             `json:"error"`
}

// 销售量
type SalesData struct {
	Classification string `json:"classification"`
	GoodsNum       int    `json:"goodsNum"`
}

// TODO
type SalesResp struct {
	CenterSalesDataList []SalesData `json:"centerSalesDataList"`
	BrandSalesDataList  []SalesData `json:"brandSalesDataList"`
	TotalCenterSalesNum int         `json:"totalCenterSalesNum"`
	TotalBrandSalesNum  int         `json:"totalBrandSalesNum"`
	Error               string      `json:"error"`
}
