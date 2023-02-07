# Boiler Plate
https://www.npmjs.com/package/angular-cli-ghpages

## Install requirements
- need angular cli 
```typescript
npm install -g @angular/cli
```
- create default project
```typescript
ng new your-angular-project --defaults
```
- add angular gh pages to host on GH
```typescript
ng add angular-cli-ghpages
```

## Adding new components from cli
- `ng` can help you create new components from the command line
- `ng generate component <xyz>`
- to create a new component it needs a 
| File | Description |
| ----------- | ----------- |
| .html | component view template |
| .scss | component common styles|
| .ts | component file |
| .spec.ts | component unit tests |

## Testing
- `ng test`
: runs tests

## Deploying to Github Pages
- `ng deploy --base_href=/Saucier720/`
- I created a bash script to deploy run `./deploy.sh` while in the app directory to run it