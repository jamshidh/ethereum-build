Running autobuild will build ethereumH in this directory and should work
seamlessly on Ubuntu.

Autobuild sets up the development environment, then calls setup (which sets up
postgres) and finally build (which itself calls cabal-install).

In the future, use "repo sync" to update ethereumH. Beware of making changes
in detached head state.
