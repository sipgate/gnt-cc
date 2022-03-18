module.exports = {
  extends: "stylelint-config-standard-scss",
  rules: {
    "selector-class-pattern": null,
    "at-rule-no-unknown": null
  },
  "overrides": [
    {
      "files": ["**/*.scss"],
      "customSyntax": "postcss-scss"
    }
  ]
}
