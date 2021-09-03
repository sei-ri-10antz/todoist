package http_test

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/google/go-cmp/cmp"
	todohttp "github.com/sei-ri-10antz/todoist/http"
)

type TestContext struct {
	w *httptest.ResponseRecorder
	r *http.Request
}

func (tc TestContext) Context(args ...gin.Param) *gin.Context {
	c, _ := gin.CreateTestContext(tc.w)
	c.Request = tc.r
	c.Params = append(c.Params, args...)
	return c
}

type TestResponse struct {
	StatusCode int
	Body       string
}

func TestService_Ping(t *testing.T) {
	type args struct {
		c TestContext
	}
	tests := []struct {
		name string
		args args
		want TestResponse
	}{
		{
			name: "ok",
			args: args{
				TestContext{
					w: httptest.NewRecorder(),
					r: httptest.NewRequest("GET", "http://any.url", nil),
				},
			},
			want: TestResponse{
				StatusCode: http.StatusOK,
				Body:       `{"ping":"pong}`,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &todohttp.Service{}
			s.Ping(tt.args.c.Context())

			resp := tt.args.c.w.Result()
			body, _ := ioutil.ReadAll(tt.args.c.w.Body)

			if diff := cmp.Diff(resp.StatusCode, tt.want.StatusCode); diff != "" {
				t.Errorf("StatusCode: -got/+want\n%s", diff)
			}
			if diff := cmp.Diff(string(body), tt.want.Body); diff != "" {
				t.Errorf("Body: -got/+want\n%s", diff)
			}
		})
	}
}
