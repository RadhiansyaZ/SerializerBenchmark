import { Elysia } from "elysia";
var Serialize = require("./helper.js");

const app = new Elysia()
  .get("/", async () => {
    let result = await Serialize.api_operations();
    return result;
  })
  .listen(3010);

console.log(
  `🦊 Elysia is running at ${app.server?.hostname}:${app.server?.port}`,
);
