[![Promo](https://brightdata.com/static/github_promo_15.png?md5=105367-daeb786e)](https://brightdata.com/?promo=github15) 

# proxy_scrap_go
basic  API sergo programs demonstrating using web scraping with Colly, GOquery, Selenium using proxy servers

This repository contains two branches:
- The `basic` branch contains the base code that will be modified in the article [Go Proxy Servers](https://brightdata.com/blog/how-tos/go-proxy-servers).
- The `main` branch is the result of the article tutorial.

This project demonstrates setting up proxy servers in [Go for web scraping](https://brightdata.com/blog/how-tos/go-proxy-servers), Proxies protect your digital identity during web scraping by using their IP address, circumventing IP bans and geoblocking.

## Installation
To use this project, you need to have Go installed on your machine. You can download and install Go from the official website: [https://golang.org/](https://golang.org/)

## Getting Started
1. Clone the repository:

    ```shell
    git clone https://github.com/shacharbd/proxy-scrape-go.git
    ```
2. Navigate to the project directory:

    ```shell
    cd proxy-scrape-go
    ```
3. Install the dependencies:

    ```shell
    go mod download
    ```
4. Run the project:

    ```shell
    go run main.go
    ```
## Usage
This project demonstrates how to use web scraping with proxy servers in Go. It utilizes the following libraries:
- [Colly](https://github.com/gocolly/colly) - a scraping framework for Go
- [Goquery](https://github.com/PuerkitoBio/goquery) - a library for parsing HTML in Go
- [Selenium](https://github.com/tebeka/selenium) - a browser automation tool

To run the project, make sure you have a valid proxy server. You can obtain proxy server details from a provider like [Bright Data](https://brightdata.com/). Once you have the proxy server details, update the `main.go` file with the appropriate proxy configuration.

## Contributing
Contributions are welcome! If you find any issues or have suggestions for improvement, please open an issue or submit a pull request.

## License
This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for more information.
