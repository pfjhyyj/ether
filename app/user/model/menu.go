package model

import (
	"github.com/pfjhyyj/ether/app/user/constants"
	"github.com/pfjhyyj/ether/common"
	"gorm.io/gorm"
)

type Menu struct {
	MenuId      uint   `gorm:"primaryKey;column:menu_id"`
	MenuType    uint   `gorm:"column:menu_type"`
	ParentId    uint   `gorm:"column:parent_id"`
	Name        string `gorm:"column:name"`
	Path        string `gorm:"column:path"`
	Locale      string `gorm:"column:locale"`
	Icon        string `gorm:"column:icon"`
	Order       int    `gorm:"column:order"`
	Description string `gorm:"column:description"`
	common.Model
}

func (m Menu) TableName() string {
	return "menu"
}

type QueryMenuParams struct {
	common.PageRequest
}

func CreateMenu(tx *gorm.DB, menu *Menu) error {
	return tx.Create(menu).Error
}

func UpdateMenu(tx *gorm.DB, menuId uint, menu *Menu) error {
	return tx.Where("menu_id = ?", menuId).Updates(menu).Error
}

func DeleteMenu(tx *gorm.DB, menuId uint) error {
	return tx.Delete("menu_id = ?", menuId).Error
}

func GetMenuByMenuId(tx *gorm.DB, menuId uint) (*Menu, error) {
	var menu Menu
	if err := tx.Where("menu_id = ?", menuId).First(&menu).Error; err != nil {
		return nil, err
	}
	return &menu, nil
}

func GetMenuByMenuIds(tx *gorm.DB, menuIds []uint) ([]*Menu, error) {
	var menus []*Menu
	query := tx.Model(&Menu{})

	if err := query.Where("menu_id in ?", menuIds).Find(&menus).Error; err != nil {
		return nil, err
	}

	return menus, nil
}

func ListMenus(tx *gorm.DB, params *QueryMenuParams) ([]*Menu, int64, error) {
	var menus []*Menu
	query := tx.Model(&Menu{})

	var total int64
	query.Count(&total)

	if params.Current > 0 && params.PageSize > 0 {
		query = query.Offset((params.Current - 1) * params.PageSize).Limit(params.PageSize)
	}

	query.Where("parent_id = 0 AND menu_type = ?", constants.MenuTypeCategory)

	if err := query.Find(&menus).Error; err != nil {
		return nil, 0, err
	}

	return menus, total, nil
}

func ListMenuTreeByMenuId(tx *gorm.DB, menuId uint) ([]*Menu, error) {
	var menus []*Menu
	query := tx.Model(&Menu{})

	query.Raw(`
		WITH RECURSIVE res AS (
			SELECT m1.menu_id, m1.menu_type, m1.parent_id, m1.name, m1.path, m1.locale, m1.icon, m1.order
			FROM menu m1
			WHERE menu_id = ?
			UNION
			SELECT m2.menu_id, m2.menu_type, m2.parent_id, m2.name, m2.path, m2.locale, m2.icon, m2.order
			FROM res m
				INNER JOIN menu m2 ON m2.parent_id = m.menu_id
		)
		SELECT *
		FROM res
		`,
		menuId,
	)

	if err := query.Find(&menus).Error; err != nil {
		return nil, err
	}

	return menus, nil
}

// ListMenuTreeByMenuIds godoc
// Get all related menus from the top to the bottom
func ListMenuTreeByMenuIds(tx *gorm.DB, menuIds []uint) ([]*Menu, error) {
	var menus []*Menu
	query := tx.Model(&Menu{})

	query.Raw(`
		WITH RECURSIVE res AS (
			SELECT m1.menu_id, m1.menu_type, m1.parent_id, m1.name, m1.path, m1.locale, m1.icon, m1.order
			FROM menu m1
			WHERE menu_id IN ?
			UNION
			SELECT m2.menu_id, m2.menu_type, m2.parent_id, m2.name, m2.path, m2.locale, m2.icon, m2.order
			FROM res m
				INNER JOIN menu m2 ON m2.parent_id = m.menu_id
		)
		SELECT *
		FROM res
		`,
		menuIds,
	)

	if err := query.Find(&menus).Error; err != nil {
		return nil, err
	}

	return menus, nil
}

// ListMenuTreeFromBottomByMenuIds godoc
// Get all related menus from the bottom to the top
func ListMenuTreeFromBottomByMenuIds(tx *gorm.DB, menuIds []uint) ([]*Menu, error) {
	var menus []*Menu
	query := tx.Model(&Menu{})

	query.Raw(`
		WITH RECURSIVE res AS (
			SELECT m1.menu_id, m1.menu_type, m1.parent_id, m1.name, m1.path, m1.locale, m1.icon, m1.order
			FROM menu m1
			WHERE menu_id IN ?
			UNION
			SELECT m2.menu_id, m2.menu_type, m2.parent_id, m2.name, m2.path, m2.locale, m2.icon, m2.order
			FROM res m
				INNER JOIN menu m2 ON m2.menu_id = m.parent_id
		)
		SELECT *
		FROM res
		`,
		menuIds,
	)

	if err := query.Find(&menus).Error; err != nil {
		return nil, err
	}

	return menus, nil
}
