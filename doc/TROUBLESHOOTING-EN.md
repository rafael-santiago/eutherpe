![troubleshooting-glyph](figures/troubleshooting-glyph.png)
# Troubleshooting

**Abstract**: Guess what? Here you find topics related to troubleshooting.

## Topics

- [I have enabled password-based login but I forgot the password!](#i-have-enabled-password-based-login-but-i-forgot-the-password)
- [I have changed the Wi-Fi password but now I cannot access Eutherpe via web](#i-have-changed-the-wi-fi-password-but-now-i-cannot-access-eutherpe-via-web)


### I have enabled password-based login but I forgot the password!

In order to solve this problem I will log in the computer that is running `Eutherpe` as 
`root` user.

If you are logged in this machine as a normal user, you can turn into `root` in several ways:

```
$ sudo su
```

or:

```
$ sudo -i
```

or still if you are not logged in, do the log in as `root`.

Once with a `root` prompt, edit the file `/etc/eutherpe/player.cache`, doing the following:

- Locate in the file the configuration `"Authenticated":true` and change it to `"Authenticated":false`.
- Still in the file, locate the configuration `"HashKey":"<a bunch of symbols>"` and replace it to
`"HashKey":""`.
- Save the changes done in `/etc/eutherpe/player.cache`.
- From the `root` `prompt` execute `systemctl restart eutherpe`.
- `Log off` from `root` session (jot down `exit` in the `root` `prompt`).
- Now browse `Eutherpe` from your `web browser`, it will not ask for password anymore.
- If you would like to reactivate the authentication and/or reset the password, the current password is from now on `music`.


[`Back`](#topics)

### I have changed the Wi-Fi password but now I cannot access Eutherpe via web

You need to log in the `Eutherpe` server as `root`.

If you are logged in this machine as a normal user, you can turn into `root` in several ways:

```
$ sudo su
```

or:

```
$ sudo -i
```

or still if you are not logged in, do the log in as `root`.

Once it done, edit the file `/etc/eutherpe/player.cache` by doing the following:

- Locate the configuration `"ESSID":"<name of your Wi-Fi network>"` and replace it to `"ESSID":""`.
- Save the changes done in `/etc/eutherpe/player.cache`.
- From the `root` `prompt` execute `systemctl restart eutherpe`.
- Browse `http://<eutherpe server ip address>:8080/eutherpe`. If you have chosen `https` instead
of `http`, use `https://<eutherpe server ip address>:8080/eutherpe`. If you have changed the
default port, replace `:8080` to `:<chosen port>`.
- Once `Eutherpe` accessed by `web` you can update all Wi-Fi credentials info, besides reactivating
it. After with all info updated, do the reboot and the access to your Wi-Fi will be restablished.

[`Back`](#topics)
