package controllers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httputil"
	"regexp"
	"strconv"

	"github.com/andypayne/goworldtime/models"
)

type worldTimeController struct {
	timeIDPat *regexp.Regexp
}

func newWorldTimeController() *worldTimeController {
	return &worldTimeController{
		//timeIDPat: regexp.MustCompile(`^/time/(\d+)/?\?(tz=(\s+))?`),
		timeIDPat: regexp.MustCompile(`^/time/(\d+)/?`),
	}
}

// https://golang.org/pkg/net/http/#Handler
func (tc worldTimeController) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	fmt.Println("Request to worldTimeController")
	//res.Write([]byte("World Time Web Service"))
	if req.URL.Path == "/times" {
		switch req.Method {
		case http.MethodGet:
			tc.getAll(res, req)
		case http.MethodPost:
			tc.post(res, req)
		default:
			res.WriteHeader(http.StatusNotImplemented)
		}
	} else {
		// Using the subgroup defined in the regexp
		matches := tc.timeIDPat.FindStringSubmatch(req.URL.Path)
		if len(matches) == 0 {
			res.WriteHeader(http.StatusNotFound)
			return
		}
		fmt.Println("matches =", matches)
		id, err := strconv.Atoi(matches[1])
		if err != nil {
			res.WriteHeader(http.StatusNotFound)
			return
		}
		switch req.Method {
		case http.MethodGet:
			tc.get(id, res)
		case http.MethodPut:
			tz := matches[2]
			fmt.Println("tz/matches[2] =", tz)
			tc.put(id, tz, res, req)
		case http.MethodDelete:
			tc.delete(id, res)
		default:
			res.WriteHeader(http.StatusNotImplemented)
		}
	}
}

func (tc *worldTimeController) getAll(res http.ResponseWriter, req *http.Request) {
	encodeResponseAsJSON(models.GetTimes(), res)
}

func (tc *worldTimeController) get(id int, res http.ResponseWriter) {
	t, err := models.GetWorldTimeByID(id)
	if err != nil {
		res.WriteHeader(http.StatusInternalServerError)
		return
	}
	encodeResponseAsJSON(t, res)
}

func (tc *worldTimeController) post(res http.ResponseWriter, req *http.Request) {
	t, err := tc.parseRequest(req)
	if err != nil {
		res.WriteHeader(http.StatusInternalServerError)
		res.Write([]byte("Unable to parse WorldTime"))
		return
	}
	t, err = models.AddTime(t)
	if err != nil {
		res.WriteHeader(http.StatusInternalServerError)
		res.Write([]byte(err.Error()))
	}
	encodeResponseAsJSON(t, res)
}

func (tc *worldTimeController) put(id int, tz string, res http.ResponseWriter, req *http.Request) {
	t, err := tc.parseRequest(req)
	if err != nil {
		res.WriteHeader(http.StatusInternalServerError)
		res.Write([]byte("Unable to parse WorldTime"))
		return
	}
	if id != t.Id {
		res.WriteHeader(http.StatusBadRequest)
		res.Write([]byte("The ID of the submitted WorldTime must match the ID in the url"))
		return
	}
	t, err = models.UpdateWorldTime(t, tz)
	if err != nil {
		res.WriteHeader(http.StatusInternalServerError)
		res.Write([]byte(err.Error()))
		return
	}
	encodeResponseAsJSON(t, res)
}

func (tc *worldTimeController) delete(id int, res http.ResponseWriter) {
	err := models.RemoveWorldTime(id)
	if err != nil {
		res.WriteHeader(http.StatusInternalServerError)
		res.Write([]byte(err.Error()))
		return
	}
	res.WriteHeader(http.StatusOK)
}

func (tc *worldTimeController) parseRequest(req *http.Request) (models.WorldTime, error) {
	reqDump, err := httputil.DumpRequest(req, true)
	if err != nil {
		fmt.Println("Error calling DumpRequest: ", err)
	} else {
		fmt.Println("request:\n", string(reqDump))
	}
	body, err := ioutil.ReadAll(req.Body)
	if err != nil {
		return models.WorldTime{}, err
	}
	var t models.WorldTime
	// https://ahmet.im/blog/golang-json-decoder-pitfalls/
	err = json.Unmarshal(body, &t)
	if err != nil {
		// UnmarshalTypeError?
		return models.WorldTime{}, err
	}
	return t, nil
}
