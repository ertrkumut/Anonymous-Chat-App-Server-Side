#ARWServer GOLANG Documents

* [Intro](https://bitbucket.org/ertrkumut/arwserver/overview#markdown-header-intro)
* [ARWObject](https://bitbucket.org/ertrkumut/arwserver/overview#markdown-header-arwobject)
* [Event Handlers](https://bitbucket.org/ertrkumut/arwserver/overview#markdown-header-event-handlers)
	* [All Events](https://bitbucket.org/ertrkumut/arwserver/overview#markdown-header-all-arwevents)
	* [Login Handler](https://bitbucket.org/ertrkumut/arwserver/overview#markdown-header-login-handler)
#Intro

ARWServer is a socket server for games. Server side writen by golang and client	side writen by c# for Unity Game Engine.

You can download in this link [Download](https://bitbucket.org/ertrkumut/arwserver/get/3d2c9cd4d70f.zip)

First of all you need to create new variable as ARWServer
#
	var arwServer *ARWServer
	func main(){
		arwServer = new(ARWServer)
	}

After that you need to add your event handler which you want to handle.	In ARWServer you can easily yo add handler like that. And you need a method for handle this event.

#
	func main(){
		arwServer = new(ARWServer)

		arwServer.AddEventHandler(&(arwServer.events.Login), Login_Event_Handler)
	}

	func Login_Event_Handler(arwObject ARWObject){
		user, error := arwObject.GetUser(arwServer)
		if error != nil{
			fmt.Println("Error", error)
			return
		}

		fmt.Println("Login Success", user.name)
	}

After all of them. You need to initialize your arwServer object and you need to start processing events.

	func main(){
		arwServer = new(ARWServer)

		arwServer.AddEventHandler(&(arwServer.events.Login), Login_Event_Handler)
		arwServer.Initialize()

		arwServer.ProcessEvents()
	}

# ARWObject

Server side and client side communicate with ARWObject. Every request need to be ARWObject format. The other way server can not understand your data and can not process current event. If you want to send same variables server to client you need to create new ARWObject variable.

	var arwObject ARWObject

I have a weapon and it's damage is 150. I want to send this information to client.

	arwObject.PutInt("weapon_damage", 150)

or

	arwObject.PutFloat("weapon_damage", 150)

You can put what ever you want in this object. In the same time you can get variables like that.

	arwobject.GetString("variable_name")
	arwobject.GetInt("variable_name")
	arwobject.GetFloat("variable_name")
	arwobject.GetBool("variable_name")

# Event Handlers

In ARWServer all event handler need to be same format. 

	func event_handler(ARWObject)

Which event you want to handle, you need to add event handler for that event.


## All ARWEvents

* Connection
* Disconnection
* Login
* Join Room
* User Enter Room
* User Exit Room
* Extension Request

## Login Handler

When user is logged in, client send Login Request to server. And server create user variable. The user who is logged in, you can easily get from arwObject.

	func Login_Event_Handler(arwObject ARWObject){
		// GetUser(arwServer) return 2 parameter. 1. Logged in user - 2 Error
		user, error := arwObject.GetUser(arwServer)
		if error != nil{
			fmt.Println("Error", error)
			return
		}

		fmt.Println("Login Success", user.name)
	}