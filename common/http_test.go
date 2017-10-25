package common

import (
	"bytes"
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
	"testing"
	"time"
)

///////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
//Test data structures, avoid import cycle
type TestupdateUserDetails struct {
	Description string `mandatory:"false" json:"description,omitempty"`
}

type listCompartmentsRequest struct {
	CompartmentID string `mandatory:"true" contributesTo:"query" name:"compartmentId"`
	Page          string `mandatory:"false" contributesTo:"query" name:"page"`
	Limit         int32  `mandatory:"false" contributesTo:"query" name:"limit"`
}

type updateUserRequest struct {
	UserID                string `mandatory:"true" contributesTo:"path" name:"userId"`
	TestupdateUserDetails `contributesTo:"body"`
	IfMatch               string `mandatory:"false" contributesTo:"header" name:"if-match"`
}

type TestcreateApiKeyDetails struct {
	Key string `mandatory:"true" json:"key,omitempty"`
}

type uploadApiKeyRequest struct {
	UserID                  string `mandatory:"true" contributesTo:"path" name:"userId"`
	TestcreateApiKeyDetails `contributesTo:"body"`
	OpcRetryToken           string `mandatory:"false" contributesTo:"header" name:"opc-retry-token"`
}

///////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

func TestHttpMarshallerInvalidStruct(t *testing.T) {
	request := http.Request{}
	err := HttpRequestMarshaller("asdf", &request)
	assert.Error(t, err, nil)
}

func TestHttpRequestMarshallerQuery(t *testing.T) {
	s := listCompartmentsRequest{CompartmentID: "ocid1", Page: "p", Limit: 23}
	request := &http.Request{}
	HttpRequestMarshaller(s, request)
	query := request.URL.Query()
	assert.True(t, query.Get("compartmentId") == "ocid1")
	assert.True(t, query.Get("page") == "p")
	assert.True(t, query.Get("limit") == "23")
}

func TestMakeDefault(t *testing.T) {
	r := MakeDefaultHttpRequest(http.MethodPost, "/one/two")
	assert.Equal(t, r.Header.Get("Content-Type"), "application/json")
	assert.NotEmpty(t, r.Header.Get("Date"))
	assert.NotEmpty(t, r.Header.Get("Opc-Client-Info"))
}

func TestHttpMarshallerSimpleHeader(t *testing.T) {
	s := updateUserRequest{UserID: "id1", IfMatch: "n=as", TestupdateUserDetails: TestupdateUserDetails{Description: "name of"}}
	request := MakeDefaultHttpRequest(http.MethodPost, "/random")
	HttpRequestMarshaller(s, &request)
	header := request.Header
	assert.True(t, header.Get("if-match") == "n=as")
}

func TestHttpMarshallerSimpleStruct(t *testing.T) {
	s := uploadApiKeyRequest{UserID: "111", OpcRetryToken: "token", TestcreateApiKeyDetails: TestcreateApiKeyDetails{Key: "thekey"}}
	request := MakeDefaultHttpRequest(http.MethodPost, "/random")
	HttpRequestMarshaller(s, &request)
	assert.True(t, strings.Contains(request.URL.Path, "111"))
}
func TestHttpMarshallerSimpleBody(t *testing.T) {
	desc := "theDescription"
	s := updateUserRequest{UserID: "id1", IfMatch: "n=as", TestupdateUserDetails: TestupdateUserDetails{Description: desc}}
	request := MakeDefaultHttpRequest(http.MethodPost, "/random")
	HttpRequestMarshaller(s, &request)
	body, _ := ioutil.ReadAll(request.Body)
	var content map[string]string
	json.Unmarshal(body, &content)
	assert.Contains(t, content, "description")
	if val, ok := content["description"]; !ok || val != desc {
		assert.Fail(t, "Should contain: "+desc)
	}
}

