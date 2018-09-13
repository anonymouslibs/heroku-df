# A minimal starting template for heroku

This is a minimal template for using Go on heroku. [go-getting-started], which is the official template for getting started with Go on heroku, hasn't been updated in a while, and hence seems to contain some *unnesseary* files.

## How can you remove files from the official template?

The reason it is possible to remove some files from [go-getting-started] is that [heroku-buildpack-go] automatically detects that you're using Go if you use one of the following package managers:

- [go modules]
- [dep]
- [govendor]
- [glide]
- [GB]
- [Godep]

# Setup

```
git clone https://github.com/barskern/go-heroku-mini-template.git
cd go-heroku-mini-template
```

Now you can choose your favorite Go dependency manager and initialize it for you current repository. With [`govendor`](https://github.com/kardianos/govendor) already installed, it would look like this:

```
govendor init
```

The next commands will install the package in `$GOPATH/bin` and run it locally using heroku, hence make sure that `$GOPATH/bin` is in your `$PATH`.

```
go install
heroku local
```

The terminal should tell you where to find your application (normally it will be at `localhost:5000`).

# Configuring your new project

Now that the example should be up and running, you can configure the repository to fit for your application. The following is a checklist of the changes you should make:

1. Move the project to a new folder and name the folder your project name
2. Change the `Procfile` to be `web: <project-folder-name>`
3. Update `app.json` with information relevant to your project (name, description, etc.)
4. Remove the current `.git` folder (`rm -rf .git`) and initialize a brand new git repository (`git init`)

To make sure everything is working after moving and reconfiguring the application, run `go install && heroku local`.

# Making a remote heroku application

When you have made some commits and you're ready to deploy your application to heroku, you can run the following commands:

*This assumes that your logged in to heroku in your terminal (`heroku login`)*

```
heroku create --region <eu or us> <project-name>
git push heroku master
```

If you want to open your newly made application in your default browser you can run:

```
heroku open
```

# Contributions and issues

Contributions and issues are more than welcome!

[heroku-buildpack-go]: https://github.com/heroku/heroku-buildpack-go
[go-getting-started]: https://github.com/heroku/go-getting-started
[go modules]: https://github.com/golang/go/wiki/Modules
[dep]: https://github.com/golang/dep
[govendor]: https://github.com/kardianos/govendor
[glide]: https://github.com/Masterminds/glide
[GB]: https://getgb.io/
[Godep]: https://github.com/tools/godep
