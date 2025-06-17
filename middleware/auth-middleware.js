module.export = function(req,res,next) {
  try {
    const headerAuth = req.headers.authorization 
  } catch (e) {
    return next(e) 
  }  
}
