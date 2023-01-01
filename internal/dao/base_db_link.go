// =================================================================================
// This is auto-generated by GoFrame CLI tool only once. Fill this file as you wish.
// =================================================================================

package dao

import (
	"github.com/sagoo-cloud/sagooiot/internal/dao/internal"
)

// internalBaseDbLinkDao is internal type for wrapping internal DAO implements.
type internalBaseDbLinkDao = *internal.BaseDbLinkDao

// baseDbLinkDao is the data access object for table base_db_link.
// You can define custom methods on it to extend its functionality as you wish.
type baseDbLinkDao struct {
	internalBaseDbLinkDao
}

var (
	// BaseDbLink is globally public accessible object for table base_db_link operations.
	BaseDbLink = baseDbLinkDao{
		internal.NewBaseDbLinkDao(),
	}
)

// Fill with you ideas below.