package repo

import (
	"context"
	"errors"

	"github.com/CutyDog/grpc-sample/services/account/internal/model"
	"gorm.io/gorm"
)

type AccountRepository interface {
	FindByID(ctx context.Context, id int64) (*model.Account, error)
	FindByUID(ctx context.Context, uid string) (*model.Account, error)
	Create(ctx context.Context, account *model.Account) error
	Update(ctx context.Context, account *model.Account) error
	Delete(ctx context.Context, id int64) error
}

type accountRepository struct {
	db *gorm.DB
}

func NewAccountRepository(db *gorm.DB) AccountRepository {
	return &accountRepository{db: db}
}

func (r *accountRepository) FindByID(ctx context.Context, id int64) (*model.Account, error) {
	var account model.Account
	if err := r.db.WithContext(ctx).First(&account, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &account, nil
}

func (r *accountRepository) FindByUID(ctx context.Context, uid string) (*model.Account, error) {
	var account model.Account
	if err := r.db.WithContext(ctx).Where("uid = ?", uid).First(&account).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &account, nil
}

func (r *accountRepository) Create(ctx context.Context, account *model.Account) error {
	return r.db.WithContext(ctx).Create(account).Error
}

func (r *accountRepository) Update(ctx context.Context, account *model.Account) error {
	return r.db.WithContext(ctx).Save(account).Error
}

func (r *accountRepository) Delete(ctx context.Context, id int64) error {
	return r.db.WithContext(ctx).Delete(&model.Account{}, id).Error
}
