<h1 align="center">Welcome to cli-password-generator-go 👋</h1>
<p>
</p>

> This is a simple command-line password generator written in Go. It allows you to generate secure passwords with various options such as length and the inclusion of special characters. The generated passwords can be customized to meet your security needs. Additionally, the program provides the option to save generated passwords for different websites or services in separate files, making it a practical tool for password management.

## Features

Generate random passwords with customizable length and special character inclusion.
Verify password length to meet security requirements.
Save generated passwords in separate files for different websites/services.
User-friendly command-line interface for ease of use.
Option to generate multiple passwords in succession.
Option to save generated passwords in a file and associate them with specific websites/services.

## Usage

Run the program with desired password length and special character options.
Verify password length and adjust it as needed.
Save generated passwords for different websites if desired.
Generate additional passwords or exit the program.
This Password Generator CLI is a helpful tool for creating and managing strong, unique passwords for your online accounts.

## Install

```sh
git clone git@github.com:JayViannay/cli-password-generator-go.git
cd cli-password-generator-go
```

## Usage

```sh
go run generatepwd.go -length=10 -special -save // generate and save a password with special char and length of 10
go run generatepwd.go // generate a password without special char and of a lenght of 10, terminal will ask if you want to save the generated password
```

## Author

👤 **JayViannay**


## Show your support

Give a ⭐️ if this project helped you!

***
_This README was generated with ❤️ by [readme-md-generator](https://github.com/kefranabg/readme-md-generator)_
