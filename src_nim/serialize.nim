import httpclient
import std/[times, json]
import options
import mummy

type
  Reply = object
    username: string
    comment: string

  Discussion = object
    username: string
    comment: string
    replies: seq[Reply]
    orders: Option[int]
    reply_count: Option[int]

  DiscussionResults = object
    discussion_results: seq[Discussion]


proc api_operation(): string {.discardable.} =
  # Create a new http client
  var client = newHttpClient()

  # Make api call and deserialize response to dictionary
  var response = client.getContent("http://localhost:8080/sample.json")
  #echo response
  let jsonResponse = parseJson(response)
  var responseData = to(jsonResponse, DiscussionResults)

  # Modify the response as expected result
  for i in 0 ..< responseData.discussion_results.len:
    responseData.discussion_results[i].orders = some(i+1)
    responseData.discussion_results[i].reply_count = some(responseData.discussion_results[i].replies.len)

  # Serialize the response to json
  let jsonStr = $ %*responseData

  return jsonStr


proc handler(request: Request) =
  case request.uri:
  of "/":
    if request.httpMethod == "GET":
      # Execution time benchmarking
      let time = cpuTime()

      # Prepare http response
      var headers: httpheaders.HttpHeaders
      headers["Content-Type"] = "application/json"
      var response = api_operation()

      # Measure duration for execution time benchmarking
      let duration = "Time taken: " & $(cpuTime() - time)
      let f = open("out.txt", fmWrite) # fmWrite is the file mode constant
      defer: f.close()
      f.writeLine(duration)

      request.respond(200, headers, response)
    else:
      request.respond(405)
  else:
    request.respond(404)


when isMainModule:
  let server = newServer(handler)
  echo "Serving on http://localhost:8081"
  server.serve(Port(8081))
