package sync

type Sync struct {
	ID        int    `json:"id"`
	SyncType  string `json:"sync_type"`
	StartDate int64  `json:"start_date"`
	EndDate   int64  `json:"end_date"`
	Message   string `json:"message"`
	Status    string `json:"status"`
	CreatedBy string `json:"created_by"`
	CreatedAt int64  `json:"created_at"`
}

type SycnStudent struct {
	Nis      string `json:"nis"`
	Name     string `json:"name"`
	Gender   string `json:"gender"`
	Ttl      string `json:"ttl"`
	Kelas    string `json:"kelas"`
	Angkatan string `json:"angkatan"`
	Alamat   string `json:"alamat"`
	NamaAyah string `json:"nama_ayah"`
	NamaIbu  string `json:"nama_ibu"`
	Phone    string `json:"phone"`
	Email    string `json:"email"`
	Status   string `json:"status"`
	Key      string `json:"key"`
}
