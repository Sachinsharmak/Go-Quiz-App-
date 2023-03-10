# Go Quiz App

# Description:

Ths presents a Simple Go quizzing App.

# Technology Used:
	Go
# Setting Up:
# Go
* ### Install Go
#### For Ubuntu: [Look here](https://go.dev/dl/go1.20.2.darwin-arm64.pkg)
#### For MacOS: [Look here](https://go.dev/dl/go1.20.2.linux-amd64.tar.gz)

>Install Go by running the following commands on your terminal:
>
1.Remove any previous Go installation by deleting the /usr/local/go folder (if it exists), then extract the archive you just downloaded into /usr/local,   creating a fresh Go tree in /usr/local/go:
>
`$ rm -rf /usr/local/go && tar -C /usr/local -xzf go1.20.2.linux-amd64.tar.gz`
>(You may need to run the command as root or through sudo).

2.Do not untar the archive into an existing /usr/local/go tree. This is known to produce broken Go installations.

3.Add `/usr/local/go/bin` to the PATH environment variable.
  You can do this by adding the following line to your $HOME/.profile or /etc/profile (for a system-wide installation):

 ```export PATH=$PATH:/usr/local/go/bin```

Note: Changes made to a profile file may not apply until the next time you log into your computer. To apply the changes immediately, just run the shell commands directly or execute them from the profile using a command such as source $HOME/.profile.

Verify that you've installed Go by opening a command prompt and typing the following command:
$ go version

Confirm that the command prints the installed version of Go.
## TO Run:
#### After Set-Up ,run:
* In the [Quiz_App Directory](https://github.com/Sachinsharmak/Go-Quiz-App-/tree/main/Quiz_App)
>
`go run main.go
`
