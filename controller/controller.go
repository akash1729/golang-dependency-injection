package controller

import (
	"fmt"
	"github/akash1729/golang-dependency-injection/model"
	"net/http"
)

// Env : Holds env vaiables for our app
type Env struct {
	DBType *model.DataBaseType
}

// ControllerFunc : The last http handler that we want to call
func (env *Env) ControllerFunc(w http.ResponseWriter, r *http.Request) {

	fmt.Println("reached the last handler")
	fmt.Println(env.DBType.MeetAndGreet())

	fmt.Fprintf(w, "let's hope it works")

	fmt.Println("finished executing the last handler")
}
