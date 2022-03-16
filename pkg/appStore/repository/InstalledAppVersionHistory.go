package appStoreRepository

import (
	"github.com/go-pg/pg"
	"go.uber.org/zap"
	"time"
)

type InstalledAppVersionHistoryRepository interface {
	CreateInstalledAppVersionHistory(model *InstalledAppVersionHistory, tx *pg.Tx) (*InstalledAppVersionHistory, error)
	UpdateInstalledAppVersionHistory(model *InstalledAppVersionHistory, tx *pg.Tx) (*InstalledAppVersionHistory, error)
	GetInstalledAppVersionHistory(id int) (*InstalledAppVersionHistory, error)
	GetInstalledAppVersionHistoryByVersionId(installAppVersionId int) ([]*InstalledAppVersionHistory, error)
}

type InstalledAppVersionHistoryRepositoryImpl struct {
	dbConnection *pg.DB
	Logger       *zap.SugaredLogger
}

func NewInstalledAppVersionHistoryRepositoryImpl(Logger *zap.SugaredLogger, dbConnection *pg.DB) *InstalledAppVersionHistoryRepositoryImpl {
	return &InstalledAppVersionHistoryRepositoryImpl{dbConnection: dbConnection, Logger: Logger}
}

type InstalledAppVersionHistory struct {
	TableName             struct{}  `sql:"installed_app_version_history" pg:",discard_unknown_columns"`
	Id                    int       `sql:"id,pk"`
	InstalledAppVersionId int       `sql:"installed_app_version_id,notnull"`
	ValuesYamlRaw         string    `sql:"values_yaml_raw"`
	Status                string    `sql:"status"`
	CreatedOn             time.Time `sql:"created_on"`
	CreatedBy             int32     `sql:"created_by"`
}

func (impl InstalledAppVersionHistoryRepositoryImpl) CreateInstalledAppVersionHistory(model *InstalledAppVersionHistory, tx *pg.Tx) (*InstalledAppVersionHistory, error) {
	err := tx.Insert(model)
	if err != nil {
		impl.Logger.Error(err)
		return model, err
	}
	return model, nil
}
func (impl InstalledAppVersionHistoryRepositoryImpl) UpdateInstalledAppVersionHistory(model *InstalledAppVersionHistory, tx *pg.Tx) (*InstalledAppVersionHistory, error) {
	err := tx.Update(model)
	if err != nil {
		impl.Logger.Error(err)
		return model, err
	}
	return model, nil
}
func (impl InstalledAppVersionHistoryRepositoryImpl) GetInstalledAppVersionHistory(id int) (*InstalledAppVersionHistory, error) {
	model := &InstalledAppVersionHistory{}
	err := impl.dbConnection.Model(model).
		Where("id = ?", id).Select()
	return model, err
}

func (impl InstalledAppVersionHistoryRepositoryImpl) GetInstalledAppVersionHistoryByVersionId(installAppVersionId int) ([]*InstalledAppVersionHistory, error) {
	var model []*InstalledAppVersionHistory
	err := impl.dbConnection.Model(&model).
		Where("installed_app_version_id.app_store_id = ?", installAppVersionId).
		Select()
	return model, err
}
