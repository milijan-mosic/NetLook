# Agent

[x] - Fetches data about CPU, RAM, SSD and network usage from device that's installed on
[x] - Sends data every second to server via HTTP
[x] - Agent should try to send the data even if the connection is closed or server is down
[] - Has configurable settings
[] - Can update settings via CLI flags or env file

## Later

    [] - Can receive settings via HTTP route
    [] - Has local database for storing data that it hasn't sent yet (in case server is down)
    [] - Sends data to server via websocket
    [] - Can be installed via Docker and AUR, or build it manually

# Server

[x] - Fetches data from agents via HTTP
[x] - Stores data into the database, when received
[x] - Has a route for serving client with data
[] - Has configurable settings
[] - Can update settings via CLI flags or env file

# Client

[x] - Takes the data from backend using SolidJS via HTTP
[] - Displays the data using D3.js -> donuts, columns, and time series (one big div per agent)
[] - Has configurable settings
[] - Can update settings via env file
[] - All three are dockerized for local development

## Later

    [] - Agents can be configured (update cycle, name, metric types)

    [] - Root account can be setup in "init" process
    [] - Auth and "update profile" pages and functionalities...

    [] - Can be installed via Docker and AUR, or build it manually
    [] - If some of the usage goes above 90%, it sends the email to owner
    [] - Can serve the data via websocket
