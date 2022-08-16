package tool

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)

var DB *gorm.DB

func init() {
	db, err := gorm.Open(mysql.New(mysql.Config{
		DriverName:                    C.SQL.Driver,
		ServerVersion:                 "",
		DSN:                           C.SQL.DSN,
		Conn:                          nil,
		SkipInitializeWithVersion:     false,
		DefaultStringSize:             255,
		DefaultDatetimePrecision:      nil,
		DisableDatetimePrecision:      false,
		DontSupportRenameIndex:        false,
		DontSupportRenameColumn:       false,
		DontSupportForShareClause:     false,
		DontSupportNullAsDefaultValue: false,
	}), &gorm.Config{
		SkipDefaultTransaction:                   false,
		NamingStrategy:                           nil,
		FullSaveAssociations:                     false,
		Logger:                                   nil,
		NowFunc:                                  nil,
		DryRun:                                   false,
		PrepareStmt:                              false,
		DisableAutomaticPing:                     false,
		DisableForeignKeyConstraintWhenMigrating: false,
		DisableNestedTransaction:                 false,
		AllowGlobalUpdate:                        false,
		QueryFields:                              false,
		CreateBatchSize:                          0,
		ClauseBuilders:                           nil,
		ConnPool:                                 nil,
		Dialector:                                nil,
		Plugins:                                  nil,
	})
	if err != nil {
		log.Fatalln(err)
	}
	DB = db
}
