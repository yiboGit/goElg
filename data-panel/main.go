package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"time"

	"eglass.com/brisk"
	"eglass.com/utils"
	"github.com/go-redis/redis"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

var (
	// ContainerPort os获取配置文件上容器端口
	ContainerPort   = os.Getenv("ContainerPort")
	mysql, sqlerror = utils.NewMysql(true, false) //TODO
	dataPane        = DataPanelInstance{
		e: echo.New(),
	}
	redisClient = utils.NewRedisClient(false)
)

type DataPanelInstance struct {
	e *echo.Echo
}

// 客户数据/客户数
func (c *DataPanelInstance) GetCustomerData() CusDataResp {

	beginDate, endDate := getDataRange()
	var (
		customerSql      = CustomerSql
		customerCaseWhen = CustomerCaseWhen
	)
	var totalNumber int
	customerSql = fmt.Sprintf(customerSql, customerCaseWhen, beginDate, endDate, customerCaseWhen)
	// 大区分布查询
	var customerDataList []CustomerData
	rows, err := mysql.Query(customerSql)
	if err != nil {
		log.Printf("get customer data has error : %v", err)
		return CusDataResp{Error: err.Error()}
	}
	defer rows.Close()
	for rows.Next() {
		var data CustomerData
		err := rows.Scan(&data.Area, &data.CusCreateNum, &data.StoreCreateNum)
		if err != nil {
			return CusDataResp{Error: err.Error()}
		}
		customerDataList = append(customerDataList, data)
		totalNumber = totalNumber + data.CusCreateNum
		log.Printf("Area: %s, CusCreateNum: %d, StoreCreateNum: %d", data.Area, data.CusCreateNum, data.StoreCreateNum)
	}
	// 类型 分布查询
	var (
		cusTypeDataList []CustomerTypeData
		customerTypeSql = CustomerTypeSql
	)
	customerTypeSql = fmt.Sprintf(customerTypeSql, beginDate, endDate)
	rowst, err := mysql.Query(customerTypeSql)
	if err != nil {
		log.Printf("get customer data has error : %v", err)
		return CusDataResp{Error: err.Error()}
	}
	defer rowst.Close()
	for rowst.Next() {
		var (
			data         CustomerTypeData
			customerType sql.NullString
		)

		err := rowst.Scan(&customerType, &data.CusCreateNum)
		if err != nil {
			return CusDataResp{Error: err.Error()}
		}
		data.Type = customerType.String
		cusTypeDataList = append(cusTypeDataList, data)
	}

	return CusDataResp{
		CustomerDataList: customerDataList,
		CusTypeDataList:  cusTypeDataList,
		TotalNumber:      totalNumber,
	}

}

// 门店数
func (c *DataPanelInstance) GetStoreData() StoreDataResp {
	var (
		totalNumber   int
		storeSql      = StoreSql
		storeCaseWhen = StoreCaseWhen
	)
	beginDate, endDate := getDataRange()
	storeSql = fmt.Sprintf(storeSql, storeCaseWhen, beginDate, endDate, storeCaseWhen)
	var storeDataList []StoreData
	rows, err := mysql.Query(storeSql)
	fmt.Println(storeSql)
	if err != nil {
		return StoreDataResp{Error: err.Error()}
	}
	defer rows.Close()
	for rows.Next() {
		var data StoreData
		err := rows.Scan(&data.Area, &data.StoreNum)
		if err != nil {
			return StoreDataResp{Error: err.Error()}
		}
		storeDataList = append(storeDataList, data)
		totalNumber += data.StoreNum
		log.Printf("Area: %s, StoreNum: %d", data.Area, data.StoreNum)

	}

	// 按照加工中心分类
	var (
		storeCenterDataList []StoreCenterData
		storeCenterSql      = StoreCenterSql
	)
	storeCenterSql = fmt.Sprintf(storeCenterSql, beginDate, endDate)
	rows, err = mysql.Query(storeCenterSql)
	if err != nil {
		return StoreDataResp{Error: err.Error()}
	}
	defer rows.Close()
	for rows.Next() {
		var (
			data       StoreCenterData
			centerName sql.NullString
		)
		err = rows.Scan(&centerName, &data.StoreNum)
		if err != nil {
			return StoreDataResp{Error: err.Error()}
		}
		data.CenterName = centerName.String
		storeCenterDataList = append(storeCenterDataList, data)
		log.Printf("CenterName: %v, StoreNum: %d", data.CenterName, data.StoreNum)
	}
	return StoreDataResp{
		StoreDataList:       storeDataList,
		StoreCenterDataList: storeCenterDataList,
		TotalNumber:         totalNumber,
	}
}

