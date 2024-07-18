# docker-ratify

A docker plugin wrapper for ratify

## Prerequisite

`docker-ratify` plugin requires [ratify](https://github.com/ratify-project/ratify/releases) being installed and available in `PATH`.

If you have already added `~/bin/` to `PATH`, you may run the following command to install `ratify v1.2.0` on Linux:
```bash
curl -L https://github.com/ratify-project/ratify/releases/download/v1.2.0/ratify_1.2.0_Linux_amd64.tar.gz | tar xvzC ~/bin/ ratify
```

## Installation

Run the following command to install on Linux:
```bash
mkdir -p ~/.docker/cli-plugins
curl -L https://github.com/ratify-project/docker-ratify/releases/download/v0.1.0/docker-ratify_0.1.0_linux_amd64.tar.gz | tar xvzC ~/.docker/cli-plugins/ docker-ratify
```

Help information can be reviewed by

```bash
docker help
```

## Example for pulling Images with Ratify

A config file is required for using ratify. Here we use an example config file:

```bash
cat > config.json <<EOF
{
    "executor": {},
    "store": {
        "version": "1.0.0",
        "plugins": [
            {
                "name": "oras",
                "cosignEnabled": true
            }
        ]
    },
    "policy": {
        "version": "1.0.0",
        "plugin": {
            "name": "regoPolicy",
            "policyPath": "",
            "policy": "package ratify.policy\ndefault valid := false\nvalid {\n not failed_verify(input)\n}\nfailed_verify(reports) {\n  [path, value] := walk(reports)\n  value == false\n  path[count(path) - 1] == \"isSuccess\"\n}"
        }
    },
    "verifier": {
        "version": "1.0.0",
        "plugins": [
            {
                "name": "notation",
                "artifactTypes": "application/vnd.cncf.notary.signature",
                "verificationCerts": [
                    "./root.crt"
                ],
                "trustPolicyDoc": {
                    "version": "1.0",
                    "trustPolicies": [
                        {
                            "name": "default",
                            "registryScopes": [
                                "*"
                            ],
                            "signatureVerification": {
                                "level": "strict"
                            },
                            "trustStores": [
                                "ca:certs"
                            ],
                            "trustedIdentities": [
                                "*"
                            ]
                        }
                    ]
                }
            }
        ]
    }
}
EOF
```

The above config uses `root.crt` as the trust anchor for ratification. Run the following command to download the root cert:

```bash
curl -Lo root.crt "http://www.microsoft.com/pkiops/certs/Microsoft%20Supply%20Chain%20RSA%20Root%20CA%202022.crt"
```

Now, we can pull an image with ratification:

```console
$ docker ratify pull -c config.json mcr.microsoft.com/oss/deislabs/ratify-base:v1.2.0
INFO[0000] Setting log level to info
INFO[0000] selected default auth provider: dockerConfig
INFO[0000] defaultPluginPath set to /home/demo/.ratify/plugins
INFO[0000] selected policy provider: regopolicy
INFO[0001] Resolve of the image completed successfully the digest is sha256:80dd14af4a7a676c962d0ca0f6e3b11a77b77826532fc863ea626317b158559c  component-type=executor go.version=go1.21.10
INFO[0002] Trust policy configuration: &{Name:default RegistryScopes:[*] SignatureVerification:{VerificationLevel:strict Override:map[]} TrustStores:[ca:certs] TrustedIdentities:[*]}  component-type=verifier go.version=go1.21.10
INFO[0002] 1 notation verification certificates loaded from path './root.crt'
WARN[0002] Invalid path '/home/demo/.ratify/ratify-certs/notation/truststore' skipped, error lstat /home/demo/.ratify/ratify-certs/notation/truststore: no such file or directory
INFO[0002] 0 notation verification certificates loaded from path '/home/demo/.ratify/ratify-certs/notation/truststore'
INFO[0002] Resolve of the image completed successfully the digest is sha256:ac86395350279f460c6bc08eb7875583c7365c423ebf9a7ac5a7a0f86f87924e  component-type=executor go.version=go1.21.10
INFO[0002] Resolve of the image completed successfully the digest is sha256:664dbce8187af59ee9a156b10f1ae66c0ab74b2d356bcce6ae3bfbffc90ddcf2  component-type=executor go.version=go1.21.10
INFO[0002] Resolve of the image completed successfully the digest is sha256:f281be7185446aa5bd346b3ee859061c95199830cfd42ce289cda2994205076f  component-type=executor go.version=go1.21.10
INFO[0002] Resolve of the image completed successfully the digest is sha256:6557162adb2a50ac98b52477ce8959858ba3bafbb94f346fcf764ac0c2aa8346  component-type=executor go.version=go1.21.10
INFO[0002] Resolve of the image completed successfully the digest is sha256:d93c4208945899f65c50f03024892a6f106344a0759eb1168e43d497d1582e40  component-type=executor go.version=go1.21.10
INFO[0002] Trust policy configuration: &{Name:default RegistryScopes:[*] SignatureVerification:{VerificationLevel:strict Override:map[]} TrustStores:[ca:certs] TrustedIdentities:[*]}  component-type=verifier go.version=go1.21.10
INFO[0002] 1 notation verification certificates loaded from path './root.crt'
WARN[0002] Invalid path '/home/demo/.ratify/ratify-certs/notation/truststore' skipped, error lstat /home/demo/.ratify/ratify-certs/notation/truststore: no such file or directory
INFO[0002] 0 notation verification certificates loaded from path '/home/demo/.ratify/ratify-certs/notation/truststore'
INFO[0002] Trust policy configuration: &{Name:default RegistryScopes:[*] SignatureVerification:{VerificationLevel:strict Override:map[]} TrustStores:[ca:certs] TrustedIdentities:[*]}  component-type=verifier go.version=go1.21.10
INFO[0002] 1 notation verification certificates loaded from path './root.crt'
WARN[0002] Invalid path '/home/demo/.ratify/ratify-certs/notation/truststore' skipped, error lstat /home/demo/.ratify/ratify-certs/notation/truststore: no such file or directory
INFO[0002] 0 notation verification certificates loaded from path '/home/demo/.ratify/ratify-certs/notation/truststore'
INFO[0002] Resolve of the image completed successfully the digest is sha256:9a330411e967bde20bd41702fa6cdb32ab27183f53cd3a17af3ebac41d3112b2  component-type=executor go.version=go1.21.10
mcr.microsoft.com/oss/deislabs/ratify-base@sha256:80dd14af4a7a676c962d0ca0f6e3b11a77b77826532fc863ea626317b158559c: Pulling from oss/deislabs/ratify-base
b2ce0e066077: Pull complete
e8d9a567199d: Pull complete
058cf3d8c2ba: Pull complete
b6824ed73363: Pull complete
7c12895b777b: Pull complete
33e068de2649: Pull complete
5664b15f108b: Pull complete
27be814a09eb: Pull complete
4aa0ea1413d3: Pull complete
da7816fa955e: Pull complete
9aee425378d2: Pull complete
0c4c7572df33: Pull complete
dff9a998dfb4: Pull complete
e03239cfd5d3: Pull complete
Digest: sha256:80dd14af4a7a676c962d0ca0f6e3b11a77b77826532fc863ea626317b158559c
Status: Downloaded newer image for mcr.microsoft.com/oss/deislabs/ratify-base@sha256:80dd14af4a7a676c962d0ca0f6e3b11a77b77826532fc863ea626317b158559c
mcr.microsoft.com/oss/deislabs/ratify-base@sha256:80dd14af4a7a676c962d0ca0f6e3b11a77b77826532fc863ea626317b158559c
```

Pulling an image without signatures will fail:

```console
$ docker ratify pull -c config.json mcr.microsoft.com/mcr/hello-world:latest
INFO[0000] Setting log level to info
INFO[0000] selected default auth provider: dockerConfig
INFO[0000] defaultPluginPath set to /home/demo/.ratify/plugins
INFO[0000] selected policy provider: regopolicy
INFO[0000] Resolve of the image completed successfully the digest is sha256:92c7f9c92844bbbb5d0a101b22f7c2a7949e40f8ea90c8b3bc396879d95e899a  component-type=executor go.version=go1.21.10
Error: no ratifications found
```
