You must use a token instead of an email and password if you auth via Apple or Google. They last for 600 mins. **Don't share them; they contain personal information.**

1. Start sniffer (I use Fiddler Classic).
2. Open https://www.nugs.net/stream/music/ in browser.
3. Find api/v1/me/subscriptions call in Fiddler.

![](https://i.imgur.com/XsRKCzY.png)

The token can also be grabbed via Chrome's developer tools

1. Open developer tools in Chrome with `ctrl+shift+i`
2. Open https://play.nugs.net in that active tab
3. Nagivate to the `Application` tab at the top
4. Nagivate to `Storage > Session Storage > https://play.nugs.net/` in the sidebar
5. Below the Key/Value table, right click `access_token` and click `Copy value`

![image](https://github.com/Sorrow446/Nugs-Downloader/assets/21085666/834e6dcc-e795-4824-bf84-6810ddc011bd)

