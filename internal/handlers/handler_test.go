package handlers

import (
	"fmt"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetItemHandler_Get(t *testing.T) {

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	repo := NewMockUebaRepository(ctrl)
	handler := New(repo)

	server := httptest.NewServer(http.HandlerFunc(handler.Get))
	defer server.Close()

	tests := []struct {
		name      string
		getParams string
		result    string
		httpCode  int
		expect    func()
	}{
		// TODO: Add test cases.
		{
			name:      "Find one id",
			getParams: "id=1",
			expect: func() {
				repo.EXPECT().GetItem("1").Return(map[string]string{
					"id": "1",
				}, nil)
			},
			result:   "[{\"id\":\"1\"}]\n",
			httpCode: http.StatusOK,
		},
		{
			name:      "Find several id",
			getParams: "id=1&id=3&id=2",
			expect: func() {
				repo.EXPECT().GetItem("1").Return(map[string]string{
					"id": "1",
				}, nil)
				repo.EXPECT().GetItem("3").Return(map[string]string{
					"id": "3",
				}, nil)
				repo.EXPECT().GetItem("2").Return(map[string]string{
					"id": "2",
				}, nil)
			},
			result:   "[{\"id\":\"1\"},{\"id\":\"3\"},{\"id\":\"2\"}]\n",
			httpCode: http.StatusOK,
		},
		{
			name:      "Trying to get absent ids",
			getParams: "id=1&id=3",
			expect: func() {
				repo.EXPECT().GetItem("1").Return(map[string]string{
					"id": "1",
				}, nil)
				repo.EXPECT().GetItem("3").Return(nil, fmt.Errorf("not found"))
			},
			result:   "[{\"id\":\"1\"}]\n",
			httpCode: http.StatusOK,
		},
		{
			name:      "Trying to get all absent ids",
			getParams: "id=1&id=3",
			expect: func() {
				repo.EXPECT().GetItem("1").
					Return(nil, fmt.Errorf("not found"))
				repo.EXPECT().GetItem("3").
					Return(nil, fmt.Errorf("not found"))
			},
			result:   "[]\n",
			httpCode: http.StatusOK,
		},
		{
			name:     "Bad request",
			httpCode: http.StatusBadRequest,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.expect != nil {
				tt.expect()
			}
		})
		res, err := http.Get(server.URL + "?" + tt.getParams)
		if err != nil {
			t.Error(err)
		}
		defer res.Body.Close()

		content, err := io.ReadAll(res.Body)
		if err != nil {
			t.Error(err)
		}

		assert.Equal(t, []byte(tt.result), content)
		assert.Equal(t, tt.httpCode, res.StatusCode)
	}
}
