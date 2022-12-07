# Steampipe


[Steampipe](https://steampipe.io/) is an open-source project that enables querying their cloud services with SQL. In this example, the steampipe service is running locally and connected to local Postgress.  

PreReq
* have AWS CLI installed and profile set since steampipe will be referencing `~/.aws/config`  

Download CLI 
```
brew tap turbot/tap
brew install steampipe
steampipe -v
```

Install PlugIn
```
steampipe plugin install steampipe
steampipe plugin install aws
```

Start Steampipe service and expose local Postgres connections. 
```
steampipe service start --show-password
```

We can query Postgres directly, which invokes steampipe API to query other providers. 

! If there are issues with accessing the resultset, Postgress does not provide an entire error message; instead, it displays an empty resultSet. A further look in `~/.steampipe/logs` will show the error message similar to `Error: operation error S3: GetBucketLogging, https response error StatusCode: 403`
