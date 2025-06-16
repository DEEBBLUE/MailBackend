class DataBase{
  constructor(base){
    this.base = base
  }
  add(key,value){
    if(!this.base[key]){
      this.base[key] = value
    }else{
      return -1 
    }
  }
  read(key){
    if(this.base[key]){
      return this.base[key]
    }else{
      return -2
    }
  }
  update(key,value){
    if(this.base[key]){
      this.base[key] = value
    }else{
      return -3
    }
  }
  remove(key){
    if(this.base[key]){
      delete this.base[key]
    }else{
      return -4
    }
  }
}

module.exports = new DataBase({}) 
