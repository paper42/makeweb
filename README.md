# Makeweb
Static site generator designed to be simple to use, but also powerful.

## Install
```sh
go get github.com/PaperMountainStudio/makeweb-cli
```

## Advantages
- fast
- cheap
- secure
- flexible
- easy to edit
- easy to migrate

## Usage
```sh
mkdir ~/makeweb_example
cd ~/makeweb_example
mkdir input
mkdir templates
touch input/index.html templates/default
```
Now you have created basic project structure. Open ```input/index.html``` in your editor and write: 
```
{
    "title": "Makeweb example"
}
---
content
```
This is your webpage content, but we also need a template. Open ```templates/default``` and write:
```html
<!DOCTYPE html>
<html>
  <head>
    <meta charset="UTF-8" />
    <meta name="viewport" content="width=device-width" />
    <title>{{.title}}</title>
  </head>
  <body>
  {{.text}}
  </body>
</html>
```
As you can see, we use ```{{.title}}``` to get title we specified in webpage content's header.
The ```{{.text}}``` is a special case, this is substituted by the content of webpage - in our case ```input/index.html``` after the ```---```.

Now everything is ready, you can generate it.
```sh
makeweb-cli
```

We can see that a new directory ```output``` has been created.
That is where our generated webpage is.
If you have python3 installed, you can use this to serve it:
```sh
cd output
python3 -m http.server
```
And type ```localhost:8000``` to your browser.
You should see the generated webpage.
