# Copyright 2018 The Kubernetes Authors.
#
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#
#     http://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.

# Genrule wrapper around the go-bindata utility.
# forked from https://github.com/kubernetes/kubernetes/blob/master/build/bindata.bzl
# IMPORTANT: Any changes to this rule may also require changes to hack/generate.sh.
# -nometadata -mode=0666 -pkg=$GOPACKAGE -o=images_node_sources.go -ignore=(\.*README\.md)|(\.*BUILD\.bazel) -prefix=./../../../ ./../../../images/node/...
def go_bindata(
        name,
        srcs,
        outs,
        prefix,
        mode = "0666",
        include_metadata = True,
        pkg = "generated",
        ignores = [],
        **kw):
    args = []
    if not include_metadata:
        args.append("-nometadata")
    if mode:
        args.append("-mode=%s" % mode)
    for ignore in ignores:
        args.extend(["-ignore", "'%s'" % ignore])
    if prefix:
        args.append("-prefix=%s" % prefix)

    native.genrule(
        name = name,
        srcs = srcs,
        outs = outs,
        cmd = """
    $(location //vendor/github.com/jteeuwen/go-bindata/go-bindata:go-bindata) \
      -o "$@" -pkg %s %s $(SRCS)
    """ % (pkg, " ".join(args)),
        tools = [
            "//vendor/github.com/jteeuwen/go-bindata/go-bindata",
        ],
        **kw
    )
