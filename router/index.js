const Router = require("express").Router
const usetController = require("../controlers/user-controller.js")

const router = new Router()

router.post("/auth",usetController.auth)
router.post("/logOut",usetController.logOut)
router.get("/refresh",usetController.refresh)

module.exports = router