// 员工数
func (c *DataPanelInstance) GetEmployeeData() EmployeeDataResp {
	beginDate, endDate := getDataRange()
	var (
		employeeData int
		employeeSql  = EmployeeSql
	)
	employeeSql = fmt.Sprintf(employeeSql, beginDate, endDate)
	rows, err := mysql.Query(employeeSql)
	if err != nil {
		return EmployeeDataResp{Error: err.Error()}
	}
	defer rows.Close()
	for rows.Next() {
		rows.Scan(&employeeData)
		log.Printf("EmployeeNum: %d", employeeData)
	}
	return EmployeeDataResp{
		EmployeeNum: employeeData,
	}
}

// 会员数
func (c *DataPanelInstance) GetMemberData() MemDataResp {
	beginDate, endDate := getDataRange()
	var (
		memberData     MemberData
		totalMemberNum int
		memberSql      = MemberSql
	)

	memberSql = fmt.Sprintf(memberSql, beginDate, endDate)
	rows, err := mysql.Query(memberSql)
	fmt.Println(memberSql)
	if err != nil {
		return MemDataResp{Error: err.Error()}
	}
	defer rows.Close()
	for rows.Next() {
		rows.Scan(&memberData.BeenFans, &memberData.UnFans, &memberData.MembershipCard, &memberData.Mall, &memberData.Records, &totalMemberNum)
		fmt.Println(memberData)
	}
	memberData.AllFans = memberData.BeenFans + memberData.UnFans
	return MemDataResp{
		MemberData:     memberData,
		TotalMemberNum: totalMemberNum,
	}
}

// 验光记录数
func (c *DataPanelInstance) GetOptometryRecord() OptometryResp {
	beginDate, endDate := getDataRange()
	var (
		optometrySql = OptometrySql
	)
	optometrySql = fmt.Sprintf(optometrySql, beginDate, endDate)
	rows, err := mysql.Query(optometrySql)
	var data OptometryData
	var optometryNum int
	if err != nil {
		return OptometryResp{Error: err.Error()}
	}
	defer rows.Close()
	for rows.Next() {
		rows.Scan(&data.OneLevel, &data.TwoLevel, &data.ThreeLevel, &data.PlusLevel, &optometryNum)
	}
	return OptometryResp{
		OptometryData:     data,
		TotalOptometryNum: optometryNum,
	}
}

// 消费记录数与销售额
func (c *DataPanelInstance) GetExpensesRecord() ExpensesResp {
	var (
		totalRecord      int
		totalRealPay     float64
		totalPrice       float64
		expensesSql      = ExpensesSql
		expensesCaseWhen = ExpensesCaseWhen
		consumptionSql   = ConsumptionSql
	)

	beginDate, endDate := getDataRange()
	expensesSql = fmt.Sprintf(expensesSql, expensesCaseWhen, beginDate, endDate, expensesCaseWhen)
	consumptionSql = fmt.Sprintf(consumptionSql, beginDate, endDate)
	// log.Println(ExpensesSql)
	var expensesDataList []ExpensesData
	rows, err := mysql.Query(expensesSql)
	if err != nil {
		return ExpensesResp{Error: err.Error()}
	}
	defer rows.Close()
	for rows.Next() {
		var data ExpensesData
		err := rows.Scan(&data.PayMethod, &data.RecordNum, &data.RealPay, &data.Price)
		if err != nil {
			return ExpensesResp{Error: err.Error()}
		}
		expensesDataList = append(expensesDataList, data)
		totalRecord += data.RecordNum
		totalRealPay += data.RealPay
		totalPrice += data.Price
		log.Printf("PayMethod : %s, RecordNum : %v, RealPay : %v, TotalPrice: %v", data.PayMethod, data.RecordNum, data.RealPay, data.Price)
	}
	rowsTwo, err := mysql.Query(consumptionSql)
	if err != nil {
		return ExpensesResp{Error: err.Error()}
	}
	defer rowsTwo.Close()
	for rowsTwo.Next() {
		var data ExpensesData
		err := rowsTwo.Scan(&data.PayMethod, &data.RecordNum, &data.RealPay, &data.Price)
		if err != nil {
			return ExpensesResp{Error: err.Error()}
		}
		expensesDataList = append(expensesDataList, data)
		totalRecord += data.RecordNum
		totalRealPay += data.RealPay
		totalPrice += data.Price
		log.Printf("PayMethod : %s, RecordNum : %v, RealPay : %v, TotalPrice: %v", data.PayMethod, data.RecordNum, data.RealPay, data.Price)
	}

	return ExpensesResp{
		ExpensesDataList: expensesDataList,
		TotalRealPay:     totalRealPay,
		TotalPrice:       totalPrice,
		TotalRecordNum:   totalRecord,
	}
}

