// app.js
const express = require("express");
const path = require("path");
const bodyParser = require("body-parser");
const fetch = global.fetch;

const app = express();
app.use(bodyParser.json());
app.use(express.static("public"));

app.post("/api/pack/calc", async (req, res) => {
  try {
    const response = await fetch("http://pack-service:8080/v1/pack/calc", {
      method: "POST",
      headers: { "Content-Type": "application/json" },
      body: JSON.stringify(req.body),
    });
    const data = await response.json();
    res.json(data);
  } catch (err) {
    res.status(500).json({ error: "Internal error" });
  }
});

app.listen(3000, () => {
  console.log("Frontend app listening on port 3000");
});