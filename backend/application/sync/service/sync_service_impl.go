package syncservice

import (
	"fmt"
	"golang-wails-reactjs/backend/domain/student"
	"golang-wails-reactjs/backend/domain/sync"
	"time"
)

type service struct {
	repo        sync.Repository
	studentRepo student.Repository
}

func New(r sync.Repository, s student.Repository) sync.Service {
	return &service{
		repo:        r,
		studentRepo: s,
	}
}

func (s *service) GetDatas() (*[]sync.Sync, error) {
	return s.repo.GetSyncs()
}

func (s *service) GetData(id int) (*sync.Sync, error) {
	return s.repo.GetSync(id)
}

func (s *service) SaveData(syncData *sync.Sync) error {
	result, _ := s.repo.GetSyncByStatus("running")
	if result.ID != 0 {
		return fmt.Errorf("sync already running")
	}

	syncData.CreatedAt = time.Now().UnixMilli()
	syncData.StartDate = time.Now().UnixMilli()
	syncData.SyncType = "upload"
	syncData.Status = "running"

	lastId, err := s.repo.CreateSync(syncData)
	if err != nil {
		return err
	}

	//sync upload data
	go s.SyncData(lastId)
	return nil
}

func (s *service) SyncData(syncId int64) {
	message := "success"
	result, err := s.studentRepo.GetStudentsBySync()
	if err != nil {
		message = err.Error()
	}

	//sync data
	for _, student := range *result {
		//send to sheet
		ttl := fmt.Sprintf("%s, %s", student.TempatLahir, student.TanggalLahir)
		err = s.repo.PostGoogleSheet("data_siswa", sync.SycnStudent{
			Nis:      student.Nis,
			Name:     student.Name,
			Gender:   student.Gender,
			Ttl:      ttl,
			Kelas:    student.Kelas,
			Angkatan: student.Angkatan,
			Alamat:   student.Alamat,
			NamaAyah: student.NamaAyah,
			NamaIbu:  student.NamaIbu,
			Phone:    student.Phone,
			Email:    student.Email,
			Status:   student.Status,
			Key:      student.Key,
		})
		if err != nil {
			message = err.Error()
			break
		}

		//update synced
		err = s.studentRepo.UpdateStudentSynced(student.ID)
		if err != nil {
			message = err.Error()
			break
		}
	}

	s.repo.UpdateSync(&sync.Sync{
		Status:  "finished",
		Message: message,
		EndDate: time.Now().UnixMilli(),
		ID:      int(syncId),
	})
}