func (c *DataPanelInstance) GetDepositData() DepositDataResp {
	var (
		totalRecords    int
		totalMoney      float64
		depositSql      = DepositSql
		depositCaseWhen = DepositCaseWhen
		depositDataList []DepositData
	)
	beginDate, endDate := getDataRange()
	depositSql = fmt.Sprintf(depositSql, depositCaseWhen, beginDate, endDate, beginDate, endDate, depositCaseWhen)
	rows, err := mysql.Query(depositSql)
	defer rows.Close()
	if err != nil {
		return DepositDataResp{Error: err.Error()}
	}
	for rows.Next() {
		var data DepositData
		rows.Scan(&data.DepositMethod, &data.DepositRecordNum, &data.DepositMoney)
		totalRecords += data.DepositRecordNum
		totalMoney += data.DepositMoney
		depositDataList = append(depositDataList, data)
		log.Printf("DepositMethod : %s, DepositRecordNum : %v, DepositMoney : %v, totalRecords: %v, totalMoney: %v \n", data.DepositMethod, data.DepositRecordNum, data.DepositMoney, totalRecords, totalMoney)
	}

	return DepositDataResp{
		DepositDataList:       depositDataList,
		TotalDepositMoney:     totalMoney,
		TotalDepositRecordNum: totalRecords,
	}
}

// 卡券相关
func (c *DataPanelInstance) GetCardVouchersData() CardVouchersResp {
	var (
		allCardSql   = AllCardSql
		cardSql      = CardSql
		cardCaseWhen = CardCaseWhen
		cardDataList []CardVouchersData
		totalCardNum int
	)
	beginDate, endDate := getDataRange()
	allCardSql = fmt.Sprintf(allCardSql, beginDate, endDate)
	cardSql = fmt.Sprintf(cardSql, cardCaseWhen, beginDate, endDate, cardCaseWhen)

	rows, err := mysql.Query(allCardSql)
	defer rows.Close()
	if err != nil {
		return CardVouchersResp{Error: err.Error()}
	}
	for rows.Next() {
		rows.Scan(&totalCardNum)
		log.Printf("TotalCardNum : %d \n", totalCardNum)
	}

	rowsTow, err := mysql.Query(cardSql)
	fmt.Printf("SQL : %s \n ", cardSql)
	defer rowsTow.Close()
	if err != nil {
		return CardVouchersResp{Error: err.Error()}
	}
	for rowsTow.Next() {
		var data CardVouchersData
		rowsTow.Scan(&data.CardEvent, &data.CardNum)
		cardDataList = append(cardDataList, data)
		log.Printf("CardEvent : %s, CardNum : %d", data.CardEvent, data.CardNum)
	}

	return CardVouchersResp{
		CardVouchersList: cardDataList,
		TotalCardNum:     totalCardNum,
	}
}

