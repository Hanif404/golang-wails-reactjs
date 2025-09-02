package sync

type Repository interface {
	GetSync(id int) (*Sync, error)
	GetSyncByStatus(status string) (*Sync, error)
	GetSyncs() (*[]Sync, error)
	CreateSync(sync *Sync) (int64, error)
	UpdateSync(sync *Sync) error
	PostGoogleSheet(sheet string, payload any) error
}
