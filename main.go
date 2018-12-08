package main

func main(){

	gateway := NewGateway()

	gateway.server.Run(":9000")
}
