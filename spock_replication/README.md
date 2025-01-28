# pgedg/Spock Deploy Test

This is a guide to deploying a cluster with two nodes in pgedge spock and testing the financial app.

## Requirements

First for all we use a venv to install nodeCtl for that we install so we'll execute the next commands:

```
$ python3 -m venv .venv
$ source .venv/bin/active

# We install the de nodeCtl

(.venv) $ python3 -c "$(curl -fsSL https://pgedge-download.s3.amazonaws.com/REPO/install.py)"

```

## Run Docker Swarm example

After the install now we execute some commands before to init test. 

For we create a directories to 