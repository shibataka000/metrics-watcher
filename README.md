# HPA Simulator

## Description
This tool output some information to know HorizontalPodAutoscaler internal behavior. For example,

- Container metrics from [metrics-server](https://github.com/kubernetes-incubator/metrics-server)

This tool support only resource metrics. This tool doesn't support custom metrics and external metrics.

This tool is developed to debug HorizontalPodAutoscaler behavior.

## Requirement
Go 1.13.3 or later

## Usage
```bash
hpasimulator
```

Output

```
```

## Install
```bash
git clone https://github.com/shibataka000/hpa-simulator
cd hpa-simulator
make install
```

## Author
[shibataka000](https://github.com/shibataka000)
