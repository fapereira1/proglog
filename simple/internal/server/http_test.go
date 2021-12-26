package server

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestProducerHandler(t *testing.T) {

	body := ProduceRequest{
		Record: Record{
			Value: []byte("Hello World"),
		},
	}

	input, err := json.Marshal(body)
	if err != nil {
		t.Fatal(err)
	}

	req := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(string(input)))
	rr := httptest.NewRecorder()

	newHttpServer().handleProduce(rr, req)

	b := rr.Body.String()
	want := ProduceResponse{Offset: 0}

	var got ProduceResponse
	json.Unmarshal([]byte(b), &got)

	assert.Equal(t, want, got)

}
