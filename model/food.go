package model
//待开发完成,先瞎写一些

/*
 * Food结构体字段(基础表)
 */
type Food struct {
	Id    							int64   `xorm:"pk autoincr" json:"id"`       //主键
	FoodName  						string  `xorm:"varchar(12)" json:"name"`     //food名称
	DelFlag                       	int        `json:"dele"`   //是否已经被删除 1表示已经删除 0表示未删除
}
