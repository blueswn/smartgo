package main

import (
	"net/http"
	"fmt"
)

func main()  {
	http.HandleFunc("/router/rest", ServeHTTP)
	http.ListenAndServe(":8002", nil)
}

//Request represents an HTTP request.
type Request struct {
	action string
	method string
	pathInfo string
	request ParameterBag
}

//Create a new Illuminate HTTP request from server variables.
func (Request) capture(r *http.Request) (Request) {
	method := r.Method
	query := r.URL.Query()

	action, ok := query["action"];
	if !ok {
		//fmt.Fprintln(w, "404");
		panic(404)
	}
	delete(query, "action")
	req := Request{}
	req.request = ParameterBag{parameters:query}
	req.action = action[(len(action) -1)]
	req.method = method
	return req;
}

//Response represents an HTTP response.
type Response struct {

}

//Sets the response content.
func (res Response) setContent()  {

}
//Sets the response status code.
func (res Response) setStatusCode()  {

}
func (res Response) setProtocolVersion()  {

}
//Sends HTTP headers and content.
func (res Response) send(w http.ResponseWriter) {
	fmt.Fprintln(w, "hello world")
	//fmt.Fprintln(w, action)
}

//ParameterBag is a container for key/value pairs.
type ParameterBag struct {
	parameters map[string][]string
}

//Application
type Application struct {
	response Response
}

func (app Application) version() (string) {
	return "1.0.0"
}

//Handle an incoming HTTP request.
func (app Application) handle(req Request) (Response) {
	return app.response
}

func ServeHTTP(w http.ResponseWriter, r *http.Request)  {
	req := Request{}
	req = req.capture(r)

	app := Application{}
	response := app.handle(req)

	//fmt.Fprintln(w, req)
	//fmt.Println(req)
	//fmt.Println(req.request.parameters)
	//fmt.Println(req.action)
	//fmt.Println(response)

	response.send(w)
	//fmt.Fprintln(w, "hello world")
	//fmt.Fprintln(w, action)
}

func router()  {
	
}

