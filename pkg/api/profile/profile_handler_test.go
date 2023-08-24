package profile

import (
	"github.com/shirrashko/BuildingAServer-step2/pkg/api/model"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
)

// MockService is a mock implementation of the profile service
type MockService struct {
	mock.Mock
}

func (m *MockService) GetProfileByID(id int) (model.UserProfile, error) {
	args := m.Called(id)
	return args.Get(0).(model.UserProfile), args.Error(1)
}

func (m *MockService) UpdateUserProfile(id int, profile model.UserProfile) error {
	args := m.Called(id, profile)
	return args.Error(0)
}

func (m *MockService) CreateNewProfile(profile model.UserProfile) (int, error) {
	args := m.Called(profile)
	return args.Int(0), args.Error(1)
}

type APITestSuite struct {
	suite.Suite
	Recorder *httptest.ResponseRecorder
	Context  *gin.Context
	Service  *MockService
}

func (s *APITestSuite) SetupTest() {
	s.Recorder = httptest.NewRecorder()
	s.Context, _ = gin.CreateTestContext(s.Recorder)
	s.Service = new(MockService)
}

func (s *APITestSuite) TestGetProfileByID() {
	// Arrange
	req, _ := http.NewRequest(http.MethodGet, "/1", nil)
	s.Context.Request = req

	mockProfile := model.UserProfile{ID: 1, Username: "john_doe", FullName: "John Doe"}
	s.Service.On("GetProfileByID", 1).Return(mockProfile, nil)

	handler := NewHandler(s.Service)

	// Act
	handler.getProfileByID(s.Context)

	// Assert
	assert.Equal(s.T(), http.StatusOK, s.Recorder.Code)
}

func TestAPITestSuite(t *testing.T) {
	suite.Run(t, new(APITestSuite))
}
