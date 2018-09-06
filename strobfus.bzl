def strobfus_impl(ctx):
    out = ctx.actions.declare_file(ctx.label.name + ".strobfus.go")
    args = ctx.actions.args()
    args.add_all("-filename", [ctx.file.src])
    args.add_all("-output", [out])
    ctx.actions.run(
        inputs = [ctx.file.src],
        outputs = [out],
        executable = ctx.executable._strobfus,
        arguments = [args],
    )
    return [
        DefaultInfo(
            files = depset([out]),
        ),
    ]

strobfus = rule(
    strobfus_impl,
    attrs = {
        "src": attr.label(
            allow_single_file = True,
            mandatory = True,
        ),
        "_strobfus": attr.label(
            cfg = "host",
            executable = True,
            default = "//:strobfus",
        ),
    }
)
