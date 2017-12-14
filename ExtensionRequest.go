package main

type ExtensionHandler func(*ARWServer, *ARWUser, ARWObject)

type ExtensionRequest struct{
	cmd string
	handler ExtensionHandler
}
