package users

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockRepository struct {
	mock.Mock
}

func (m *MockRepository) Save(user *Users) error {
	args := m.Called(user)
	return args.Error(0)
	// return nil
}

func (m *MockRepository) FindByEmail(email string) (*Users, error) {
	args := m.Called(email)
	return args.Get(0).(*Users), args.Error(1)
	// return nil, nil
}

func (m *MockRepository) FindByPhoneNumber(phone string) (*Users, error) {
	args := m.Called(phone)
	return args.Get(0).(*Users), args.Error(1)
	// return nil, nil
}

func (m *MockRepository) FindByEmailOrPhone(emailOrPhone string) (*Users, error) {
	args := m.Called(emailOrPhone)
	return args.Get(0).(*Users), args.Error(1)
	// return nil, nil
}

func TestRegister(t *testing.T) {
	mockRepo := new(MockRepository)
	service := NewService(mockRepo)

	tests := []struct {
		name          string
		input         RegisterInput
		mockBehavior  func()
		expectedError string
	}{
		{
			name: "success register with email",
			input: RegisterInput{
				Name:        "Jefri",
				Email:       "jefri@mail.com.com",
				PhoneNumber: "",
				Password:    "password123",
			},
			mockBehavior: func() {
				mockRepo.On("FindByEmail", "jefri@mail.com.com").Return(nil, nil)
				mockRepo.On("Save", mock.Anything).Return(nil)
			},
			expectedError: "",
		},
		{
			name: "email already registered",
			input: RegisterInput{
				Name:        "Jefri",
				Email:       "jefri@mail.com.com",
				PhoneNumber: "",
				Password:    "password123",
			},
			mockBehavior: func() {
				mockRepo.On("FindByEmail", "jefri@mail.com.com").Return(&Users{}, nil)
			},
			expectedError: "email already registered",
		},
		{
			name: "phone number already registered",
			input: RegisterInput{
				Name:        "Jefri",
				Email:       "",
				PhoneNumber: "123456789",
				Password:    "password123",
			},
			mockBehavior: func() {
				mockRepo.On("FindByPhoneNumber", "123456789").Return(&Users{}, nil)
			},
			expectedError: "phone number already registered",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mockBehavior()

			user, err := service.Register(tt.input)

			if tt.expectedError != "" {
				assert.Error(t, err)
				assert.EqualError(t, err, tt.expectedError)
			} else {
				assert.NoError(t, err)
				assert.NotNil(t, user)
			}

			mockRepo.AssertExpectations(t)
		})
	}
}

func TestLogin(t *testing.T) {
	mockRepo := new(MockRepository)
	service := NewService(mockRepo)

	tests := []struct {
		name          string
		input         LoginInput
		mockBehavior  func()
		expectedError string
		expectedToken string
	}{
		{
			name: "success login",
			input: LoginInput{
				EmailOrPhone: "jefri@mail.com.com",
				Password:     "password123",
			},
			mockBehavior: func() {
				mockRepo.On("FindByEmailOrPhone", "jefri@mail.com.com").Return(&Users{
					Password: "$2a$12$SZI.9ERg1Zzp.DYuYtnLMuG7yXfwHb8iz/izTKwQ1sz/Z5FgUz5Bi",
				}, nil)
			},
			expectedError: "",
			expectedToken: "some-jwt-token",
		},
		{
			name: "user not found",
			input: LoginInput{
				EmailOrPhone: "unknown@example.com",
				Password:     "password123",
			},
			mockBehavior: func() {
				mockRepo.On("FindByEmailOrPhone", "unknown@example.com").Return(nil, nil)
			},
			expectedError: "email or phone number not found",
		},
		{
			name: "invalid password",
			input: LoginInput{
				EmailOrPhone: "jefri@mail.com.com",
				Password:     "wrongpassword",
			},
			mockBehavior: func() {
				mockRepo.On("FindByEmailOrPhone", "jefri@mail.com.com").Return(&Users{
					Password: "$2a$12$SZI.9ERg1Zzp.DYuYtnLMuG7yXfwHb8iz/izTKwQ1sz/Z5FgUz5Bi",
				}, nil)
			},
			expectedError: "invalid password",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mockBehavior()

			response, err := service.Login(tt.input)

			if tt.expectedError != "" {
				assert.Error(t, err)
				assert.EqualError(t, err, tt.expectedError)
			} else {
				assert.NoError(t, err)
				assert.NotNil(t, response)
				assert.Equal(t, tt.expectedToken, response["token"])
			}

			mockRepo.AssertExpectations(t)
		})
	}
}
