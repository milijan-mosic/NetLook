- Agent

  [x] - Fetches data about CPU, RAM, SSD and network usage from device that's installed on
  [x] - Sends data every second to server via HTTP

  [] - Sends data to server via websocket
  [x] - Agent should try to send the data even if the connection is closed or server is down

  - Later

    [] - Has configurable settings
    [] - Can update settings via CLI flags
    [] - Can receive settings via HTTP route

    [] - Has local database for storing data that it hasn't sent yet (in case server is down)
    [] - Can be installed via Docker and AUR, or build it manually

- Server

  [x] - Fetches data from agents via HTTP

  [] - Fetches data from the client via websocket and stores it in InfluxDB
  [] - Can serve the data via websocket

  - Later

    [] - Takes the data from backend using SolidJS via websocket
    [] - Displays the data using D3.js -> donuts, columns, and time series (one big div per agent)
    [] - Agents can be configured (update cycle, name, metric types)

    [] - Root account can be setup in "init" process
    [] - Has auth. and main root account can be updated (email and/or password)

    [] - Can be installed via Docker and AUR, or build it manually
    [] - If some of the usage goes above 90%, it sends the email to owner