// 全部易吸粉数据，与时间无关
func (c *DataPanelInstance) GetYiFansData() YiFansResp {
	var (
		activitySql   = ActivityNumSql
		channelNumSql = ChannelNumSql
		fansNumSql    = FansNumSql
		_, endDate    = getDataRange()
		ingData       YiFansData
		endData       YiFansData
		totalChannel  int
	)

	aMonthAgo := fmt.Sprintf("%s 00:00:00", time.Now().AddDate(0, 0, -30).Format("2006-01-02"))
	activitySql = fmt.Sprintf(activitySql, aMonthAgo, endDate)
	channelNumSql = fmt.Sprintf(channelNumSql, aMonthAgo, endDate)
	fansNumSql = fmt.Sprintf(fansNumSql, aMonthAgo, endDate)
	rows, err := mysql.Query(activitySql)
	fmt.Printf(activitySql)
	defer rows.Close()
	if err != nil {
		return YiFansResp{Error: err.Error()}
	}
	for rows.Next() {
		var (
			state       string
			activityNum int
		)
		rows.Scan(&state, &activityNum)
		if state == "进行中" {
			ingData.ActivityState = state
			ingData.ActivityNum = activityNum
		} else {
			endData.ActivityState = state
			endData.ActivityNum = activityNum
		}
		log.Printf("ActivityState : %s, ActivityNum : %d", state, activityNum)
	}

	rowsTwo, err := mysql.Query(channelNumSql)
	fmt.Printf(channelNumSql)
	defer rowsTwo.Close()
	if err != nil {
		return YiFansResp{Error: err.Error()}
	}
	for rowsTwo.Next() {
		var (
			state       string
			eChannelNum int
		)
		rowsTwo.Scan(&state, &eChannelNum)
		if state == "进行中" {
			ingData.EChannelNum = eChannelNum
		} else {
			endData.EChannelNum = eChannelNum
		}
		totalChannel += eChannelNum
		log.Printf("ActivityState : %s,  EChannelNum : %d", state, eChannelNum)
	}

	rowsThree, err := mysql.Query(fansNumSql)
	fmt.Printf(fansNumSql)
	defer rowsThree.Close()
	if err != nil {
		return YiFansResp{Error: err.Error()}
	}
	for rowsThree.Next() {
		var (
			state      string
			actFansNum int
		)
		rowsThree.Scan(&state, &actFansNum)
		fmt.Printf("fansNumber: %d \n", actFansNum)
		if state == "进行中" {
			ingData.ActFansNum = actFansNum
		} else {
			endData.ActFansNum = actFansNum
		}
		fmt.Printf("ActivityState : %s,  ActFansNum : %d \n", state, actFansNum)
	}

	return YiFansResp{
		YiFansDataList: []YiFansData{ingData, endData},
		TotalChannel:   totalChannel,
	}
}

func (c *DataPanelInstance) GetPosterData() PosterResp {
	var (
		posterNumSql       = PosterNumSql
		channelNumSql      = ChanNumSql
		fansNumSql         = FanNumSql
		beginDate, endDate = getDataRange()
		posterData         PosterData
	)
	posterNumSql = fmt.Sprintf(posterNumSql, beginDate, endDate)
	channelNumSql = fmt.Sprintf(channelNumSql, beginDate, endDate)
	fansNumSql = fmt.Sprintf(fansNumSql, beginDate, endDate)
	rows, err := mysql.Query(posterNumSql)
	fmt.Println(posterNumSql)
	if err != nil {
		return PosterResp{Error: err.Error()}
	}
	defer rows.Close()
	for rows.Next() {
		rows.Scan(&posterData.ProsterCresteNum)
		fmt.Printf("ProsterCresteNum: %d, ChanCreateNum: %d, NewFansNum: %d", posterData.ProsterCresteNum, posterData.ChanCreateNum, posterData.NewFansNum)
	}

	rowst, err := mysql.Query(channelNumSql)
	if err != nil {
		return PosterResp{Error: err.Error()}
	}
	defer rowst.Close()
	for rowst.Next() {
		// &posterData.NewFansNum
		rowst.Scan(&posterData.ChanCreateNum)
		fmt.Printf(" ChanCreateNum: %d", posterData.ChanCreateNum)

	}

	rowsts, err := mysql.Query(fansNumSql)
	if err != nil {
		return PosterResp{Error: err.Error()}
	}
	defer rowsts.Close()
	for rowsts.Next() {
		// &posterData.NewFansNum
		rowsts.Scan(&posterData.NewFansNum)
		fmt.Printf(" NewFansNum: %d", posterData.NewFansNum)

	}
	return PosterResp{
		PosterData: posterData,
	}
}

