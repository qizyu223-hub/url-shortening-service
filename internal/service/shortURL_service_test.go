package service

import (
	"github.com/stretchr/testify/require"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"testing"
	"url-shortening-service/internal/model"
	"url-shortening-service/internal/repository"
)

//func startPostgres(t *testing.T) *gorm.DB {
//	ctx := context.Background()
//	req := testcontainers.ContainerRequest{
//		Image:        "postgres:15-alpine",
//		Env:          map[string]string{"POSTGRES_PASSWORD": "pwd", "POSTGRES_DB": "testdb"},
//		ExposedPorts: []string{"5432/tcp"},
//		WaitingFor:   wait.ForListeningPort("5432/tcp"),
//	}
//	pgC, err := testcontainers.GenericContainer(ctx, testcontainers.GenericContainerRequest{
//		ContainerRequest: req,
//		Started:          true,
//	})
//	require.NoError(t, err)
//
//	endpoint, err := pgC.Endpoint(ctx, "")
//	require.NoError(t, err)
//
//	dsn := "postgres://postgres:pwd@" + endpoint + "/testdb?sslmode=disable"
//	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
//	require.NoError(t, err)
//
//	require.NoError(t, db.AutoMigrate(&model.ShortURL{}))
//
//	t.Cleanup(func() {
//		_ = pgC.Terminate(ctx)
//	})
//	return db
//}

func startSqlite(t *testing.T) *gorm.DB {
	sqliteDB, err := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	require.NoError(t, err)
	err = sqliteDB.AutoMigrate(&model.ShortURL{})
	require.NoError(t, err)
	return sqliteDB
}

func TestCreateAndGet(t *testing.T) {
	//db := startPostgres(t)
	db := startSqlite(t)
	repo := repository.NewURLShorteningRepository(db)
	svc := NewURLShorteningService(repo)

	tests := []struct {
		name    string
		url     string
		want    int
		wantErr bool
	}{
		{"ok", "http://example.com", 1, false},
		{"ok2", "http://example/2", 2, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := svc.Create(tt.url)
			if tt.wantErr {
				require.Error(t, err)
				return
			}
			require.NoError(t, err)
			require.Equal(t, tt.want, int(got.ID))
			require.NotEmpty(t, got.ShortCode)

			fetched, err := svc.GetByShortCode(got.ShortCode)
			require.NoError(t, err)
			require.Equal(t, tt.url, fetched.URL)
		})
	}
}
