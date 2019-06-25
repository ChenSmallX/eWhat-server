package dbcontroller

import (
	//"errors"
	"fmt"
)

type tableName int

const (
	users              tableName = 0 // 用户表
	avatar             tableName = 1 // 用户头像表
	foodWarehouse      tableName = 2 // 菜品库表
	foodWarehousePhoto tableName = 3 // 菜品库图片表
	food               tableName = 4 // 菜品表
	foodPhoto          tableName = 5 // 菜品图片表
	diary              tableName = 6 // 美食日记表
)

func (table tableName) String() (string, error) {
	switch table {
	case users:
		return "user", nil
	case avatar:
		return "avator", nil
	case foodWarehouse:
		return "foodWarehouse", nil
	case foodWarehousePhoto:
		return "foodWarehousePhoto", nil
	case food:
		return "food", nil
	case foodPhoto:
		return "foodPhoto", nil
	case diary:
		return "diary", nil
	default:
		var err error
		err = fmt.Errorf("No this table")
		return "", err
	}
}
