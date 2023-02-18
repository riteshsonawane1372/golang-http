package main

import (

  "log"
	"net/http"

	"github.com/go-openapi/loads"
	"github.com/go-openapi/runtime/middleware"
	"github.com/riteshsonawane/http-restAPI/pkg/swagger/server/restapi"
	"github.com/riteshsonawane/http-restAPI/pkg/swagger/server/restapi/operations"
)

func main(){


    //Init Swagger

    swagger,err:=loads.Analyzed(restapi.SwaggerJSON,"")
    if err!=nil{
      log.Fatal(err)
    }

    api:=operations.NewHelloAPIAPI(swagger)
    server:=restapi.NewServer(api)

    defer func(){
        err:=server.Shutdown()
        if err!=nil{
          log.Fatal(err)
        }
    }()

    server.Port=8000

  
  
    api.CheckHealthHandler = operations.CheckHealthHandlerFunc(Health)

    api.GetHelloUserHandler = operations.GetHelloUserHandlerFunc(GetHelloUser)

    api.GetGopherNameHandler = operations.GetGopherNameHandlerFunc(GetGopherByName)

    // Start server 
    if err := server.Serve(); err != nil {
        log.Fatalln(err)
    }
}

//Health route returns OK
func Health(operations.CheckHealthParams) middleware.Responder {
    return operations.NewCheckHealthOK().WithPayload("OK")
}

//Returns Hello + UserName
func GetHelloUser(user operations.GetHelloUserParams) middleware.Responder {
  return operations.NewGetHelloUserOK().WithPayload("Hello " + user.User + "!")
}


func GetGopherByName(gopher operations.GetGopherNameParams) middleware.Responder{

  var URL string

  if gopher.Name!=""{
    URL="https://github.com/scraly/gophers/raw/main/" + gopher.Name + ".png"
  }else{
    URL = "https://github.com/scraly/gophers/raw/main/dr-who.png"
  }

  response,err:=http.Get(URL)
  if err!=nil{
    log.Fatal(err)
  }

  return operations.NewGetGopherNameOK().WithPayload(response.Body)

}






  





