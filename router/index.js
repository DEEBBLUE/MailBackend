const Router = require("express").Router
const token = require('../service/token.service.js')
const orm = require('../database/orm.js')
const middleware = require("../middleware/auth-middleware.js")

const deltatime = 30*24*60*60*1000 

const router = new Router()

router.post("/reg",function(req,res){

  const { login,passwd } = req.body
  const createUserRes = orm.createUser(login,passwd)

  if(createUserRes === true) {
    const { accessToken,refreshToken } = token.generateToken({lg: login}) 
    orm.addToken(login,refreshToken)    

    res.cookie("refreshToken", refreshToken,{ maxAge: deltatime, httpOnly: true })
    return res.json({accessToken,refreshToken})
  }
})

router.post("/login",function(req,res){

  const { login,passwd } = req.body
  const getUser = orm.findUser(login,passwd)

  if(getUser === true){
    const { accessToken,refreshToken } = token.generateToken({lg: login}) 
    orm.addToken(login,refreshToken)    

    res.cookie("refreshToken", refreshToken,{ maxAge: deltatime, httpOnly: true })
    return res.json({accessToken,refreshToken})
  }
})

router.post("/logOut",function(req,res){
  const { login } = req.body
  orm.deleteToken(login)

  res.clearCookie("refreshToken")
  return res.json("200")   
})

router.get("/refresh",function(req,res){
  const { refreshToken } = req.cookies;
  const userData = token.validRefreshTonke(refreshToken);
  if(orm.fingToken(userData,this.refresh)) {

    const { accessToken,refreshToken } = token.generateToken({lg: login}) 
    orm.addToken(login,refreshToken)    

    res.cookie("refreshToken", refreshToken,{ maxAge: deltatime, httpOnly: true })
    return res.json({accessToken,refreshToken})
  }
})

router.get("/fuzzing",middleware.authMiddleware,function(req,res){
  console.log("qur")
  const search = req.query.search
  const result = "asdfasdfasdf " + search
  const list = [
    result,
    result,
    result,
    result,
    result
  ]
  return res.json(list) 
})

module.exports = router
