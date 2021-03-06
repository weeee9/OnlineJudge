package routers

import (
	"OnlineJudge/controllers"
	"github.com/astaxie/beego"
)

func init() {
	// List problems (or have a seperate controller as IndexController?)
	beego.Router("/", &controllers.ProblemController{}, "*:List")

	// On User -- Includes user profiles, settings and login and sign up routes
	beego.AutoRouter(&controllers.UserController{})

	// On Problems
	beego.Router("/problem/create", &controllers.ProblemController{}, "*:Create;post:SaveProblem")
	beego.Router("/problem/:id/addtest", &controllers.ProblemController{}, "*:AddTestCase;post:SaveTestCase")
	beego.Router("/problem/:id", &controllers.ProblemController{}, "*:ProblemById")
	beego.Router("/problem/:type/:page", &controllers.ProblemController{}, "*:ProblemByCategory")
	beego.Router("/problem/:id/submit", &controllers.ProblemController{}, "post:SaveSubmission") // ->ProblemController(notes that user has tried solving problem)->ExecController(seek for helper to exec)->ProblemController(get result info & build on it)
	beego.Router("/problem/:id/run", &controllers.ProblemController{}, "post:RunCode")
	beego.Router("/test", &controllers.ProblemController{}, "post:Test")
	// beego.Router("/problem/:id/edit", &controllers.ProblemController{}, "post:SaveProblem;*:Edit")
}
