package prettyKit

import (
	"github.com/jedib0t/go-pretty/v6/table"
	//"github.com/scylladb/termtables"
)

// NewTableWriter 格式化输出为 table 格式.
/*
PS: 最好不要出现中文，显示会有点错乱.
*/
var NewTableWriter func() table.Writer = table.NewWriter

//// CreateTable
///*
//Deprecated: Use NewTableWriter instead（scylladb/termtables长时间未更新了）.
//
//(1) AddHeader()添加头部信息；
//(2) AddRow()逐行添加数据；
//(3) Render()返回渲染后的表格字符串.
//*/
//var CreateTable func() *termtables.Table = termtables.CreateTable
