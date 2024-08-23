import falcon
import falcon.asgi
import time
import timeit
import requests
import orjson



# Falcon follows the REST architectural style, meaning (among
# other things) that you think in terms of resources and state
# transitions, which map to HTTP verbs.
#
def api_operations():
        result = requests.get('http://127.0.0.1:8080/sample.json')
        res_obj = orjson.loads(result.text)
        i = 1
        for discussion_result in res_obj['discussion_results']:
            discussion_result['orders'] = i
            discussion_result['reply_count'] = len(discussion_result['replies'])
            i = i + 1

        return orjson.dumps(res_obj)

class ThingsResource:

    async def on_get(self, req, resp):

        t = time.process_time()
        """Handles GET requests"""
        res = api_operations()

        resp.status = falcon.HTTP_200  # This is the default status
        resp.text = res

        elapsed_time = time.process_time() - t
        f = open("out.txt", "w")
        f.write(f"Elapsed time: {elapsed_time*1000}")
        f.close()

# print(timeit.timeit(api_operations))

# falcon.asgi.App instances are callable ASGI apps...
# in larger applications the app is created in a separate file
app = falcon.asgi.App()

# Resources are represented by long-lived class instances
things = ThingsResource()

# things will handle all requests to the '/things' URL path
app.add_route('/', things)
