const express = require("express");
const cors = require("cors");
const cookieParser = require("cookie-parser");
const router = require("./router/index.js");

const app = express()

const CorsConfig ={
  origin: 'http://localhost:5173',
  credentials: true,
  optionsSuccessStatus: 200,
}

app.use(cors(CorsConfig));
app.use(express.json());
app.use(cookieParser());
app.use("/",router);

const start = async () => {
  try{
    app.listen(5000,() => console.log("Server started on Port=5000"))
  }catch(e){
    console.log(e)
  }
}
start()
