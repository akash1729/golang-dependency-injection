package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"

	"github/akash1729/golang-dependency-injection/controller"
	"github/akash1729/golang-dependency-injection/model"
)

func main() {

	port := os.Getenv("PORT")
	if port == "" {
		port = "8000" //localhost
	}

	dbType := model.DataBaseType("PostgreSQL")

	env := &controller.Env{DBType: &dbType}

	router := NewRouter(env)

	router.Use(firstMiddleware(env))
	router.Use(secondMiddleware(env))

	// starting server
	err := http.ListenAndServe(":"+port, router)
	if err != nil {
		log.Fatal(err)
	}

}

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

type Routes []Route

func NewRouter(env *controller.Env) *mux.Router {

	router := mux.NewRouter().StrictSlash(true)

	var routes = Routes{
		Route{
			"Index",
			"GET",
			"/defaultRoute",
			env.ControllerFunc,
		},
	}

	for _, route := range routes {

		//handler := route.HandlerFunc
		router.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(route.HandlerFunc)
	}
	return router
}

// We will wrap the middleware handler and pass the env, Since we can't break the signature of the handler
// We can use env as a receiver here but, then we need to keep the env, controllers and middleware in the same package
// In this approch we can split our app into different packages and including routing package
func firstMiddleware(env *controller.Env) func(http.Handler) http.Handler {

	//  middleware functionality of gorilla/mux will pass the next handler in the next argument
	return func(next http.Handler) http.Handler {

		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

			fmt.Println("started first middleware")

			// access dependency variable inside middleware
			fmt.Println(env.DBType.MeetAndGreet())

			// continue chaining
			next.ServeHTTP(w, r)
			fmt.Println("ended first middleware")
		})
	}
}

// We will wrap the middleware handler and pass the env, Since we can't break the signature of the handler
// We can use env as a receiver here but, then we need to keep the env, controllers and middleware in the same package
// In this approch we can split our app into different packages and including routing package
func secondMiddleware(env *controller.Env) func(http.Handler) http.Handler {

	return func(next http.Handler) http.Handler {

		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

			fmt.Println("started second middleware")
			// access dependecy variable inside second middleware

			fmt.Println(env.DBType.MeetAndGreet())

			// continue chaining
			next.ServeHTTP(w, r)
			fmt.Println("ended second middleware")
		})
	}
}
