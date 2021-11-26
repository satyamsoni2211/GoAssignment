# RPC sample bot for validating blockchain adress (ethereum)

This is a simulation bot with very basic functionality consisting of greeting and validation of block adress.

To run the code, you will need to install certain dependencies. Go to the directory where code is cloned and run 

```bash
go mod download
```

Once all the modules are downloaded, you will have to start [RPC](https://pkg.go.dev/net/rpc) server. Server will be using TCP protocol and wwill run default on port **1234**.

```bash
go run main.go
```

Once you run the server, if everything works fine you can run your client script in `client` folder.

To run client script, go to `client` folder and run the below command:

```bash
go run client.go
```

you shall see output something like below:

```sh
╭─satyamsoni@Vedikas-MacBook-Air ~/Documents/GoAssignment/rpc_bot/client ‹master*› 
╰─$ go run client.go
< Hi there from RPC bot, you can greet me or ask me to validate eth adress, /q to quit
> 
```

Enjoy playing with the bot ...

You can follow
me [![GitHub](https://img.shields.io/badge/github-%23121011.svg?style=for-the-badge&logo=github&logoColor=white)](https://github.com/satyamsoni2211/) [![Twitter](https://img.shields.io/twitter/url/https/twitter.com/cloudposse.svg?style=social&label=Follow%20%40satyam_soni1306)](https://twitter.com/satyam_soni1306) [![LinkedIn](https://img.shields.io/badge/linkedin-%230077B5.svg?style=for-the-badge&logo=linkedin&logoColor=white)](https://www.linkedin.com/in/satyam-soni-ba648192/)