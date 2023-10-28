## Manifest.json

```bash
[
  {
    "Config": "d68fcc334f453a8c889c682226e6c6dc39694eaa4af54fdc8cc03bba03fdbb1c.json",
    "RepoTags": [
      "test:latest"
    ],
    "Layers": [
      "7b473dec0fa9e1cd2ffeb04ca39b125972ca0927000ccd033404674671768b8a/layer.tar",
      "247460e92ef67fdf643394f29fdfdfcec2fde609010ec63e3b7bee779e1a4846/layer.tar"
    ]
  }
]
```


```bash

@zakhaev26 ➜ /workspaces/devops/docker-training/03_02_before (main) $ tar -x -O -f ./hello-world.tar.gz |jq .
1
{
  "id": "247460e92ef67fdf643394f29fdfdfcec2fde609010ec63e3b7bee779e1a4846",
  "parent": "7b473dec0fa9e1cd2ffeb04ca39b125972ca0927000ccd033404674671768b8a",
  "created": "2022-11-29T21:16:36.165868669Z",
  "container_config": {
    "Hostname": "",
    "Domainname": "",
    "User": "",
    "AttachStdin": false,
    "AttachStdout": false,
    "AttachStderr": false,
    "Tty": false,
    "OpenStdin": false,
    "StdinOnce": false,
    "Env": [
      "PATH=/usr/local/sbin:/usr/local/bin:/usr/sbin:/usr/bin:/sbin:/bin"
    ],
    "Cmd": [
      "/bin/sh",
      "-c",
      "#(nop) COPY file:5c850056aa49a5a99cda22a7b820dc934348738ac919683db00729574acbeb11 in / "
    ],
    "Image": "sha256:46331d942d6350436f64e614d75725f6de3bb5c63e266e236e04389820a234c4",
    "Volumes": null,
    "WorkingDir": "",
    "Entrypoint": null,
    "OnBuild": null,
    "Labels": null
  },
  "docker_version": "20.10.3",
  "config": {
    "Hostname": "",
    "Domainname": "",
    "User": "",
    "AttachStdin": false,
    "AttachStdout": false,
    "AttachStderr": false,
    "Tty": false,
    "OpenStdin": false,
    "StdinOnce": false,
    "Env": [
      "PATH=/usr/local/sbin:/usr/local/bin:/usr/sbin:/usr/bin:/sbin:/bin"
    ],
    "Cmd": [
      "/hello"
    ],
    "Image": "sha256:46331d942d6350436f64e614d75725f6de3bb5c63e266e236e04389820a234c4",
    "Volumes": null,
    "WorkingDir": "",
    "Entrypoint": null,
    "OnBuild": null,
    "Labels": null
  },
  "architecture": "arm64",
  "variant": "v8",
  "os": "linux"
}

```