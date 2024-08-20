import axios from "axios";
const api_operations = async () => {
  // Initiate execution time tracking
  // console.time("execution");
  // Call an external API via axios
  const instance = axios.create({
    baseURL: "http://127.0.0.1:8080",
  });
  let request = await instance.get("/sample.json", {
    timeout: 600000,
  });

  //  modify the response
  let cnt = 1;
  var result = request.data.discussion_results;
  request.data.discussion_results.forEach((element) => {
    element.orders = cnt;
    element.reply_count = element.replies.length;
    cnt++;
  });
  // Measure the execution time
  // console.timeEnd("execution");
  return result;
};
export { api_operations };
