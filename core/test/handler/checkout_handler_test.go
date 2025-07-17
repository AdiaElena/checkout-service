package handler_test

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"

	"github.com/AdiaElena/checkout-service/core/src/dto"
	"github.com/AdiaElena/checkout-service/core/src/handler"
	"github.com/AdiaElena/checkout-service/core/src/service"
	"github.com/AdiaElena/checkout-service/core/test/testdata/mocks"
)

func TestHandler_Checkout(t *testing.T) {
	gin.SetMode(gin.TestMode)

	tests := []struct {
		name      string
		body      string
		setupMock func(m *mocks.CheckoutServiceMock)
		wantCode  int
		wantError string // for 4xx/5xx
		wantTotal *int   // for 200
	}{
		{
			name:      "empty basket",
			body:      `{"skus":[]}`,
			setupMock: func(m *mocks.CheckoutServiceMock) { m.GetTotalPriceFunc = func() (int, error) { return 0, nil } },
			wantCode:  http.StatusOK,
			wantTotal: ptrInt(0),
		},
		{
			name:      "happy path",
			body:      `{"skus":["A","B","B","C"]}`,
			setupMock: func(m *mocks.CheckoutServiceMock) { m.GetTotalPriceFunc = func() (int, error) { return 115, nil } },
			wantCode:  http.StatusOK,
			wantTotal: ptrInt(115),
		},
		{
			name:      "invalid JSON",
			body:      `{"skus":[A]}`, // malformed
			setupMock: func(m *mocks.CheckoutServiceMock) {},
			wantCode:  http.StatusBadRequest,
			wantError: "invalid JSON",
		},
		{
			name:      "lowercase SKU",
			body:      `{"skus":["a"]}`,
			setupMock: func(m *mocks.CheckoutServiceMock) {},
			wantCode:  http.StatusBadRequest,
			// we expect the validator to fire on uppercase tag
			wantError: "failed on the 'uppercase' tag",
		},
		{
			name:      "multi-char SKU",
			body:      `{"skus":["AA"]}`,
			setupMock: func(m *mocks.CheckoutServiceMock) {},
			wantCode:  http.StatusBadRequest,
			// we expect the validator to fire on len=1 tag
			wantError: "failed on the 'len' tag",
		},
		{
			name: "scan error",
			body: `{"skus":["X"]}`,
			setupMock: func(m *mocks.CheckoutServiceMock) {
				m.ScanFunc = func(s string) error { return fmt.Errorf("invalid SKU: %q", s) }
			},
			wantCode:  http.StatusBadRequest,
			wantError: `invalid SKU: "X"`,
		},
		{
			name: "total error",
			body: `{"skus":["A"]}`,
			setupMock: func(m *mocks.CheckoutServiceMock) {
				m.GetTotalPriceFunc = func() (int, error) { return 0, fmt.Errorf("db error") }
			},
			wantCode:  http.StatusInternalServerError,
			wantError: "db error",
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			w := httptest.NewRecorder()
			c, _ := gin.CreateTestContext(w)
			c.Request = httptest.NewRequest("POST", "/checkout", strings.NewReader(tc.body))
			c.Request.Header.Set("Content-Type", "application/json")

			svcMock := mocks.NewCheckoutServiceMock()
			tc.setupMock(svcMock)

			h := handler.NewHandler(func() service.ICheckout { return svcMock })
			h.Checkout(c)

			assert.Equal(t, tc.wantCode, w.Code)

			if tc.wantCode == http.StatusOK {
				var resp dto.CheckoutResponse
				err := json.Unmarshal(w.Body.Bytes(), &resp)
				assert.NoError(t, err)
				assert.Equal(t, *tc.wantTotal, resp.Total)
			} else {
				var errResp dto.ErrorResponse
				err := json.Unmarshal(w.Body.Bytes(), &errResp)
				assert.NoError(t, err)
				assert.Contains(t, errResp.Error, tc.wantError)
			}
		})
	}
}

func ptrInt(v int) *int { return &v }
