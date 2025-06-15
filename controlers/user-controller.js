const token = require('../service/token.service.js')
const deltatime = 30*24*60*60*1000 


class UserController{
  async auth(req,res,next){
    try{
      const { login,passwd } = req.body
      const { accessToken,refreshToken } = token.generateToken({lg: login}) 
      console.log(refreshToken)
      res.cookie("refreshToken", refreshToken,{ maxAge: deltatime, httpOnly: true })
      return res.json(accessToken)
    }catch(e){
      console.log(e)
    }
  }
  async logOut(req,res,next){
    try{

    }catch(e){

    }
  }
  async refresh(req,res,next){
    try{
      res.json(['123',"456"])
    }catch(e){

    }
  }
}

module.exports = new UserController()
