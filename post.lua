wrk.method = "POST"
wrk.body = "{\"query\": \"query{users{name}}\"}"
wrk.headers["Content-Type"] = "application/json"