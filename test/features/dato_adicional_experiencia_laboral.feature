Feature: Validate API responses
  EXPERIENCIA_LABORAL_CRUD
  probe JSON reponses


Scenario Outline: To probe route code response  /dato_adicional_experiencia_laboral
  When I send "<method>" request to "<route>" where body is json "<bodyreq>"
  Then the response code should be "<codres>"      

  Examples: 
  |method|route                                 |bodyreq               |codres       |
  |GET   |/v1/dato_adicional_experiencia_laboral|./files/req/Vacio.json|200 OK       |
  |GET   |/v1/dato_adicional_experiencia_labora |./files/req/Vacio.json|404 Not Found|
  |POST  |/v1/dato_adicional_experiencia_labora |./files/req/Vacio.json|404 Not Found|
  |PUT   |/v1/dato_adicional_experiencia_labora |./files/req/Vacio.json|404 Not Found|
  |DELETE|/v1/dato_adicional_experiencia_labora |./files/req/Vacio.json|404 Not Found|


Scenario Outline: To probe response route /dato_adicional_experiencia_laboral       
  When I send "<method>" request to "<route>" where body is json "<bodyreq>"
  Then the response code should be "<codres>"      
  And the response should match json "<bodyres>"

  Examples:
  |method|route                                 |bodyreq               |codres         |bodyres                |
  |GET   |/v1/dato_adicional_experiencia_laboral|./files/req/Vacio.json|200 OK         |./files/res2/Vok1.json |
  |POST  |/v1/dato_adicional_experiencia_laboral|./files/req/Vacio.json|400 Bad Request|./files/res2/Ierr1.json|