func (c *DataPanelInstance) GetProductsData() ProductsResp {
	var (
		goodsLibSql        = GoodsLibSql
		scoreGoodsSql      = ScoreGoodsSql
		shopCartSql        = ShopCartSql
		productData        ProductsData
		beginDate, endDate = getDataRange()
		totalPro           int
	)

	goodsLibSql = fmt.Sprintf(goodsLibSql, beginDate, endDate)
	rows, err := mysql.Query(goodsLibSql)
	if err != nil {
		return ProductsResp{Error: err.Error()}
	}
	defer rows.Close()
	for rows.Next() {
		rows.Scan(&productData.GoodsLibNum)
		totalPro += productData.GoodsLibNum
		log.Printf("GoodsLibNum : %d \n", productData.GoodsLibNum)
	}

	scoreGoodsSql = fmt.Sprintf(scoreGoodsSql, beginDate, endDate)
	rowsx, err := mysql.Query(scoreGoodsSql)
	defer rowsx.Close()
	if err != nil {
		return ProductsResp{Error: err.Error()}
	}
	for rowsx.Next() {
		rowsx.Scan(&productData.ScoreGoodNum)
		totalPro += productData.ScoreGoodNum
		log.Printf("ScoreGoodNum : %d \n", productData.ScoreGoodNum)
	}

	shopCartSql = fmt.Sprintf(shopCartSql, beginDate, endDate)
	rowsy, err := mysql.Query(shopCartSql)
	defer rowsy.Close()
	if err != nil {
		return ProductsResp{Error: err.Error()}
	}
	for rowsy.Next() {
		rowsy.Scan(&productData.ShopCartNum)
		totalPro += productData.ShopCartNum
		log.Printf("ShopCartNum : %d \n", productData.ShopCartNum)
	}

	return ProductsResp{
		ProductsData:  productData,
		TotalProducts: totalPro,
	}
}

func (c *DataPanelInstance) GetOrderData() OrderResp {
	var (
		payStatusSql       = PayStatusSql
		checkStatusSql     = CheckStatusSql
		orderDataList      []OrderData
		beginDate, endDate = getDataRange()
		totalOrderNum      int
		totalOrderMoney    float64
	)

	payStatusSql = fmt.Sprintf(payStatusSql, beginDate, endDate)
	checkStatusSql = fmt.Sprintf(checkStatusSql, beginDate, endDate)

	rows, err := mysql.Query(payStatusSql)
	defer rows.Close()
	if err != nil {
		return OrderResp{Error: err.Error()}
	}
	for rows.Next() {
		var data OrderData
		rows.Scan(&data.OrderStatus, &data.OrderNum, &data.OrderMoney)
		log.Printf("OrderStatus: %s, OrderNum : %d, OrderMoney : %f \n", data.OrderStatus, data.OrderNum, data.OrderMoney)
		orderDataList = append(orderDataList, data)
	}

	rowst, err := mysql.Query(checkStatusSql)
	defer rowst.Close()
	if err != nil {
		return OrderResp{Error: err.Error()}
	}
	for rowst.Next() {
		var data OrderData
		rowst.Scan(&data.OrderStatus, &data.OrderNum, &data.OrderMoney)
		log.Printf("OrderStatus: %s, OrderNum : %d, OrderMoney : %f \n", data.OrderStatus, data.OrderNum, data.OrderMoney)
		orderDataList = append(orderDataList, data)
		totalOrderNum += data.OrderNum
		totalOrderMoney += data.OrderMoney
	}

	return OrderResp{
		OrderDataList:   orderDataList,
		TotalOrderNum:   totalOrderNum,
		TotalOrderMoney: totalOrderMoney,
	}
}

