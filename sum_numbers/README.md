# Using golang to sum first 100 natural numbers

This programme uses `goroutine` and `channel` for summing first 100 natural numbers **series by series**.

```
# first series 10 20 30 40 50 60 70 80 90 100
# seconds series 11 21 31 41 51 61 71 81 91
... so on
```

Function `sumNumbers` is responsible to add series and return over **write** only channel which is then consumed by main function which further adds up the sums.

To run the script, simple run the below command:

```sh
go run main.go
```

This programme does not require any dependencies to be installed except go compiler ;)

You can follow
me [![GitHub](https://img.shields.io/badge/github-%23121011.svg?style=for-the-badge&logo=github&logoColor=white)](https://github.com/satyamsoni2211/) [![Twitter](https://img.shields.io/twitter/url/https/twitter.com/cloudposse.svg?style=social&label=Follow%20%40satyam_soni1306)](https://twitter.com/satyam_soni1306) [![LinkedIn](https://img.shields.io/badge/linkedin-%230077B5.svg?style=for-the-badge&logo=linkedin&logoColor=white)](https://www.linkedin.com/in/satyam-soni-ba648192/)