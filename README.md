
# exmo-trading
<p>the application allows you to automate trading on the (https://exmo.me/)</p>

## installation

### via git

<p>you need to have installed golang, git.</p>

    git clone https://github.com/mikhailbuslaev/exmo-trading.git

### via docker

<p>you need to have installed docker.</p>

    docker pull mbesl/exmo-trading:latest
    docker run -it mbesl/exmo-trading

## run

    cd exmo-trading
    go run exmo-trader.go

## run at server

    cd exmo-trading
    nohup go run exmo-trader.go &

## setting

<p>by default, the application starts 2 data processors and one trade 
for each of these processors, but you can add as many data processors and traders as you like.</p>

## example: add new trader

### 1 stage - make new trader config

    cd exmo-trading/configs/trader-configs
    mkfile trader-config.yaml

<p>trader-config.yaml example:</p>
