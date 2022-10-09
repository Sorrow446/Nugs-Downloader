You must use a token instead of an email and password if you auth via Apple or Google. They last for 600 mins. **Don't share them; they contain personal information.**

1. Start sniffer (I use Fiddler Classic).
2. Open https://www.nugs.net/stream/music/ in browser.
3. Find api/v1/me/subscriptions call in Fiddler.

![](https://i.imgur.com/XsRKCzY.png)
