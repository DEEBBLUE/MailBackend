const db = require("./database.js")

class User{
  constructor(id,login,password,refresh){
    this.id = id
    this.login = login
    this.password = password
    this.refresh = refresh
  }
}

class MyOrm{
  constructor(db){
    this.localDb = db
    this.userCount = 0
  }

  createUser(login,password,refresh){
    if(this.localDb.read(login) === -2) {
      const user = new User(
        id = this.userCount,
        login = login,
        password = password,
        refresh = refresh
      )

      this.localDb.add(login,user)

      this.userCount += 1

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
}
module.exports = new MyOrm(db)  
