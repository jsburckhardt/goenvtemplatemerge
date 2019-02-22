# GOENVTEMPLATEMERGE

goenvtemplatemerge helps updating config/templates that contain place holders base on the environment variables in the system. Place holders are of the form {{.ENVIRONMENT_VARIABLE_NAME}}

## Getting Started

These instructions will get you a copy of the project up and running on your local machine for development and testing purposes. See deployment for notes on how to deploy the project on a live system.

### Installing dependecies

```
dep ensure
```

### Build program
```
go build
```
or
```
go install
```

## Examples:
### Windows -> PowerShell
```
$env:HASH = "1.14"
$env:NGINX_HOST = "myhost.com"
$env:NGINX_PORT = "80"

goenvtemplatemerge update -templatefile sampleTemplate.yaml
```


### Linux -> Ubuntu
```
export HASH=1.14
export NGINX_HOST=myhost.com
export NGINX_PORT=80

goenvtemplatemerge update -templatefile sampleTemplate.yaml --debug
```

## Build Status

* [![Build status](https://juanburckhardt.visualstudio.com/goenvtemplatemerge/_apis/build/status/goenvtemplatemerge-Go%20(preview)-CI)](https://juanburckhardt.visualstudio.com/goenvtemplatemerge/_build/latest?definitionId=7)

---
## License

[![License](http://img.shields.io/:license-mit-blue.svg?style=flat-square)](http://badges.mit-license.org)

- **[MIT license](http://opensource.org/licenses/mit-license.php)**
