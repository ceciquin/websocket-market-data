## Tabla de contenidos
* [Informacion General](#nformacion-general)
* [Technologies](#tecnologias)
* [Setup](#setup)

## Informacion General
- desarrollo de un servidor-cliente que publica y procesa precios de mercado.
- colaboradora : Cecilia Maria Quintana Amarilla 
- ejecucion del programa : terminal o linea de comandos
- SO : Windows 10 o macOS 
- codigo fuente y sus Archivos y carpetas principales: 


  ├──cmd/
    |  ├──client/
    |  |  ├──myclient.go
    |  ├──server/
    |  |  ├──myserver.go
    |  | 
  ├──pkg/
    |  ├──client/
    |  |  ├──client.go
    |  ├──server/
    |  |  ├──market.go
    |  |  ├──server.go
  ├──test/
       └──myclient_test.go
       
- Descripcion del projecto : “websocket-market-data” es un programa diseñado para transmitir datos de mercados en una conxion dual entre el cliente y el servidor utilizando la conexion websocket. Esto nos va a permitir tener una bidereccionalidad de la data. 


## Tecnologias
Este projecto fue creado usando las siguientes tecnologias:

* go1.20.1 darwin/amd64
*la conexion entre el servidor y el cliente se realizo usando una conexion dual de tipo websocket, usando el modulo **golang.org/x/net/websocket**

	
## Setup
Para correr de manera local esta aplicacion es necesario contar con **Golang** y correr los siguientes comandos:

primera terminal, desde **market-data-websocket-connection/cmd/server**
```
$ go run myserver.go 

```

primera terminal, desde **market-data-websocket-connection/cmd/client**
```
$ go run myclient.go 

```

**Nota:** luego de correr los comandos, hay que presionar ctrl + c en la terminal para terminar la conexion.
