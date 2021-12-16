
from invoke import task

@task
def update(c):
    """ Update
    """
    c.run("go get -u -v github.com/vit1251/go-ncursesw@main", pty=True, echo=True)
#    c.run("go get -u -v github.com/vit1251/go-ncursesw@245d690f6bbbda0af23a2e33910e343ed312c633", pty=True, echo=True)

