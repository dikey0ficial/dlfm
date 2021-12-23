# DLFM
*it was a fork of [DiscordLastfmScrobbler](https://github.com/edvardpotter/DiscordLastfmScrobbler)*
### DOWNLOAD [HERE](https://github.com/dikey0ficial/dlfm/releases "Releases")

This app will update your Discord "Listening to" status with when you are scrobbling something to Last.FM

![Profile screenshot](https://i.imgur.com/SbtvFJa.png "Profile screenshot") *profile screenshot*


![Userlist screenshot](https://i.imgur.com/x5mWIXR.png "Screenshot from list of users") *screenshot from list of users*



To run it, you need to edit config and complete three things:
## **1. Last.FM API Key**

Go to the [Last.FM API Page](https://www.last.fm/api/account/create) and sign in with your existing Last.FM username and password. It should bring you to the **Create API account** page and ask you for a few things.

It doesn't really matter what you put in most of the fields, but it should probably look something like this:

![LastFM Create API Account Screenshot](https://i.imgur.com/VQYa8nr.png?1)

After clicking Submit you should get a confirmation page with two items: *API Key* and *Shared Secret*. The API Key is the only one you need for this, but I recommend you save both for future use just in case, as they don't actually provide a way to retrieve these later.

![LastFM API Account Created Screenshot](https://i.imgur.com/1Qb7LeO.png "don't ask why names aren't same")

Copy and paste the API Key value into the config file in the `api_key = xxx` line

Then, feel field `username` with your last.fm username

Also you should feel field `check_interval` with integer number of seconds script should check new tracks. Recomended values: 1-5 seconds (if you are listening to grindcore, write 0)

## **2. Discord User Token**

For this one you'll need to use the desktop app - it won't work on mobile or in web.

If you are using the desktop app:

- Press **Ctrl+R** (or **Cmd+R** on Mac)
- Click the "*Application*" tab
- Click and expand the "*Local Storage*" section
- Click on the only entry in this section, "*https://discord.com/*"
- Press **Ctrl+Shift+I** (or **Cmd+Shift+I** on Mac)
- Wait few seconds
- Right click -> Edit Value in the field to the right of "*token*"
- Copy and paste the token value into the config file on the `token = xxx` line and remove the quotation marks from it.

![Desktop Token](https://i.imgur.com/sPs0New.png)

## **3. App settings**

You can change there *title* and *endless* options.

If you want to change title of status (for example, "Listening to Deezer" or "Listening to Youtube Music"), change "title" value
Example:

![Example of changed title](https://i.imgur.com/9OShK3U.png)
In title you can use tags, which will be replaced with other value.

List of tags:

| Tag         |             Replaced with |
|:------------|--------------------------:|
| {{artist}}  | Track's artist(s) name(s) |
| {{album}}   | Track's album name        |
| {{name}}    | Track name                |

Also if you want to ignore not critic errors, change "*endless_mode*" to "true", else feel it with "false"

## When you're done

Save your config file as "*config.ini*" and it should look something like this:

![Finished Config File](https://i.imgur.com/vP1WXeD.png)

Now just run the executable. It should connect to Discord and immediately start setting your "*Playing*" status to whatever you're listening to on Last.FM

If it's working, it will look like this:

![Running Executable](https://i.imgur.com/S2LehkW.png)