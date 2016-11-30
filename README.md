**httptest is a simple http client for performing integration tests on http endpoints.**

Currently you can use it to test:

 - Response Status Code
 - Response Body
 - Content Type
 - Request method

httptest accepts a json or yaml manifest which defines your expectations for the service you would like to test.

If everything goes well the process will exit with code 0. It will exit with code 1 on any failure.

An example JSON manifest might look like this:

    {
      "api_version": 1,
      "manifest_name": "exampleServer",
      "host_name": "exampleServer",
      "host_port": 8080,
      "paths": [
        {
          "path": "/health",
          "method": "GET",
          "content_type": "text/plain; charset=utf-8",
          "response_status": 200,
          "response_body": "\"I'm pretty sure there's a lot more to life than being really, really, ridiculously good looking.\" - Derek Zoolander\n"
        },
        {
          "path": "/health",
          "method": "GET",
          "content_type": "application/json",
          "response_status": 200,
          "response_body": "{\"Body\":\"\\\"I'm pretty sure there's a lot more to life than being really, really, ridiculously good looking.\\\" - Derek Zoolander\\n\"}"
        },
        {
          "path": "/yummydata",
          "method": "POST",
          "content_type": "application/json",
          "response_status": 201,
          "response_body": "{\"Body\":\"hello world\"}"
        },
        {
          "path": "/should-404",
          "method": "GET",
          "content_type": "text/plain; charset=utf-8",
          "response_status": 404
        }
      ]
    }

An example YAML manifest can be found in this repo at:

    test/fixtures/fixtures.yml

An example of how to use httptest for integration testing with docker-compose and CircleCI can be found in this repo at `scripts/integration-test.sh`

**Example integration test usage:**

    docker run -ti \
      --net httptest \
      --link exampleServer_dev_local:exampleServer \
      -v $root_dir:/opt/$PROJECT_NAME \
      -e CONFIG_FILE=/opt/$PROJECT_NAME/test/fixtures/fixtures.json \
      i0nw/httptest:latest \
      dockerize -wait http://exampleServer:8080/health -timeout 120s httptest

The excellent dockerize (https://github.com/jwilder/dockerize) is included in the docker image for httptest, making it easy to override the default command and wait for another service to be ready before testing (as above)

The path to the httptest manifest can be passed as a flag `-config-file` or an environment variable `CONFIG_FILE` for convenience.

The docker image for httptest can be found at:
 https://hub.docker.com/r/i0nw/httptest/
It's very lightweight (Around 32mb) and is probably the easiest way to use httptest.

Alternatively binaries for linux-amd64 and darwin-amd64 are bundled with each release.

Suggestions and pull requests to make this tool better welcome!
