import Benchmark from "benchmark";
import { api_operations } from "./helper.js";

const suite = new Benchmark.Suite();

suite
  .add("api_ops", async () => await api_operations())
  .on("cycle", function (event) {
    console.log(String(event.target));
  })
  .on("complete", function () {
    console.log("Fastest is " + this.filter("fastest").map("name"));
  })
  .run({ async: true });
