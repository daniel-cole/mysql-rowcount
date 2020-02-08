# Parallel MySQL Row Count

A simple tool for anyone who might need to count MySQL database rows accurately and with a high degree of parallelism.
Useful for counting rows in large databases and comparing the output.

You can specify the configuration file by setting the following environment variable: `ROWCOUNT_CONFIG=<config file>`. 
Defaults to `rowcount-config.yaml`

The output of the program is stored in the output file and is sorted by `database.tablename`.
This can be used to compare against executions on other databases. i.e. `diff <output1> <output2>`

:warning: If you run this with a high number of workers in will load up your database :warning:
:warning: Don't run on a live production database if you care about performance :warning:

Example configuration options:
```
user: "row_count"
password: "secret_password"
address: "10.12.55.4:3306"
net: "tcp"
nativepasswords: true
databasestoignore:
  - "percona"
  - "mysql"
  - "performance_schema"
  - "information_schema"
databasestoinclude: []
tablestoignore:
  - "mysql"
maxworkers: 1000
outputfile: "rowcount.txt"
```
[rowcount-config.yaml.example](rowcount-config.yaml.example)