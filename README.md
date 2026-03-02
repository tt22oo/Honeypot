# Honeypot
This is Honeypot with Go   
It collects attacker's connection data and inputs, sends them to the server, and stores them in MongoDB.

## Structure
* **honeypot/** : telnet, ssh honeypot
* **server/** : server for log collection

## Settings
**honeypot/configs.json:**
```json
{
    "name": "honeypot1",
    "key": "test@1234",
    "report_url": "http://127.0.0.1:8080/honeypot/report",
    "telnet": {
        "start": true,
        "addrs": ["0.0.0.0:2323"]
    },
    "ssh": {
        "start": true,
        "version": "SSH-2.0-OpenSSH_8.2p1 Ubuntu-4ubuntu0.2",
        "addrs": ["0.0.0.0:2222"]
    }
}
```
* `name`: this is honeypot's name. it's using to identify.
* `key`: it's using to authenticate requests to the report server.
* `report_url`: it's report server's url
* `telnet`: it's telnet's config.
  * `start`: it's whether to start telnet honeypot.
  * `addrs`: it's list of addresses and ports to listen on.
* `ssh`: it's ssh's config.
  * `start`: it's whether to start ssh honeypot.
  * `version`: it's ssh version string presented to attackers.
  * `addrs`: it's list of addresses and ports to listen on.

**server/configs.json:**
```json
{
    "listen": ":8080",
    "key": "test@1234",
    "mongo_db": "mongodb://localhost:27017",
    "database_name": "honeypot"
}
```
* `listen`: it's address and port for the report server.
* `key`: it's using to authenticate incoming reports.
* `mongo_db`: it's mongodb connection uri.
* `database_name`: it's name of the mongodb database using to store reports.
