# AliyunDrive ECC Signature

> Welcome star, issue and pull requests.

## What can it do

1. Generate random `ECC` key-pair to create a new session.
2. Get the file’s download url without `invalid X-Device-Id`.

## How to use

1. Set the `deviceId` in line 7.
2. Set the `userId` in line 8.
3. Set the `authorization` in line 34.
4. (Optional) Set the private key in line 11,12. (Only if you want to use custom key-pair.)

## Some details

1. `deviceId` could be found in `DevTools` -> `Application` -> `Local Storage` -> `cna`.

2. `userId` could be found in `DevTools` -> `Application` -> `Local Storage` -> `token` -> `user_id`.

3. `authorization` could be found in `Network`, choose any request, `Request Headers` -> `authorization`.

4. `private key` could be found in `DevTools` -> `Application` -> `IndexedDB` -> `ALIYUN_DRIVE_CLIENT_SIGNATURE` -> `signature` -> `privateKey`.

   `private key` will be generated automatically, so there’s no need to specificate.

## Last

There’re some problems need to be solved.

* [ ] How to generate a `deviceId`?
* [ ] How to generate `authorization` from cookie or local storage?
* [ ] How often is `nonce` updated? (Could be frozen?)

## Test

You can use `ali_renew_test.py` to test `nonce` update.

If `nonce` is 0, it will create a session automatically.

If `nonce` is larger than 0, it will renew the session.

**Notice:** If it fails, try to increase the interval.

Now it could pass almost 1000 rounds :thumbsup:

## LICENSE

Anti-996-License.
