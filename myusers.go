package main

import (
	"database/sql"
	"errors"
	"flag"
	"fmt"
	"strconv"
	"strings"

	_ "github.com/go-sql-driver/mysql"
)

const appName = "My users go"

var version = "0.0.0"

// ParseMyVersion parses a major and minor numbers from mysql version string.
func ParseMyVersion(v string) (uint, uint, error) {
	versionInfo := strings.Split(v, ".")
	var e error
	if len(version) >= 2 {
		if major, err := strconv.ParseUint(versionInfo[0], 10, 64); err == nil {
			if minor, err := strconv.ParseUint(versionInfo[1], 10, 64); err == nil {
				return uint(major), uint(minor), nil
			} else {
				e = err
			}
		} else {
			e = err
		}
	} else {
		e = errors.New("Check MySql version string. It should contain a major and minor numbers")
	}
	return 0, 0, e
}

func must(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	var myuser string
	var mypass string
	var myserver string
	var myport string
	var mydb string
	var showVersion bool
	flag.StringVar(&myuser, "u", "root", "user name")
	flag.StringVar(&mypass, "p", "12345", "user password")
	flag.StringVar(&myserver, "s", "127.0.0.1", "db server ip")
	flag.StringVar(&myport, "-port", "3306", "db server port")
	flag.StringVar(&mydb, "d", "mysql", "db name with users")
	flag.BoolVar(&showVersion, "v", false, "show version information and exit")
	flag.Parse()
	if showVersion {
		fmt.Println(appName, version)
		return
	}
	connFmt := "%s:%s@tcp(%s:%s)/%s"
	myVersionQ := "SELECT SUBSTRING_INDEX(VERSION(), '-', 1);"
	usersPasswordQ := "SELECT `User`, `Host`, `Password` FROM `user`;"
	usersAuthStringQ := "SELECT `User`, `Host`, `Authentication_String` as `Password` FROM `user`;"
	createUserByPasswordFmt := "CREATE USER `%s`@`%s` IDENTIFIED BY PASSWORD '%s';\n"
	createUserWithAuthPluginFmt := "CREATE USER `%s`@`%s` IDENTIFIED WITH MYSQL_NATIVE_PASSWORD AS '%s';\n"
	showGrantsFmt := "SHOW GRANTS FOR `%s`@`%s`;"
	outputGrantsFmt := "%s;\n"
	flushPrivilegesQ := "FLUSH PRIVILEGES;"
	cannotDetermineVersionError := errors.New("Cannot determine version of MySql")
	connStr := fmt.Sprintf(connFmt, myuser, mypass, myserver, myport, mydb)
	db, err := sql.Open("mysql", connStr)
	must(err)
	defer db.Close()
	err = db.Ping()
	must(err)
	myVersion, err := db.Query(myVersionQ)
	must(err)
	defer myVersion.Close()
	var majorMinor string
	if myVersion.Next() {
		if err := myVersion.Scan(&majorMinor); err != nil {
			panic(err)
		}
	} else {
		panic(cannotDetermineVersionError)
	}
	major, minor, err := ParseMyVersion(majorMinor)
	must(err)
	usersQ := usersPasswordQ
	createUserFmt := createUserByPasswordFmt
	if major > 5 || (major == 5 && minor > 6) {
		usersQ = usersAuthStringQ
		createUserFmt = createUserWithAuthPluginFmt
	}
	users, err := db.Query(usersQ)
	must(err)
	defer users.Close()
	for users.Next() {
		var name string
		var host string
		var password string
		if err := users.Scan(&name, &host, &password); err != nil {
			panic(err)
		}
		fmt.Printf(createUserFmt, name, host, password)
		grantsQ := fmt.Sprintf(showGrantsFmt, name, host)
		grants, err := db.Query(grantsQ)
		must(err)
		defer grants.Close()
		for grants.Next() {
			var grant string
			if err := grants.Scan(&grant); err != nil {
				panic(err)
			}
			fmt.Printf(outputGrantsFmt, grant)
		}
	}
	fmt.Println(flushPrivilegesQ)
}
