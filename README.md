# IM-gossip-space
Instant messaging chat room written by Go and relevant techniques. Up to 100,000 online users supported.

## System design

- **Resources**     
  - Messages contained text, emojis, figures and voice messages.
![Resources](desgin1.png)

- **Database**          
  - Use WebSocket and HTTP to transfer data.
  - Application server read and write to database (MySQL).
![Database](design2.png)

## Installation & Run
```
# bash
```

## Test

## Deployment

## Contribution
Feel free to fork and make pull requests. Raise issues or comments are welcomed.

## References
- [Instant Messaging System](https://cloud.tencent.com/developer/article/1658166)
- [Gorilla WebSocket](https://github.com/gorilla/websocket)
- [go-redis](https://github.com/go-redis/redis)
- [Vue.js](https://vuejs.org/)
- [Ajax](https://developer.mozilla.org/en-US/docs/Web/Guide/AJAX)
- [Nginx](https://www.nginx.com/)

## Acknowledgement
: heart : Friendly Gophers.

## License
- [GNU General Public License v3.0](https://github.com/acse-sm321/IM-gossip-space/blob/main/LICENSE)