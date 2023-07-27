package handlers

import (
	"context"
	mockservices "gateway/mock"
	"net/http"
	"net/http/httptest"
	"testing"

	pb "gateway/pb"

	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"google.golang.org/grpc"
)

func TestUserSignup(t *testing.T) {

	testCase := map[string]struct {
		input         interface{}
		buildStub     func(handlerMock *mockservices.MockMyServiceClient)
		checkResponse func(t *testing.T, responseRecorder *httptest.ResponseRecorder)
	}{
		"success": {
			buildStub: func(mockservices *mockservices.MockMyServiceClient) {
				req := &pb.Request{
					Data: "Mydata",
				}
				mockservices.EXPECT().MyMethod(context.Background(), req)
			},
			checkResponse: func(t *testing.T, responseRecorder *httptest.ResponseRecorder) {
				assert.Equal(t, http.StatusOK, responseRecorder.Code)

			},
		},
	}

	for testName, test := range testCase {
		test := test
		t.Run(testName, func(t *testing.T) {
			t.Parallel()
			ctrl := gomock.NewController(t)
			mockservices := mockservices.NewMockMyServiceClient(ctrl)
			test.buildStub(mockservices)

			authConn, err := grpc.Dial("localhost:5050", grpc.WithInsecure())
			if err != nil {

				t.Fatalf("Failed to dial gRPC server: %v", err)
			}

			authHandler := NewAuthenticationHandler(authConn)

			server := gin.Default()
			server.GET("/service1/healthcheck", authHandler.HealthCheck)

			mockRequest, err := http.NewRequest(http.MethodGet, "/service1/healthcheck", nil)
			assert.NoError(t, err)
			responseRecorder := httptest.NewRecorder()

			server.ServeHTTP(responseRecorder, mockRequest)

			test.checkResponse(t, responseRecorder)

		})

	}
}
