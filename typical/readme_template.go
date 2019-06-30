package typical

const readmeTemplate = `<!-- Autogenerated by Typical-Go; Modify 'typical/readme_template.go' to add more content -->
# {{ .Name}}

{{ .Description}}

## Getting Started

This is intruction to start working with the project:
1. Install go 
2. Clone the project 

## Usage 

There is no specific requirement to run the application. Please find the binary at the /bin folder. 

TODO: show list of application usage if any

## Typical Task

The project is empowered by typical-go to provide day-to-day development/admnistration task

TODO: show list of typical task

## Deployment 

There is no specific requirement to deploy the application. Please find the stable binary at release page

## Configuration

Please set the configuration using [Environment Variable](https://12factor.net/config). 
{{ .ConfigDoc }}
`