# Crusher - hold still, I'm scanning you for signs of life

Crusher is a monitor process that can verify the health of a container.

Crusher can either perform a port check:

```
crusher --port-check=8080
```

or a URL check:

```
crusher --url-check=http://127.0.0.1:8080/
```

## Getting Crusher

Crusher is available as part of:

docker:///onsi/away-team

(see the [away-team](https://github.com/1701-diego/away-team) repo)

Or as a stand-alone tarball at:

http://onsi-public.s3.amazonaws.com/crusher.tar.gz

Examples of using Crusher are in [Laforge](https://github.com/1701-diego/laforge)