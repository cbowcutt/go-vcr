package http

import(
	"net/http"
)

type Recorder interface {
	writeToFile(*http.Request)
}

type HttpRecorder struct {
	handler http.Handler
	recorder Recorder 
}


func (r *HttpRecorder)  ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// write contents of request to file
	handler.ServeHttp(w, r)
	r.recorder.WriteToFile(r, )
	// pass request to 
}