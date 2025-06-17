const token = require('../service/token.service.js')
const orm = require('../database/orm.js')
const deltatime = 30*24*60*60*1000 

class UserController{

  async register(req,res,next){
    try{
      const { login,passwd } = req.body

      const createUserRes = orm.createUser(login,passwd,refreshToken)

      if(createUserRes === true) {

        const { accessToken,refreshToken } = token.generateToken({lg: login}) 

        orm.addToken(login,refreshToken)    

        res.cookie("refreshToken", refreshToken,{ maxAge: deltatime, httpOnly: true })

        return res.json({accessToken,refreshToken})

      }else{
        return res.json("User already exist") 
      }
    }catch(e){
      next(e)
    }
  }

  async logIn(req,res,next) {
    try {
      const { login,passwd } = req.body

      const getUser = orm.findUser(login,passwd)

      if(getUser === true){

        const { accessToken,refreshToken } = token.generateToken({lg: login}) 
 
        orm.addToken(login,refreshToken)    

        res.cookie("refreshToken", refreshToken,{ maxAge: deltatime, httpOnly: true })

        return res.json({accessToken,refreshToken})

      }
    } catch (e) {
      next(e) 
    }  
  }

  async logOut(req,res,next){
    try{
      const { login } = req.body
      orm.deleteToken(login)

      res.clearCookie("refreshToken")
      return res.json("200")
    }catch(e){
      next(e)
    }
  }

  async refresh(req,res,next){
    try{
      const { refreshToken } = req.cookies;
      const userData = token.validRefreshTonke(refreshToken);
      if(orm.fingToken(userData,this.refresh)) {
        const { accessToken,refreshToken } = token.generateToken({lg: login}) 
 
        orm.addToken(login,refreshToken)    

        res.cookie("refreshToken", refreshToken,{ maxAge: deltatime, httpOnly: true })

        return res.json({accessToken,refreshToken})
      }
    }catch(e){
      next(e)
    }
  }
}

module.exports = new UserController()
