# Angular skeleton project

This repository can be used for quick bootstrapping new Angular applications. 

It is rougly based on angular-seed, but is refactored to use Grunt for compilation of JS files, and to use ui-route 
instead of ng router. If is based on angular-bootstrap and bootstrap-material-design as a basic theme.

## Running locally

First install dependencies with 

	npm install
	bower install

Then run `grunt watch` to watch changes to local files

### Using NPM
Run npm start to start a local webserver on localhost:8000. Visit
the site in you browser on http://localhost:8000/app
 
### Docker
 
Run docker-compose up to star an nginx webserver as a dockercontainer. Then attach to localhost (or you docker virtual 
machine) to access your site.

## Compiling JS and CSS
CSS and JS are compiled using Grunt. Files are concatened, so all js files are places in src/js. All .js files in this
folder (And all subfolders) are automatically included in your app. We use bower for js dependency management. 
Bower scripts needs to be manualle added to the Grundfile.js

We use compass for our CSS. The default grunt task includes compiling CSS.

To watch for changes, run `grunt watch`

## Deploymetns with Docker
A simple docker/Dockerfile is included that copies (instead of mounting) all files in app/. This can be used to package the app
as a finished docker container just run `docker build -t pco/my-cool-appname -f docker/Dockerfiles .` to builds the docker
 container. Then run `docker run pco/my-cool-appname`
