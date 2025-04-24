package rt_profile

import (
	"errors"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/abiyyu03/siruta/config"
	"github.com/abiyyu03/siruta/entity/model"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func setupTestDB(t *testing.T) (*gorm.DB, sqlmock.Sqlmock, func()) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)

	gormDB, err := gorm.Open(postgres.New(postgres.Config{
		Conn: db,
	}), &gorm.Config{})
	assert.NoError(t, err)

	return gormDB, mock, func() {
		db.Close()
	}
}

func TestRegister_Success(t *testing.T) {
	db, mock, close := setupTestDB(t)
	defer close()

	repo := &RTProfileRegisterRepository{}

	// Simulasi referal code yang valid
	referalCode := "ABC123"
	profileId := "profile-001"
	rtProfile := &model.RTProfile{}

	// Mock referal code lookup
	mock.ExpectBegin()
	mock.ExpectQuery(`SELECT .* FROM "referal_codes"`).
		WithArgs(referalCode, sqlmock.AnyArg()).
		WillReturnRows(sqlmock.NewRows([]string{"profile_id"}).AddRow(profileId))

	// Mock create RTProfile
	mock.ExpectExec(`INSERT INTO "rt_profiles"`).
		WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectCommit()

	// Override config.DB dengan test DB
	oldDB := config.DB
	config.DB = db
	defer func() { config.DB = oldDB }()

	res, err := repo.Register(rtProfile, referalCode)
	assert.NoError(t, err)
	assert.Equal(t, profileId, res.RWProfileId)
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestRegister_InvalidReferalCode(t *testing.T) {
	db, mock, close := setupTestDB(t)
	defer close()

	repo := &RTProfileRegisterRepository{}

	referalCode := "INVALID"
	rtProfile := &model.RTProfile{}

	mock.ExpectBegin()
	mock.ExpectQuery(`SELECT .* FROM "referal_codes"`).
		WithArgs(referalCode, sqlmock.AnyArg()).
		WillReturnError(errors.New("referal not found"))
	mock.ExpectRollback()

	oldDB := config.DB
	config.DB = db
	defer func() { config.DB = oldDB }()

	res, err := repo.Register(rtProfile, referalCode)
	assert.Nil(t, res)
	assert.Error(t, err)
	assert.Equal(t, "referal not found", err.Error())
	assert.NoError(t, mock.ExpectationsWereMet())
}