func (c *DataPanelInstance) GetWorkOrderData() WorkOrderResp {
	var (
		beginData, endData = getDataRange()
		sourceOrderSql     = SourceOrderSql
		centerOrderSql     = CenterOrderSql
		statusOrderSql     = StatusOrderSql
		sourceOrderList    []WorkOrderData
		centerOrderList    []WorkOrderData
		statusOrderList    []WorkOrderData
		totalSourceNum     int
		totalCentereNum    int
		totalStatusNum     int
	)

	sourceOrderSql = fmt.Sprintf(sourceOrderSql, beginData, endData)
	centerOrderSql = fmt.Sprintf(centerOrderSql, beginData, endData)
	statusOrderSql = fmt.Sprintf(statusOrderSql, beginData, endData)

	rows, err := mysql.Query(sourceOrderSql)
	if err != nil {
		return WorkOrderResp{
			Error: err.Error(),
		}
	}
	defer rows.Close()
	for rows.Next() {
		var data WorkOrderData
		rows.Scan(&data.Classification, &data.WorkOrderNum)
		log.Printf("Classification : %s, WorkOrderNum : %v \n", data.Classification, data.WorkOrderNum)
		sourceOrderList = append(sourceOrderList, data)
		totalSourceNum += data.WorkOrderNum
	}

	rowsTwo, err := mysql.Query(centerOrderSql)
	if err != nil {
		return WorkOrderResp{
			Error: err.Error(),
		}
	}
	defer rowsTwo.Close()
	for rowsTwo.Next() {
		var data WorkOrderData
		rowsTwo.Scan(&data.Classification, &data.WorkOrderNum)
		log.Printf("Classification : %s, WorkOrderNum : %v \n", data.Classification, data.WorkOrderNum)
		centerOrderList = append(centerOrderList, data)
		totalCentereNum += data.WorkOrderNum
	}

	rowsThree, err := mysql.Query(statusOrderSql)
	if err != nil {
		return WorkOrderResp{
			Error: err.Error(),
		}
	}
	defer rowsThree.Close()
	for rowsThree.Next() {
		var data WorkOrderData
		rowsThree.Scan(&data.Classification, &data.WorkOrderNum)
		log.Printf("Classification : %s, WorkOrderNum : %v \n", data.Classification, data.WorkOrderNum)
		statusOrderList = append(statusOrderList, data)
		totalStatusNum += data.WorkOrderNum
	}

	return WorkOrderResp{
		SourceOrderDataList: sourceOrderList,
		CenterOrderDataList: centerOrderList,
		StatusOrderDataList: statusOrderList,
		TotalCenterOrderNum: totalCentereNum,
		TotalSourceOrderNum: totalSourceNum,
		TotalStatusOrderNum: totalStatusNum,
	}
}

func (c *DataPanelInstance) GetDeliverOrderData() DeliverOrderResp {
	var (
		beginData, endData = getDataRange()
		centerDeliverSql   = CenterDeliverSql
		statusDeliverSql   = StatusDeliverSql
		centerDeliverList  []DeliverOrderData
		statusDeliverList  []DeliverOrderData
		totalCenterNum     int
		totalStatusNum     int
	)

	centerDeliverSql = fmt.Sprintf(centerDeliverSql, beginData, endData)
	statusDeliverSql = fmt.Sprintf(statusDeliverSql, beginData, endData)

	rows, err := mysql.Query(centerDeliverSql)
	if err != nil {
		return DeliverOrderResp{
			Error: err.Error(),
		}
	}
	defer rows.Close()
	for rows.Next() {
		var data DeliverOrderData
		rows.Scan(&data.Classification, &data.DeliverOrderNum)
		log.Printf("Classification : %s, DeliverOrderNum : %v \n", data.Classification, data.DeliverOrderNum)
		centerDeliverList = append(centerDeliverList, data)
		totalCenterNum += data.DeliverOrderNum
	}

	rowsTwo, err := mysql.Query(statusDeliverSql)
	if err != nil {
		return DeliverOrderResp{
			Error: err.Error(),
		}
	}
	defer rowsTwo.Close()
	for rowsTwo.Next() {
		var data DeliverOrderData
		rowsTwo.Scan(&data.Classification, &data.DeliverOrderNum)
		log.Printf("Classification : %s, DeliverOrderNum : %v \n", data.Classification, data.DeliverOrderNum)
		statusDeliverList = append(statusDeliverList, data)
		totalStatusNum += data.DeliverOrderNum
	}

	return DeliverOrderResp{
		CenterDeliverDataList: centerDeliverList,
		StatusDeliverDataList: statusDeliverList,
		TotalCenterDeliverNum: totalCenterNum,
		TotalStatusDeliverNum: totalStatusNum,
	}
}

