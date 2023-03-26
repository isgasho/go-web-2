package model

// CasbinRuleTable 模型
type CasbinRuleTable struct {
	Id    uint   `gorm:"primaryKey;autoIncrement"`
	PType string `gorm:"size:100;uniqueIndex:uk_index;column:ptype;comment:策略类型"` // 多个字段联合唯一
	V0    string `gorm:"size:100;uniqueIndex:uk_index;comment:角色关键字"`
	V1    string `gorm:"size:100;uniqueIndex:uk_index;comment:资源名称"`
	V2    string `gorm:"size:100;uniqueIndex:uk_index;comment:请求类型"`
	V3    string `gorm:"size:100;uniqueIndex:uk_index"`
	V4    string `gorm:"size:100;uniqueIndex:uk_index"`
	V5    string `gorm:"size:100;uniqueIndex:uk_index"`
}

// TableName 定义 CasbinRuleTable 表名
func (c *CasbinRuleTable) TableName() string {
	return "casbin_rule"
}

// CasbinRule 这个模型用于 CasbinRuleTable 的字段对应，后续操作这个 model 就行了，字段更有意义
type CasbinRule struct {
	PType   string `json:"ptype" gorm:"column:ptype" description:"策略类型"`
	Keyword string `json:"keyword" gorm:"column:v0" description:"角色关键字"`
	Path    string `json:"path" gorm:"column:v1" description:"API路径"`
	Method  string `json:"method" gorm:"column:v2" description:"访问方法"`
}

// TableName 让它表名称和 CasbinRuleTable 一致，便于操作
func (c *CasbinRule) TableName() string {
	return "casbin_rule"
}
