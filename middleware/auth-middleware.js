const token = module.require("../service/token.service.js")

class Middleware{
  authMiddleware (req,res,next) {
    const headerAuth = req.headers.authorization 

    if(headerAuth === undefined){
      return res.status(401)
    }

    const accessToken = headerAuth.split(' ')[1];
    console.log(accessToken)

    if(accessToken === undefined){
      return res.status(401)
    }

    const userData = token.validAccessToken(accessToken)
    if(userData === undefined){
      return res.status(401)
    }

    next()
  }

}

module.exports = new Middleware()
