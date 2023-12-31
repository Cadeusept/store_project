package repository

import (
	"store-project/internal/models"
	"testing"
	"time"

	sqlmock "github.com/zhashkevych/go-sqlxmock"
)

func TestDelivery_ChangeStatusFromNew(t *testing.T) {
	db, mock, err := sqlmock.Newx()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	repo := NewTransactionsPostgres(db)

	tests := []struct {
		name        string
		repo        *TransactionsPostgres
		transaction models.Transaction
		id          int64
		status      string
		mock        func()
		want        int64
		wantErr     bool
	}{
		{
			name: "StatusError",
			repo: repo,
			transaction: models.Transaction{
				UserId:    1,
				UserEmail: "fakeemail@gmail.com",
				Amount:    100,
				Currency:  "rubles",
			},
			id:     1,
			status: statusError,
			mock: func() {
				mock.ExpectExec("UPDATE transactions tt SET").WillReturnResult(sqlmock.NewResult(1, 1))
			},
			want:    1,
			wantErr: false,
		},
		{
			name: "StatusSuccess",
			repo: repo,
			transaction: models.Transaction{
				Id:        1,
				UserId:    1,
				UserEmail: "fakeemail@gmail.com",
				Amount:    100,
				Currency:  "rubles",
				Created:   time.Now(),
				Changed:   time.Now(),
				Status:    "NEW",
			},
			id:     1,
			status: statusSuccess,
			mock: func() {
				mock.ExpectExec("UPDATE transactions tt SET").WillReturnResult(sqlmock.NewResult(1, 1))
			},
			want:    1,
			wantErr: false,
		},
		{
			name: "StatusFailed",
			repo: repo,
			transaction: models.Transaction{
				Id:        1,
				UserId:    1,
				UserEmail: "fakeemail@gmail.com",
				Amount:    100,
				Currency:  "rubles",
				Created:   time.Now(),
				Changed:   time.Now(),
				Status:    "NEW",
			},
			id:     1,
			status: statusFailed,
			mock: func() {
				mock.ExpectExec("UPDATE transactions tt SET").WillReturnResult(sqlmock.NewResult(1, 1))
			},
			want:    1,
			wantErr: false,
		},
		{
			name: "StatusCancelled",
			repo: repo,
			transaction: models.Transaction{
				Id:        1,
				UserId:    1,
				UserEmail: "fakeemail@gmail.com",
				Amount:    100,
				Currency:  "rubles",
				Created:   time.Now(),
				Changed:   time.Now(),
				Status:    "NEW",
			},
			id:     1,
			status: statusCancelled,
			mock: func() {
				mock.ExpectExec("UPDATE transactions tt SET").WillReturnResult(sqlmock.NewResult(1, 0))
			},
			wantErr: true,
		},
		{
			name: "WrongStatus",
			repo: repo,
			transaction: models.Transaction{
				Id:        1,
				UserId:    1,
				UserEmail: "fakeemail@gmail.com",
				Amount:    100,
				Currency:  "rubles",
				Created:   time.Now(),
				Changed:   time.Now(),
				Status:    "NEW",
			},
			id:     1,
			status: "abracadabra",
			mock: func() {
				mock.ExpectExec("UPDATE transactions tt SET").WillReturnResult(sqlmock.NewResult(1, 0))
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mock()
			tt.repo.Create(tt.transaction)
			err := tt.repo.ChangeStatus(tt.id, tt.status)
			if (err != nil) != tt.wantErr {
				t.Errorf("Got error new = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}

func TestDelivery_ChangeStatusFromError(t *testing.T) {
	db, mock, err := sqlmock.Newx()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	repo := NewTransactionsPostgres(db)

	tests := []struct {
		name        string
		repo        *TransactionsPostgres
		transaction models.Transaction
		id          int64
		status      string
		mock        func()
		want        int64
		wantErr     bool
	}{
		{
			name: "StatusError",
			repo: repo,
			transaction: models.Transaction{
				UserId:    1,
				UserEmail: "fakeemail@gmail.com",
				Amount:    100,
				Currency:  "rubles",
			},
			id:     1,
			status: statusError,
			mock: func() {
				mock.ExpectExec("UPDATE transactions tt SET").WillReturnResult(sqlmock.NewResult(1, 0))
			},
			wantErr: true,
		},
		{
			name: "StatusSuccess",
			repo: repo,
			transaction: models.Transaction{
				Id:        1,
				UserId:    1,
				UserEmail: "fakeemail@gmail.com",
				Amount:    100,
				Currency:  "rubles",
				Created:   time.Now(),
				Changed:   time.Now(),
				Status:    "NEW",
			},
			id:     1,
			status: statusSuccess,
			mock: func() {
				mock.ExpectExec("UPDATE transactions tt SET").WillReturnResult(sqlmock.NewResult(1, 0))
			},
			wantErr: true,
		},
		{
			name: "StatusFailed",
			repo: repo,
			transaction: models.Transaction{
				Id:        1,
				UserId:    1,
				UserEmail: "fakeemail@gmail.com",
				Amount:    100,
				Currency:  "rubles",
				Created:   time.Now(),
				Changed:   time.Now(),
				Status:    "NEW",
			},
			id:     1,
			status: statusFailed,
			mock: func() {
				mock.ExpectExec("UPDATE transactions tt SET").WillReturnResult(sqlmock.NewResult(1, 1))
			},
			wantErr: false,
		},
		{
			name: "StatusCancelled",
			repo: repo,
			transaction: models.Transaction{
				Id:        1,
				UserId:    1,
				UserEmail: "fakeemail@gmail.com",
				Amount:    100,
				Currency:  "rubles",
				Created:   time.Now(),
				Changed:   time.Now(),
				Status:    "NEW",
			},
			id:     1,
			status: statusCancelled,
			mock: func() {
				mock.ExpectExec("UPDATE transactions tt SET").WillReturnResult(sqlmock.NewResult(1, 0))
			},
			wantErr: true,
		},
		{
			name: "WrongStatus",
			repo: repo,
			transaction: models.Transaction{
				Id:        1,
				UserId:    1,
				UserEmail: "fakeemail@gmail.com",
				Amount:    100,
				Currency:  "rubles",
				Created:   time.Now(),
				Changed:   time.Now(),
				Status:    "NEW",
			},
			id:     1,
			status: "abracadabra",
			mock: func() {
				mock.ExpectExec("UPDATE transactions tt SET").WillReturnResult(sqlmock.NewResult(1, 0))
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.repo.Create(tt.transaction)
			tt.repo.ChangeStatus(tt.id, statusError)
			tt.mock()
			err := tt.repo.ChangeStatus(tt.id, tt.status)
			if (err != nil) != tt.wantErr {
				t.Errorf("Got error new = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}

func TestDelivery_ChangeStatusFromSuccess(t *testing.T) {
	db, mock, err := sqlmock.Newx()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	repo := NewTransactionsPostgres(db)

	tests := []struct {
		name        string
		repo        *TransactionsPostgres
		transaction models.Transaction
		id          int64
		status      string
		mock        func()
		want        int64
		wantErr     bool
	}{
		{
			name: "StatusError",
			repo: repo,
			transaction: models.Transaction{
				UserId:    1,
				UserEmail: "fakeemail@gmail.com",
				Amount:    100,
				Currency:  "rubles",
			},
			id:     1,
			status: statusError,
			mock: func() {
				mock.ExpectExec("UPDATE transactions tt SET").WillReturnResult(sqlmock.NewResult(1, 0))
			},
			wantErr: true,
		},
		{
			name: "StatusSuccess",
			repo: repo,
			transaction: models.Transaction{
				Id:        1,
				UserId:    1,
				UserEmail: "fakeemail@gmail.com",
				Amount:    100,
				Currency:  "rubles",
				Created:   time.Now(),
				Changed:   time.Now(),
				Status:    "NEW",
			},
			id:     1,
			status: statusSuccess,
			mock: func() {
				mock.ExpectExec("UPDATE transactions tt SET").WillReturnResult(sqlmock.NewResult(1, 0))
			},
			wantErr: true,
		},
		{
			name: "StatusFailed",
			repo: repo,
			transaction: models.Transaction{
				Id:        1,
				UserId:    1,
				UserEmail: "fakeemail@gmail.com",
				Amount:    100,
				Currency:  "rubles",
				Created:   time.Now(),
				Changed:   time.Now(),
				Status:    "NEW",
			},
			id:     1,
			status: statusFailed,
			mock: func() {
				mock.ExpectExec("UPDATE transactions tt SET").WillReturnResult(sqlmock.NewResult(1, 0))
			},
			wantErr: true,
		},
		{
			name: "StatusCancelled",
			repo: repo,
			transaction: models.Transaction{
				Id:        1,
				UserId:    1,
				UserEmail: "fakeemail@gmail.com",
				Amount:    100,
				Currency:  "rubles",
				Created:   time.Now(),
				Changed:   time.Now(),
				Status:    "NEW",
			},
			id:     1,
			status: statusCancelled,
			mock: func() {
				mock.ExpectExec("UPDATE transactions tt SET").WillReturnResult(sqlmock.NewResult(1, 0))
			},
			wantErr: true,
		},
		{
			name: "WrongStatus",
			repo: repo,
			transaction: models.Transaction{
				Id:        1,
				UserId:    1,
				UserEmail: "fakeemail@gmail.com",
				Amount:    100,
				Currency:  "rubles",
				Created:   time.Now(),
				Changed:   time.Now(),
				Status:    "NEW",
			},
			id:     1,
			status: "abracadabra",
			mock: func() {
				mock.ExpectExec("UPDATE transactions tt SET").WillReturnResult(sqlmock.NewResult(1, 0))
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.repo.Create(tt.transaction)
			tt.repo.ChangeStatus(tt.id, statusSuccess)
			tt.mock()
			err := tt.repo.ChangeStatus(tt.id, tt.status)
			if (err != nil) != tt.wantErr {
				t.Errorf("Got error new = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}

func TestDelivery_ChangeStatusFromFailed(t *testing.T) {
	db, mock, err := sqlmock.Newx()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	repo := NewTransactionsPostgres(db)

	tests := []struct {
		name        string
		repo        *TransactionsPostgres
		transaction models.Transaction
		id          int64
		status      string
		mock        func()
		want        int64
		wantErr     bool
	}{
		{
			name: "StatusError",
			repo: repo,
			transaction: models.Transaction{
				UserId:    1,
				UserEmail: "fakeemail@gmail.com",
				Amount:    100,
				Currency:  "rubles",
			},
			id:     1,
			status: statusError,
			mock: func() {
				mock.ExpectExec("UPDATE transactions tt SET").WillReturnResult(sqlmock.NewResult(1, 0))
			},
			wantErr: true,
		},
		{
			name: "StatusSuccess",
			repo: repo,
			transaction: models.Transaction{
				Id:        1,
				UserId:    1,
				UserEmail: "fakeemail@gmail.com",
				Amount:    100,
				Currency:  "rubles",
				Created:   time.Now(),
				Changed:   time.Now(),
				Status:    "NEW",
			},
			id:     1,
			status: statusSuccess,
			mock: func() {
				mock.ExpectExec("UPDATE transactions tt SET").WillReturnResult(sqlmock.NewResult(1, 0))
			},
			wantErr: true,
		},
		{
			name: "StatusFailed",
			repo: repo,
			transaction: models.Transaction{
				Id:        1,
				UserId:    1,
				UserEmail: "fakeemail@gmail.com",
				Amount:    100,
				Currency:  "rubles",
				Created:   time.Now(),
				Changed:   time.Now(),
				Status:    "NEW",
			},
			id:     1,
			status: statusFailed,
			mock: func() {
				mock.ExpectExec("UPDATE transactions tt SET").WillReturnResult(sqlmock.NewResult(1, 0))
			},
			wantErr: true,
		},
		{
			name: "StatusCancelled",
			repo: repo,
			transaction: models.Transaction{
				Id:        1,
				UserId:    1,
				UserEmail: "fakeemail@gmail.com",
				Amount:    100,
				Currency:  "rubles",
				Created:   time.Now(),
				Changed:   time.Now(),
				Status:    "NEW",
			},
			id:     1,
			status: statusCancelled,
			mock: func() {
				mock.ExpectExec("UPDATE transactions tt SET").WillReturnResult(sqlmock.NewResult(1, 0))
			},
			wantErr: true,
		},
		{
			name: "WrongStatus",
			repo: repo,
			transaction: models.Transaction{
				Id:        1,
				UserId:    1,
				UserEmail: "fakeemail@gmail.com",
				Amount:    100,
				Currency:  "rubles",
				Created:   time.Now(),
				Changed:   time.Now(),
				Status:    "NEW",
			},
			id:     1,
			status: "abracadabra",
			mock: func() {
				mock.ExpectExec("UPDATE transactions tt SET").WillReturnResult(sqlmock.NewResult(1, 0))
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.repo.Create(tt.transaction)
			tt.repo.ChangeStatus(tt.id, statusFailed)
			tt.mock()
			err := tt.repo.ChangeStatus(tt.id, tt.status)
			if (err != nil) != tt.wantErr {
				t.Errorf("Got error new = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}

func TestDelivery_ChangeStatusFromCancelled(t *testing.T) {
	db, mock, err := sqlmock.Newx()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	repo := NewTransactionsPostgres(db)

	tests := []struct {
		name        string
		repo        *TransactionsPostgres
		transaction models.Transaction
		id          int64
		status      string
		mock        func()
		want        int64
		wantErr     bool
	}{
		{
			name: "StatusError",
			repo: repo,
			transaction: models.Transaction{
				UserId:    1,
				UserEmail: "fakeemail@gmail.com",
				Amount:    100,
				Currency:  "rubles",
			},
			id:     1,
			status: statusError,
			mock: func() {
				mock.ExpectExec("UPDATE transactions tt SET").WillReturnResult(sqlmock.NewResult(1, 0))
			},
			wantErr: true,
		},
		{
			name: "StatusSuccess",
			repo: repo,
			transaction: models.Transaction{
				Id:        1,
				UserId:    1,
				UserEmail: "fakeemail@gmail.com",
				Amount:    100,
				Currency:  "rubles",
				Created:   time.Now(),
				Changed:   time.Now(),
				Status:    "NEW",
			},
			id:     1,
			status: statusSuccess,
			mock: func() {
				mock.ExpectExec("UPDATE transactions tt SET").WillReturnResult(sqlmock.NewResult(1, 0))
			},
			wantErr: true,
		},
		{
			name: "StatusFailed",
			repo: repo,
			transaction: models.Transaction{
				Id:        1,
				UserId:    1,
				UserEmail: "fakeemail@gmail.com",
				Amount:    100,
				Currency:  "rubles",
				Created:   time.Now(),
				Changed:   time.Now(),
				Status:    "NEW",
			},
			id:     1,
			status: statusFailed,
			mock: func() {
				mock.ExpectExec("UPDATE transactions tt SET").WillReturnResult(sqlmock.NewResult(1, 0))
			},
			wantErr: true,
		},
		{
			name: "StatusCancelled",
			repo: repo,
			transaction: models.Transaction{
				Id:        1,
				UserId:    1,
				UserEmail: "fakeemail@gmail.com",
				Amount:    100,
				Currency:  "rubles",
				Created:   time.Now(),
				Changed:   time.Now(),
				Status:    "NEW",
			},
			id:     1,
			status: statusCancelled,
			mock: func() {
				mock.ExpectExec("UPDATE transactions tt SET").WillReturnResult(sqlmock.NewResult(1, 0))
			},
			wantErr: true,
		},
		{
			name: "WrongStatus",
			repo: repo,
			transaction: models.Transaction{
				Id:        1,
				UserId:    1,
				UserEmail: "fakeemail@gmail.com",
				Amount:    100,
				Currency:  "rubles",
				Created:   time.Now(),
				Changed:   time.Now(),
				Status:    "NEW",
			},
			id:     1,
			status: "abracadabra",
			mock: func() {
				mock.ExpectExec("UPDATE transactions tt SET").WillReturnResult(sqlmock.NewResult(1, 0))
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.repo.Create(tt.transaction)
			tt.repo.CancelTransactionById(tt.id)
			tt.mock()
			err := tt.repo.ChangeStatus(tt.id, tt.status)
			if (err != nil) != tt.wantErr {
				t.Errorf("Got error new = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}

func TestDelivery_Cancel(t *testing.T) {
	db, mock, err := sqlmock.Newx()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	repo := NewTransactionsPostgres(db)

	tests := []struct {
		name        string
		repo        *TransactionsPostgres
		transaction models.Transaction
		id          int64
		status      string
		mock        func()
		want        int64
		wantErr     bool
	}{
		{
			name: "StatusNew",
			repo: repo,
			transaction: models.Transaction{
				UserId:    1,
				UserEmail: "fakeemail@gmail.com",
				Amount:    100,
				Currency:  "rubles",
			},
			id:     1,
			status: statusNew,
			mock: func() {
				mock.ExpectExec("UPDATE transactions tt SET").WillReturnResult(sqlmock.NewResult(1, 1))
			},
			wantErr: false,
		},
		{
			name: "StatusError",
			repo: repo,
			transaction: models.Transaction{
				UserId:    1,
				UserEmail: "fakeemail@gmail.com",
				Amount:    100,
				Currency:  "rubles",
			},
			id:     1,
			status: statusError,
			mock: func() {
				mock.ExpectExec("UPDATE transactions tt SET").WillReturnResult(sqlmock.NewResult(1, 1))
			},
			wantErr: false,
		},
		{
			name: "StatusSuccess",
			repo: repo,
			transaction: models.Transaction{
				Id:        1,
				UserId:    1,
				UserEmail: "fakeemail@gmail.com",
				Amount:    100,
				Currency:  "rubles",
				Created:   time.Now(),
				Changed:   time.Now(),
				Status:    "NEW",
			},
			id:     1,
			status: statusSuccess,
			mock: func() {
				mock.ExpectExec("UPDATE transactions tt SET").WillReturnResult(sqlmock.NewResult(1, 0))
			},
			wantErr: true,
		},
		{
			name: "StatusFailed",
			repo: repo,
			transaction: models.Transaction{
				Id:        1,
				UserId:    1,
				UserEmail: "fakeemail@gmail.com",
				Amount:    100,
				Currency:  "rubles",
				Created:   time.Now(),
				Changed:   time.Now(),
				Status:    "NEW",
			},
			id:     1,
			status: statusFailed,
			mock: func() {
				mock.ExpectExec("UPDATE transactions tt SET").WillReturnResult(sqlmock.NewResult(1, 0))
			},
			wantErr: true,
		},
		{
			name: "StatusCancelled",
			repo: repo,
			transaction: models.Transaction{
				Id:        1,
				UserId:    1,
				UserEmail: "fakeemail@gmail.com",
				Amount:    100,
				Currency:  "rubles",
				Created:   time.Now(),
				Changed:   time.Now(),
				Status:    "NEW",
			},
			id:     1,
			status: statusCancelled,
			mock: func() {
				mock.ExpectExec("UPDATE transactions tt SET").WillReturnResult(sqlmock.NewResult(1, 0))
			},
			wantErr: true,
		},
		{
			name: "WrongStatus",
			repo: repo,
			transaction: models.Transaction{
				Id:        1,
				UserId:    1,
				UserEmail: "fakeemail@gmail.com",
				Amount:    100,
				Currency:  "rubles",
				Created:   time.Now(),
				Changed:   time.Now(),
				Status:    "NEW",
			},
			id:     1,
			status: "abracadabra",
			mock: func() {
				mock.ExpectExec("UPDATE transactions tt SET").WillReturnResult(sqlmock.NewResult(1, 0))
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.repo.Create(tt.transaction)
			if tt.status != statusNew {
				tt.repo.ChangeStatus(tt.id, tt.status)
			}
			tt.mock()
			err := tt.repo.CancelTransactionById(tt.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("Got error new = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}

func TestDelivery_CheckStatusById(t *testing.T) {
	db, mock, err := sqlmock.Newx()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	repo := NewTransactionsPostgres(db)

	tests := []struct {
		name        string
		repo        *TransactionsPostgres
		transaction models.Transaction
		id          int64
		status      string
		mock        func()
		want        string
		wantErr     bool
	}{
		{
			name: "StatusNew",
			repo: repo,
			transaction: models.Transaction{
				UserId:    1,
				UserEmail: "fakeemail@gmail.com",
				Amount:    100,
				Currency:  "rubles",
			},
			id:     1,
			status: statusNew,
			mock: func() {
				rows := sqlmock.NewRows([]string{"id"}).AddRow(statusNew)
				mock.ExpectQuery("SELECT tt.stat FROM transactions").WithArgs(1).WillReturnRows(rows)
			},
			want:    statusNew,
			wantErr: false,
		},
		{
			name: "StatusError",
			repo: repo,
			transaction: models.Transaction{
				UserId:    1,
				UserEmail: "fakeemail@gmail.com",
				Amount:    100,
				Currency:  "rubles",
			},
			id:     1,
			status: statusError,
			mock: func() {
				rows := sqlmock.NewRows([]string{"id"}).AddRow(statusError)
				mock.ExpectQuery("SELECT tt.stat FROM transactions").WithArgs(1).WillReturnRows(rows)
			},
			want:    statusError,
			wantErr: false,
		},
		{
			name: "StatusSuccess",
			repo: repo,
			transaction: models.Transaction{
				Id:        1,
				UserId:    1,
				UserEmail: "fakeemail@gmail.com",
				Amount:    100,
				Currency:  "rubles",
				Created:   time.Now(),
				Changed:   time.Now(),
				Status:    "NEW",
			},
			id:     1,
			status: statusSuccess,
			mock: func() {
				rows := sqlmock.NewRows([]string{"id"}).AddRow(statusSuccess)
				mock.ExpectQuery("SELECT tt.stat FROM transactions").WithArgs(1).WillReturnRows(rows)
			},
			want:    statusSuccess,
			wantErr: false,
		},
		{
			name: "StatusFailed",
			repo: repo,
			transaction: models.Transaction{
				Id:        1,
				UserId:    1,
				UserEmail: "fakeemail@gmail.com",
				Amount:    100,
				Currency:  "rubles",
				Created:   time.Now(),
				Changed:   time.Now(),
				Status:    "NEW",
			},
			id:     1,
			status: statusFailed,
			mock: func() {
				rows := sqlmock.NewRows([]string{"id"}).AddRow(statusFailed)
				mock.ExpectQuery("SELECT tt.stat FROM transactions").WithArgs(1).WillReturnRows(rows)
			},
			want:    statusFailed,
			wantErr: false,
		},
		{
			name: "StatusCancelled",
			repo: repo,
			transaction: models.Transaction{
				Id:        1,
				UserId:    1,
				UserEmail: "fakeemail@gmail.com",
				Amount:    100,
				Currency:  "rubles",
				Created:   time.Now(),
				Changed:   time.Now(),
				Status:    "NEW",
			},
			id:     1,
			status: statusCancelled,
			mock: func() {
				rows := sqlmock.NewRows([]string{"id"}).AddRow(statusCancelled)
				mock.ExpectQuery("SELECT tt.stat FROM transactions").WithArgs(1).WillReturnRows(rows)
			},
			want:    statusCancelled,
			wantErr: false,
		},
		{
			name: "WrongId",
			repo: repo,
			transaction: models.Transaction{
				Id:        1,
				UserId:    1,
				UserEmail: "fakeemail@gmail.com",
				Amount:    100,
				Currency:  "rubles",
				Created:   time.Now(),
				Changed:   time.Now(),
				Status:    "NEW",
			},
			id:     3,
			status: "abracadabra",
			mock: func() {
				rows := sqlmock.NewRows([]string{"id"}).AddRow("")
				mock.ExpectQuery("SELECT tt.stat FROM transactions").WithArgs(3).WillReturnRows(rows)
			},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.repo.Create(tt.transaction)
			if tt.status != statusNew {
				tt.repo.ChangeStatus(tt.id, tt.status)
			}
			tt.mock()
			got, err := tt.repo.CheckStatusById(tt.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("Got error new = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if err == nil && got != tt.want {
				t.Errorf("Got = %v, want %v", got, tt.want)
			}
		})
	}
}
