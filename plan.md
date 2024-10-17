- Client

  [x] - Fetches data about CPU, RAM, SSD and network usage from device that's installed on
  [x] - Sends data every second to server via HTTP

  [] - Sends data every X seconds to server via websocket
  [] - Can be installed via Docker and AUR, or build it manually
  [] - Has configurable settings
  [] - Can update settings via CLI flags
  [] - Can receive settings via HTTP route
  [] - Has local database for storing data that it hasn't sent yet (in case server is down)

- Server

  [] - Fetches data from the client via websocket and stores it in InfluxDB
  [] - Takes the data using frontend microframework via websocket
  [] - Displays the data using D3.js -> donuts, columns, and time series (one big div per agent)
  [] - If some of the usage goes above 90%, it sends the email to owner
  [] - Agents can be configured (update cycle, name, metric types)
  [] - Can be installed via Docker and AUR, or build it manually
  [] - Has auth. and main root account can be updated (email and/or password)
  [] - Root account can be setup in "init" process
