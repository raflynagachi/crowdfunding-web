package controllers

type Controller struct {
	AuthController
	UserController
}

func RegisterController(auth AuthController, user UserController) Controller {
	return Controller{
		AuthController: auth,
		UserController: user,
	}
}
