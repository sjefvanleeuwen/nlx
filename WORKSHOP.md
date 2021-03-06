# NLX Workshop

During this workshop we'll setup an NLX outway and request a service that is available on the NLX demo network.

To run the outway, we need a certificate for TLS communications with other NLX services.

Lets start by creating a directory structure for this workshop.

```bash
mkdir -p $HOME/nlx-workshop/certs
```

Change directory to the nlx-workshop folder.

```bash
cd $HOME/nlx-workshop
```
We need to create a key and certificate signing request (csr) for our organization.

```bash
openssl req -utf8 -nodes -sha256 -newkey rsa:4096 -keyout certs/org.key -out certs/org.csr
```

You will need to provide some bogus information to the questions asked. One field is of importance and is used in logs; the "Organization Name" field. Please use your name or make up a random name. e.g. "Louis-van-Gaal". Make sure this is something unique so you can easily find your own logs later.

View the contents of `certs/org.csr`:

```bash
cat certs/org.csr
```

Go to https://certportal.demo.nlx.io and paste the contents of the ` you've just created. Click "Request Certificate" and copy the resulting certificate text.

Create a file, `certs/org.crt`, and paste the certificate. Or click download and move the `certificate.crt` to `certs/org.crt`.

__In a production environment one would not be able to create their own certificates, but instead retrieves them from the central authority, which asserts the Organization Name field is properly set. The certportal is only available in test and demo environments.__

We now have a certificate by which we can identify ourselves to other NLX services, but we still need the rootCA certificate, which we use to verify other NLX services.

```bash
wget https://certportal.demo.nlx.io/root.crt -O certs/root.crt
```

We're now ready to start the outway. We do this by running the outway in docker. First, download the latest version of the container image.

```bash
docker pull nlxio/outway:latest
```

Next, we want docker will to start a container based on this image.

```bash
docker run \
    --tty --interactive \
    --volume=$HOME/nlx-workshop/certs:/certs \
    --publish=12018:12018 \
    nlxio/outway:latest \
    /usr/local/bin/nlx-outway \
    --log-type=development \
    --log-level=debug \
    --directory-address=directory.demo.nlx.io:1984 \
    --tls-nlx-root-cert=/certs/root.crt \
    --tls-org-cert=/certs/org.crt \
    --tls-org-key=/certs/org.key \
    --disable-logdb
```

We give docker several arguments:

- `-log-*` setup developmeng logs so we can see what's going on
- `--tty --interactive` tells docker to show us the output of the process.
- `--volume` tells docker to make the `certs` folder, where we just put the certificates, available inside the container.
- `--publish` opens port 12018 for traffic. We will send requests to this port.
- `nlxio/outway:latest` is the name of our docker image (`nlxio/outway`) as stored in the docker registry and the version we want to use (`latest`). The `--` tells docker that all arguments after this one are meant for the outway process, not for docker itself.
- `/usr/local/bin/nlx-outway` is the binary in the container that we want to execute.

All arguments after the image and executable name are passed to the process (nlx-outway) itself..

- `--directory-address` is set to point the outway to the demo environment directory service. This means we can use this outway to connect to services in the demo environment.
- `--tls-*` tells outway where to find the root, org-cert and org-key files.
- `--disable-logdb` tells outway to not write request log records, because we have not set up a database for it. Logs are still written at services we will be using.

Make a request through the outway.

```bash
curl "localhost:12018/demo-organization/demo-api/"
```
