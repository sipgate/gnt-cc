# gnt-cc Frontend

This is the frontend to gnt-cc.

## Howto build & run

### Development

The application will be available at [http://localhost:3000](http://localhost:3000).

```shell
# Install dependencies
npm install

# Start development server
npm run start

# lint application
npm run lint

# fix linting errors
npm run lint:fix
```

**Note for VSCode:** The ESLint plugin for VSCode requires the following options to be set in `../.vscode/settings.json` for proper linting and to fix errors on save.

```json
{
  "editor.codeActionsOnSave": {
    "source.fixAll.eslint": true
  },
  "eslint.workingDirectories": [
    "./web"
  ]
}
```

### Testing (TODO)

```shell
npm run test
```

### Build (TODO)

```shell
npm run build
```
