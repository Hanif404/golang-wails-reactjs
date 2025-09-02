package syncrepository

import (
	"database/sql"
	"golang-wails-reactjs/backend/domain/sync"
	"golang-wails-reactjs/backend/pkg/google"
)

type repo struct {
	db *sql.DB
}

func New(db *sql.DB) sync.Repository {
	return &repo{
		db,
	}
}

func (r *repo) GetSyncs() (*[]sync.Sync, error) {
	rows, err := r.db.Query("SELECT id, syncType, createdAt, createdBy, startDate, endDate, message, status FROM sync_sessions")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var syncs []sync.Sync
	for rows.Next() {
		var s sync.Sync
		err := rows.Scan(&s.ID, &s.SyncType, &s.CreatedAt, &s.CreatedBy, &s.StartDate, &s.EndDate, &s.Message, &s.Status)
		if err != nil {
			return nil, err
		}
		syncs = append(syncs, s)
	}

	return &syncs, nil
}

func (r *repo) GetSync(id int) (*sync.Sync, error) {
	var s sync.Sync
	err := r.db.QueryRow("SELECT id, syncType, createdAt, createdBy, startDate, endDate, message, status FROM sync_sessions WHERE id = ?", id).Scan(&s.ID, &s.SyncType, &s.CreatedAt, &s.CreatedBy, &s.StartDate, &s.EndDate, &s.Message, &s.Status)
	if err != nil {
		return &sync.Sync{}, err
	}
	return &s, nil
}

func (r *repo) GetSyncByStatus(status string) (*sync.Sync, error) {
	var s sync.Sync
	err := r.db.QueryRow("SELECT id, syncType, createdAt, createdBy, startDate, endDate, message, status FROM sync_sessions WHERE status = ?", status).Scan(&s.ID, &s.SyncType, &s.CreatedAt, &s.CreatedBy, &s.StartDate, &s.EndDate, &s.Message, &s.Status)
	if err != nil {
		return &sync.Sync{}, err
	}
	return &s, nil
}

func (r *repo) CreateSync(s *sync.Sync) (int64, error) {
	result, err := r.db.Exec("INSERT INTO sync_sessions (syncType, createdAt, createdBy, startDate, endDate, message, status) VALUES (?, ?, ?, ?, ?, ?, ?)", s.SyncType, s.CreatedAt, s.CreatedBy, s.StartDate, s.EndDate, s.Message, s.Status)
	if err != nil {
		return 0, err
	}

	lastID, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}
	return lastID, nil
}

func (r *repo) UpdateSync(s *sync.Sync) error {
	_, err := r.db.Exec("UPDATE sync_sessions SET endDate = ?, message = ?, status = ? WHERE id = ?", s.EndDate, s.Message, s.Status, s.ID)
	if err != nil {
		return err
	}
	return nil
}

func (r *repo) PostGoogleSheet(sheet string, payload any) error {
	_, err := google.PostResponse("AKfycbxDDhQMTjfKQ5ZFBQYy8-mI9OnumVJqCgXE1AhrCy-pVHNS0a4TYcIfiAMVR0ibXdBfAg", "?sheet="+sheet, payload)
	if err != nil {
		return err
	}
	return nil
}
