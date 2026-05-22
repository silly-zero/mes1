package models

import (
	"fmt"
	"time"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

// SeedData 填充初始演示数据
func SeedData() {
	seedUsers()
	seedMaterials()
	seedInboundOrders()
	seedStocks()
	seedOutboundOrders()
	seedMSDRecords()
}

func seedUsers() {
	var count int64
	DB.Model(&User{}).Count(&count)
	if count <= 1 { // 如果只有默认的 user 账号，则补充一些
		users := []User{
			{Username: "admin1", Password: hashPassword("123456"), Nickname: "张三", Role: RoleAdmin},
			{Username: "admin2", Password: hashPassword("123456"), Nickname: "李四", Role: RoleAdmin},
		}
		for _, u := range users {
			DB.FirstOrCreate(&u, User{Username: u.Username})
		}
		fmt.Println("Seed users created.")
	}
}

func seedMaterials() {
	var count int64
	DB.Model(&Material{}).Count(&count)
	if count == 0 {
		materials := []Material{
			{MaterialCode: "33004070002", MaterialName: "原材料2", Spec: "规格A", Unit: "pcs", MaterialType: "普通物料", ABCClass: "C"},
			{MaterialCode: "11002204007001", MaterialName: "成品 (通用类)", Spec: "规格B", Unit: "set", MaterialType: "成品", ABCClass: "C"},
			{MaterialCode: "33004070001", MaterialName: "sku物料", Spec: "规格C", Unit: "pcs", MaterialType: "普通物料", ABCClass: "C"},
			{MaterialCode: "33004070003", MaterialName: "原材料3", Spec: "规格D", Unit: "pcs", MaterialType: "普通物料", ABCClass: "C"},
		}
		DB.Create(&materials)
		fmt.Println("Seed materials created.")
	}
}

func seedInboundOrders() {
	var count int64
	DB.Model(&InboundOrder{}).Count(&count)
	// 如果记录数较少或没有，我们重新填充以展示新字段
	if count < 3 {
		DB.Unscoped().Session(&gorm.Session{AllowGlobalUpdate: true}).Delete(&InboundOrder{})
		now := time.Now().Format("2006-01-02 15:04:05")
		orders := []InboundOrder{
			{
				InboundNo:       "MESSC202246303042401250911000001",
				IsUrgent:        true,
				ConfirmType:     "来料入库确认",
				ReceiveNo:       "W202346303032926051900001",
				ReceiptTypeCode: "lyz004",
				ReceiptTypeName: "客供收料",
				MaterialCode:    "33004070002",
				TotalQty:        200,
				QualifiedQty:    0,
				UnqualifiedQty:  0,
				IsQCConfirmed:   false,
				IsInbound:       false,
				Status:          "未完成",
				OrgCode:         "2023463030329",
				OrgName:         "萧梓彬",
				Source:          "手动创建",
				ModifierID:      "2023463030329",
				ModifierName:    "萧梓彬",
				ModifyTime:      &now,
			},
			{
				InboundNo:       "MESSC2023463030109260509000004",
				IsUrgent:        false,
				ConfirmType:     "来料入库确认",
				ReceiveNo:       "W202346303032926051900002",
				ReceiptTypeCode: "lyz005",
				ReceiptTypeName: "客退收料",
				QCNo:            "QC260512000018",
				MaterialCode:    "11002204007001",
				TotalQty:        51,
				QualifiedQty:    50,
				UnqualifiedQty:  1,
				IsQCConfirmed:   true,
				QCUser:          "古波",
				QCTime:          &now,
				IsInbound:       false,
				Status:          "未完成",
				OrgCode:         "2023463030329",
				OrgName:         "萧梓彬",
				Source:          "手动创建",
				ModifierID:      "2023463030329",
				ModifierName:    "萧梓彬",
				ModifyTime:      &now,
			},
			{
				InboundNo:       "MESSC202405200001",
				IsUrgent:        false,
				ConfirmType:     "来料入库确认",
				ReceiveNo:       "W202346303012826051200016",
				ReceiptTypeCode: "0001",
				ReceiptTypeName: "初始化收料",
				QCNo:            "QC202405200001",
				MaterialCode:    "33004070001",
				TotalQty:        100,
				QualifiedQty:    100,
				UnqualifiedQty:  0,
				IsQCConfirmed:   true,
				QCUser:          "admin1",
				QCTime:          &now,
				IsInbound:       true,
				InboundUser:     "admin1",
				InboundTime:     &now,
				Status:          "已完成",
				OrgCode:         "2023463030128",
				OrgName:         "韦之猛",
				Source:          "手动创建",
				ModifierID:      "2023463030128",
				ModifierName:    "韦之猛",
				ModifyTime:      &now,
			},
		}
		DB.Create(&orders)
		fmt.Println("Seed inbound orders created.")
	}
}

func seedStocks() {
	var count int64
	DB.Model(&Stock{}).Count(&count)
	if count == 0 {
		stocks := []Stock{
			{MaterialCode: "33004070002", Quantity: 500, Warehouse: "原料仓"},
			{MaterialCode: "11002204007001", Quantity: 120, Warehouse: "成品仓"},
			{MaterialCode: "33004070001", Quantity: 300, Warehouse: "原料仓"},
			{MaterialCode: "33004070003", Quantity: 50, Warehouse: "原料仓"},
		}
		DB.Create(&stocks)
		fmt.Println("Seed stocks created.")
	}
}

func seedOutboundOrders() {
	var count int64
	DB.Model(&OutboundOrder{}).Count(&count)
	if count == 0 {
		now := time.Now().Add(-24 * time.Hour).Format("2006-01-02 15:04:05")
		orders := []OutboundOrder{
			{
				OutboundNo:   "OUT202405210001",
				MaterialCode: "33004070002",
				Quantity:     50,
				Warehouse:    "原料仓",
				Receiver:     "王五",
				Status:       "completed",
				Operator:     "admin1",
				OutboundTime: &now,
			},
			{
				OutboundNo:   "OUT202405210002",
				MaterialCode: "11002204007001",
				Quantity:     10,
				Warehouse:    "成品仓",
				Receiver:     "赵六",
				Status:       "completed",
				Operator:     "admin2",
				OutboundTime: &now,
			},
		}
		DB.Create(&orders)
		fmt.Println("Seed outbound orders created.")
	}
}

func seedMSDRecords() {
	var count int64
	DB.Model(&MSDRecord{}).Count(&count)
	if count == 0 {
		records := []MSDRecord{
			{
				MaterialCode:   "33004070002",
				LotNo:          "LOT20240522001",
				MSLLevel:       "3",
				TotalFloorLife: 168,
				RemainingLife:  168,
				Status:         "InBag",
				Warehouse:      "电子料仓",
				Operator:       "admin1",
			},
			{
				MaterialCode:   "11002204007001",
				LotNo:          "LOT20240522002",
				MSLLevel:       "4",
				TotalFloorLife: 72,
				RemainingLife:  48.5,
				Status:         "Exposed",
				Warehouse:      "电子料仓",
				Operator:       "admin2",
			},
		}
		DB.Create(&records)
		fmt.Println("Seed MSD records created.")
	}
}

func hashPassword(pwd string) string {
	hash, _ := bcrypt.GenerateFromPassword([]byte(pwd), bcrypt.DefaultCost)
	return string(hash)
}
