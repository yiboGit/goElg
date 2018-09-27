package main

import (
	"github.com/labstack/echo"
)

func (dataPane *DataPanelInstance) bind() {
	dataPane.e.Add("POST", "/DataPanel/GetAllDataPanel", func(c echo.Context) error {

		result := dataPane.GetAllDataPanel()

		if result.Error != "" {
			return c.JSON(500, result)
		}
		return c.JSON(200, result)
	})
	dataPane.e.Add("POST", "/DataPanel/GetCardVouchersData", func(c echo.Context) error {

		result := dataPane.GetCardVouchersData()

		if result.Error != "" {
			return c.JSON(500, result)
		}
		return c.JSON(200, result)
	})
	dataPane.e.Add("POST", "/DataPanel/GetCustomerData", func(c echo.Context) error {

		result := dataPane.GetCustomerData()

		if result.Error != "" {
			return c.JSON(500, result)
		}
		return c.JSON(200, result)
	})
	dataPane.e.Add("POST", "/DataPanel/GetDataPanel", func(c echo.Context) error {

		result := dataPane.GetDataPanel()

		if result.Error != "" {
			return c.JSON(500, result)
		}
		return c.JSON(200, result)
	})
	dataPane.e.Add("POST", "/DataPanel/GetDeliverOrderData", func(c echo.Context) error {

		result := dataPane.GetDeliverOrderData()

		if result.Error != "" {
			return c.JSON(500, result)
		}
		return c.JSON(200, result)
	})
	dataPane.e.Add("POST", "/DataPanel/GetDepositData", func(c echo.Context) error {

		result := dataPane.GetDepositData()

		if result.Error != "" {
			return c.JSON(500, result)
		}
		return c.JSON(200, result)
	})
	dataPane.e.Add("POST", "/DataPanel/GetEmployeeData", func(c echo.Context) error {

		result := dataPane.GetEmployeeData()

		if result.Error != "" {
			return c.JSON(500, result)
		}
		return c.JSON(200, result)
	})
	dataPane.e.Add("POST", "/DataPanel/GetExpensesRecord", func(c echo.Context) error {

		result := dataPane.GetExpensesRecord()

		if result.Error != "" {
			return c.JSON(500, result)
		}
		return c.JSON(200, result)
	})
	dataPane.e.Add("POST", "/DataPanel/GetMemberData", func(c echo.Context) error {

		result := dataPane.GetMemberData()

		if result.Error != "" {
			return c.JSON(500, result)
		}
		return c.JSON(200, result)
	})
	dataPane.e.Add("POST", "/DataPanel/GetOptometryRecord", func(c echo.Context) error {

		result := dataPane.GetOptometryRecord()

		if result.Error != "" {
			return c.JSON(500, result)
		}
		return c.JSON(200, result)
	})
	dataPane.e.Add("POST", "/DataPanel/GetOrderData", func(c echo.Context) error {

		result := dataPane.GetOrderData()

		if result.Error != "" {
			return c.JSON(500, result)
		}
		return c.JSON(200, result)
	})
	dataPane.e.Add("POST", "/DataPanel/GetPosterData", func(c echo.Context) error {

		result := dataPane.GetPosterData()

		if result.Error != "" {
			return c.JSON(500, result)
		}
		return c.JSON(200, result)
	})
	dataPane.e.Add("POST", "/DataPanel/GetProductsData", func(c echo.Context) error {

		result := dataPane.GetProductsData()

		if result.Error != "" {
			return c.JSON(500, result)
		}
		return c.JSON(200, result)
	})
	dataPane.e.Add("POST", "/DataPanel/GetSalesData", func(c echo.Context) error {

		result := dataPane.GetSalesData()

		if result.Error != "" {
			return c.JSON(500, result)
		}
		return c.JSON(200, result)
	})
	dataPane.e.Add("POST", "/DataPanel/GetStoreData", func(c echo.Context) error {

		result := dataPane.GetStoreData()

		if result.Error != "" {
			return c.JSON(500, result)
		}
		return c.JSON(200, result)
	})
	dataPane.e.Add("POST", "/DataPanel/GetWorkOrderData", func(c echo.Context) error {

		result := dataPane.GetWorkOrderData()

		if result.Error != "" {
			return c.JSON(500, result)
		}
		return c.JSON(200, result)
	})
	dataPane.e.Add("POST", "/DataPanel/GetYiFansData", func(c echo.Context) error {

		result := dataPane.GetYiFansData()

		if result.Error != "" {
			return c.JSON(500, result)
		}
		return c.JSON(200, result)
	})

}
