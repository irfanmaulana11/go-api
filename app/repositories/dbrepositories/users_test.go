package dbrepositories

import (
	"context"
	"errors"
	"regexp"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/jinzhu/gorm"
)

func TestDBRepository_GetUserByEmail(t *testing.T) {
	dbMock, mock, _ := sqlmock.New()
	db, _ := gorm.Open("postgres", dbMock)

	timeMock := time.Date(2023, 10, 06, 0, 0, 0, 0, time.UTC)

	type args struct {
		ctx   context.Context
		email string
	}

	columns := []string{
		"id", "name", "email", "password", "is_active", "created_at", "updated_at",
	}

	tests := []struct {
		name    string
		args    args
		wantErr bool
		mock    func(args)
	}{
		{
			name: "Negative | Record not found",
			args: args{
				ctx:   context.TODO(),
				email: "ipann1297@gmail.com",
			},
			wantErr: true,
			mock: func(a args) {
				mock.ExpectBegin()
				mock.ExpectQuery(regexp.QuoteMeta(`SELECT * FROM "users" WHERE (email = $1)`)).WillReturnError(errors.New("record not found"))
				mock.ExpectRollback()
			},
		},
		{
			name: "Positive | Record exist",
			args: args{
				ctx:   context.TODO(),
				email: "ipann1297@gmail.com",
			},
			wantErr: false,
			mock: func(a args) {
				mock.ExpectBegin()
				rows := mock.NewRows(columns).AddRow(1, "irfan", "ipann1297@gmail.com", "123123123123", true, timeMock, timeMock)
				mock.ExpectQuery(regexp.QuoteMeta(`SELECT * FROM "users" WHERE (email = $1)`)).WillReturnRows(rows)
				mock.ExpectCommit()
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &DBRepository{
				db: db,
			}
			tt.mock(tt.args)
			_, err := r.GetUserByEmail(tt.args.ctx, tt.args.email)
			if (err != nil) != tt.wantErr {
				t.Errorf("DBRepository.GetUserByEmail() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}
