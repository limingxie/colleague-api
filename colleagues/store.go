package colleagues

import (
	"context"
	"errors"
	"time"

	"github.com/hublabs/colleague-api/factory"
	"github.com/hublabs/colleague-api/tenants"
	"github.com/hublabs/common/api"
)

type Store struct {
	Id         int64           `json:"id"`
	TenantCode string          `json:"tenantCode"`
	Code       string          `json:"code"`
	Name       string          `json:"name"`
	Province   string          `json:"province"`
	City       string          `json:"city"`
	District   string          `json:"district"`
	Detail     string          `json:"detail"`
	Enable     bool            `json:"enable"`
	CreatedAt  time.Time       `json:"-" xorm:"created"`
	UpdatedAt  time.Time       `json:"-" xorm:"updated"`
	Brands     []tenants.Brand `json:"brands" xorm:"-"`
}

type StoreJsonView struct {
	Id     int64           `json:"id"`
	Code   string          `json:"code"`
	Name   string          `json:"name"`
	Role   string          `json:"role"`
	Brands []tenants.Brand `json:"brands"`
}

func (store *Store) CreateStore(ctx context.Context) error {
	s, err := Store{}.GetStoreByCode(ctx, store.TenantCode, store.Code)
	if err != nil {
		return err
	}

	if s.Id != 0 {
		return api.ErrorInvalidFields.New(errors.New("storeCode already exists"))
	}

	_, err = factory.DB(ctx).Insert(store)
	if err != nil {
		return err
	}

	return nil
}

func (store *Store) UpdateStore(ctx context.Context) error {
	s, err := Store{}.GetStoreById(ctx, store.TenantCode, store.Id)
	if err != nil {
		return err
	}

	if s.Id == 0 {
		return api.ErrorHasExisted.New(errors.New("Invalid store(storeId)"))
	}

	if _, err := factory.DB(ctx).ID(store.Id).
		Cols("name, province, city, district, detail").
		Update(store); err != nil {
		return err
	}

	return nil
}

func (Store) DeleteStore(ctx context.Context, tenantCode string, storeId int64) error {
	s, err := Store{}.GetStoreById(ctx, tenantCode, storeId)
	if err != nil {
		return err
	}

	if s.Id == 0 {
		return api.ErrorHasExisted.New(errors.New("Invalid store(storeId)"))
	}

	s.Enable = false
	if _, err := factory.DB(ctx).ID(s.Id).Cols("enable").
		Update(&s); err != nil {
		return err
	}

	return nil
}

func (Store) GetStoreAddressById(ctx context.Context, storeId int64) (map[string]interface{}, error) {
	var store Store
	if _, err := factory.DB(ctx).
		Where("store.id = ? ", storeId).
		And("store.enable = ? ", true).
		Get(&store); err != nil {
		return nil, err
	}

	return map[string]interface{}{
		"id":       store.Id,
		"province": store.Province,
		"city":     store.City,
		"district": store.District,
		"detail":   store.Detail,
	}, nil
}

func (Store) GetStoreAndBrandsByStoreId(ctx context.Context, storeId int64) (Store, error) {
	var store Store
	if _, err := factory.DB(ctx).
		Where("store.id = ? ", storeId).
		And("store.enable = ? ", true).
		Get(&store); err != nil {
		return Store{}, err
	}

	brands, err := Store{}.GetBrandsByStoreId(ctx, store.Id)
	if err != nil {
		return Store{}, err
	}

	store.Brands = brands

	return store, nil
}

func (Store) GetStoreByCode(ctx context.Context, tenantCode, storeCode string) (Store, error) {
	var store Store
	if has, err := factory.DB(ctx).Table("store").
		Where("store.code = ? ", storeCode).
		And("store.tenant_code = ? ", tenantCode).
		And("store.enable = ? ", true).
		Get(&store); err != nil {
		return Store{}, err
	} else if has {
		return store, nil
	} else {
		return Store{}, nil
	}
}

func (Store) GetStoreById(ctx context.Context, tenantCode string, storeId int64) (Store, error) {
	var store Store
	if has, err := factory.DB(ctx).
		Where("store.id = ? ", storeId).
		And("store.tenant_code = ? ", tenantCode).
		And("store.enable = ? ", true).
		Get(&store); err != nil {
		return Store{}, err
	} else if has {
		return store, nil
	} else {
		return Store{}, nil
	}

}
