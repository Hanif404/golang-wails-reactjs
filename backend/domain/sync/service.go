package sync

type Service interface {
	GetData(id int) (*Sync, error)
	GetDatas() (*[]Sync, error)
	SaveData(sync *Sync) error
}
