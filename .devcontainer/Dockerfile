FROM gitrhythm/devbase:latest

RUN sudo apt-get update && \
    sudo apt-get upgrade && \
    brew upgrade

ENV GOENV_ROOT=$HOME/.goenv
ENV PATH $GOENV_ROOT/bin:$PATH

RUN git clone https://github.com/syndbg/goenv.git ${GOENV_ROOT}

RUN echo 'eval "$(goenv init -)"' >> ~/.profile && \
    echo 'export PATH="$PATH:$GOPATH/bin"' >> ~/.profile && \
    # fish用
    echo 'eval (goenv init - | source)' >> ~/.config/fish/config.fish && \
    echo 'set -x PATH $GOPATH/bin $PATH' >> ~/.config/fish/config.fish

# 最新のgoをインストール
RUN goenv install "$(goenv install -l | egrep "^  [0-9]+.[0-9]+.[0-9]+\$" | tail -n 1 | tr -d '  ')" && \
    goenv global $(goenv versions | tr -d '  ')
