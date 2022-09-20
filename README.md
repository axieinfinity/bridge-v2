# Bridge v2

Bridge v2 is the second version of [bridge](https://github.com/axieinfinity/bridge) which is used for validators and relayers listen to events triggered from Ronin or other chains connected to Ronin (eg: Ethereum)

# How to run
## Requirements

- Ronin RPC url is used to listen/trigger events from Ronin chain
- Ethereum RPC url (alchemy, infura, etc.) is used to listen/trigger events from Ethereum
- Postgres db to store events and tasks

## Install and run
### Using docker
There is a `docker-compose.yaml` file in `docker` directory. Modify `.env` file and run `bridge` with the following command
```
docker-compose -f docker/docker-compose.yaml --env-file .env up -d
```
### Manually
    
Install postgres database, create `bridge` database

build and install bridge

```
make bridge
```

then run bridge with 

```
bridge --config <path-to-config>
```

to run cleaner, use `cleaner` command

```
bridge cleaner --config <path-to-config>
```

## Configuration

Config file can be found in `config` directory. There are 3 main components in configuration: listeners, database and cleaner

### Listeners
List all chains that bridge is listening to. Each name reflects to a specific function defined [here](https://github.com/axieinfinity/bridge-v2/blob/master/internal/init_listeners.go).

For example: `Ronin` reflects with function `InitRonin`

Therefore, do not change the name, otherwise the program cannot run properly.

#### 1. chainId
chain's identity (ronin: 0x7e4, ethereum: 0x1)

#### 2. rpcUrl 
RPC Url of chain which is used to query new events or submit transaction to.

#### 3. blockTime
time of new block is generated which is used to periodically to listen to new events from new block,

#### 4. safeBlockRange
safe block range which guarantees that reorg cannot happen.

#### 5. maxTasksQuery
maximum number of pending/processing tasks queried from database

#### 6. transactionCheckPeriod
period of checking whether a transaction is mined or not by querying its transaction's receipt. If receipt is found, it will try 3 more times to make sure the transaction is not replaced because of reorg.

#### 7. secret
stores private key of validator and relayer. These fields can be empty and passed via environment variables through 2 variables: `RONIN_VALIDATOR_KEY`, `RONIN_RELAYER_KEY` and Ethereum are: `ETHEREUM_VALIDATOR_KEY`, `ETHEREUM_RELAYER_KEY`
##### syntax: `<key>`
##### example: `xxxx4563e6591c1eba4b932a3513006cb5bcd1a6f69c32295dxxxx`
if KMS is used, set these environments
 - `RONIN_VALIDATOR_KMS_KEY_TOKEN_PATH`
 - `RONIN_VALIDATOR_KMS_SSL_CERT_PATH`
 - `RONIN_VALIDATOR_KMS_SERVER_ADDR`
 - `RONIN_VALIDATOR_KMS_SOURCE_ADDR`
 - `RONIN_VALIDATOR_KMS_SIGN_TIMEOUT`

replace `RONIN` with `ETHEREUM`, `VALIDATOR` with `RELAYER` when needed

format in config file
```json
"secret": {
  "validator": {
    "kmsConfig": {
      "keyTokenPath": "./key",
      "sslCertPath": "./ssl.crt",
      "kmsServerAddr": "127.0.0.1:1234",
      "kmsSourceAddr": ":5000",
      "signTimeout": 3000
    }
  },
  "relayer": {
    "plainPrivateKey": ""
  }
}
```

#### 8. fromHeight
Initially, bridge uses this property to load data from this block. After that, bridge will store latest processed block into `processed_block` table and use value from this table to continue.

#### 9. processWithinBlocks
This property guarantees that bridge does not process too far. Specifically, when `latestBlock - processWithinBlocks > fromHeight`, bridge `latestBlock - processWithinBlocks` instead of `fromHeight` to process.

#### 10. contracts
Stores a map (pair) of name and contract address, which can be used during processing tasks or jobs of a listener. For example, in `Ronin` listener, 2 contracts which are ronin gateway contract (at `Gateway`) and eth gateway contract (at `EthGateway`) are used:
```json
{
"Gateway": "0x03d1F13c7391F6B5A651143a31034cf728A93694",
"EthGateway": "0x3b6371EB912bFd5C0E249A16000ffbC6B881555A"
}
```

#### 11. subscriptions

Includes all subscriptions bridge is observing in a listener. Each subscription contains subscription's name and subscription's config.

- to: indicates the receiver/contract address that bridge uses as one of the condition to trigger a subscription
- type: there are 2 types, `0` is `transaction event` and `1` is `log's event`
- handler: when tx/event is matched given condition, subscription used `handler's abi` and `handler's name` to unpack tx/log's data to get encoded data.
- callbacks: list all callbacks function when data is decoded. This is a map (pair) where key is the listener's name and value is the function that is called in that listener. For example:

```json
{
  "callbacks": {
    "Ronin": "StoreMainchainWithdrawCallback"
  }
}
```

bridge will trigger function `StoreMainchainWithdrawCallback` in `RoninListener`

### Database

Database configuration is defined within key `database`. Basic properties include: host, port, user, password and dbName.

```json
{
  "database": {
    "host": "localhost",
    "user": "postgres",
    "password": "example",
    "dbName": "bridge",
    "port": 5432
  }
}
```

### Cleaner

cleaner configuration indicates time to clean data from database. It is defined within key `cleaner`, it is the map where key is the function name and the value is the config for this function.

```json
{
  "cleaner": {
    "ClearSuccessTasks": {
      "cron": "0 0 1 * *",
      "removeAfter": 7776000,
      "skipIfLessThan": 1000,
      "description": "triggered at 00:00 on day-of-month 1"
    }
  }
}
```

Above sample will trigger `ExecClearSuccessTasks` function and start cleaning data every first day of month.
- cron: cron job
- removeAfter: delete data that is created before this time, 7776000 (seconds) is 90 days
- skipIfLessThan: skip cleaning data if there is less than 1000 records left over
- description: describe the task.
