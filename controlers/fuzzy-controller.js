class FuzzyController{
  async fuzzing(req,res,next){
    try{
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
    }catch(e){
      console.log(e)
    } 
  }
}

module.exports = new FuzzyController()
