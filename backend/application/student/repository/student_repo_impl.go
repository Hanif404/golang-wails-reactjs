package studentrepository

import (
	"database/sql"
	"golang-wails-reactjs/backend/domain/student"
)

type repo struct {
	db *sql.DB
}

func New(db *sql.DB) student.Repository {
	return &repo{
		db,
	}
}

func (r *repo) GetStudents() (*[]student.Student, error) {
	rows, err := r.db.Query("SELECT id, nis, namaLengkap, jenisKelamin, kelas, angkatan, namaAyah, namaIbu, phoneOrtu, status FROM data_siswa")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var students []student.Student
	for rows.Next() {
		var s student.Student
		err := rows.Scan(&s.ID, &s.Nis, &s.Name, &s.Gender, &s.Kelas, &s.Angkatan, &s.NamaAyah, &s.NamaIbu, &s.Phone, &s.Status)
		if err != nil {
			return nil, err
		}
		students = append(students, s)
	}

	return &students, nil
}

func (r *repo) GetStudent(id int) (*student.Student, error) {
	var s student.Student
	err := r.db.QueryRow("SELECT * FROM data_siswa WHERE id = ?", id).Scan(&s.ID, &s.Nis, &s.Name, &s.Gender, &s.TempatLahir, &s.TanggalLahir, &s.Kelas, &s.Angkatan, &s.Alamat, &s.NamaAyah, &s.NamaIbu, &s.Phone, &s.Email, &s.Status, &s.Synced, &s.CreatedBy, &s.CreatedAt)
	if err != nil {
		return &student.Student{}, err
	}
	return &s, nil
}

func (r *repo) CreateStudent(s *student.Student) error {
	_, err := r.db.Exec("INSERT INTO data_siswa (NIS, namaLengkap, jenisKelamin, tempatLahir, tglLahir, kelas, angkatan, alamat, namaAyah, namaIbu, phoneOrtu, email, status, createdBy, createdAt, keyData) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)", s.Nis, s.Name, s.Gender, s.TempatLahir, s.TanggalLahir, s.Kelas, s.Angkatan, s.Alamat, s.NamaAyah, s.NamaIbu, s.Phone, s.Email, s.Status, s.CreatedBy, s.CreatedAt, s.Key)
	if err != nil {
		return err
	}
	return nil
}

func (r *repo) UpdateStudent(s *student.Student) error {
	_, err := r.db.Exec("UPDATE data_siswa SET NIS = ?, namaLengkap = ?, tempatLahir = ?, tglLahir = ?, kelas = ?, angkatan = ?, alamat = ?, namaAyah = ?, namaIbu = ?, phoneOrtu = ?, email = ?, status = ?, synced = ?  WHERE id = ?", s.Nis, s.Name, s.Gender, s.TempatLahir, s.TanggalLahir, s.Kelas, s.Angkatan, s.Alamat, s.NamaAyah, s.NamaIbu, s.Phone, s.Email, s.Status, 0, s.ID)
	if err != nil {
		return err
	}
	return nil
}

func (r *repo) DeleteStudent(id int) error {
	_, err := r.db.Exec("DELETE FROM data_siswa WHERE id = ?", id)
	if err != nil {
		return err
	}
	return nil
}

func (r *repo) GetStudentsBySync() (*[]student.Student, error) {
	rows, err := r.db.Query("SELECT id, nis, namaLengkap, jenisKelamin, tempatLahir, tglLahir, kelas, angkatan, alamat, namaAyah, namaIbu, phoneOrtu, email, status, keyData FROM data_siswa where synced=0")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var students []student.Student
	for rows.Next() {
		var s student.Student
		err := rows.Scan(&s.ID, &s.Nis, &s.Name, &s.Gender, &s.TempatLahir, &s.TanggalLahir, &s.Kelas, &s.Angkatan, &s.Alamat, &s.NamaAyah, &s.NamaIbu, &s.Phone, &s.Email, &s.Status, &s.Key)
		if err != nil {
			return nil, err
		}
		students = append(students, s)
	}

	return &students, nil
}

func (r *repo) UpdateStudentSynced(id int) error {
	_, err := r.db.Exec("UPDATE data_siswa SET synced = ?  WHERE id = ?", 1, id)
	if err != nil {
		return err
	}
	return nil
}
