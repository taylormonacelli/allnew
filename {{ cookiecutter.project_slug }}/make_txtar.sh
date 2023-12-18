rm -rf /tmp/{{ cookiecutter.project_slug }}

rm -f /tmp/{{ cookiecutter.project_slug }}.tar
rm -f /tmp/filelist.txt

{
    rg --files . \
        | grep -vE '{{ cookiecutter.project_slug }}$' \
        | grep -v README.org \
        | grep -v make_txtar.sh \
        | grep -v go.sum \
        | grep -v go.mod \
        | grep -v Makefile \
        # | grep -v cmd/main.go \
        # | grep -v options/options.go \
        # | grep -v {{ cookiecutter.project_slug }}.go \

} | tee /tmp/filelist.txt
tar -cf /tmp/{{ cookiecutter.project_slug }}.tar -T /tmp/filelist.txt
mkdir -p /tmp/{{ cookiecutter.project_slug }}
tar xf /tmp/{{ cookiecutter.project_slug }}.tar -C /tmp/{{ cookiecutter.project_slug }}
rg --files /tmp/{{ cookiecutter.project_slug }}
txtar-c /tmp/{{ cookiecutter.project_slug }} | pbcopy