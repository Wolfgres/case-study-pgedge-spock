# On premise install pgedge/spock

This is a small tutorial if you want install pgedge/spock on premise for the next tutorial we'll use a: 

- Rocky Linux 9.3
- PostgreSQL 16
- pgedge/spock
- Python

## Installation

For first time we enable some configurations in Operation System. 

https://docs.pgedge.com/platform/prerequisites

### Prerequisites

We need have passwordless in our machines

```
ssh-keygen -t rsa
cd ~/.ssh
cat id_rsa.pub >> authorized_keys
chmod 700 ~/.ssh && chmod 600 ~/.ssh/authorized_keys
```

we install the public key to all nodes in our cluster to avoid request password.

```
ssh-copy-id -i .ssh/id_rsa.pub spock@nodeb
```

Turnoff our firewall and disable firewalld to allow connections

```
sudo systemctl disable firewalld
sudo systemctl stop firewalld
```

And add in we username in our sudoers file to avoid writing the password for each sudo command: 

```
%spock ALL = (ALL) NOPASSWD: ALL
```

In this case we user is **spock**, if you use other set it. 

We need validate that we have binaries like **tar** because maybe raise a error when  we start installation. For any missing execute  this command: 

```
sudo dnf install tar vim python3.9-pip 

```

### Install PgEdge CLI

Now we install the CLI for pgegde:

```
python3 -c "$(curl -fsSL https://pgedge-download.s3.amazonaws.com/REPO/install.py)"
```

Now we need create a configuration to cluster. This configuration we can create with a json template, for that we execute the next commands:

1. Create template for our cluster

```
cd pgedge

./pgedge cluster json-template wfg_cluster wolfgres_db 2 wolfgres_admin 123456 16 5432

```

The above command generate a json file with the template configuration, this file is generated in the cluster/wfg_cluster/wfg_cluster.json. We need edit the json file with configuration that we need use for that: 

```
vim cluster/wfg_cluster/wfg_cluster.json


# set IP each node for example here:
..
       "name": "n1",
      "is_active": "on",
      "public_ip": "X.X.X.X",
      "private_ip": "X.X.X.X",
      "port": "5432",
      "path": "/home/spock/wfg_cluster/n1",
      "backrest": {
        "stanza": "demo_stanza",
        "repo1-path": "/var/lib/pgbackrest",
        "repo1-retention-full": "7",
        "log-level-console": "info",
        "repo1-cipher-type": "aes-256-cbc"
      }

..

```

and after that we validate the json file with this command: 

```
./pgedge cluster json-template wfg_cluster
```

if all going well we can init the cluster with this command: 

```
./pgedge cluster init wfg_cluster
```

a few minutes after the executed command



## Fisrt Test

We create the table schema in both nodes: 

```

CREATE TABLE test_mm_rep (
    id serial PRIMARY KEY,
    node_insert text,
    data_stuff text
);

```

after that we can add table to default replication in both servers:

```
SELECT spock.repset_add_table('default','test_mm_rep', true);
SELECT spock.repset_add_seq('default','test_mm_rep_id_seq', true);

SELECT * FROM spock.replication_set_table ;
SELECT * FROM spock.replication_set_seq ;

```




