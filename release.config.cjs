module.exports = {
  plugins: [
    "@semantic-release/commit-analyzer",
    "@semantic-release/release-notes-generator",
    [
      "@semantic-release/changelog",
      {
        changelogFile: "CHANGELOG.md",
      },
    ],
    [
      "@semantic-release/github",
      {
        assets: [
          {
            path: "packages/gopeg",
            name: "gopeg_${nextRelease.version}_linux_x64",
          },
          {
            path: "packages/gopeg_x64.exe",
            name: "gopeg_${nextRelease.version}_windows_x64.exe",
          },
          {
            path: "packages/gopeg_macos",
            name: "gopeg_${nextRelease.version}_macos_universal",
          },
        ],
      },
    ],
    [
      "@semantic-release/git",
      {
        assets: ["CHANGELOG.md"],
        message:
          "chore(release): ${nextRelease.version} [skip ci]\n\n${nextRelease.notes}",
      },
    ],
  ],
};
