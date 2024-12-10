# Databases

## Start up

Run mysql server with: `mysql.server start`
Indexes are for faster searches, they can slow down updates because indexes need to be updated

create a new user that is not root for security

GRANT keyword can be used to give other users privileges.

Database drivers are used to convert Go commands to MySQL commands.

you can specify versions for packages you download by adding an @v(some-number) to the end of the package name.

//indirect on a line in a go.mod means that the package isn't directly imported by any code.

`go mod verify` double checks correctness of downloaded packages

You (or someone else in the future) can run go mod download to download the exact versions of all the packages that your project needs.
`go get -u <package-name> updates a package`
`go get -u <package-name>@none removes a package`
mysql parseTime flag converts SQL date and time fields to Go time.
sql.db returns a connection pool that automatically scales with usage.