func (c *DataPanelInstance) GetSalesData() SalesResp {
	var (
		beginData, endData = getDataRange()
		brandSalesSql      = BrandSalesSql
		centerSalesSql     = CenterSalesSql
		brandSalesList     []SalesData
		centerSalesList    []SalesData
		totalBrandNum      int
		totalCenterNum     int
	)
	brandSalesSql = fmt.Sprintf(brandSalesSql, beginData, endData)
	centerSalesSql = fmt.Sprintf(centerSalesSql, beginData, endData)

	rows, err := mysql.Query(brandSalesSql)
	if err != nil {
		return SalesResp{
			Error: err.Error(),
		}
	}
	defer rows.Close()
	for rows.Next() {
		var data SalesData
		rows.Scan(&data.Classification, &data.GoodsNum)
		log.Printf("Classification : %s, GoodsNum : %v \n", data.Classification, data.GoodsNum)
		brandSalesList = append(brandSalesList, data)
		totalBrandNum += data.GoodsNum
	}

	rowsTwo, err := mysql.Query(centerSalesSql)
	if err != nil {
		return SalesResp{
			Error: err.Error(),
		}
	}
	defer rowsTwo.Close()
	for rowsTwo.Next() {
		var data SalesData
		rowsTwo.Scan(&data.Classification, &data.GoodsNum)
		log.Printf("Classification : %s, GoodsNum : %v \n", data.Classification, data.GoodsNum)
		brandSalesList = append(brandSalesList, data)
		totalCenterNum += data.GoodsNum
	}

	return SalesResp{
		BrandSalesDataList:  brandSalesList,
		CenterSalesDataList: centerSalesList,
		TotalBrandSalesNum:  totalBrandNum,
		TotalCenterSalesNum: totalCenterNum,
	}
}

// GetAllDataPanel 全部数据
func (c *DataPanelInstance) GetAllDataPanel() DataPanelResp {
	var resp DataPanelResp
	result, err := redisClient.Get("AllDataPanel").Result()
	if err == nil {
		byteResult := []byte(result)
		json.Unmarshal(byteResult, &resp)
		return resp
	}
	if err != redis.Nil {
		return DataPanelResp{
			Error: err.Error(),
		}
	}
	resp = c.getAllData()
	fmt.Printf("resp : %v \n", resp)
	bytes, _ := json.Marshal(resp)
	err = redisClient.Set("AllDataPanel", string(bytes), 6*time.Hour).Err()
	if err != nil {
		fmt.Printf("error : %v \n", err)
		return DataPanelResp{
			Error: err.Error(),
		}
	}

	return resp

}
func (c *DataPanelInstance) GetDataPanel() DataPanelResp {
	return c.getAllData()
}

func (c *DataPanelInstance) getAllData() DataPanelResp {
	var resp DataPanelResp
	resp.CustomerDataResp = c.GetCustomerData()
	resp.StoreDataResp = c.GetStoreData()
	resp.EmployeeDataResp = c.GetEmployeeData()
	resp.MemDataResp = c.GetMemberData()
	resp.OptometryResp = c.GetOptometryRecord()
	resp.ExpensesResp = c.GetExpensesRecord()
	resp.DepositDataResp = c.GetDepositData()
	resp.ProductsResp = c.GetProductsData()
	resp.CardVouchersResp = c.GetCardVouchersData()
	resp.PosterResp = c.GetPosterData()
	resp.YiFansResp = c.GetYiFansData()
	resp.OrderResp = c.GetOrderData()
	resp.DeliverOrderResp = c.GetDeliverOrderData()
	resp.WorkOrderResp = c.GetWorkOrderData()
	resp.SalesResp = c.GetSalesData()

	return resp
}

type ServerInstance struct{}

func (s ServerInstance) Service() error {
	defer mysql.Close()

	e := dataPane.e
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	dataPane.bind()
	// 8999  ":8999"
	return e.Start(fmt.Sprintf(":%s", ContainerPort))
}

func main() {
	if sqlerror != nil {
		log.Printf("err %v", sqlerror)
	}
	server := ServerInstance{}
	brisk.HandleServiceLifeCycle(server)
}

// 获取并格式化时间
func getDataRange() (string, string) {
	yesterday := time.Now().AddDate(0, 0, -1).Format("2006-01-02")
	beginDate := yesterday + " 00:00:00"
	endDate := yesterday + " 23:59:59"
	return beginDate, endDate
}

// func main() {
// 	if sqlerror != nil {
// 		log.Printf("err %v", sqlerror)
// 	}
// 	log.Println("begin")
// 	defer mysql.Close()
// 	resp := dataPane.GetExpensesRecord()
// 	log.Printf("%v", resp)
// }
