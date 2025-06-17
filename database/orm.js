const db = require("./database.js")

class User{
  constructor(login,password,refresh){
    this.login = login
    this.password = password
    this.refresh = refresh
  }
}

class MyOrm{
  constructor(db){
    this.localDb = db
  }

  createUser(login,password){
    if(this.localDb.read(login) === -2) {
      let refresh = ""
      const user = new User(
        login = login,
        password = password,
        refresh = refresh 
      )

      this.localDb.add(login,user)

      return true 
    }else{
      return "Error: user already exist"
    }
  }

  findUser(login,password){
    const user = this.localDb.read(login)
    if( user == -2){
      return "Error: user not found"
    }else{
      if( user.password == password) {
        return true 
      }else{
        return "Error: login or password incoerct"
      }
    }
  }

  addToken(login,refresh){
    let user = this.localDb.read(login)
    user.refresh = refresh

    this.localDb.update(login,user)
  }

  deleteToken(login){
    let user = this.localDb.read(login)
    user.refresh = ""

    this.localDb.update(login,user)
  }

  fingToken(login,token){
    let user = this.localDb.read(login)
    
    if(user.refresh === token){
      return true;
    }else{
      return false;
    }
  }

}
module.exports = new MyOrm(db)  
