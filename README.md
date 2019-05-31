**jellz/deploy is a deploy server written in Go that makes deploying projects simple.**

### How does it work?
jellz/deploy uses webhooks to know when a change has been made to a project. Once the server is notified by the webhook, it will pull the latest code from the repository (e.g. GitHub).

### Installation
- Get the **[latest release](https://github.com/jellz/deploy/releases/latest)**.
- Fill `.env.example` with your information and rename the file to `.env`.
- Fill `projects.json.example` with your projects and rename the file to `projects.json`.
- Run the binary. It is a good idea to run it with GNU Screen or tmux so it keeps running even after you exit the terminal.

**It is important to note jellz/deploy only works on Linux at the moment.**