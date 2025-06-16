module.export = function(err,req,res,next){
  console.log(err) 
  if(err.status === 401){
    return res.status(401).json("Not auth") 
  }
  return res.status(500)
}
