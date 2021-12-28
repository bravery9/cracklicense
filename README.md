# cobalt strike license
> 参考大佬们的文章，完成一次CobaltStrike破解

### cobaltstrike4.3 cobaltstrike.auth 结构
``` plain text
00 00 00 00  | -889274157
00 00        | 之后的字节数（总长度 -6 ）
00 00 00 00  | 29999999
00 00 00 00  | watermark
00           | >43

00           | Authorization():var8 = 16
00*16        | len = var8 CS4.0 用来对dll解密的密钥 

00           | Authorization():var10 = 16
00*16        | len = var10 CS4.1 用来对dll解密的密钥 

00           | Authorization():var12 = 16
00*16        | len = var12 CS4.2 用来对dll解密的密钥 

00           | Authorization():var14 = 16
00*16        | len = var14 CS4.3 用来对dll解密的密钥 

00           | Authorization():var14 = 16
00*16        | len = var14 CS4.4 用来对dll解密的密钥 
```

### Cobalt Strike 4.X Sleeve解密密钥
``` go
var cs40 = []int8{27, -27, -66, 82, -58, 37, 92, 51, 85, -114, -118, 28, -74, 103, -53, 6}
var cs41 = []int8{-128, -29, 42, 116, 32, 96, -72, -124, 65, -101, -96, -63, 113, -55, -86, 118}
var cs42 = []int8{-78, 13, 72, 122, -35, -44, 113, 52, 24, -14, -43, -93, -82, 2, -89, -96}
var cs43 = []int8{58, 68, 37, 73, 15, 56, -102, -18, -61, 18, -67, -41, 88, -83, 43, -103}
var cs44 = []int8{94, -104, 25, 74, 1, -58, -76, -113, -91, -126, -90, -87, -4, -69, -110, -42}

// String
// 4.0 1be5be52c6255c33558e8a1cb667cb06
// 4.1 80e32a742060b884419ba0c171c9aa76
// 4.2 b20d487addd4713418f2d5a3ae02a7a0
// 4.3 3a4425490f389aeec312bdd758ad2b99
// 4.4 5e98194a01c6b48fa582a6a9fcbb92d6
```


### crack license 使用
1. 参考Moriarty的视频，先将cobaltstrike进行反编译
2. 编译
``` bash
go build cracklicense.go
```
3. 运行`./cracklicense`
4. 使用openssl生成非对称密钥
``` bash
openssl genrsa -out pk.pem 2048
openssl rsa -in pk.pem -pubout -out authkey.pub -outform DER
```
5.使用私钥对cobaltstrike.auth.unsign签名
``` bash
openssl rsautl -sign -inkey pk.pem -in cobaltstrike.auth.unsign -out cobaltstrike.auth
```
6. 计算authkey.pub的md5值
``` bash
openssl dgst -md5 authkey.pub
```
7.将authkey.pub放入cobaltstrike中
``` bash
cp authkey.pub src/resources/authkey.pub
```
8. 修改AuthCrypt.load()中的MD5值
9. `Build` -> `Build Artifacts` -> `Build`


### 参考
> 排序不分前后
1. https://eviladan0s.github.io/2021/08/30/cobalt-strike-crack/
2. https://ca3tie1.github.io/post/cobaltstrike40-wu-hook-man-li-cracked-license-si-lu/
3. https://githubplus.com/lovechoudoufu/cobaltstrike4.4_cdf
4. https://www.cnblogs.com/ssooking/p/12535998.html
5. https://mp.weixin.qq.com/s/Pneu8R0zoG0ONyFXF9VLpg
6. https://blog.pr0ph3t.com/posts/Cobalt-Strike4.3%E7%A0%B4%E8%A7%A3%E6%97%A5%E8%AE%B0/