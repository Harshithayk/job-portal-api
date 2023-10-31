package services

import (
	"errors"
	"project/internal/model"
	"reflect"
	"testing"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type Mokeusersigup struct{}

func (m *Mokeusersigup) CreateUser(m1 model.User) (model.User, error) {
	if m1.UserName == "" {
		return model.User{}, errors.New("users name is empty")
	}
	if m1.Email == "" {
		return model.User{}, errors.New("users EMAIL is empty")
	}
	if m1.PasswordHash == "" {
		return model.User{}, errors.New("users PasswordHash is empty")
	}
	return model.User{UserName: m1.UserName, Email: m1.Email}, nil
}

func (m *Mokeusersigup) FetchUserByEmail(s string) (model.User, error) {
	if s == "" {
		return model.User{}, errors.New("OUTPUT IS NOT PROPER")
	}
	return model.User{UserName: "harsgi", Email: "ykharshi@gmail.com", PasswordHash: "$2a$10$votXUqKwkXe6l5.2aVKSU.08QEPzZYuXy47OP7JuHebrZSppBlYSW"}, nil
}

func TestService_UserSignup(t *testing.T) {
	// type args struct {
	// 	nu model.UserSignup
	// }
	tests := []struct {
		name    string
		s       *Service
		nu      model.UserSignup
		want    model.User
		wantErr bool
	}{
		{name: "case pass",
			s:       &Service{r: &Mokeusersigup{}},
			nu:      model.UserSignup{UserName: "harshi", Email: "ykharshi@gmail.com"},
			want:    model.User{UserName: "harshi", Email: "ykharshi@gmail.com"},
			wantErr: true,
		},
		{name: "case fail",
			s:       &Service{r: &Mokeusersigup{}},
			nu:      model.UserSignup{UserName: "", Email: "ykharshi@gmail.com"},
			want:    model.User{},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.s.UserSignup(tt.nu)
			if (err == nil) != tt.wantErr {
				t.Errorf("Service.UserSignup() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Service.UserSignup() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestService_Userlogin(t *testing.T) {
	// type args struct {
	// 	l model.UserLogin
	// }
	tests := []struct {
		name    string
		s       *Service
		nu      model.UserLogin
		want    jwt.RegisteredClaims
		wantErr bool
	}{
		{name: "case pass",
			s:       &Service{r: &Mokeusersigup{}},
			nu:      model.UserLogin{Email: "ykharshi@gmail.com", Password: "dfsgdy532"},
			want:    jwt.RegisteredClaims{Issuer: "service project", Subject: "0", Audience: jwt.ClaimStrings{"users"}, ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour)), IssuedAt: jwt.NewNumericDate(time.Now())},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.s.Userlogin(tt.nu)
			if (err == nil) != tt.wantErr {
				t.Errorf("Service.Userlogin() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Service.Userlogin() = %v, want %v", got, tt.want)
			}
		})
	}
}
