# Demo App for proving Shell Canaries on a distroless/scratch image 

## Overview 
Distroless or Scratch containers are useful for reducing the amount of attack surface for an application. Typically in a distroless container there is only your application, the runtime that interacts with the kernel and some support files ( like /etc/passwd, timezone files etc ). 

Not having a shell for instance makes it a little harder for an attacker to execute code. But what if instead of a shell you packaged a canary. A shell binary that alerted you that an attacker is trying to run commands in your container.

This repo is the POC!GTFO for that idea. A simple exploitable server with canaries built in.

## Acknowledgements

I'd like to acknowledge Geert Baeke fro this [article](https://blog.baeke.info/2021/03/28/distroless-or-scratch-for-go-apps/) from which the way to build a scratch container is mercilessly stolen.

I'd also like to acknowledge [Thinkst](https://blog.baeke.info/2021/03/28/distroless-or-scratch-for-go-apps/) for providing the inspiration for this idea and the canary token service that makes this work.

## Disclaimers

This is not good code. This is a quick and dirty POC. It should not be used anywhere near production and it's really not been tested. Also this demo was designed to work on Linux, it's probably not going to work well on Windows or MacOS.  

IF by any chance you happen to be a security researcher who has come to disclose a serious vulnerability in the server app in this codebase, please be aware this is a deliberately vulnerable application. 

## The important bit. How this all works.

To get this working you will need [golang](https://go.dev/doc/install), [docker](https://docs.docker.com/engine/install/) and probably curl.

You will also need to create a web bug at [canarytokens.org](https://canarytokens.org/nest/) as a destination to send your alerts to.

To create the image run `./build.sh`. 

To run the image you need two paramters to pass in.
<table>
  <tr><td><b>Parameter</b></td><td><b>Description</b></td></tr>
  <tr><td>TOKEN</td><td>This is you token identifier which is everything in your web bug token after `https://canarytokens.com/`. often it's something like `terms/<somerandomstuff>/index.jsp` </td></tr>
  <tr><td>IDENT</td><td>This is some kind of unique identifier so you know what container it's come from. If you were doing this in an actual deployment it might be the Pod Name.</td></tr>
</table>

To run the server you execute the following command:<br>
`sudo docker run -p 127.0.0.1:8080:8080/tcp -e TOKEN=<TOKEN VALUE> -e IDENT=<CONTAINER IDENTIFIER> canary-test:latest`

Once running to test it up and working the following query should return a payload with the current time.<br>
`curl 'http://127.0.0.1:8080/tool/time'`

To trigger the canary you can run:<br>
`'http://127.0.0.1:8080/tool/..%2f..%2f..%2fbin%2fbash'`

Or another way:<br>
`sudo docker exec -it <container id> /bin/dash`

Have fun.







