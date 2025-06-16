const express = require("express");
const cors = require("cors");
const cookieParser = require("cookie-parser");
const errorMid = require("./middleware/error-midleware.js")
const router = require("./router/index.js");

const app = express()

app.use(express.json());
app.use(cookieParser());
app.use(cors());
app.use("/",router);
app.use(errorMid)

const start = async () => {
  try{
    app.listen(5000,() => console.log("Server started on Port=5000"))
  }catch(e){
    console.log(e)
  }
}
start()
