job("BE rpcpublic docker image") {
    startOn {
        gitPush {
            branchFilter {
                +"refs/heads/main"
                +"refs/heads/prod"
            }
            pathFilter {
                +"be/**"
            }
        }
    }

    docker {
        env["SSH_PRV_KEY"] = Secrets("builder_ssh_prv_key")
        env["SSH_PUB_KEY"] = Secrets("builder_ssh_pub_key")

        beforeBuildScript {
            content = """
                export BRANCH=${'$'}(echo ${'$'}JB_SPACE_GIT_BRANCH | cut -d'/' -f 3)
                export VERSION=${'$'}(grep 'version: ' ./be/lib/config/default.yaml | grep -E '[0-9.]+' -o)
            """
        }

        build {
            file = "be/cmd/rpcpublic/Dockerfile"
            labels["vendor"] = "artDecoction"
            args["SSH_PRV_KEY"] = "\$SSH_PRV_KEY"
            args["SSH_PUB_KEY"] = "\$SSH_PUB_KEY"
        }

        push("artdecoction.registry.jetbrains.space/p/wt/services/be-rpcpublic") {
            tags("\$VERSION-\$BRANCH-build-\$JB_SPACE_EXECUTION_NUMBER")
        }
    }
}
