package service

import (
	"context"
	"testing"

	"github.com/kwtryo/go-sample/model"
	"github.com/kwtryo/go-sample/store"
	"github.com/kwtryo/go-sample/testutil"
	"github.com/stretchr/testify/assert"
	"golang.org/x/crypto/bcrypt"
)

func TestRegisterUser(t *testing.T) {
	type want struct {
		user *model.User
	}
	type test struct {
		fr   *model.FormRequest
		want want
	}

	tests := map[string]test{
		// 正常系
		"ok": {
			fr: getTestFormRequest(t),
			want: want{
				user: testutil.GetTestUser(t),
			},
		},
	}
	for n, tst := range tests {
		tst := tst
		ctx := context.Background()
		t.Run(n, func(t *testing.T) {
			t.Parallel()

			moqDb := &DBConnectionMock{}
			moqRepo := &UserRepositoryMock{}
			moqRepo.RegisterUserFunc =
				func(ctx context.Context, db store.DBConnection, u *model.User) (*model.User, error) {
					return u, nil
				}

			us := &UserService{
				DB:   moqDb,
				Repo: moqRepo,
			}
			got, err := us.RegisterUser(ctx, tst.fr)
			if err != nil {
				t.Fatalf("unexpected error: %v: ", err)
			}

			if err := bcrypt.CompareHashAndPassword([]byte(got.Password), []byte(tst.want.user.Password)); err != nil {
				t.Fatalf("password is wrong: %v", err)
			}
		})
	}
}

func TestGetUser(t *testing.T) {
	type want struct {
		user *model.User
	}
	type test struct {
		userName string
		want     want
	}

	testUser := testutil.GetTestUser(t)
	tests := map[string]test{
		// 正常系
		"ok": {
			userName: "testUser",
			want: want{
				user: testUser,
			},
		},
	}

	for n, tst := range tests {
		tst := tst
		ctx := context.Background()
		t.Run(n, func(t *testing.T) {
			t.Parallel()

			moqDb := &DBConnectionMock{}
			moqRepo := &UserRepositoryMock{}
			moqRepo.GetUserByUserNameFunc =
				func(ctx context.Context, db store.DBConnection, userName string) (*model.User, error) {
					return testutil.GetTestUser(t), nil
				}

			us := &UserService{
				DB:   moqDb,
				Repo: moqRepo,
			}
			got, err := us.GetUser(ctx, tst.userName)
			if err != nil {
				t.Fatalf("unexpected error: %v: ", err)
			}
			assert.Equal(t, tst.want.user, got)
		})
	}
}

// テスト用FormRequest構造体を返す
func getTestFormRequest(t *testing.T) *model.FormRequest {
	t.Helper()

	return &model.FormRequest{
		Name:     "testUserFullName",
		UserName: "testUser",
		Password: "testPassword",
		Role:     "admin",
		Email:    "test@example.com",
		Address:  "testAddress",
		Phone:    "000-0000-0000",
		Website:  "ttp://test.com",
		Company:  "testCompany",
	}
}
