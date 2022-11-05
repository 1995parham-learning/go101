package handler_test

import (
	"bytes"
	"mime/multipart"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/1995parham-learning/go101/hello-multipart/handler"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/suite"
)

type HandlerTestSuite struct {
	suite.Suite
}

func (suite *HandlerTestSuite) TestFormValue() {
	require := suite.Require()
	e := echo.New()

	body := new(bytes.Buffer)

	reqWriter := multipart.NewWriter(body)
	require.NoError(reqWriter.WriteField("k1", "Hello World"))
	reqWriter.Close()

	req := httptest.NewRequest(http.MethodPost, "/", body)
	req.Header.Set(echo.HeaderContentType, reqWriter.FormDataContentType())

	rec := httptest.NewRecorder()

	ctx := e.NewContext(req, rec)

	require.NoError(handler.Handler(ctx))
}

func TestHandler(t *testing.T) {
	t.Parallel()

	suite.Run(t, new(HandlerTestSuite))
}