func TestHttpMarshalerAll(t *testing.T) {
	desc := "theDescription"
	s := struct {
		Id      string                `contributesTo:"path"`
		Name    string                `contributesTo:"query" name:"name"`
		When    time.Time             `contributesTo:"query" name:"when"`
		Income  float32               `contributesTo:"query" name:"income"`
		Male    bool                  `contributesTo:"header" name:"male"`
		Details TestupdateUserDetails `contributesTo:"body"`
	}{
		"101", "tapir", time.Now(), 3.23, true, TestupdateUserDetails{Description: desc},
	}
	request := MakeDefaultHttpRequest(http.MethodPost, "/")
	HttpRequestMarshaller(s, &request)
	var content map[string]string
	body, _ := ioutil.ReadAll(request.Body)
	json.Unmarshal(body, &content)
	when := s.When.Format(time.RFC3339)
	assert.True(t, request.URL.Path == "/101")
	assert.True(t, request.URL.Query().Get("name") == s.Name)
	assert.True(t, request.URL.Query().Get("income") == strconv.FormatFloat(float64(s.Income), 'f', 6, 32))
	assert.True(t, request.URL.Query().Get("when") == when)
	assert.Contains(t, content, "description")
	if val, ok := content["description"]; !ok || val != desc {
		assert.Fail(t, "Should contain: "+desc)
	}
}

func TestHttpMarshalerPointers(t *testing.T) {

	var n *string = new(string)
	*n = "theName"
	s := struct {
		Name *string `contributesTo:"query" name:"name"`
	}{
		n,
	}
	request := MakeDefaultHttpRequest(http.MethodPost, "/random")
	HttpRequestMarshaller(s, &request)
	assert.NotNil(t, request)
	assert.True(t, request.URL.Query().Get("name") == *s.Name)
}

func TestHttpMarshalerUntaggedFields(t *testing.T) {
	s := struct {
		Name  string `contributesTo:"query" name:"name"`
		AList []string
		AMap  map[string]int
		TestupdateUserDetails
	}{
		"theName", []string{"a", "b"}, map[string]int{"a": 1, "b": 2},
		TestupdateUserDetails{Description: "n"},
	}
	request := &http.Request{}
	e := HttpRequestMarshaller(s, request)
	assert.NoError(t, e)
	assert.NotNil(t, request)
	assert.True(t, request.URL.Query().Get("name") == s.Name)
}
func TestHttpMarshalerPathTemplate(t *testing.T) {
	urlTemplate := "/name/{userId}/aaa"
	s := uploadApiKeyRequest{UserID: "111", OpcRetryToken: "token", TestcreateApiKeyDetails: TestcreateApiKeyDetails{Key: "thekey"}}
	request := MakeDefaultHttpRequest(http.MethodPost, urlTemplate)
	e := HttpRequestMarshaller(s, &request)
	assert.NoError(t, e)
	assert.Equal(t, "/name/111/aaa", request.URL.Path)
}

func TestHttpMarshalerFunnyTags(t *testing.T) {
	s := struct {
		Name  string `contributesTo:"quer" name:"name"`
		AList []string
		AMap  map[string]int
		TestupdateUserDetails
	}{
		"theName", []string{"a", "b"}, map[string]int{"a": 1, "b": 2},
		TestupdateUserDetails{Description: "n"},
	}
	request := &http.Request{}
	e := HttpRequestMarshaller(s, request)
	assert.Error(t, e)
}

func TestHttpMarshalerUnsupportedTypes(t *testing.T) {
	s1 := struct {
		Name string         `contributesTo:"query" name:"name"`
		AMap map[string]int `contributesTo:"query" name:"theMap"`
	}{
		"theName", map[string]int{"a": 1, "b": 2},
	}
	s2 := struct {
		Name  string   `contributesTo:"query" name:"name"`
		AList []string `contributesTo:"query" name:"theList"`
	}{
		"theName", []string{"a", "b"},
	}
	s3 := struct {
		Name                  string `contributesTo:"query" name:"name"`
		TestupdateUserDetails `contributesTo:"query" name:"str"`
	}{
		"theName", TestupdateUserDetails{Description: "a"},
	}
	var n *string = new(string)
	col := make([]int, 10)
	*n = "theName"
	s4 := struct {
		Name *string `contributesTo:"query" name:"name"`
		Coll *[]int  `contributesTo:"query" name:"coll"`
	}{
		n, &col,
	}

	lst := []interface{}{s1, s2, s3, s4}
	for _, l := range lst {
		request := &http.Request{}
		e := HttpRequestMarshaller(l, request)
		Debugln(e)
		assert.Error(t, e)
	}
}

