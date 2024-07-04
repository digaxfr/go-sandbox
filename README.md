# go-sandbox

In an effort to "do more development" and get a bit more better with Go, for things that I would
write in Bash or Ansible, I will attempt to write it out in Go first even if it does not make any
logical sense to.

## Ideas

Let us start off with a simple Ansible clone. Make remote calls via SSH. This should give me enough
use cases to write code. Keep it simple at first, target one host only.

### Features

* Use golang flag library, no external.
  * Decided that usage of kong is acceptable. Doing this all in flag is a bit cumbersome.

* SSH connection
  * Make a ping/pong test
  * Make remote calls, e.g. ls
  * Template out files
