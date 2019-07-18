# Googlon

## Pre requisites

Install go or docker

For go instllation follow the instructions from this link:
`https://golang.org/doc/install`

For docker installation follow the instructions from this link:
`https://docs.docker.com/install/`

### Running application

If go was installed, go to root folder and build the project with:

`go build`

Then you can run the application with:

`./goolon`

If docker was installed, run the following command:
`
`docker build -t googlon -f Dockerfile .`

then to run the container just run the following command:

`docker run -d -p 8080:8080 googlon`

#### Testing the application

For testing that container is running or application is running, run the following command

`curl -X GET http://localhost:8080/ping`

This will tell if the application is running or not.

To test some test cases, run the following command, after the Header, we can put the text that will be processed

```curl -X POST -H "Content-Type: text/plain" eorq gn mgnoqdjh odi nqeduhc cpgyx cjxdl uspxjgdlei jnqul hpngdocxw wisly ndepfswml oux umrocw dje exsij sxwc icy wydjx eq sinprhlwdm hwmjp cndlxuqhfg dimxwofheq phlyf qhiwfyudsp ey xwqmucp ec epqdonschm ipoufr syjqfr isxl mqgwclf ciex pysfieucx oylfmierg hygu ch orl rjoh peqfw ljsf nqc hwsqc yrshie mwrpcg pxiomfqwrd rxsfecil goi ylnxequr jrcdiye ojgelhisc fwerpnxogs gnl up yjgq hsfywx rnpjw lx qljs hjwsqel srhm mfginrqulc wpxgyh qugre mjrliogucs iupdjfnhm muoplcrfi wxmyhipou umwrc hdroml irplxyofsm emnd ncjysiwdh cgr qrj qpngufxjli ufidryq nsdohjim yqmgihus jimlxfsygc phdgqir yedhpws fqexsi xwule hgcmqix edmpwfnsuh fsxcerumow rijpsfgm pmldxecqf ypwocmxni ywcr" http://localhost:8080/process```