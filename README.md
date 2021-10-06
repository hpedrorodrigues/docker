# Tiny Docker implementation

An extremely simple implementation of "docker run" command to help me learn more
about docker internals.

This project basically can pull an image from Dockerhub and execute commands in it.
During this process, it uses [chroot](https://en.wikipedia.org/wiki/Chroot),
[kernel namespaces](https://en.wikipedia.org/wiki/Linux_namespaces) and the
[Docker registry API](https://docs.docker.com/registry/spec/api/).

This project was inspired by
[challenge](https://codecrafters.io/challenges/docker) of CodeCrafters.

And based on Docker registry API docs:
- [Token Authentication](https://docs.docker.com/registry/spec/auth/token/)
- [Pulling an image manifest](https://docs.docker.com/registry/spec/api/#pulling-an-image-manifest)
- [Pulling a layer](https://docs.docker.com/registry/spec/api/#pulling-a-layer)

### Running this project

You can run this project using the [run.sh](./run.sh) script

e.g.
```bash
./run.sh run ubuntu:latest ls -la .
```

Or you can use the [Dockerfile](./Dockerfile) available in the root directory

e.g.
```bash
docker build -t tiny-docker . \
  && docker run --cap-add='SYS_ADMIN' tiny-docker run ubuntu:latest ls -la .
```

> Note: The `--cap-add='SYS_ADMIN' flag is required to create
> [PID Namespaces](https://man7.org/linux/man-pages/man7/pid_namespaces.7.html).
