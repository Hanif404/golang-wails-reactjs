package student

type Student struct {
	ID           int    `json:"id"`
	Nis          string `json:"nis" validate:"required"`
	Name         string `json:"name" validate:"required"`
	Gender       string `json:"gender" validate:"required"`
	Phone        string `json:"phone" validate:"required"`
	Email        string `json:"email" validate:"required,email"`
	Kelas        string `json:"kelas" validate:"required"`
	Angkatan     string `json:"angkatan" validate:"required"`
	Alamat       string `json:"alamat" validate:"required"`
	TempatLahir  string `json:"tempat_lahir" validate:"required"`
	TanggalLahir string `json:"tanggal_lahir" validate:"required"`
	NamaAyah     string `json:"nama_ayah" validate:"required"`
	NamaIbu      string `json:"nama_ibu" validate:"required"`
	Status       string `json:"status" validate:"required"`
	Synced       bool   `json:"synced"`
	CreatedBy    string `json:"created_by"`
	CreatedAt    int64  `json:"created_at"`
	Key          string `json:"key"`
}
