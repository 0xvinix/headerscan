<h1 align="center">
  <img src="https://user-images.githubusercontent.com/51921808/181158239-daf26cb4-15c5-4e2a-8b45-d97d75003850.png" width="500px"></a>
  <br>
</h1>
<p align="center">
  <a href="#installation">🔧 Installation</a> |
  <a href="#usage">⚙️ Usage</a> | 
  <a href="#features">✨ Features</a>
</p>

[🕷️] Make sure a website has all basic security headers.

 ## Installation 

```sh
git clone https://github.com/0xvinix/headerscan
cd headerscan/
go build
./headerscan
```

  <a href="https://asciinema.org/a/wTXH0nzantiWIyEhUUh4w8xWN" target="_blank"><img width="50%" src="https://asciinema.org/a/wTXH0nzantiWIyEhUUh4w8xWN.svg" /></a>
  <br>

## Usage

```sh
./headerscan https://example.com
```

  <a href="https://asciinema.org/a/8BwXbHS6JfTT2t8q3C3K9P8uy" target="_blank"><img width="50%" src="https://asciinema.org/a/8BwXbHS6JfTT2t8q3C3K9P8uy.svg" /></a>
  <br>

## Features

- [✔️] Check for web server/technology exposure
- [✔️] Check X-Frame-Options
- [✔️] Check X-XSS-Protection
- [✔️] Check Strict-Transport-Security 
- [✔️] Check X-Content-Type-Options
- [✔️] Check Referrer-Policy
- [✔️] Check Content-Security-Policy
