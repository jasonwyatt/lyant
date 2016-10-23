#!/usr/local/bin/python
import subprocess
import os
import getpass

ENV_FILENAME = 'env.sh'

def do_setup():
    global ENV_FILENAME
    env = {}

    neo4j_setup(env)

    f = open(ENV_FILENAME, 'w')
    write_env(f, env)
    f.close()

    print("\nYou're good!\nBe sure to run `source ./env.sh` in your shell before starting lyant!")

def write_env(f, env={}):
    write_section_header(f, "Lyant Environment Variables")
    for name, value in env.iteritems():
        f.write("export %s=%s\n" % (name, value))
    f.write("\n")

def write_section_header(f, title):
    f.write("#"*80+"\n")
    f.write("# %s\n" % title)
    f.write("#"*80+"\n\n")

def neo4j_setup(env = {}):
    return_code = subprocess.call('neo4j status', stdout=None, stderr=None, shell=True)
    if return_code != 0:
        print("Can't find neo4j, installing...")
        subprocess.call('brew install neo4j')
        print("Installed. Available at: http://localhost:7474")
        env['LYANT_NEO4J_URL'] = 'http://localhost:7474'
        env['LYANT_NEO4J_USER'] = 'neo4j'
        env['LYANT_NEO4J_PASS'] = 'neo4j'
    else:
        print("Let's set up your neo4j environment variables...")
        env['LYANT_NEO4J_URL'] = get_user_env_var("URL", "LYANT_NEO4J_URL", "http://localhost:7474")
        env['LYANT_NEO4J_USER'] = get_user_env_var("Username", "LYANT_NEO4J_USER", "neo4j")
        env['LYANT_NEO4J_PASS'] = get_user_env_var_password("Password", "LYANT_NEO4J_PASS", "neo4j")

def get_user_env_var(title, env_name, default):
    return raw_input(
        "%s (%s: %s):" % (
            title,
            'current' if os.environ.get(env_name) else 'default',
            os.environ.get(env_name, default)
        )) or default

def get_user_env_var_password(title, env_name, default):
    return getpass.getpass("%s (leave empty for %s):" % (title, 'current' if os.environ.get(env_name) else 'default')) or default

if __name__ == '__main__':
    do_setup()