////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
//Response Unmarshaling
////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
// ListRegionsResponse wrapper for the ListRegions operation
type listRegionsResponse struct {

	// The []Region instance
	Items []int `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact Oracle about a
	// particular request, please provide the request ID.
	OpcRequestID string `presentIn:"header" name:"opcrequestid"`
}

type listUsersResponse struct {
	Items        []int     `presentIn:"body"`
	OpcRequestID string    `presentIn:"header" name:"opcrequestid"`
	OpcNextPage  int32     `presentIn:"header" name:"opcnextpage"`
	SomeUint     uint      `presentIn:"header" name:"someuint"`
	SomeBool     bool      `presentIn:"header" name:"somebool"`
	SomeTime     time.Time `presentIn:"header" name:"sometime"`
	SomeFloat    float64   `presentIn:"header" name:"somefloat"`
}

func TestUnmarshalResponse_StringHeader(t *testing.T) {
	header := http.Header{}
	opcId := "111"
	header.Set("OpcrequestId", opcId)
	r := http.Response{Header: header}
	s := listRegionsResponse{}
	err := UnmarshalResponse(&r, &s)
	assert.NoError(t, err)
	assert.Equal(t, s.OpcRequestID, opcId)

}

func TestUnmarshalResponse_MixHeader(t *testing.T) {
	header := http.Header{}
	opcId := "111"
	nextPage := int32(333)
	someuint := uint(12)
	somebool := true
	sometime := time.Now()
	somefloat := 2.556

	header.Set("OpcrequestId", opcId)
	header.Set("opcnextpage", strconv.FormatInt(int64(nextPage), 10))
	header.Set("someuint", strconv.FormatUint(uint64(someuint), 10))
	header.Set("somebool", strconv.FormatBool(somebool))
	header.Set("sometime", FormatTime(sometime))
	header.Set("somefloat", strconv.FormatFloat(somefloat, 'f', 3, 64))

	r := http.Response{Header: header}
	s := listUsersResponse{}
	err := UnmarshalResponse(&r, &s)
	assert.NoError(t, err)
	assert.Equal(t, s.OpcRequestID, opcId)
	assert.Equal(t, nextPage, s.OpcNextPage)
	assert.Equal(t, someuint, s.SomeUint)
	assert.Equal(t, somebool, s.SomeBool)
	assert.Equal(t, sometime.Format(time.RFC3339), s.SomeTime.Format(time.RFC3339))

}

type rgn struct {
	Key  string `mandatory:"false" json:"key,omitempty"`
	Name string `mandatory:"false" json:"name,omitempty"`
}

func TestUnmarshalResponse_SimpleBody(t *testing.T) {
	sampleResponse := `{"key" : "FRA","name" : "eu-frankfurt-1"}`
	header := http.Header{}
	opcId := "111"
	header.Set("OpcrequestId", opcId)
	s := struct {
		Rg rgn `presentIn:"body"`
	}{}
	r := http.Response{Header: header}
	bodyBuffer := bytes.NewBufferString(sampleResponse)
	r.Body = ioutil.NopCloser(bodyBuffer)
	err := UnmarshalResponse(&r, &s)
	assert.NoError(t, err)
	assert.Equal(t, "eu-frankfurt-1", s.Rg.Name)
}

func TestUnmarshalResponse_SimpleBodyList(t *testing.T) {
	sampleResponse := `[{"key" : "FRA","name" : "eu-frankfurt-1"},{"key" : "IAD","name" : "us-ashburn-1"}]`
	header := http.Header{}
	opcId := "111"
	header.Set("OpcrequestId", opcId)
	s := struct {
		Items []rgn `presentIn:"body"`
	}{}
	r := http.Response{Header: header}
	bodyBuffer := bytes.NewBufferString(sampleResponse)
	r.Body = ioutil.NopCloser(bodyBuffer)
	err := UnmarshalResponse(&r, &s)
	assert.NoError(t, err)
	assert.NotEmpty(t, s.Items)
	assert.Equal(t, "eu-frankfurt-1", s.Items[0].Name)
	assert.Equal(t, "IAD", s.Items[1].Key)
}

func TestUnmarshalResponse_SimpleBodyPtr(t *testing.T) {
	sampleResponse := `{"key" : "FRA","name" : "eu-frankfurt-1"}`
	header := http.Header{}
	opcId := "111"
	header.Set("OpcrequestId", opcId)
	s := struct {
		Rg *rgn `presentIn:"body"`
	}{}
	r := http.Response{Header: header}
	bodyBuffer := bytes.NewBufferString(sampleResponse)
	r.Body = ioutil.NopCloser(bodyBuffer)
	err := UnmarshalResponse(&r, &s)
	assert.NoError(t, err)
	assert.Equal(t, "eu-frankfurt-1", s.Rg.Name)
}

type testRnUnexported struct {
	Key  string `mandatory:"false" json:"key,omitempty"`
	Name string `mandatory:"false" json:"name,omitempty"`
}

type TestRn struct {
	Key  string `mandatory:"false" json:"key,omitempty"`
	Name string `mandatory:"false" json:"name,omitempty"`
}

type listRgRes struct {
	testRnUnexported `presentIn:"body"`
	OpcRequestID     string `presentIn:"header" name:"opcrequestid"`
}

type listRgResEx struct {
	TestRn       `presentIn:"body"`
	OpcRequestID string `presentIn:"header" name:"opcrequestid"`
}

func TestUnmarshalResponse_BodyAndHeaderUnex(t *testing.T) {
	sampleResponse := `{"key" : "FRA","name" : "eu-frankfurt-1"}`
	header := http.Header{}
	opcId := "111"
	header.Set("OpcrequestId", opcId)
	s := listRgRes{}
	r := http.Response{Header: header}
	bodyBuffer := bytes.NewBufferString(sampleResponse)
	r.Body = ioutil.NopCloser(bodyBuffer)
	err := UnmarshalResponse(&r, &s)
	assert.NoError(t, err)
	assert.Equal(t, opcId, s.OpcRequestID)
	assert.Equal(t, "", s.Name)
	assert.Equal(t, "", s.Key)
}

func TestUnmarshalResponse_BodyAndHeader(t *testing.T) {
	sampleResponse := `{"key" : "FRA","name" : "eu-frankfurt-1"}`
	header := http.Header{}
	opcId := "111"
	header.Set("OpcrequestId", opcId)
	s := listRgResEx{}
	r := http.Response{Header: header}
	bodyBuffer := bytes.NewBufferString(sampleResponse)
	r.Body = ioutil.NopCloser(bodyBuffer)
	err := UnmarshalResponse(&r, &s)
	assert.NoError(t, err)
	assert.Equal(t, opcId, s.OpcRequestID)
	assert.Equal(t, "eu-frankfurt-1", s.Name)
	assert.Equal(t, "FRA", s.Key)
}

////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
// BaseClient
////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
func TestBaseClient_prepareRequest(t *testing.T) {
	r := MakeDefaultHttpRequest(http.MethodPost, "/random")
	c := NewClient()
	c.ApiVersion = "v1"
	e := c.prepareRequest(&r)
	assert.NoError(t, e)
	assert.Equal(t, "/v1/random", r.URL.Path)

}