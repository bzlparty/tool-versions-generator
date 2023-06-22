# `TOOL_VERSIONS` generator

This tool generates a `TOOL_VERSIONS` dict (in Starlark) to be used in a [`versions.bzl`](https://github.com/bazel-contrib/rules-template/blob/main/mylang/private/versions.bzl) file from release assets of a github repository.

Build:

```bash
go build github.com/bzlparty/tool-versions-generator/cmd/tvg
```

For [`oak`](https://github.com/thesephist/oak):

```bash
./tvg --repo thesephist/oak --platform linux,darwin 
```

Output:

```starlark
# This file was generated from https://github.com/thesephist/oak/releases
TOOL_VERSIONS = { 
    "0.1": { 
        "darwin": "sha384-IFGYZUufyvXU3EZvNteFaBBbQlAw1ZPJvDCCUViHiApuLDkyANXZvKJsWy8FM1Cb", 
        "linux": "sha384-SJmbqfFWdU7m4xS0Xq2kfjGEfq+41cCmvOfDgOpq4j9wLvlZYwQlhBCceBTjrCpX", 
    }, 
    "0.2": { 
        "darwin": "sha384-ifyKA6CsaHTkRfwQyjlTKPC6nmEETshBVYwXEme8olnzvmWUDePod1S/JjaXbxjz", 
        "linux": "sha384-7mX7LWx1clH1MUtSoLQIk5ncc2hAhixt3UIGSJyTGI9g0P7XDiy1E855SrikQTw0", 
    }, 
    "0.3": { 
        "darwin": "sha384-wT8x5iglKnGm5+nryH2SbfFGQh5yq7X5qjtkMXYmVtVAhTvPY09ptklW9tZqoa+f", 
        "linux": "sha384-jUe08kY/F7TQKBGzsGLNXTNAsGKq3zD0Glo7K7HuHzfU9cA6PjK8aOSl8TVqlToc", 
    }, 
}
```
