Feature: Validate API responses
  EXPERIENCIA_LABORAL_CRUD
  probe JSON reponses


Scenario Outline: To probe route code response  /tipo_dedicacion
  When I send "<method>" request to "<route>" where body is json "<bodyreq>"
  Then the response code should be "<codres>"      

  Examples: 
  |method|route              |bodyreq               |codres       |
  |GET   |/v1/tipo_dedicacion|./files/req/Vacio.json|200 OK       |
  |GET   |/v1/tipo_dedicacio |./files/req/Vacio.json|404 Not Found|
  |POST  |/v1/tipo_dedicacio |./files/req/Vacio.json|404 Not Found|
  |PUT   |/v1/tipo_dedicacio |./files/req/Vacio.json|404 Not Found|
  |DELETE|/v1/tipo_dedicacio |./files/req/Vacio.json|404 Not Found|

Scenario Outline: To probe response route /tipo_dedicacion       
  When I send "<method>" request to "<route>" where body is json "<bodyreq>"
  Then the response code should be "<codres>"      
  And the response should match json "<bodyres>"

  Examples:
  |method|route              |bodyreq               |codres         |bodyres                |
  |GET   |/v1/tipo_dedicacion|./files/req/Vacio.json|200 OK         |./files/res6/Vok1.json |
  |POST  |/v1/tipo_dedicacion|./files/req/Yt1.json  |201 Created    |./files/res6/Vok2.json |
  |POST  |/v1/tipo_dedicacion|./files/req/Vacio.json|400 Bad Request|./files/res6/Ierr1.json|
  |POST  |/v1/tipo_dedicacion|./files/req/Nt1.json  |400 Bad Request|./files/res6/Ierr2.json|
  |POST  |/v1/tipo_dedicacion|./files/req/Nt2.json  |400 Bad Request|./files/res6/Ierr3.json|
  |POST  |/v1/tipo_dedicacion|./files/req/Nt3.json  |400 Bad Request|./files/res6/Ierr4.json|
  |POST  |/v1/tipo_dedicacion|./files/req/Nt4.json  |400 Bad Request|./files/res6/Ierr5.json|
  |POST  |/v1/tipo_dedicacion|./files/req/Nt5.json  |400 Bad Request|./files/res6/Ierr6.json| 
  |POST  |/v1/tipo_dedicacion|./files/req/Nt6.json  |400 Bad Request|./files/res6/Ierr7.json| 
  |POST  |/v1/tipo_dedicacion|./files/req/Yt2.json  |400 Bad Request|./files/res6/Ierr8.json| 
  |PUT   |/v1/tipo_dedicacion|./files/req/Yt2.json  |200 OK         |./files/res6/Vok2.json |
  |GETID |/v1/tipo_dedicacion|./files/req/Vacio.json|200 OK         |./files/res6/Vok2.json |
  |DELETE|/v1/tipo_dedicacion|./files/req/Vacio.json|200 OK         |./files/res6/Ino.json  |
  |DELETE|/v1/tipo_dedicacion|./files/req/Vacio.json|404 Not Found  |./files/res6/Ierr9.json|
