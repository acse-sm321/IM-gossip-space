![example workflow](https://github.com/acse-sm321/IM-gossip-space/actions/workflows/test.yml/badge.svg)
# IM-gossip-space
Instant messaging chat room written by Go and relevant techniques.

## Live Demo
![]()

## System Design
- **Resources**     
  - Messages contained text, emojis, figures and voice messages.
![Resources](desgin1.png)

- **Database**          
  - Use WebSocket and HTTP to transfer data.
  - Application server read and write to database (MySQL).
![Database](design2.png)

- **Other design**    

This project developed as MVC model.

## Installation & Run
```
$ git clone https://github.com/acse-sm321/IM-gossip-space.git
$ cd download_folder/
$ go build main -o appserver
$ ./appserver
$ curl "http://localhost:8080/"
```

## Test
Run the pre-defined Go unit test.
```
$ go test ./
```
use cURL to check the HTTP responses.

## Deployment
TODO: How to deploy on a Linux server, set up Nginx reverse proxy
```
$ vim build.sh
rm -rf ./release
mkdir  release
go build -o chat
chmod +x ./chat
cp chat ./release/
cp favicon.ico ./release/
cp -arf ./asset ./release/
cp -arf ./view ./release/
```
Notice: Do not deploy this project to your server, take this project as an reference of chat room application instead.

## Contribution
Please fork and make pull requests. Raise issues or comments are welcomed.

## References
- [Instant Messaging System](https://cloud.tencent.com/developer/article/1658166)
- [Gorilla WebSocket](https://github.com/gorilla/websocket)
- [go-redis](https://github.com/go-redis/redis)
- [gorm](https://gorm.io/index.html)
- [Vue.js](https://vuejs.org/)
- [Ajax](https://developer.mozilla.org/en-US/docs/Web/Guide/AJAX)
- [Nginx](https://www.nginx.com/)
- [oss]()

## Update & Troubleshooting
```
#  Update details
- 1. xorm is not maintained anymore, will use gorm instead.
- 2. try to use go-redis to improve the performance
- 3. try to deloy using Nginx and see if it works

#  Trouble shooting
- 1.
```

## Acknowledgement
:heartpulse: Friendly Gophers.



## License & Copyright
- [GNU General Public License v3.0](https://github.com/acse-sm321/IM-gossip-space/blob/main/LICENSE)