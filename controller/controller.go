package controller

import (
	"fmt"
	"net/http"

	"github.com/akash1729/golang-dependency-injection/model"
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
