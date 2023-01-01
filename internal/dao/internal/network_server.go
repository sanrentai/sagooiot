// ==========================================================================
// Code generated by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// NetworkServerDao is the data access object for table network_server.
type NetworkServerDao struct {
	table   string               // table is the underlying table name of the DAO.
	group   string               // group is the database configuration group name of current DAO.
	columns NetworkServerColumns // columns contains all the column names of Table for convenient usage.
}

// NetworkServerColumns defines and stores column names for table network_server.
type NetworkServerColumns struct {
	Id        string //
	Name      string //
	Types     string // tcp/udp
	Addr      string //
	Register  string // 注册包
	Heartbeat string // 心跳包
	Protocol  string // 协议
	Devices   string // 默认设备
	Status    string //
	CreatedAt string //
	UpdatedAt string //
	CreateBy  string //
	Remark    string // 备注
}

//  networkServerColumns holds the columns for table network_server.
var networkServerColumns = NetworkServerColumns{
	Id:        "id",
	Name:      "name",
	Types:     "types",
	Addr:      "addr",
	Register:  "register",
	Heartbeat: "heartbeat",
	Protocol:  "protocol",
	Devices:   "devices",
	Status:    "status",
	CreatedAt: "created_at",
	UpdatedAt: "updated_at",
	CreateBy:  "create_by",
	Remark:    "remark",
}

// NewNetworkServerDao creates and returns a new DAO object for table data access.
func NewNetworkServerDao() *NetworkServerDao {
	return &NetworkServerDao{
		group:   "default",
		table:   "network_server",
		columns: networkServerColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *NetworkServerDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *NetworkServerDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *NetworkServerDao) Columns() NetworkServerColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *NetworkServerDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *NetworkServerDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *NetworkServerDao) Transaction(ctx context.Context, f func(ctx context.Context, tx *gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}