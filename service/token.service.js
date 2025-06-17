const jwt = require("jsonwebtoken")

const access_key = "asdfa3fd:1fadf453@"
const refresh_key = "Fsdfa3fd:1fadf453@"

class TokenService{
  generateToken(payload){
    const accessToken = jwt.sign(payload, access_key,{expiresIn: "30m"}) 
    const refreshToken = jwt.sign(payload, refresh_key,{expiresIn: "30d"}) 
    
    console.log(accessToken)
    
    return {
      accessToken,
      refreshToken
    }
  }
  validRefreshTonke(refreshToken){
    const userData = jwt.verify(refreshToken,refresh_key);
    return userData
  }
  validAccessToken(accessToken){
    const userData = jwt.verify(accessToken,access_key);
    return userData
  }
}

module.exports = new TokenService();
