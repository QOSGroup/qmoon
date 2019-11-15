module.exports = {
    title: "Qmoon Data Server",
    description: "Documentation for the Qmoon.",
    dest: "./dist",
    base: "/qmoon/",
    markdown: {
        lineNumbers: true
    },
    themeConfig: {
        lastUpdated: "Last Updated",
        nav: [{ text: "Back to Qmoon", link: "https://qmoon.network" }],
        sidebar: [
            {
                title: "Introduction",
                collapsable: false,
                children: [
                    "/",
                ]
            },
            {
                title: "Started",
                collapsable: false,
                children: [
                    "/started/installation.md",
                ]
            },
            {
                title: "Base API",
                collapsable: false,
                children: [
                    "/apis/base.md",
                    "/apis/blockchain.md",
                    "/apis/block.md",
                    "/apis/txs.md",
                    "/apis/tx.md",
                    "/apis/sequence.md",
                    "/apis/validators.md",
                    "/apis/validator.md",
                    "/apis/peers.md",
                    "/apis/account_query.md",
                    "/apis/account_tx.md",
                ]
            },
            {
                title: "Plugin API",
                collapsable: false,
                children: [
                    "/plugins/transfer.md",
                    "/plugins/atm.md",
                ]
            },
            {
                title: "Developer",
                collapsable: false,
                children: [
                    "/developer/developer.md",
                ]
            }
        ]
    }
}
